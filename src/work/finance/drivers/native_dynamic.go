package drivers

/*
#include <stdlib.h>
*/
import "C"

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/finance/types"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"
)

// ======================================================================
// 本文件依赖同目录下的 dynloader_unix.go / dynloader_windows.go 里提供的：
//   loadSDK, closeSDK
//   cSDK, cSlice, cMedia  类型别名
//   cNewSdk, cInit, cDestroy
//   cNewSlice, cFreeSlice, cGetChatData, cDecryptData
//   cNewMedia, cFreeMedia, cGetMediaDataCall, cMediaData, cMediaOutIndex, cMediaIsFinish
// ======================================================================

// ----------------------------------------------------------------------------
// 小工具：若传的是目录，按当前平台补库文件名（为了容错；factory 通常已解析为文件）
// ----------------------------------------------------------------------------

func resolveLibPathIfDir(path string) (string, error) {
	if path == "" {
		return "", errors.New("libPath 不能为空")
	}
	l := strings.ToLower(path)
	if strings.HasSuffix(l, ".so") || strings.HasSuffix(l, ".dll") || strings.HasSuffix(l, ".dylib") {
		return path, nil
	}
	// 传的是目录：按平台补文件名
	var name string
	switch types.SDKPlatform(runtime.GOOS + "_" + runtime.GOARCH) {
	case types.SDKPlatformLinuxAMD64, types.SDKPlatformLinuxARM64:
		name = "libWeWorkFinanceSdk_C.so"
	case types.SDKPlatformDarwinARM64:
		name = "libWeWorkFinanceSdk_C.so"
	case types.SDKPlatformWindowsAMD64:
		name = "WeWorkFinanceSdk.dll"
	default:
		return "", fmt.Errorf("未支持的平台: %s_%s", runtime.GOOS, runtime.GOARCH)
	}
	return filepath.Join(path, name), nil
}

// ----------------------------------------------------------------------------
// 客户端
// ----------------------------------------------------------------------------

type nativeClient struct{ sdk *cSDK }

// NewDynamicWithPath 由 factory 调用：传入最终库路径（目录或文件），以及 CorpID/CorpSecret
func NewDynamicWithPath(libPath, corpID, corpSecret string) (*nativeClient, error) {
	if strings.TrimSpace(corpID) == "" || strings.TrimSpace(corpSecret) == "" {
		return nil, fmt.Errorf("corpID/corpSecret 不能为空")
	}
	lp, err := resolveLibPathIfDir(libPath)
	if err != nil {
		return nil, err
	}
	if err := loadSDK(lp); err != nil {
		return nil, fmt.Errorf("加载 SDK 动态库失败: %v (path=%s)", err, lp)
	}

	cCorp := C.CString(corpID)
	cSec := C.CString(corpSecret)
	defer C.free(unsafe.Pointer(cCorp))
	defer C.free(unsafe.Pointer(cSec))

	sdk := cNewSdk()
	if sdk == nil {
		return nil, errors.New("NewSdk 返回空指针")
	}
	if ret := int(cInit(sdk, cCorp, cSec)); ret != 0 {
		cDestroy(sdk)
		return nil, fmt.Errorf("Init 失败: ret=%d", ret)
	}
	return &nativeClient{sdk: sdk}, nil
}

func (c *nativeClient) Close() {
	if c == nil || c.sdk == nil {
		return
	}
	cDestroy(c.sdk)
	c.sdk = nil
	// 如需在进程生命周期内卸载库，可在进程退出时调用 closeSDK()
}

// ----------------------------------------------------------------------------
// GetChatData：接口要求返回 []*types.ResponseChatData
// SDK 回包是 {errcode, errmsg, chatdata:[...] }，这里直接把 chatdata 切片取出返回
// ----------------------------------------------------------------------------

