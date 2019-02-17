package main

//import "fmt"
//import "os"
//import "flag"
//import "std"

// dump hex
func _PpdL( ___Vlen int , ___Vbuf *[]byte) {
    __Xlen := len( *___Vbuf )
    if __Xlen > ___Vlen { __Xlen = ___Vlen }
    _Ppf( "(%d)", __Xlen )
    __Xlen --
    for __Vi := 0 ; __Vi <= __Xlen ; __Vi ++ {
        if __Vi == __Xlen {
            _Ppf( "%02x", (*___Vbuf)[__Vi] )
        } else {
            _Ppf( "%02x ", (*___Vbuf)[__Vi] )
        }
    }
} // _PpdL

//func _FcopyByte( &((*___VacceptTCP).Vbuf2) , &((*___VacceptTCP).Vbuf), (*___VacceptTCP).Vlen )
//(*___VacceptTCP).Vbuf2 = make([]byte , (*___VacceptTCP).Vlen ); copy( (*___VacceptTCP).Vbuf2 , (*___VacceptTCP).Vbuf )
func _FcopyByte( ___dst *[]byte , ___src *[]byte , ___len int ) {
    __len2 := cap( (*___src) ) 
    if ( ___len > __len2 ) { ___len = __len2 }
    (*___dst) = make([]byte , ___len )
    copy( (*___dst) , (*___src) )
}// _FcopyByte

func _PpdLN( ___Vlen int , ___Vbuf *[]byte) {
    _PpdL( ___Vlen , ___Vbuf )
    _Pn()
} // _PpdLN
func _Fpd( ___Vlen int , ___Vbuf *[]byte) {
    _Fph()
    _PpdL( ___Vlen , ___Vbuf )
} // _Fpd
func _FpdN( ___Vlen int , ___Vbuf *[]byte) {
    _Fpd( ___Vlen , ___Vbuf )
    _Pn()
} // _FpdN

func _FmakeByte( ___Vlen int , ___VbyteSlice interface{} ) []byte {
    __Vbyte := make( []byte , ___Vlen )
    copy( __Vbyte , ___VbyteSlice )
    return __Vbyte
} // _FmakeByte
