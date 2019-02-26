package main

import (
    //"encoding/binary"
    //"fmt"
    "math/rand"
    "sync"
    "time"
)

const (
    _VsizeOfRandBuf_block     =   1024 * 8 // * 16
    _VsizeOfRandBuf_byte      =   _VsizeOfRandBuf_block * 16
)
type _TgenRand struct {
    buf             []byte
    remain          uint32
    lock            sync.Mutex
}

var (
    _VgenRand   _TgenRand
)

func _FreGenRandBuf() {
    if 0 == len(_VgenRand.buf) {
        _VgenRand.buf = make( []byte , _VsizeOfRandBuf_byte )
    }
    __Vk := _FgenTimeRand16byte()
    __Viv :=_VgenRand.buf[16:32]

    __Vtmp, __Verr := _FencAesCbc__only( &__Vk , &__Viv , &(_VgenRand.buf) )
    _FerrExit( " 371911 " , __Verr )
    _VgenRand.buf = __Vtmp[16:]
    _FnotEqExit( " 371911 " , _VsizeOfRandBuf_byte , len( _VgenRand.buf ) )

    _VgenRand . remain = _VsizeOfRandBuf_byte
} // _FreGenRandBuf

// gen rand byte slice , size is N
func _FgenRand_nByte( ___Vlen uint32 ) ([]byte) {
    var __Vout []byte

    if ( ___Vlen == 0 ) {
        return __Vout
    }

    _VgenRand.lock.Lock()
    for {
        if ( _VgenRand . remain == 0 ) {
            _FreGenRandBuf()
        }

    }
    _VgenRand.lock.Unlock()

    return __Vout
} // _FgenRand_nByte

//	__Vbyte = _FappendRandLen2byteArr( &__Vbyte , 0 , 16 )
// func math/rand.Uint32() uint32
func _FappendRandLen2byteArr( ___Vbyte *[]byte , ___Vmin , ___Vmax uint32 ) ([]byte) {
    __Vi := rand.Uint32()
    __Vj := ___Vmax - ___Vmin
    __Vk := uint32( __Vj )
    __Vl := __Vi % __Vk
    __Vm := __Vl + ___Vmin

    __Vlen := byte( __Vl )
    __Vlen += byte(___Vmin)

    _Fpt( " 389111 :" , __Vi , __Vj , __Vk, __Vl , __Vm, __Vlen , ___Vmin, ___Vmax)

    if ( __Vlen == 0 ) {
        return *___Vbyte
    }

    __Vb := make( []byte , __Vlen )

    rand.Read( __Vb )

    return append( *___Vbyte , __Vb ... )
} // _FappendRandLen2byteArr

func _FgenTimeRand16byte() ([]byte) {
    var __Vb []byte
    __Vb = []byte(time.Now().String())
    return _Fbase_1101b__gen_md5Only( & __Vb )
} // _FgenTimeRand16byte
