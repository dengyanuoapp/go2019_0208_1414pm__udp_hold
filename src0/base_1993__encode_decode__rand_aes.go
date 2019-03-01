package main

import (
	//"time"
	"crypto/md5"
)

var (
	//_VencAesRand_iv128__now  []byte = make([]byte, 16)
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
	__Vtmp[0] = byte((__VlenIn >> 8) & 0xFF)

	//__Vtmp = append(__Vtmp, (*___VbyteIn)...)
	//__Vtmp = append(__Vtmp, _FmakeByte16(md5.Sum(*___VbyteIn))...)
	copy(__Vtmp[2:], *___VbyteIn)
	_FpfhexN(&__Vtmp, 48, " 192392 Vtmp : ")
	//copy(__Vtmp[2+__VlenIn:], _FmakeByte16(md5.Sum(*___VbyteIn)))
	copy(__Vtmp[2+__VlenIn:], _FmakeByte16(md5.Sum(__Vtmp[:2+__VlenIn])))
	_FpfhexN(&__Vtmp, 48, " 192393 Vtmp : ")

	if 2 == 3 {
		_FpfhexN(&__Vtmp, 48, " 192394 pre  : ")
		_FpfhexN(&__Vtmp, 28, " 192394 pre  : ")
		_FpfhexN(&__Vtmp, 30, " 192394 pre  : ")
		_FpfhexN(&__Vtmp, 32, " 192394 pre  : ")
		_FpfhexN(&__Vtmp, 33, " 192394 pre  : ")
		_FpfhexN(&__Vtmp, 34, " 192394 pre  : ")
		_FpfhexlastN(&__Vtmp, 28, " 192394 post : ")
		_FpfhexlastN(&__Vtmp, 30, " 192394 post : ")
		_FpfhexlastN(&__Vtmp, 33, " 192394 post : ")
		_FpfhexlastN(&__Vtmp, 34, " 192394 post : ")
	}

	__Vtmp2, __Verr := _FencAesCbc__only___(___Vkey, &__Viv, &__Vtmp)
	_FerrExit(" 192395 ", __Verr)

	_FpfhexlastN(&__Vtmp2, 40, " 192396 Vtmp : ")

	__Vout2 := _FappendRandPAT0_15(&__Vtmp2)

	_FpfhexlastN(&__Vout2, 40, " 192397 Vtmp : ")

	return __Vout2, __Verr
} // _FencAesRand__only

func _FdecAesRand__only(___Vkey *[]byte, ___VbyteIn *[]byte) ([]byte, error) {

	__Vout2, __Verr := _FdecAesCbc__only___(___Vkey, ___VbyteIn)
	_FerrExit(" 392395 ", __Verr)

	_Fex1(" 83819101939 debug ")

	_FpfhexlastN(&__Vout2, 40, " 392397 Vout : ")
	return __Vout2, __Verr
} // _FdecAesRand__only

func _FencAesRandExit(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	__Vb, __Verr := _FencAesRand__only(___Vkey, ___VbyteIn)
	_FerrExit(___VeMsg+" 292391 ", __Verr)
	return __Vb
} // _FencAesRandExit

func _FdecAesRandExit(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	__Vb, __Verr := _FdecAesRand__only(___Vkey, ___VbyteIn)
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
	_FpfN(" 738191 : using key :%x", __Vkey)

	__VbI := []byte(___VstrIn)
	__Vt1 := _FaesRand_test__encode(" 182812 ", &__Vkey, &__VbI)
	// len(2) + string(27) + md5(16) == 2+27+16 == 45 ///// origin_text(45)  --> encAES
	// // text(45) + padadd(3) == 48 == 3 * 16 as aes_data_load(48)
	// // iv(16) + aes_data_load(48) == 16 + 48 == 64(aesENtext)
	// // 64(aesENtext) + random(0-15 byte) == 64 + N byte == 64 + 8(for example) --> 72 byte( aes_ran_gen_Crypt_text(72) )

	_FpfhexN(&__Vt1, 48, " 738194 : genRandSecText ")
	__VbyteO := _FaesRand_test__decode(" 182814 ", &__Vkey, &__Vt1)
	_FpfhexN(&__VbyteO, 48, " 738195 : plain_text ")
	__VstrO := string(__VbyteO)

	_FnotEqExit(" 738198 : comp error.", ___VstrIn, __VstrO)
	_FpfN(" 738199 : comp ok:%s", ___VstrIn)
} // _FaesRand_test__en_de_Exit
