package main

import "fmt"
import "os"

type _TdebugPrint bool

var _P _TdebugPrint

func ( ___b _TdebugPrint ) n( ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Println( ___V ) 
    }
} // _P.n

func ( ___b _TdebugPrint ) X( ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Println( ___V ) 
    }
    os.Exit(1)
} // _P.X

func ( ___b _TdebugPrint ) f( ___Vfmt string , ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Printf( ___Vfmt , ___V ) 
    }
} // _P.f

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