func (c *nativeClient) GetChatData(seq uint64, limit uint32, proxy, passwd string, timeoutSec int) ([]*types.ResponseChatData, error) {
	if c == nil || c.sdk == nil {
		return nil, errors.New("client not initialized")
	}
	if limit == 0 || limit > 1000 {
		return nil, fmt.Errorf("invalid limit=%d (1..1000)", limit)
	}

	cProxy := C.CString(strings.TrimSpace(proxy))
	cPass := C.CString(strings.TrimSpace(passwd))
	defer C.free(unsafe.Pointer(cProxy))
	defer C.free(unsafe.Pointer(cPass))

	out := cNewSlice()
	if out == nil {
		return nil, errors.New("NewSlice failed")
	}
	defer cFreeSlice(out)

	ret := int(cGetChatData(c.sdk, C.ulonglong(seq), C.uint(limit), cProxy, cPass, C.int(timeoutSec), out))
	if ret != 0 {
		return nil, fmt.Errorf("GetChatData failed: ret=%d", ret)
	}

	body := C.GoBytes(unsafe.Pointer(out.buf), C.int(out.len))

	// 先解到 envelope，再把 chatdata 切片拿出来
	var envelope struct {
		ErrCode  int                       `json:"errcode"`
		ErrMsg   string                    `json:"errmsg"`
		ChatData []*types.ResponseChatData `json:"chatdata"`
	}
	if err := json.Unmarshal(body, &envelope); err != nil {
		return nil, fmt.Errorf("unmarshal chatdata envelope 失败: %w, raw=%s", err, string(body))
	}
	if envelope.ErrCode != 0 {
		return envelope.ChatData, fmt.Errorf("wecom api error: errcode=%d, errmsg=%s", envelope.ErrCode, envelope.ErrMsg)
	}
	return envelope.ChatData, nil
}

// ----------------------------------------------------------------------------
// DecryptData（接口签名版）：
// 1) 用 RSA 私钥（PKCS#1/PKCS#8）解开 encrypt_random_key（base64），得到 randomKey（二进制）
// 2) 调 C 的 DecryptData(randomKey, encryptMsg) 得到明文 JSON
// 3) 反序列化为 types.ChatMessage
// ----------------------------------------------------------------------------

func (c *nativeClient) DecryptData(encryptRandomKey string, encryptMsg string, specificPrivateKey string) (*types.ChatMessage, error) {
	if c == nil || c.sdk == nil {
		return nil, errors.New("client not initialized")
	}
	if strings.TrimSpace(encryptRandomKey) == "" || strings.TrimSpace(encryptMsg) == "" || strings.TrimSpace(specificPrivateKey) == "" {
		return nil, fmt.Errorf("encryptRandomKey/encryptMsg/specificPrivateKey 不能为空")
	}

	// a) 解析 PEM 私钥（支持 PKCS#1 / PKCS#8）
	privKey, err := parseRSAPrivateKeyFromPEM([]byte(specificPrivateKey))
	if err != nil {
		return nil, err
	}

	// b) base64 decode encrypt_random_key，再用 RSA/PKCS1 v1.5 解开得到 randomKey
	cipherBytes, err := base64.StdEncoding.DecodeString(encryptRandomKey)
	if err != nil {
		return nil, fmt.Errorf("encrypt_random_key base64 解码失败: %w", err)
	}
	randomKey, err := rsa.DecryptPKCS1v15(nil, privKey, cipherBytes)
	if err != nil {
		return nil, fmt.Errorf("RSA PKCS1 解密失败: %w", err)
	}

	// c) 调 C 的 DecryptData（注意 C 原型不需要 sdk 指针）
	cRand := (*C.char)(unsafe.Pointer(&randomKey[0]))
	cMsg := C.CString(encryptMsg)
	defer C.free(unsafe.Pointer(cMsg))

	out := cNewSlice()
	if out == nil {
		return nil, errors.New("NewSlice failed")
	}
	defer cFreeSlice(out)

	if ret := int(cDecryptData(cRand, cMsg, out)); ret != 0 {
		return nil, fmt.Errorf("DecryptData failed: ret=%d", ret)
	}

	plain := C.GoBytes(unsafe.Pointer(out.buf), C.int(out.len))

	// d) 反序列化为强类型
	var msg types.ChatMessage
	if err := json.Unmarshal(plain, &msg); err != nil {
		return nil, fmt.Errorf("反序列化 ChatMessage 失败: %w, raw=%s", err, string(plain))
	}
	return &msg, nil
}

