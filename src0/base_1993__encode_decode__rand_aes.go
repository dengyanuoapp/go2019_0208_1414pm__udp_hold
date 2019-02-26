package main

import (
	"time"
	"crypto/md5"
)

var (
	_VencAesRand_iv128   []byte = make([]byte, 16)
	_VencAesRand_last128 []byte = make([]byte, 16)
)

func _FencAesRand__gen_iv() {
	var __Vb []byte
	__Vb = append(_VencAesRand_iv128, []byte(time.Now().String())...)
	//_FpfN( " 192391 Vb len (%d) ", len( __Vb ) )

	__Vb = _Fbase_1101b__gen_md5Only(append(__Vb))
	//_FpfN( " 192392 Vb len (%d) , %0x", len( __Vb ), __Vb )

	for __Vi := 0; __Vi < 16; __Vi++ {
		__Vb[__Vi] = (__Vb[__Vi] ^ _VencAesRand_last128[__Vi])
	}

	_VencAesRand_iv128 = __Vb
} // _FencAesRand__gen_iv

func _FencAesRand_only(___Vkey *[]byte, ___VbyteIn *[]byte) ([]byte, error) {
    var __Vb    []byte
	_FpfN(" 192393 key (%d) %x , byte (%d) %x , %s ", len(*___Vkey), *___Vkey, len(*___VbyteIn), *___VbyteIn, string(*___VbyteIn))

	_FencAesRand__gen_iv()

    __Vlen := len(*___VbyteIn)

    __Vb = make( []byte, 2 )
    __Vb[1] = byte( __Vlen & 0xFF )
    __Vlen >>= 8
    __Vb[0] = byte( __Vlen & 0xFF )

    __Vb    = append( __Vb , (*___VbyteIn)              ...)
    __Vb    = append( __Vb , _FmakeByte16(md5.Sum(*___VbyteIn))  ...)
	//__Vb := _FmakeByte32(
	_FpfN(" 192394 byte (%d) %x , %s ", len(__Vb), __Vb, string(__Vb))

	__Vout, __Verr := _FencAesCbc__only(___Vkey, &_VencAesRand_iv128, &__Vb)

	if len(__Vout) > 16 {
		copy(_VencAesRand_last128, __Vout[:16])
	}

	_FpfN(" 192395 byte (%d) %x , %s ", len(__Vout), __Vout, string(__Vout))

	__Vout = _FappendRand2( &__Vout , 0 , 16 )

	_FpfN(" 192396 byte (%d) %x , %s ", len(__Vout), __Vout, string(__Vout))

	return __Vout, __Verr
} // _FencAesRand_only

func _FencAesRandExit(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	__Vb, __Verr := _FencAesRand_only(___Vkey, ___VbyteIn)
	_FerrExit(___VeMsg+" 192399 ", __Verr)
	return __Vb
} // _FencAesRandExit
