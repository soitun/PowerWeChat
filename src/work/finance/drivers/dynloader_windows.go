//go:build windows

package drivers

/*
#include <stdint.h>
#include <stdlib.h>
#include <windows.h>
#include <string.h>

// —— 不透明类型前置声明 ——
typedef struct WeWorkFinanceSdk_t WeWorkFinanceSdk_t;
typedef struct MediaData           MediaData_t;

typedef struct {
    char* buf;
    int   len;
} Slice_t;

// 注意：若厂商库用的是 __stdcall，需要把下面的 __cdecl 全部改掉
typedef WeWorkFinanceSdk_t* (__cdecl *fnNewSdk)();
typedef int  (__cdecl *fnInit)(WeWorkFinanceSdk_t*, const char*, const char*);
typedef void (__cdecl *fnDestroySdk)(WeWorkFinanceSdk_t*);

typedef Slice_t* (__cdecl *fnNewSlice)();
typedef void     (__cdecl *fnFreeSlice)(Slice_t*);
typedef int      (__cdecl *fnGetChatData)(WeWorkFinanceSdk_t*, unsigned long long, unsigned int,
                                          const char*, const char*, int, Slice_t*);
typedef int      (__cdecl *fnDecryptData)(const char*, const char*, Slice_t*);

// MediaData
typedef int         (__cdecl *fnGetMediaData)(WeWorkFinanceSdk_t*, const char* indexbuf,
                                              const char* sdkFileid, const char* proxy, const char* passwd,
                                              int timeout, MediaData_t* media_data);
typedef MediaData_t* (__cdecl *fnNewMediaData)();
typedef void         (__cdecl *fnFreeMediaData)(MediaData_t*);
typedef char*        (__cdecl *fnGetOutIndexBuf)(MediaData_t*);
typedef char*        (__cdecl *fnGetData)(MediaData_t*);
typedef int          (__cdecl *fnGetIndexLen)(MediaData_t*);
typedef int          (__cdecl *fnGetDataLen)(MediaData_t*);
typedef int          (__cdecl *fnIsMediaDataFinish)(MediaData_t*);

static HMODULE g_handle = NULL;
static DWORD   g_last_err = 0;

static fnNewSdk            pNewSdk            = NULL;
static fnInit              pInit              = NULL;
static fnDestroySdk        pDestroySdk        = NULL;
static fnNewSlice          pNewSlice          = NULL;
static fnFreeSlice         pFreeSlice         = NULL;
static fnGetChatData       pGetChatData       = NULL;
static fnDecryptData       pDecryptData       = NULL;

// Media
static fnGetMediaData      pGetMediaData      = NULL;
static fnNewMediaData      pNewMediaData      = NULL;
static fnFreeMediaData     pFreeMediaData     = NULL;
static fnGetOutIndexBuf    pGetOutIndexBuf    = NULL;
static fnGetData           pGetData           = NULL;
static fnGetIndexLen       pGetIndexLen       = NULL;
static fnGetDataLen        pGetDataLen        = NULL;
static fnIsMediaDataFinish pIsMediaDataFinish = NULL;

static FARPROC wf_sym(const char* n){ FARPROC p=GetProcAddress(g_handle,n); if(!p){ g_last_err=GetLastError(); } return p; }

static int wf_resolve() {
    g_last_err = 0;
    pNewSdk      = (fnNewSdk)      wf_sym("NewSdk");
    pInit        = (fnInit)        wf_sym("Init");
    pDestroySdk  = (fnDestroySdk)  wf_sym("DestroySdk");
    pNewSlice    = (fnNewSlice)    wf_sym("NewSlice");
    pFreeSlice   = (fnFreeSlice)   wf_sym("FreeSlice");
    pGetChatData = (fnGetChatData) wf_sym("GetChatData");
    pDecryptData = (fnDecryptData) wf_sym("DecryptData");

    // Media
    pGetMediaData      = (fnGetMediaData)      wf_sym("GetMediaData");
    pNewMediaData      = (fnNewMediaData)      wf_sym("NewMediaData");
    pFreeMediaData     = (fnFreeMediaData)     wf_sym("FreeMediaData");
    pGetOutIndexBuf    = (fnGetOutIndexBuf)    wf_sym("GetOutIndexBuf");
    pGetData           = (fnGetData)           wf_sym("GetData");
    pGetIndexLen       = (fnGetIndexLen)       wf_sym("GetIndexLen");
    pGetDataLen        = (fnGetDataLen)        wf_sym("GetDataLen");
    pIsMediaDataFinish = (fnIsMediaDataFinish) wf_sym("IsMediaDataFinish");

    return (pNewSdk && pInit && pDestroySdk && pNewSlice && pFreeSlice && pGetChatData && pDecryptData
            && pGetMediaData && pNewMediaData && pFreeMediaData && pGetOutIndexBuf && pGetData
            && pGetIndexLen && pGetDataLen && pIsMediaDataFinish) ? 0 : -1;
}

static int WfLoad(const char* libpath) {
    if (g_handle) return 0;
    g_last_err = 0;
    g_handle = LoadLibraryA(libpath);
    if (!g_handle) { g_last_err = GetLastError(); return -1; }
    return wf_resolve();
}
static void WfClose() {
    if (g_handle) { FreeLibrary(g_handle); g_handle=NULL; }

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
static DWORD WfLastError() { return g_last_err; }

static WeWorkFinanceSdk_t* WfNewSdk() { return pNewSdk(); }
static int  WfInit(WeWorkFinanceSdk_t* s, const char* a, const char* b) { return pInit(s, a, b); }
static void WfDestroy(WeWorkFinanceSdk_t* s) { pDestroySdk(s); }
static Slice_t* WfNewSlice(){ return pNewSlice(); }
static void WfFreeSlice(Slice_t* s){ pFreeSlice(s); }
static int  WfGetChatData(WeWorkFinanceSdk_t* s, unsigned long long seq, unsigned int limit, const char* proxy, const char* pass, int timeout, Slice_t* out) {
    return pGetChatData(s, seq, limit, proxy, pass, timeout, out);
}
static int  WfDecryptData(const char* rk, const char* enc, Slice_t* out) {
    return pDecryptData(rk, enc, out);
}

// Media 稳定包装
static MediaData_t* WfNewMediaData() { return pNewMediaData(); }
static void        WfFreeMediaData(MediaData_t* m) { pFreeMediaData(m); }
static int         WfGetMediaData(WeWorkFinanceSdk_t* s, const char* indexbuf,
                                  const char* fileid, const char* proxy, const char* pass,
                                  int timeout, MediaData_t* m) {
    return pGetMediaData(s, indexbuf, fileid, proxy, pass, timeout, m);
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

func loadSDK(libPath string) error {
	cpath := C.CString(libPath)
	defer C.free(unsafe.Pointer(cpath))
	if ret := int(C.WfLoad(cpath)); ret != 0 {
		return fmt.Errorf("LoadLibrary failed: winerr=%d", int(C.WfLastError()))
	}
	return nil
}
func closeSDK() { C.WfClose() }

// 与 unix 侧一致的别名/薄封装
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

// Media
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
