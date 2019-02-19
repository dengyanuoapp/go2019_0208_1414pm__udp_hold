package main

import "fmt"
import "os"
import "flag"
//import "std"

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

