package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

var (
	// you must define a Var and name it : _Vconfig , it must be save as UDP config
	_VS = &_Vself   // _Tself
	_VC = &_Vconfig // _Tconfig
)

var (
	_Ppf func(___Vfmt string, ___Vpara ...interface{}) (int, error) = fmt.Printf
	_Ppt func(___Vpara ...interface{}) (int, error)                 = fmt.Print
	//_Ppn        func (                  ___Vpara ...interface{}) (int, error)   = fmt.Println
	_Pdefault func() = flag.PrintDefaults
)

func _Pn() {
	_Ppf("\n")
} // _Pn

func _Perr(___Verr error, ___Vfmt string, ___Vpara ...interface{}) {
	_Ppf(___Vfmt, ___Vpara...)
	_Ppt(___Verr)
	_Pn()
} // _Perr

func _Ppn(___Vpara ...interface{}) (int, error) {
	__Vi, __Verr := _Ppt(___Vpara...)
	_Pn()
	return __Vi, __Verr
} // _Ppn

func _Fpf(___Vfmt string, ___Vpara ...interface{}) (int, error) {
	_Fph()
	return _Ppf(___Vfmt, ___Vpara...)
} // _Fpf
func _Spf(___Vfmt string, ___Vpara ...interface{}) string {
	return _Pspf(___Vfmt, ___Vpara...)
} // _Spf

func _Pspf(___Vfmt string, ___Vpara ...interface{}) string {
	return fmt.Sprintf(___Vfmt, ___Vpara...)
} // _Pspf

func _Fpt(___Vpara ...interface{}) (int, error) {
	_Fph()
	return _Ppt(___Vpara...)
} // _Fpt
func _Spt(___Vpara ...interface{}) string {
	return _Sph() + fmt.Sprint(___Vpara...)
} // _Spt

//func log.Fatal(v ...interface{})
//func log.Fatalf(format string, v ...interface{})
func _Fex(___Vstr string, ___V interface{}) {
	_Ppn(___Vstr)
	if ___V != nil {
		_Ppn(___V)
	}
	os.Exit(1)
} // _Fex

func _Fex1(___Vstr string) {
	_Ppn(___Vstr)
	os.Exit(1)
} // _Fex1

func _FpfN(___Vfmt string, ___Vpara ...interface{}) (int, error) {
	_Fph()
	__Vn, __Verr := _Ppf(___Vfmt, ___Vpara...)
	_Pn()
	return __Vn, __Verr
} // _FpfN
func _CpfN(___Vfmt string, ___Vpara ...interface{}) {
	__Vs := _Sph()
	__Vs += _Spf(___Vfmt+"\n", ___Vpara...)
	//_Clog <- __Vs
} // _CpfN

func _FpfNex(___Vfmt string, ___Vpara ...interface{}) {
	_Fph()
	_Ppf("\n\n"+___Vfmt+"\n\n\n\n", ___Vpara...)
	os.Exit(1)
} // _FpfNex

func _Fph() {
	if "" != _Vself.ProjName {
		_Ppf("%s:", _Vself.ProjName)
	}
} // _Fph
func _Sph() string {
	if "" != _Vself.ProjName {
		return _Pspf("%s:", _Vself.ProjName)
	}
	return ""
} // _Fph

//flag.PrintDefaults()
func _FPargs() {
	_Fph()
	_Ppn(" cmd paras : [ ", flag.Args(), " ]")
	_Pdefault()
} // _FPargs

func _FdebugPrintTest() {

	_Ppn(" debug is off 1")

} // _FdebugPrintTest

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
	return fmt.Sprintf("%s:%d", file, line)
} // _FgetFuncName2

func _FgetFuncName3() string {
	__Vstr1 := _FgetFrame(1).Function
	__Vstr2 := string([]byte(__Vstr1)[strings.LastIndexByte(__Vstr1, '.')+1:])
	return __Vstr2
} // _FgetFuncName3

var __VfpndbI641 int64
var __VfpndbI642 int64
var __VpfNdbMap map[string]int

func _FpfNdb(___Vfmt string, ___Vpara ...interface{}) {
	__Vstr1 := _FgetFrame(1).Function
	__Vstr2 := string([]byte(__Vstr1)[strings.LastIndexByte(__Vstr1, '.')+1:])
	//_FpfN(___Vfmt+" : %s", ___Vpara, _FgetFuncName3())
	_Fpf(___Vfmt, ___Vpara...)
	__VfpndbI642 = __VfpndbI641
	__VfpndbI641 = _FtimeI64()
	if 0 == __VfpndbI642 {
		__VfpndbI642 = __VfpndbI641
	}
	_, _, __Vline, _ := runtime.Caller(1)
	__Vstr3 := _Pspf("%s_%d", __Vstr2, __Vline)

	if nil == __VpfNdbMap {
		__VpfNdbMap = make(map[string]int)
	}
	__Vi3, __Vok3 := __VpfNdbMap[__Vstr3]
	if true == __Vok3 {
		__Vi3 += 1
	} else {
		__Vi3 = 1
	}

	__VpfNdbMap[__Vstr3] = __Vi3

	//_Ppf("%11d:%2d: %s_%d\n\n", __VfpndbI641, __VfpndbI641-__VfpndbI642, __Vstr2, __Vline)
	_Ppf("%11d:%2d: %s(%d)\n\n", __VfpndbI641, __VfpndbI641-__VfpndbI642, __Vstr2, __Vi3)
}

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
