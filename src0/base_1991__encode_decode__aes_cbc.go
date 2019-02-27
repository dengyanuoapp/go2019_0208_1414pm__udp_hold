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
	var __Vlen, __Vlen2, __Vi, __Volen int
	var __Vtmp, __Vout []byte

	_FpfN(" 132811 : len In (%d) , key %x , iv %x", len(*___VbyteIn), *___Vkey, *___Viv)

	__Vlen = len(*___VbyteIn)
	__Vi = __Vlen & 0xF

	if 0 != __Vi {
		__Vlen2 = __Vlen + 16 - __Vi
		__Vtmp = make([]byte, __Vlen2)
		copy(__Vtmp, (*___VbyteIn))
		__Volen = __Vlen2 + 16

		_FpfN(" 132812 : use tmp")
	} else {
		__Volen = __Vlen + 16
	}

	__Vout = make([]byte, __Volen)
	copy(__Vout, (*___Viv)[0:16])

	__Vblock, __Verr := aes.NewCipher(*___Vkey)
	_FerrExit(" 132813 ", __Verr)

	__Vmode := cipher.NewCBCEncrypter(__Vblock, *___Viv)
	if 0 == __Vi {
		__Vmode.CryptBlocks(__Vout[16:], *___VbyteIn)
	} else {
		__Vmode.CryptBlocks(__Vout[16:], __Vtmp)
	}

	_FpfN(" 132819 : len Out (%d) ", len(__Vout))
	return __Vout, nil
} // _FencAesCbc__only___

func _FencAesCbcExit(___Vkey *[]byte, ___Viv *[]byte, ___VbyteIn *[]byte) []byte {
	__Vbyte, __Verr := _FencAesCbc__only___(___Vkey, ___Viv, ___VbyteIn)
	_FerrExit(" 182811 ", __Verr)
	return __Vbyte
} // _FencAesCbcExit
