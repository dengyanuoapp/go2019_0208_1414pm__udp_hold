package main

import (
	"reflect"
	"runtime"
	"strings"
	"sync"
	//"fmt"
)

func _FgetFuncName1(___Va interface{}) string {
	//  func FuncForPC(pc uintptr) *Func
	//func (f *Func) Name() string
	__Vstr1 := runtime.FuncForPC(reflect.ValueOf(___Va).Pointer()).Name()
	__Vstr2 := string([]byte(__Vstr1)[strings.LastIndexByte(__Vstr1, '.')+1:])
	return __Vstr2
} // _FgetFuncName1

func _FgetFuncName2() string {
	// func Caller(skip int) (pc uintptr, file string, line int, ok bool)
	_, file, line, _ := runtime.Caller(1)
	return _Spf("%s:%d", file, line)
} // _FgetFuncName2

func _FgetFuncName3() string {
	__Vstr1 := _FgetFrame(1).Function
	__Vstr2 := string([]byte(__Vstr1)[strings.LastIndexByte(__Vstr1, '.')+1:])
	return __Vstr2
} // _FgetFuncName3

var __VfpndbI641 int64
var __VfpndbI642 int64
var __VpfNdbMap map[string]int = make(map[string]int)
var __VpfNdbMaPmux sync.Mutex

func _FgetFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and _FgetFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
} // _FgetFrame

func _FpfNonce(___Vfmt string, ___Vpara ...interface{}) {
	__Vstr1 := _FgetFrame(1).Function
	__Vstr2 := string([]byte(__Vstr1)[strings.LastIndexByte(__Vstr1, '.')+1:])
	__VfpndbI642 = __VfpndbI641
	__VfpndbI641 = _FtimeI64()
	if 0 == __VfpndbI642 {
		__VfpndbI642 = __VfpndbI641
	}
	_, _, __Vline, _ := runtime.Caller(1)
	__Vstr3 := _Spf("%s_%d", __Vstr2, __Vline)

	__VpfNdbMaPmux.Lock()
	__Vi3, __Vok3 := __VpfNdbMap[__Vstr3]
	if true == __Vok3 {
		__Vi3 += 1
	} else {
		__Vi3 = 1
	}
	__VpfNdbMap[__Vstr3] = __Vi3
	__VpfNdbMaPmux.Unlock()

	if __Vi3 == 1 {
		//_FpfN(___Vfmt+" : %s", ___Vpara, _FgetFuncName3())
		_FpfN("prONCE :"+___Vfmt, ___Vpara...)
		//_Ppf("%11d:%2d: %s_%d\n\n", __VfpndbI641, __VfpndbI641-__VfpndbI642, __Vstr2, __Vline)
		//_Ppf("%11d:%2d: %s(%d)\n\n", __VfpndbI641, __VfpndbI641-__VfpndbI642, __Vstr2, __Vi3)
	}
}

func _SpfNdb(___Vfmt string, ___Vpara ...interface{}) string {
	__Vstr1 := _FgetFrame(2).Function
	__Vstr2 := string([]byte(__Vstr1)[strings.LastIndexByte(__Vstr1, '.')+1:])
	__VfpndbI642 = __VfpndbI641
	__VfpndbI641 = _FtimeI64()
	if 0 == __VfpndbI642 {
		__VfpndbI642 = __VfpndbI641
	}
	_, _, __Vline, _ := runtime.Caller(1)
	__Vstr3 := _Spf("%s_%d", __Vstr2, __Vline)

	__VpfNdbMaPmux.Lock()
	__Vi3, __Vok3 := __VpfNdbMap[__Vstr3]
	if true == __Vok3 {
		__Vi3 += 1
	} else {
		__Vi3 = 1
	}
	__VpfNdbMap[__Vstr3] = __Vi3
	__VpfNdbMaPmux.Unlock()

	__Vso := _Spf("(%11d:%d) %s(%d):", __VfpndbI641, __VfpndbI641-__VfpndbI642, __Vstr2, __Vi3)
	__Vso += _Spf(___Vfmt, ___Vpara...)
	return __Vso
}

func _CFpfNdb(___Vfmt string, ___Vpara ...interface{}) {
	__Vstr1 := _SpfNdb(___Vfmt, ___Vpara...) + "\n\n"
	_Ppf(__Vstr1)
	_CpfS(&__Vstr1)
}
func _CpfNdb(___Vfmt string, ___Vpara ...interface{}) {
	__Vstr1 := _SpfNdb(___Vfmt, ___Vpara...) + "\n\n"
	_CpfS(&__Vstr1)
}
func _FpfNdb(___Vfmt string, ___Vpara ...interface{}) {
	__Vstr1 := _SpfNdb(___Vfmt, ___Vpara...) + "\n\n"
	_Ppf(__Vstr1)
}

/*
func _FpfNdb(___Vfmt string, ___Vpara ...interface{}) {
	__Vstr1 := _FgetFrame(1).Function
	__Vstr2 := string([]byte(__Vstr1)[strings.LastIndexByte(__Vstr1, '.')+1:])
	__VfpndbI642 = __VfpndbI641
	__VfpndbI641 = _FtimeI64()
	if 0 == __VfpndbI642 {
		__VfpndbI642 = __VfpndbI641
	}
	_, _, __Vline, _ := runtime.Caller(1)
	__Vstr3 := _Spf("%s_%d", __Vstr2, __Vline)

	__VpfNdbMaPmux.Lock()
	__Vi3, __Vok3 := __VpfNdbMap[__Vstr3]
	if true == __Vok3 {
		__Vi3 += 1
	} else {
		__Vi3 = 1
	}
	__VpfNdbMap[__Vstr3] = __Vi3
	__VpfNdbMaPmux.Unlock()

	//_FpfN(___Vfmt+" : %s", ___Vpara, _FgetFuncName3())
	_Fpf(___Vfmt, ___Vpara...)
	//_Ppf("%11d:%2d: %s_%d\n\n", __VfpndbI641, __VfpndbI641-__VfpndbI642, __Vstr2, __Vline)
	_Ppf("%11d:%2d: %s(%d)\n\n", __VfpndbI641, __VfpndbI641-__VfpndbI642, __Vstr2, __Vi3)
}
*/
