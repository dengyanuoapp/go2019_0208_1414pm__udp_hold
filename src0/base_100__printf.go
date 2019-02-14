package main

import "fmt"

type _TdebugPrint bool

var _DebugPrint _TdebugPrint

func ( ___b _TdebugPrint ) Println( ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Println( ___V ) 
    }
} // Println

func ( ___b _TdebugPrint ) Printf( ___Vfmt string , ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Printf( ___Vfmt , ___V ) 
    }
} // Printf


func _DebugPrintTest() {
    __Vp := _DebugPrint 

    _DebugPrint = false
    _DebugPrint.Println( " debug is off " )
    _DebugPrint = true
    _DebugPrint.Println( " debug is on " )

    _DebugPrint = __Vp 
} // _DebugPrintTest
