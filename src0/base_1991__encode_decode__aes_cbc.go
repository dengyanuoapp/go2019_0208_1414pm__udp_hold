package main

import (
	"crypto/aes"
	"crypto/cipher"
)

// block is 128 bit , 16 byte
// key   is 256 bit , 32 byte : fixed use 256 aes , not 128 , 196
type _Taes struct {
	name string
	key  []byte
} // _Taes

// https://golang.org/pkg/crypto/cipher/#example_NewCBCEncrypter
func _FencAesCbc__only___(___Vkey *[]byte, ___Viv *[]byte, ___VbyteIn *[]byte) ([]byte, error) {
	var (
		__VlenIn, __VoLen, __VtLen, __VlenAdd, __Vi int
		__Vtmp, __Vout                              []byte
		__VneedPat                                  bool
	)

	//_FpfN(" 132811 _FencAesCbc__only___ : len In (%d) , key %x , iv %x", len(*___VbyteIn), *___Vkey, *___Viv)

	__VlenIn = len(*___VbyteIn)
	_FnotEqExit(" 132811 ", 16, __VlenIn)
	_FnotEqExit(" 132813 ", 16, len(*___Viv))

	__Vi = __VlenIn & 0xF
	__VneedPat = (0 != __Vi)
	if __VneedPat {
		__VlenAdd = 16 - __Vi
		__VtLen = __VlenIn + __VlenAdd
		__Vtmp = make([]byte, __VtLen)
		__Vpat := _FgenRand_nByte__(uint16(__VlenAdd))
		copy(__Vtmp[__VlenIn:], __Vpat)
		copy(__Vtmp, *___VbyteIn)
		__VoLen = __VtLen + 16

		_FpfN(" 132815 : add pat %d", __VlenAdd)
	} else {
		__VoLen = __VlenIn + 16
	}

	__Vout = make([]byte, __VoLen)

	copy(__Vout[0:16], (*___Viv))

	__Vblock, __Verr := aes.NewCipher(*___Vkey)
	_FerrExit(" 132813 ", __Verr)

	__Vmode := cipher.NewCBCEncrypter(__Vblock, *___Viv)
	if __VneedPat {
		__Vmode.CryptBlocks(__Vout[16:], __Vtmp)
	} else {
		__Vmode.CryptBlocks(__Vout[16:], (*___VbyteIn))
	}

	if 3 == 3 {
		_FpfhexN(&__Vout, 32, " 132819 _FencAesCbc__only___ : lenIn %d , dataOut: ", __VlenIn)
	}

	return __Vout, nil
} // _FencAesCbc__only___

func _FencAesCbcExit(___Vkey *[]byte, ___Viv *[]byte, ___VbyteIn *[]byte) []byte {
	__Vbyte, __Verr := _FencAesCbc__only___(___Vkey, ___Viv, ___VbyteIn)
	_FerrExit(" 182811 ", __Verr)
	return __Vbyte
} // _FencAesCbcExit
