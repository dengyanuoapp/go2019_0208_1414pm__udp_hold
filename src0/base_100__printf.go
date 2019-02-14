package main

import "fmt"
import "os"
import "flag"

type _TdebugPrint bool

var (
    _P _TdebugPrint
    _PT _TdebugPrint = true
    _PF _TdebugPrint = false
    _VprojectName string
)

func ( ___b _TdebugPrint ) n( ___V ... interface{} ) {
    if ( ___b ) {
        if ( "" == _VprojectName ) {
            fmt.Println( ___V ) 
        } else {
            fmt.Println( _VprojectName , ___V ) 
        }
    }
} // _P.n

func ( ___b _TdebugPrint ) X( ___V ... interface{} ) {
    ___b . n( ___V ) 
    os.Exit(1)
} // _P.X

func ( ___b _TdebugPrint ) f( ___Vfmt string , ___V ... interface{} ) {
    if ( ___b ) {
        if ( "" == _VprojectName ) {
            fmt.Printf( ___Vfmt , ___V ) 
        } else {
            fmt.Printf( _VprojectName + ":" + ___Vfmt , ___V ) 
        }
    }
} // _P.f

//flag.PrintDefaults()
func ( ___b _TdebugPrint ) PrintArgs(){
    ___b.n( " cmd paras : [ " , flag.Args() , " ]")
    if (___b) {
        flag.PrintDefaults()
    }
} // PrintArgs

func _FdebugPrintTest() {
    __Vp := _P 

    _P = false
    _P.n( " debug is off 1" )
    _P = true
    _P.n( " debug is on 2" )
    _P = false
    _P.n( " debug is off 3" )
    _P = true
    _P.n( " debug is on 4" )

    _P = __Vp 
} // _FdebugPrintTest
