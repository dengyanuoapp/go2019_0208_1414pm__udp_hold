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

var _PPpf func (format string, a ...interface{}) (n int, err error) = fmt.Printf

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

func ( ___b _TdebugPrint ) prI( ___V ... interface{} ) {
    if ( ___b ) {
        //fmt.Println( ___V )
        /*
        fmt.Print( "<" )
        fmt.Print( ___V )
        fmt.Print( ">" )
        */
        //for __Vidx , __Varg := range ___V {
        /*
        for _ , __Varg := range ___V {
            //fmt.Printf( " {%d:%T,%x} " , __Vidx , __Varg , __Varg )
            fmt.Printf( "%s" , __Varg )
        }
        */
        fmt.Print( ___V )
    }
} // _P.prI
func ( ___b _TdebugPrint ) pr( ___V ... interface{} ) {
    ___b.PH()
    ___b.prI( ___V )
} // _P.pr
func ( ___b _TdebugPrint ) prN( ___V ... interface{} ) {
    ___b.pr( ___V )
    ___b.PN()
} // _P.prN

func ( ___b _TdebugPrint ) EX( ___V ... interface{} ) {
    ___b . pr( ___V )
    os.Exit(1)
} // _P.EX

func ( ___b _TdebugPrint ) pfI( ___Vfmt string , ___V ...interface{} ) {
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
    if (___b) {
        _PT.PH()
        _PT.prI( " cmd paras : [ " , flag.Args() , " ]")
        flag.PrintDefaults()
    }
} // PrintArgs

func _FdebugPrintTest() {
    __Vp := _P

    _P = false
    _P.pr( " debug is off 1" )
    _P = true
    _P.pr( " debug is on 2" )
    _P = false
    _P.pr( " debug is off 3" )
    _P = true
    _P.pr( " debug is on 4" )

    _P = __Vp
} // _FdebugPrintTest
