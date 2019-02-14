package main

import "fmt"
import "os"
import "flag"
//import "std"

var (
    _VprojectName string
)

var (
    _Ppf        func (format string, a ...interface{}) (n int, err error)   = fmt.Printf
    _PpfN       func (format string, a ...interface{}) (n int, err error)   = _PpfN___
    _Ppt        func (a ...interface{}) (n int, err error)                  = fmt.Print
    _Ppn        func (a ...interface{}) (n int, err error)                  = fmt.Println
    _Pdefault   func () ()                                                  = flag.PrintDefaults
)

func _Pex( ___Vstr string , ___V ... interface{} ) {
    _Ppn( ___Vstr )
    if ___V != nil {
        _Ppn( ___V )
    }
    os.Exit(1)
} // _P.EX

func _PpfN___(format string, a ...interface{}) (n int, err error)  {
    //_n, err := _Ppf( format , a )
    //_n, err := std.Output(2, fmt.Sprintf(format, a...))
    //fmt.Sprintf(format, a)
    //os.Stdout.Write( _n )
    //_n, err := std.Output(2, fmt.Sprintf(format, a...))
    //std.Output(2, fmt.Sprintf(format, a))
    //os.Stdout.Write( []byte(fmt.Sprintf(format, a)) )
    //aa:= []byte(fmt.Sprintf(format, a...)) 
    //os.Stdout.Write( aa[2:] )
    _n, _err := _Ppf( format , a... )
    _Ppn()
    //return 0 , nil
    return _n , _err
} // _PpfN

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
