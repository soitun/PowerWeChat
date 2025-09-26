//go:build linux || darwin

package drivers

/*
#cgo linux  LDFLAGS: -ldl
#cgo darwin LDFLAGS:

#include <stdint.h>
#include <stdlib.h>
#include <dlfcn.h>
#include <string.h>

// —— 不透明类型前置声明（非常重要，避免退化为 unsafe.Pointer） ——
typedef struct WeWorkFinanceSdk_t WeWorkFinanceSdk_t;
typedef struct MediaData           MediaData_t;

// 与官方一致的最小结构（仅用于 Slice）
typedef struct {
    char* buf;
    int   len;
} Slice_t;

// —— 函数指针类型 ——
// 基本 SDK
typedef WeWorkFinanceSdk_t* (*fnNewSdk)();
typedef int  (*fnInit)(WeWorkFinanceSdk_t*, const char*, const char*);
typedef void (*fnDestroySdk)(WeWorkFinanceSdk_t*);

// Slice
typedef Slice_t* (*fnNewSlice)();
typedef void     (*fnFreeSlice)(Slice_t*);
typedef int      (*fnGetChatData)(WeWorkFinanceSdk_t*, unsigned long long, unsigned int,
                                  const char*, const char*, int, Slice_t*);
typedef int      (*fnDecryptData)(const char*, const char*, Slice_t*);

// MediaData 相关
typedef int         (*fnGetMediaData)(WeWorkFinanceSdk_t*, const char* indexbuf,
                                      const char* sdkFileid, const char* proxy, const char* passwd,
                                      int timeout, MediaData_t* media_data);
typedef MediaData_t* (*fnNewMediaData)();
typedef void         (*fnFreeMediaData)(MediaData_t*);
typedef char*        (*fnGetOutIndexBuf)(MediaData_t*);
typedef char*        (*fnGetData)(MediaData_t*);
typedef int          (*fnGetIndexLen)(MediaData_t*);
typedef int          (*fnGetDataLen)(MediaData_t*);
typedef int          (*fnIsMediaDataFinish)(MediaData_t*);

// —— 动态库全局句柄与符号 ——
static void* g_handle = NULL;
static const char* g_last_err = NULL;

static fnNewSdk            pNewSdk            = NULL;
static fnInit              pInit              = NULL;
static fnDestroySdk        pDestroySdk        = NULL;
static fnNewSlice          pNewSlice          = NULL;
static fnFreeSlice         pFreeSlice         = NULL;
static fnGetChatData       pGetChatData       = NULL;
static fnDecryptData       pDecryptData       = NULL;
static fnGetMediaData      pGetMediaData      = NULL;
static fnNewMediaData      pNewMediaData      = NULL;
static fnFreeMediaData     pFreeMediaData     = NULL;
static fnGetOutIndexBuf    pGetOutIndexBuf    = NULL;
static fnGetData           pGetData           = NULL;
static fnGetIndexLen       pGetIndexLen       = NULL;
static fnGetDataLen        pGetDataLen        = NULL;
static fnIsMediaDataFinish pIsMediaDataFinish = NULL;

static int wf_resolve() {
    g_last_err = NULL;

    pNewSdk            = (fnNewSdk)            dlsym(g_handle, "NewSdk");
    pInit              = (fnInit)              dlsym(g_handle, "Init");
    pDestroySdk        = (fnDestroySdk)        dlsym(g_handle, "DestroySdk");
    pNewSlice          = (fnNewSlice)          dlsym(g_handle, "NewSlice");
    pFreeSlice         = (fnFreeSlice)         dlsym(g_handle, "FreeSlice");
    pGetChatData       = (fnGetChatData)       dlsym(g_handle, "GetChatData");
    pDecryptData       = (fnDecryptData)       dlsym(g_handle, "DecryptData");

    // —— 媒体相关符号 ——
    pGetMediaData      = (fnGetMediaData)      dlsym(g_handle, "GetMediaData");
    pNewMediaData      = (fnNewMediaData)      dlsym(g_handle, "NewMediaData");
    pFreeMediaData     = (fnFreeMediaData)     dlsym(g_handle, "FreeMediaData");
    pGetOutIndexBuf    = (fnGetOutIndexBuf)    dlsym(g_handle, "GetOutIndexBuf");
    pGetData           = (fnGetData)           dlsym(g_handle, "GetData");
    pGetIndexLen       = (fnGetIndexLen)       dlsym(g_handle, "GetIndexLen");
    pGetDataLen        = (fnGetDataLen)        dlsym(g_handle, "GetDataLen");
    pIsMediaDataFinish = (fnIsMediaDataFinish) dlsym(g_handle, "IsMediaDataFinish");

    if (!pNewSdk || !pInit || !pDestroySdk || !pNewSlice || !pFreeSlice || !pGetChatData || !pDecryptData
        || !pGetMediaData || !pNewMediaData || !pFreeMediaData || !pGetOutIndexBuf || !pGetData
        || !pGetIndexLen || !pGetDataLen || !pIsMediaDataFinish) {
        g_last_err = dlerror();
        return -1;
    }
    return 0;
}

// 加载 & 卸载
static int WfLoad(const char* libpath) {
    g_last_err = NULL;
    if (g_handle) return 0;
    g_handle = dlopen(libpath, RTLD_LAZY | RTLD_LOCAL);
    if (!g_handle) { g_last_err = dlerror(); return -1; }
    return wf_resolve();
}
static void WfClose() {
    if (g_handle) { dlclose(g_handle); g_handle = NULL; }

    // 基础符号
    pNewSdk       = NULL;
    pInit         = NULL;
    pDestroySdk   = NULL;
    pNewSlice     = NULL;
    pFreeSlice    = NULL;
    pGetChatData  = NULL;
    pDecryptData  = NULL;

    // 媒体相关符号
    pGetMediaData      = NULL;
    pNewMediaData      = NULL;
    pFreeMediaData     = NULL;
    pGetOutIndexBuf    = NULL;
    pGetData           = NULL;
    pGetIndexLen       = NULL;
    pGetDataLen        = NULL;
    pIsMediaDataFinish = NULL;
}


static const char* WfLastError() { return g_last_err ? g_last_err : ""; }

// —— 对 Go 暴露的“稳定”C 包装 ——（避免在 Go 侧接触函数指针）
static WeWorkFinanceSdk_t* WfNewSdk() { return pNewSdk(); }
static int  WfInit(WeWorkFinanceSdk_t* sdk, const char* c1, const char* c2) { return pInit(sdk, c1, c2); }
static void WfDestroy(WeWorkFinanceSdk_t* sdk) { pDestroySdk(sdk); }
static Slice_t* WfNewSlice() { return pNewSlice(); }
static void WfFreeSlice(Slice_t* s) { pFreeSlice(s); }
static int  WfGetChatData(WeWorkFinanceSdk_t* sdk, unsigned long long seq, unsigned int limit,
                           const char* proxy, const char* pass, int timeout, Slice_t* out) {
    return pGetChatData(sdk, seq, limit, proxy, pass, timeout, out);
}
static int  WfDecryptData(const char* rk, const char* enc, Slice_t* out) {
    return pDecryptData(rk, enc, out);
}

// —— MediaData 稳定包装 ——
static MediaData_t* WfNewMediaData() { return pNewMediaData(); }
static void        WfFreeMediaData(MediaData_t* m) { pFreeMediaData(m); }
static int         WfGetMediaData(WeWorkFinanceSdk_t* sdk, const char* indexbuf,
                                  const char* sdkFileid, const char* proxy, const char* passwd,
                                  int timeout, MediaData_t* media_data) {
    return pGetMediaData(sdk, indexbuf, sdkFileid, proxy, passwd, timeout, media_data);
}
static const char* WfMediaOutIndex(MediaData_t* m, int* plen) {
    if (!m) { *plen=0; return NULL; }
    *plen = pGetIndexLen(m);
    return pGetOutIndexBuf(m);
}
static const char* WfMediaDataPtr(MediaData_t* m, int* plen) {
    if (!m) { *plen=0; return NULL; }
    *plen = pGetDataLen(m);
    return pGetData(m);
}
static int WfMediaIsFinish(MediaData_t* m) { return pIsMediaDataFinish(m); }

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// 供 Go 调用的薄封装
func loadSDK(libPath string) error {
	cpath := C.CString(libPath)
	defer C.free(unsafe.Pointer(cpath))
	if ret := int(C.WfLoad(cpath)); ret != 0 {
		return fmt.Errorf("dlopen failed: %s", C.GoString(C.WfLastError()))
	}
	return nil
}
func closeSDK() { C.WfClose() }

// 暴露 C 侧类型别名（仅在 drivers 包内使用）
type cSDK = C.WeWorkFinanceSdk_t
type cSlice = C.Slice_t
type cMedia = C.MediaData_t

func cNewSdk() *cSDK                      { return C.WfNewSdk() }
func cInit(sdk *cSDK, a, b *C.char) C.int { return C.WfInit(sdk, a, b) }
func cDestroy(sdk *cSDK)                  { C.WfDestroy(sdk) }
func cNewSlice() *cSlice                  { return C.WfNewSlice() }
func cFreeSlice(s *cSlice)                { C.WfFreeSlice(s) }
func cGetChatData(sdk *cSDK, seq C.ulonglong, limit C.uint, proxy, pass *C.char, timeout C.int, out *cSlice) C.int {
	return C.WfGetChatData(sdk, seq, limit, proxy, pass, timeout, out)
}
func cDecryptData(randomKey, encMsg *C.char, out *cSlice) C.int {
	return C.WfDecryptData(randomKey, encMsg, out)
}

// MediaData 辅助封装
func cNewMedia() *cMedia   { return C.WfNewMediaData() }
func cFreeMedia(m *cMedia) { C.WfFreeMediaData(m) }
func cGetMediaDataCall(sdk *cSDK, index, fileID, proxy, pass *C.char, timeout C.int, m *cMedia) C.int {
	return C.WfGetMediaData(sdk, index, fileID, proxy, pass, timeout, m)
}
func cMediaOutIndex(m *cMedia) (ptr *C.char, n int) {
	var ln C.int
	p := C.WfMediaOutIndex(m, &ln)
	return (*C.char)(p), int(ln)
}
func cMediaData(m *cMedia) (ptr *C.char, n int) {
	var ln C.int
	p := C.WfMediaDataPtr(m, &ln)
	return (*C.char)(p), int(ln)
}
func cMediaIsFinish(m *cMedia) int { return int(C.WfMediaIsFinish(m)) }
