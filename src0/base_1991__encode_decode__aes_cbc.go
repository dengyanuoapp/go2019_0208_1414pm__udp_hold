package main

import (
	"crypto/aes"
	"crypto/cipher"
	"sync"
)

// block is 128 bit , 16 byte
// key   is 256 bit , 32 byte : fixed use 256 aes , not 128 , 196
//type _Taes struct {
//	name string
//	key  []byte
//} // _Taes

// https://golang.org/pkg/crypto/cipher/#example_NewCBCEncrypter
var ___VencAesCbc__only__mux sync.Mutex

func _FencAesCbc__only___(___Vkey []byte, ___Viv []byte, ___VbyteIn []byte, ___VtraceInt int) ([]byte, error) {

	defer ___VencAesCbc__only__mux.Unlock()
	___VencAesCbc__only__mux.Lock()

	var __VoLen int
	__VtLen := 0
	//_FpfN(" 132811 _FencAesCbc__only___ : len In (%d) , key %x , iv %x", len(*___VbyteIn), *___Vkey, *___Viv)

	__VlenInTmp2 := len(___VbyteIn)
	_FnotEqExit(" 132819 01 ", 32, len(___Vkey))
	_FnotEqExit(" 132819 02 ", 16, len(___Viv))

	__Vkey2 := ___Vkey
	__Viv2 := ___Viv

	__Vremain := __VlenInTmp2 & 0xF
	__VneedPat := (0 != __Vremain)
	__VtBufEN := []byte{}
	if __VneedPat {
		__VlenAdd := 16 - __Vremain
		__VtLen = __VlenInTmp2 + __VlenAdd
		if 0 != (__VtLen & 0xF) {
			_Fex(_Spf(" 132819 04 : why not ZERO : __VlenInTmp2 %d , __Vremain %d, __VneedPat %T, __VlenAdd %d , __VtLen %d?",
				__VlenInTmp2, __Vremain, __VneedPat, __VlenAdd, __VtLen))
		}

		__VtBufEN = make([]byte, __VtLen)

		__Vpat := _FgenRand_nByte__(uint16(__VlenAdd))
		copy(__VtBufEN[__VlenInTmp2:], __Vpat)
		copy(__VtBufEN, ___VbyteIn)
		__VoLen = __VtLen + 16

		//_FpfN(" 132815 : add pat %d", __VlenAdd)
	} else {
		__VtBufEN = ___VbyteIn
		__VoLen = __VlenInTmp2 + 16 // origin Len + iv(16)
	}

	__Vout4 := make([]byte, __VoLen-16)
	__Vout5 := make([]byte, __VoLen)

	__VblockEN, __Verr := aes.NewCipher(__Vkey2)
	_FerrExit(" 132819 03 ", __Verr)

	__VmodeEN := cipher.NewCBCEncrypter(__VblockEN, __Viv2)
	__VmodeEN.CryptBlocks(__Vout4, __VtBufEN)

	copy(__Vout5[0:16], __Viv2)
	copy(__Vout5[16:], __Vout4)

	__VblockEN = nil
	__VmodeEN = nil

	//_FpfNhex(&__Vout5, 32, " 132819 08 _FencAesCbc__only___ : lenIn %d , dataOut: ", __VlenInTmp2)
	__VinSideKey := __VtBufEN[7:39]
	___CpfN(" 132819 09 Ti:%d , aesENC(noRandPat) inM5{%x} outM5{%x} INfirst20<%x> in<%x> out:<%x> insideKey is <%x> ",
		___VtraceInt,
		_Fmd5__5x(&___VbyteIn),
		_Fmd5__5x(&__Vout5),
		__VtBufEN[:20],
		__VtBufEN,
		__Vout5,
		__VinSideKey,
	)

	return __Vout5, nil
} // _FencAesCbc__only___

func _FencAesCbcExit(___Vkey []byte, ___Viv []byte, ___VbyteIn []byte) []byte {
	__Vbyte, __Verr := _FencAesCbc__only___(___Vkey, ___Viv, ___VbyteIn, 0)
	_FerrExit(" 182811 ", __Verr)
	return __Vbyte
} // _FencAesCbcExit

var ___V_FdecAesCbc__only___mux sync.Mutex

func _FdecAesCbc__only___(___Vkey []byte, ___VbyteIn []byte, ___VtraceIntDE int) ([]byte, error) {

	defer ___V_FdecAesCbc__only___mux.Unlock()
	___V_FdecAesCbc__only___mux.Lock()

	__VoutNull := []byte{}
	if 32 != len(___Vkey) {
		_FpfNonce(" 838181 key len error (not equals to 32): %d:%s", len(___Vkey), String5s(&___Vkey))
		return __VoutNull, nil
	}

	__VlenIn := len(___VbyteIn)
	if __VlenIn < 32 {
		return __VoutNull, nil
	}
	__VdataEnd := (__VlenIn & 0xFFFFFFF0)
	//_FpfN(" 838180 __VlenIn %d , __VdataEnd %d ", __VlenIn, __VdataEnd)
	__Vout2 := make([]byte, __VdataEnd-16)
	//_FpfNhex(___VbyteIn, 82, " 838181 dataIn ")

	__Viv := (___VbyteIn)[:16]
	__VcipherText := (___VbyteIn)[16:__VdataEnd]

	__VblockDE, __Verr := aes.NewCipher(___Vkey) // func NewCipher(key []byte) (cipher.Block, error) // import "crypto/aes"
	_FerrExit(" 838182 ", __Verr)

	if 2 == 3 {
		_FpfNhex(&__Viv, 82, " 838183 iv ")
		_FpfNhex(&__VcipherText, 82, " 838184 cipherText ")
	}

	__VmodeDE := cipher.NewCBCDecrypter(__VblockDE, __Viv)
	_FnullExit(" 838186 ", __VmodeDE)

	// CryptBlocks(dst, src []byte)
	__VmodeDE.CryptBlocks(__Vout2, __VcipherText)
	//_FpfNhex(&__Vout2, 82, " 838189 out ")

	__VblockDE = nil
	__VmodeDE = nil
	return __Vout2, nil
} // _FdecAesCbc__only___
