package main

import (
	//"time"
	"crypto/md5"
)

var (
	_VencAesRand_iv128__now  []byte = make([]byte, 16)
	_VencAesRand_iv128__last []byte = make([]byte, 16)
)

func _FencAesRand__gen_iv__by_timeMd5() {
	_VencAesRand_iv128__now := _FgenMd5_now1__(&_VencAesRand_iv128__last)
	_VencAesRand_iv128__last = _VencAesRand_iv128__now
} // _FencAesRand__gen_iv__by_timeMd5

func _FencAesRand_only(___Vkey *[]byte, ___VbyteIn *[]byte) ([]byte, error) {
	var __Vb []byte
	_FpfN(" 192393 key (%d) %x , byteIn (%d) %x , %s ", len(*___Vkey), *___Vkey, len(*___VbyteIn),
		*___VbyteIn, string(*___VbyteIn))

	_FencAesRand__gen_iv__by_timeMd5()
	_FpfN(" 192394 gen new iv (%d) %x , old iv (%d) %x", len(_VencAesRand_iv128__now), _VencAesRand_iv128__now,
		len(_VencAesRand_iv128__last), _VencAesRand_iv128__last)

	__Vlen := len(*___VbyteIn)

	__Vb = make([]byte, 2)
	__Vb[1] = byte(__Vlen & 0xFF)
	__Vlen >>= 8
	__Vb[0] = byte(__Vlen & 0xFF)

	__Vb = append(__Vb, (*___VbyteIn)...)
	__Vb = append(__Vb, _FmakeByte16(md5.Sum(*___VbyteIn))...)
	//__Vb := _FmakeByte32(
	_FpfN(" 192395 byte (%d) %x , %s ", len(__Vb), __Vb, string(__Vb))

	__Vout, __Verr := _FencAesCbc__only(___Vkey, &_VencAesRand_iv128__now, &__Vb)
	_FerrExit(" 192396 ", __Verr)

	if len(__Vout) > 16 {
		copy(_VencAesRand_iv128__last, __Vout[:16])
	}

	_FpfN(" 192397 byte (%d) %x , %s ", len(__Vout), __Vout, string(__Vout))

	__Vout = _FappendRandLen2byteArr(&__Vout, 0, 16)

	_FpfN(" 192398 byte (%d) %x , %s ", len(__Vout), __Vout, string(__Vout))

	return __Vout, __Verr
} // _FencAesRand_only

func _FencAesRandExit(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	__Vb, __Verr := _FencAesRand_only(___Vkey, ___VbyteIn)
	_FerrExit(___VeMsg+" 292391 ", __Verr)
	return __Vb
} // _FencAesRandExit

func _FdecAesRandExit(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	__Vb, __Verr := _FencAesRand_only(___Vkey, ___VbyteIn)
	_FerrExit(___VeMsg+" 292395 ", __Verr)
	return __Vb
} // _FdecAesRandExit

func _FaesRand_test__encode(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	return _FencAesRandExit(___VeMsg+" 481911 ", ___Vkey, ___VbyteIn)
} // _FaesRand_test__encode

func _FaesRand_test__decode(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	return _FdecAesRandExit(___VeMsg+" 481913 ", ___Vkey, ___VbyteIn)
} // _FaesRand_test__decode

func _FaesRand_test__en_de_Exit(___VstrIn string) {
	__Vkey := _FgenRand_nByte__(16)

	__VbI := []byte(___VstrIn)
	__Vt1 := _FaesRand_test__encode(" 182812 ", &__Vkey, &__VbI)
	__VstrO := string(_FaesRand_test__decode(" 182814 ", &__Vkey, &__Vt1))

	_FnotEqExit(" 182818 : comp error.", ___VstrIn, __VstrO)
	_FpfN(" 182819 : comp ok.")
} // _FaesRand_test__en_de_Exit