// parseRSAPrivateKeyFromPEM：支持 PKCS#1 与 PKCS#8（未加密）
func parseRSAPrivateKeyFromPEM(pemBytes []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, fmt.Errorf("解析私钥失败：PEM decode 为空")
	}
	// PKCS#1
	if k, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return k, nil
	}
	// PKCS#8
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err == nil {
		if rk, ok := key.(*rsa.PrivateKey); ok {
			return rk, nil
		}
		return nil, fmt.Errorf("PKCS#8 不是 RSA 私钥")
	}
	return nil, fmt.Errorf("不支持的私钥格式或解析失败")
}

// ----------------------------------------------------------------------------
// GetMediaData：媒体分片拉取（图片/文件等）
// - indexBuf：首次可以传空串；之后传上一次返回的 OutIndexBuf
// - 自动循环直到 is_finish == 1，期间把 data 片段累加到一起再返回
// ----------------------------------------------------------------------------

func (c *nativeClient) GetMediaData(indexBuf, sdkFileId, proxy, passwd string, timeout int) (*types.MediaData, error) {
	if c == nil || c.sdk == nil {
		return nil, fmt.Errorf("client not initialized")
	}
	if strings.TrimSpace(sdkFileId) == "" {
		return nil, fmt.Errorf("sdkFileId 不能为空")
	}

	cProxy := C.CString(proxy)
	cPass := C.CString(passwd)
	defer C.free(unsafe.Pointer(cProxy))
	defer C.free(unsafe.Pointer(cPass))

	var acc []byte
	nextIndex := indexBuf
	const maxLoop = 10000
	loop := 0

	for {
		loop++
		if loop > maxLoop {
			return nil, fmt.Errorf("GetMediaData 循环超过 %d 次，可能 SDK 异常", maxLoop)
		}

		var cIndex *C.char
		if nextIndex != "" {
			cIndex = C.CString(nextIndex)
		} else {
			cIndex = C.CString("")
		}
		cFile := C.CString(sdkFileId)

		md := cNewMedia()
		if md == nil {
			C.free(unsafe.Pointer(cIndex))
			C.free(unsafe.Pointer(cFile))
			return nil, fmt.Errorf("NewMediaData 返回空指针")
		}

		ret := int(cGetMediaDataCall(c.sdk, cIndex, cFile, cProxy, cPass, C.int(timeout), md))

		// 释放本轮入参
		C.free(unsafe.Pointer(cIndex))
		C.free(unsafe.Pointer(cFile))

		if ret != 0 {
			cFreeMedia(md)
			return nil, fmt.Errorf("GetMediaData 调用失败：ret=%d", ret)
		}

		// 读取分片数据
		if p, n := cMediaData(md); n > 0 && p != nil {
			chunk := C.GoBytes(unsafe.Pointer(p), C.int(n))
			if len(chunk) > 0 {
				acc = append(acc, chunk...)
			}
		}

		// 下次索引
		var outIdx string
		if p, n := cMediaOutIndex(md); n > 0 && p != nil {
			outIdx = C.GoStringN(p, C.int(n))
		}

		finish := cMediaIsFinish(md)
		cFreeMedia(md)

		if finish == 1 {
			return &types.MediaData{
				Data:        acc,
				OutIndexBuf: outIdx, // 一般完成时为空；保留以便上层断点续拉
				IsFinish:    true,   // 若你的 types.MediaData 用的是 int/uint8，请改成 1
			}, nil
		}

		// 未完成则继续
		nextIndex = outIdx
	}
}
