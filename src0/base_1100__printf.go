package main

import (
    "fmt"
    "os"
    "flag"
    "reflect"
    "runtime"
)

var (
    _VprojectName string
)

var (
    _Ppf        func (___Vfmt string,   ___Vpara ...interface{}) (int, error)   = fmt.Printf
    _Ppt        func (                  ___Vpara ...interface{}) (int, error)   = fmt.Print
    //_Ppn        func (                  ___Vpara ...interface{}) (int, error)   = fmt.Println
    _Pdefault   func ()                                                         = flag.PrintDefaults
)

func _Pn(){
    _Ppf("\n")
} // _Pn

func _Perr( ___Verr error , ___Vfmt string , ___Vpara ... interface{} ) {
    _Ppf( ___Vfmt , ___Vpara ... )
    _Ppt( ___Verr )
    _Pn()
} // _Perr

func _Ppn(___Vpara ...interface{}) (int, error)   {
    __Vi, __Verr := _Ppt( ___Vpara ... )
    _Pn()
    return __Vi, __Verr
} // _Ppn

func _Fpf        (___Vfmt string,   ___Vpara ...interface{}) (int, error)   {
    _Fph()
    return _Ppf( ___Vfmt , ___Vpara ... )
} // _Fpf

func _Pspf        (___Vfmt string,   ___Vpara ...interface{}) string {
    return fmt.Sprintf( ___Vfmt , ___Vpara ... )
} // _Pspf

func _Fpt        (                  ___Vpara ...interface{}) (int, error)   {
    _Fph()
    return _Ppt(           ___Vpara ... )
} // _Fpt

//func log.Fatal(v ...interface{})
//func log.Fatalf(format string, v ...interface{})
func _Fex( ___Vstr string , ___V interface{} ) {
    _Ppn( ___Vstr )
    if ___V != nil {
        _Ppn( ___V )
    }
    os.Exit(1)
} // _Fex

func _FpfN(___Vfmt string, ___Vpara ...interface{}) (int, error)  {
    _Fph()
    __Vn , __Verr := _Ppf( ___Vfmt , ___Vpara... )
    _Pn()
    return __Vn , __Verr ;
} // _FpfN

func _Fph( ) {
    if "" != _VprojectName {
        _Ppf( "%s:" , _VprojectName )
    }
} // _Fph

//flag.PrintDefaults()
func _FPargs(){
    _Fph()
    _Ppn( " cmd paras : [ " , flag.Args() , " ]")
    _Pdefault()
} // _FPargs

func _FdebugPrintTest() {

    _Ppn( " debug is off 1" )

} // _FdebugPrintTest


func _FgetFuncName1(___Va interface{}) string {
    //  func FuncForPC(pc uintptr) *Func
    //func (f *Func) Name() string
    return runtime.FuncForPC(reflect.ValueOf( ___Va ).Pointer()).Name()
} // _FgetFuncName1

func _FgetFuncName2() string {
    // func Caller(skip int) (pc uintptr, file string, line int, ok bool)
    _, file, line, _ := runtime.Caller(1)
    return fmt.Sprintf( "%s:%d", file, line )
} // _FgetFuncName2

func _FgetFuncName3() string {
    return _FgetFrame(1).Function
} // _FgetFuncName3


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
