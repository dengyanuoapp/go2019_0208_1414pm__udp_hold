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
    _Ppn        func (                  ___Vpara ...interface{}) (int, error)   = fmt.Println
    _Pdefault   func () ()                                                      = flag.PrintDefaults
)

func _Fex( ___Vstr string , ___V ... interface{} ) {
    _Ppn( ___Vstr )
    if ___V != nil {
        _Ppn( ___V )
    }
    os.Exit(1)
} // _P.EX

func _FpfN(___Vfmt string, ___Vpara ...interface{}) (int, error)  {
    _n, _err := _Ppf( ___Vfmt , ___Vpara... )
    _Ppn()
    //return 0 , nil
    return _n , _err
} // _PpfN

func _Fph( ) {
    if "" != _VprojectName {
        _Ppf( "%s:" , _VprojectName )
    }
} // PH

//flag.PrintDefaults()
func _FPargs(){
    _Fph()
    _Ppn( " cmd paras : [ " , flag.Args() , " ]")
    _Pdefault()
} // PrintArgs

func _FdebugPrintTest() {

    _Ppn( " debug is off 1" )

} // _FdebugPrintTest
