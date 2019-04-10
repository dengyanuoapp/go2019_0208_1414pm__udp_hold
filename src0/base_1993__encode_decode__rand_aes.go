package main

import (
	//"time"
	"bytes"
	"crypto/md5"
	"fmt"
)

//var (
//	//_VencAesRand_iv128__now  []byte = make([]byte, 16)
//	_VencAesRand_iv128__last []byte = make([]byte, 16)
//)
//
//// func _FencAesRand__gen_iv__by_timeMd5() {
//// 	_VencAesRand_iv128__now := _FgenMd5_nowB___(&_VencAesRand_iv128__last)
//// 	_VencAesRand_iv128__last = _VencAesRand_iv128__now
//// } // _FencAesRand__gen_iv__by_timeMd5

func _FencAesRand__only(___Vkey *[]byte, ___VbyteIn *[]byte) ([]byte, error) {
	var (
		__Vout, __Vtmp []byte
	)

	if 32 != len(*___Vkey) {
		_FpfNhex(___Vkey, 32, " 192390 key error len")
		return __Vout, nil
	}

	__Viv := _FgenRand_nByte__(16)
	__VlenIn := len(*___VbyteIn)
	//_FpfNhex(___VbyteIn, 40, " 192391 iv %x , byteIn ", __Viv)

	__Vtmp = make([]byte, 2+__VlenIn+16) // 2 byte len , data , 16byteMd5

	__Vtmp[1] = byte(__VlenIn & 0xFF)        // byte 1 : low byte of the uint16
	__Vtmp[0] = byte((__VlenIn >> 8) & 0xFF) // byte 0 : high byte of the uint16

	//__Vtmp = append(__Vtmp, (*___VbyteIn)...)
	//__Vtmp = append(__Vtmp, _FmakeByte16(md5.Sum(*___VbyteIn))...)
	copy(__Vtmp[2:], *___VbyteIn)
	//_FpfNhex(&__Vtmp, 48, " 192392 Vtmp : ")
	//copy(__Vtmp[2+__VlenIn:], _FmakeByte16(md5.Sum(*___VbyteIn)))
	copy(__Vtmp[2+__VlenIn:], _FmakeByte16(md5.Sum(__Vtmp[:2+__VlenIn])))
	//_FpfNhex(&__Vtmp, 48, " 192393 Vtmp : ")

	if 2 == 3 {
		_FpfNhex(&__Vtmp, 48, " 192394 pre  : ")
		_FpfNhex(&__Vtmp, 28, " 192394 pre  : ")
		_FpfNhex(&__Vtmp, 30, " 192394 pre  : ")
		_FpfNhex(&__Vtmp, 32, " 192394 pre  : ")
		_FpfNhex(&__Vtmp, 33, " 192394 pre  : ")
		_FpfNhex(&__Vtmp, 34, " 192394 pre  : ")
		_FpfhexlastN(&__Vtmp, 28, " 192394 post : ")
		_FpfhexlastN(&__Vtmp, 30, " 192394 post : ")
		_FpfhexlastN(&__Vtmp, 33, " 192394 post : ")
		_FpfhexlastN(&__Vtmp, 34, " 192394 post : ")
	}

	__Vtmp2, __Verr := _FencAesCbc__only___(___Vkey, &__Viv, &__Vtmp)
	_FerrExit(" 192395 ", __Verr)

	//_FpfhexlastN(&__Vtmp2, 40, " 192396 Vtmp : ")

	__Vout = _FappendRandPAT0_15(&__Vtmp2)

	//_FpfhexlastN(&__Vout, 40, " 192397 Vtmp : ")

	return __Vout, __Verr
} // _FencAesRand__only

