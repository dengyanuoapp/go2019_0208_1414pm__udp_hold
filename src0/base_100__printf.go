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

func ( ___b _TdebugPrint ) H( ) {
    if ( "" != _VprojectName ) {
        ___b.nint( _VprojectName , ":") 
    }
} // H
func ( ___b _TdebugPrint ) N( ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Println() 
    }
} // _P.N

func ( ___b _TdebugPrint ) nint( ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Println( ___V ) 
    }
} // _P.nint
func ( ___b _TdebugPrint ) n( ___V ... interface{} ) {
    ___b.H()
    ___b.nint( ___V ) 
} // _P.n

func ( ___b _TdebugPrint ) X( ___V ... interface{} ) {
    ___b . n( ___V ) 
    os.Exit(1)
} // _P.X

func ( ___b _TdebugPrint ) fint( ___Vfmt string , ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Printf( ___Vfmt , ___V ) 
    }
} // _P.fint

func ( ___b _TdebugPrint ) f( ___Vfmt string , ___V ... interface{} ) {
    ___b.H()
    ___b.fint( ___Vfmt , ___V ) 
} // _P.f


func ( ___b _TdebugPrint ) dint( ___Vlen , ___Vmax int , ___Vbuf []byte ) {
    if ( ___b ) {
        if ( ___Vlen > ___Vmax ) { ___Vlen = ___Vmax  }
        for __Vi := 0 ; __Vi < ___Vlen ; __Vi ++ { 
            if __Vi == ___Vlen -1  {
                fmt.Printf( "%02x" , ___Vbuf[__Vi] ) 
            } else {
                fmt.Printf( "%02x " , ___Vbuf[__Vi] ) 
            }
        }
    }
} // _P.dint
func ( ___b _TdebugPrint ) dintN( ___Vlen , ___Vmax int , ___Vbuf []byte ) {
    if ( ___b ) {
        _PT.dint( ___Vlen , ___Vmax , ___Vbuf )
        _PT.N()
    }
} // _P.dintN
func ( ___b _TdebugPrint ) d( ___Vlen , ___Vmax int , ___Vbuf []byte ) {
    ___b.H()
    ___b.dint( ___Vlen , ___Vmax , ___Vbuf )
} // _P.d
func ( ___b _TdebugPrint ) dN( ___Vlen , ___Vmax int , ___Vbuf []byte ) {
    ___b.d(  ___Vlen , ___Vmax , ___Vbuf )
    ___b.N()
} // _P.dN

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
