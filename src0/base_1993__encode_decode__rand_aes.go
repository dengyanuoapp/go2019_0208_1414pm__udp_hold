package main

import (
	"time"
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
	_FpfN(" 192393 key (%d) %x , byte (%d) %x , %s ", len(*___Vkey), *___Vkey, len(*___VbyteIn), *___VbyteIn, string(*___VbyteIn))

	_FencAesRand__gen_iv()

	__Vbyte, __Verr := _FencAesCbc__only(___Vkey, &_VencAesRand_iv128, ___VbyteIn)

	if len(__Vbyte) > 16 {
		copy(_VencAesRand_last128, __Vbyte[:16])
	}

	return __Vbyte, __Verr
} // _FencAesRand_only

func _FencAesRandExit(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	__Vb, __Verr := _FencAesRand_only(___Vkey, ___VbyteIn)
	_FerrExit(___VeMsg+" 192395 ", __Verr)
	return __Vb
} // _FencAesRandExit
