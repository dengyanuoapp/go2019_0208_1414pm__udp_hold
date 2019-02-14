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

func ( ___b _TdebugPrint ) PH( ) {
    if "" != _VprojectName && ___b  {
        fmt.Printf( "%s:" , _VprojectName )
    }
} // PH
func ( ___b _TdebugPrint ) PN( ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Printf("\n")
    }
} // _P.PN

func ( ___b _TdebugPrint ) pnI( ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Println( ___V )
    }
} // _P.pnI
func ( ___b _TdebugPrint ) pn( ___V ... interface{} ) {
    ___b.PH()
    ___b.pnI( ___V )
} // _P.pn

func ( ___b _TdebugPrint ) EX( ___V ... interface{} ) {
    ___b . pn( ___V )
    os.Exit(1)
} // _P.EX

func ( ___b _TdebugPrint ) pfI( ___Vfmt string , ___V ... interface{} ) {
    if ( ___b ) {
        fmt.Printf( ___Vfmt , ___V )
    }
} // _P.pfI

func ( ___b _TdebugPrint ) pf( ___Vfmt string , ___V ... interface{} ) {
    ___b.PH()
    ___b.pfI( ___Vfmt , ___V )
} // _P.pf
func ( ___b _TdebugPrint ) pfN( ___Vfmt string , ___V ... interface{} ) {
    ___b.pf( ___Vfmt , ___V )
    ___b.PN()
} // _P.pfN


func ( ___b _TdebugPrint ) ddI( ___Vlen , ___Vmax int , ___Vbuf []byte ) {
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
} // _P.ddI
func ( ___b _TdebugPrint ) ddIN( ___Vlen , ___Vmax int , ___Vbuf []byte ) {
    if ( ___b ) {
        _PT.ddI( ___Vlen , ___Vmax , ___Vbuf )
        _PT.PN()
    }
} // _P.ddIN
func ( ___b _TdebugPrint ) dd( ___Vlen , ___Vmax int , ___Vbuf []byte ) {
    ___b.PH()
    ___b.ddI( ___Vlen , ___Vmax , ___Vbuf )
} // _P.dd
func ( ___b _TdebugPrint ) ddN( ___Vlen , ___Vmax int , ___Vbuf []byte ) {
    ___b.dd(  ___Vlen , ___Vmax , ___Vbuf )
    ___b.PN()
} // _P.ddN

//flag.PrintDefaults()
func ( ___b _TdebugPrint ) PrintArgs(){
    ___b.pn( " cmd paras : [ " , flag.Args() , " ]")
    if (___b) {
        flag.PrintDefaults()
    }
} // PrintArgs

func _FdebugPrintTest() {
    __Vp := _P

    _P = false
    _P.pn( " debug is off 1" )
    _P = true
    _P.pn( " debug is on 2" )
    _P = false
    _P.pn( " debug is off 3" )
    _P = true
    _P.pn( " debug is on 4" )

    _P = __Vp
} // _FdebugPrintTest
