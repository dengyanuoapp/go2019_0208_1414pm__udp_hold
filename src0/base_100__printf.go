package main

import "fmt"
import "os"
import "flag"

var (
    _VprojectName string
)

var (
    _Ppf func (format string, a ...interface{}) (n int, err error)  = fmt.Printf
    _Ppt func (a ...interface{}) (n int, err error)                 = fmt.Print
    _Ppn func (a ...interface{}) (n int, err error)                 = fmt.Println
    _Pdefault func () ()                                            = flag.PrintDefaults
)

func _Pex( ___Vstr string , ___V ... interface{} ) {
    _Ppn( ___Vstr )
    if ___V != nil {
        _Ppn( ___V )
    }
    os.Exit(1)
} // _P.EX

func _Pph( ) {
    if "" != _VprojectName {
        _Ppf( "%s:" , _VprojectName )
    }
} // PH

//flag.PrintDefaults()
func _Pargs(){
    _Pph()
    _Ppn( " cmd paras : [ " , flag.Args() , " ]")
    _Pdefault()
} // PrintArgs

func _FdebugPrintTest() {

    _Ppn( " debug is off 1" )

} // _FdebugPrintTest