func _FdecAesRand__only(___Vkey *[]byte, ___VbyteIn *[]byte) ([]byte, error) {
	var (
		__Vb0, __Vb1, __Vb2, __Vb3, __Vlen int
	)

	if nil == ___Vkey {
		return nil, fmt.Errorf(" 392391 01 : why key len error nil ? ")
	}
	if nil == ___VbyteIn {
		return nil, fmt.Errorf(" 392391 02 : why input len error nil ? ")
	}
	if 0 == len(*___Vkey) {
		return nil, fmt.Errorf(" 392391 03 : why key len zero ? ")
	}
	if 0 == len(*___VbyteIn) {
		return nil, fmt.Errorf(" 392391 04 : why input len zero ? ")
	}

	_CpfN("392391 05 : AESdec input key:%s , (%d){%x}", String5(___Vkey), len(*___VbyteIn), _FgenMd5__5(___VbyteIn))

	//_FpfNhex(___VbyteIn, 20, " 392391 06 : %x : ", ___Vkey)
	__VdeO, __Verr := _FdecAesCbc__only___(___Vkey, ___VbyteIn)
	_FerrExit(" 392391 07", __Verr)

	__Vlen = len(__VdeO)
	//_FpfNhex(&__VdeO, 50, " 392391 09 : %d : ", __Vlen)

	if __Vlen < 32 {
		return nil, fmt.Errorf(" 392391 10 : why len error ? (%d) ")
	}

	__Vb1 = int(__VdeO[1]) // byte 1 : low byte of the uint16
	__Vb0 = int(__VdeO[0]) // byte 0 : high byte of the uint16
	__Vb2 = (__Vb0 << 8) | __Vb1
	__Vb3 = __Vb2 + 2

	//_FpfN(" 392393 vb0 %d, vb1 %d , vb2 %d, vb3 %d, vlen %d ", __Vb0, __Vb1, __Vb2, __Vb3, __Vlen)
	//_FpfhexlastN(&__VdeO, 128, " 392393 Vtmp : ")
	//_FtrueExit(" 392393 ", __Vb3 > __Vlen)

	if __Vb3+16 > __Vlen {
		return nil, fmt.Errorf(" 392392 04: len error , this is NOT the data for me(using my key). Vb3:%d , Vlen:%d vb0,1:%x,%x",
			__Vb3, __Vlen, __VdeO[0], __VdeO[1])
	}

	__Vmd5InPack := __VdeO[__Vb3 : __Vb3+16]
	__Vmd5calc := _FmakeByte16(md5.Sum(__VdeO[:__Vb3]))
	if false == bytes.Equal(__Vmd5InPack, __Vmd5calc) {
		_FpfhexlastN(&__Vmd5InPack, 16, " 392392 05 ")
		_FpfhexlastN(&__Vmd5calc, 16, " 392392 06 ")
		_FpfhexlastN(&__VdeO, 160, " 392392 07 ")
		return nil, fmt.Errorf(" 392392 08: md5 error,it is a fake package(maybe random package) ")
	}

	__Vout2 := make([]byte, __Vb2)
	copy(__Vout2, __VdeO[2:__Vb3])

	__Vtmp3 := __VdeO[:__Vb3]
	_CpfN("392392 09 : AESdec out key:%s , in(%d){%x} outBufall(%d){%x} dec-without-2byteLen(%d){%x} dec-with-2byteLen(origin)(%d){%x} ", String5(___Vkey),
		len(*___VbyteIn), _FgenMd5__5(___VbyteIn),
		len(__VdeO), _FgenMd5__5(&__VdeO),
		len(__Vout2), _FgenMd5__5(&__Vout2),
		len(__Vtmp3), _FgenMd5__5(&__Vtmp3),
	)

	//_FpfhexlastN(&__Vout2, 40, " 392399 Vout : ")
	return __Vout2, __Verr
} // _FdecAesRand__only

func _FencAesRandExit(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) []byte {
	__Vb, __Verr := _FencAesRand__only(___Vkey, ___VbyteIn)
	_FerrExit(___VeMsg+" 292391 ", __Verr)
	return __Vb
} // _FencAesRandExit

func _FencAesRandExit2(___VeMsg string, ___Vkey *[]byte, ___VbyteIn *[]byte) ([]byte, []byte) {
	__Vb1, __Verr := _FencAesRand__only(___Vkey, ___VbyteIn)
	_FerrExit(___VeMsg+" 292392 01", __Verr)
	__Vb2, __Verr := _FencAesRand__only(___Vkey, ___VbyteIn)
	_FerrExit(___VeMsg+" 292392 02", __Verr)
	return __Vb1, __Vb2
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

func _FaesRand_test__en_de_Exit(___VstrIn string, ___VloopAmount int) {
	for ___VloopAmount > 0 {
		___VloopAmount--
		__Vkey := _FgenRand_nByte__(32)
		//_FpfNhex(&__Vkey, 40, " 738191 : using key ")

		__VbI := []byte(___VstrIn)
		__Vt1 := _FaesRand_test__encode(" 182812 ", &__Vkey, &__VbI)
		// len(2) + string(27) + md5(16) == 2+27+16 == 45 ///// origin_text(45)  --> encAES
		// // text(45) + padadd(3) == 48 == 3 * 16 as aes_data_load(48)
		// // iv(16) + aes_data_load(48) == 16 + 48 == 64(aesENtext)
		// // 64(aesENtext) + random(0-15 byte) == 64 + N byte == 64 + 8(for example) --> 72 byte( aes_ran_gen_Crypt_text(72) )

		//_FpfNhex(&__Vt1, 48, " 738194 : genRandSecText ")
		__VbyteO := _FaesRand_test__decode(" 182814 ", &__Vkey, &__Vt1)
		//_FpfNhex(&__VbyteO, 48, " 738195 : plain_text ")
		__VstrO := string(__VbyteO)

		_FnotEqExit(" 738198 : comp error.", ___VstrIn, __VstrO)
		_FpfN(" 738199 : comp_ok:%d : %s", ___VloopAmount, ___VstrIn)
	}
} // _FaesRand_test__en_de_Exit
