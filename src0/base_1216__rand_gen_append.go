package main

import (
	//"encoding/binary"
	//"fmt"
	"math/rand"
)

//	__Vbyte = _FappendRand2( &__Vbyte , 0 , 16 )
// func math/rand.Uint32() uint32
func _FappendRand2( ___Vbyte *[]byte , ___Vmin , ___Vmax byte ) ([]byte) {
    __Vlen := byte( rand.Uint32() )
    if ( __Vlen == 0 ) {
        return *___Vbyte
    }

    __Vb := make( []byte , __Vlen )

    rand.Read( __Vb )

    return append( *___Vbyte , __Vb ... )
}
