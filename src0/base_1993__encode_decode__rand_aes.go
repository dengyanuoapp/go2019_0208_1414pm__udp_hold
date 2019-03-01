package main

import (
//"time"
//"crypto/md5"
)

var (
	_VencAesRand_iv128__now  []byte = make([]byte, 16)
	_VencAesRand_iv128__last []byte = make([]byte, 16)
)

// func _FencAesRand__gen_iv__by_timeMd5() {
// 	_VencAesRand_iv128__now := _FgenMd5_now1___(&_VencAesRand_iv128__last)
// 	_VencAesRand_iv128__last = _VencAesRand_iv128__now
// } // _FencAesRand__gen_iv__by_timeMd5

func _FencAesRand__only(___Vkey *[]byte, ___VbyteIn *[]byte) ([]byte, error) {
	var __Vtmp []byte

	__Viv := _FgenRand_nByte__(16)
	__VlenIn := len(*___VbyteIn)
	_FpfhexN(___VbyteIn, 20, " 192391 key %x , iv %x , byteIn ", *___Vkey, __Viv)

	__Vtmp = make([]byte, 2+__VlenIn+16) // 2 byte len , data , 16byteMd5

	__Vtmp[1] = byte(__VlenIn & 0xFF)
	__VlenIn >>= 8
	__Vtmp[0] = byte(__VlenIn & 0xFF)

	//__Vtmp = append(__Vtmp, (*___VbyteIn)...)
	//__Vtmp = append(__Vtmp, _FmakeByte16(md5.Sum(*___VbyteIn))...)
	copy(__Vtmp[2:], *___VbyteIn)
	//copy(__Vtmp[2+__VlenIn:], _FmakeByte16(md5.Sum(*___VbyteIn)))
	if 2 == 3 {
		_FpfhexN(&__Vtmp, 48, " 192392 pre  : ")
		_FpfhexN(&__Vtmp, 28, " 192392 pre  : ")
		_FpfhexN(&__Vtmp, 30, " 192392 pre  : ")
		_FpfhexN(&__Vtmp, 32, " 192392 pre  : ")
		_FpfhexN(&__Vtmp, 33, " 192392 pre  : ")
		_FpfhexN(&__Vtmp, 34, " 192392 pre  : ")
		_FpfhexlastN(&__Vtmp, 28, " 192393 post : ")
		_FpfhexlastN(&__Vtmp, 30, " 192393 post : ")
		_FpfhexlastN(&__Vtmp, 33, " 192393 post : ")
		_FpfhexlastN(&__Vtmp, 34, " 192393 post : ")
	}

	_FpfhexN(&__Vtmp, 16, " 192395 Vtmp ")

	__Vout, __Verr := _FencAesCbc__only___(___Vkey, &_VencAesRand_iv128__now, &__Vtmp)
	_FerrExit(" 192396 ", __Verr)

	if len(__Vout) > 16 {
		copy(_VencAesRand_iv128__last, __Vout[:16])
	}

	_FpfN(" 192397 byte (%d) %x , %s ", len(__Vout), __Vout, string(__Vout))

	__Vout = _FappendRandLen2byteArr(&__Vout, 0, 16)

	_FpfN(" 192398 byte (%d) %x , %s ", len(__Vout), __Vout, string(__Vout))

	return __Vout, __Verr
} // _FencAesRand__only

func _FencAesRandExit(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	__Vb, __Verr := _FencAesRand__only(___Vkey, ___VbyteIn)
	_FerrExit(___VeMsg+" 292391 ", __Verr)
	return __Vb
} // _FencAesRandExit

func _FdecAesRandExit(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	__Vb, __Verr := _FencAesRand__only(___Vkey, ___VbyteIn)
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
	_FpfN(" 182819 : comp ok:%s", ___VstrIn)
} // _FaesRand_test__en_de_Exit
