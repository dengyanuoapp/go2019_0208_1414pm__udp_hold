package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// you must define a Var and name it : _Vconfig , it must be save as UDP config
	_VS = &_Vself   // _Tself
	_VC = &_Vconfig // _Tconfig
)

var (
	_Ppf      func(___Vfmt string, ___Vpara ...interface{}) (int, error) = fmt.Printf
	_Ppt      func(___Vpara ...interface{}) (int, error)                 = fmt.Print
	_Ppln     func(___Vpara ...interface{}) (int, error)                 = fmt.Println
	_Pdefault func()                                                     = flag.PrintDefaults
)

func _FdebugPrintTest() {
	_Ppln(" debug is off 1")
} // _FdebugPrintTest

//flag.PrintDefaults()
func _FPargs() {
	_Fph()
	_Ppln(" cmd paras : [ ", flag.Args(), " ]")
	_Pdefault()
} // _FPargs

func _Fph() {
	if "" != _Vself.ProjName {
		_Ppf("%s:", _Vself.ProjName)
	}
} // _Fph

func _Fn() {
	_Ppf("\n")
} // _Pn

func _Fpf(___Vfmt string, ___Vpara ...interface{}) {
	_Fph()
	_Ppf(___Vfmt, ___Vpara...)
} // _Fpf

func _Fpt(___Vpara ...interface{}) {
	_Fph()
	_Ppt(___Vpara...)
} // _Fpt

//func log.Fatal(v ...interface{})
//func log.Fatalf(format string, v ...interface{})
func _Fex(___Vstr string) {
	_Ppln(___Vstr)
	os.Exit(1)
} // _Fex
func _Fex1(___Vstr string, ___Vobj interface{}) {
	_Ppf("%s : %v", ___Vstr, ___Vobj)
	os.Exit(1)
} // _Fex1

func __FpfN(___Vfmt string, ___Vpara ...interface{}) {
}
func ___FpfN(___Vfmt string, ___Vpara ...interface{}) {
}
func _CFpfN(___Vfmt string, ___Vpara ...interface{}) {
	_FpfNt(___Vfmt, ___Vpara...)
	_CpfNt(___Vfmt, ___Vpara...)
}
func _FpfNt(___Vfmt string, ___Vpara ...interface{}) {
	_Fph()
	_Ppf("%d ", _FtimeInt())
	_Ppf(___Vfmt+"\n", ___Vpara...)
}
func _FpfN(___Vfmt string, ___Vpara ...interface{}) {
	//_Fph()
	//_Ppf(___Vfmt+"\n", ___Vpara...)
	__Vs := _Sph()
	__Vs += _Spf(___Vfmt+"\n", ___Vpara...)
	_Ppf(__Vs)
} // _FpfN

func _FpfNex(___Vfmt string, ___Vpara ...interface{}) {
	_Fph()
	_Ppf("\n\n"+___Vfmt+"\n\n\n\n", ___Vpara...)
	os.Exit(1)
} // _FpfNex
