package main

import (
	//"encoding/binary"
	//"fmt"
	"math/rand"
)

//	__Vbyte = _FappendRand2( &__Vbyte , 0 , 16 )
// func math/rand.Uint32() uint32
func _FappendRand2( ___Vbyte *[]byte , ___Vmin , ___Vmax uint32 ) ([]byte) {
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
}
