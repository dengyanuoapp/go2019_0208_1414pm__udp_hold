package main

import (
	//"bytes"
	//"os"
	//"io/ioutil"
	"crypto/md5"
	"crypto/sha256"
	//"log"
	//"encoding/hex"
	"time"
	//"strconv"
	"encoding/binary"
)

type _Tb256 struct {
	b256 []byte
	A1   uint64
	A2   uint64
	A3   uint64
	A4   uint64
}
type _Tb128 struct {
	b128 []byte
	A1   uint64
	A2   uint64
}

func _Fbase_1101a__gen_shaOnly(___VbyteIn *[]byte) []byte {
	__Vb := _FmakeByte32(sha256.Sum256(*___VbyteIn))
	//_FpfN( " 289892 _Fbase_1101a__gen_shaOnly %d [%x] (%x)" , len( ___VbyteIn ) , md5.Sum( ___VbyteIn ) , __Vb )
	return __Vb
} // _Fbase_1101a__gen_shaOnly

func _Fbase_1101b__gen_md5Only(___VbyteIn *[]byte) []byte {
	__Vb := _FmakeByte16(md5.Sum(*___VbyteIn))
	//_FpfN( " 289891 _Fbase_1101b__gen_md5Only %d [%x] (%x)" , len( ___VbyteIn ) , md5.Sum( ___VbyteIn ) , __Vb )
	return __Vb
} // _Fbase_1101b__gen_md5Only

func (___Vsha *_Tb256) _Fbase_1101__gen_shaT(___VshaInput *[]byte) {
	___Vsha.b256 = _Fbase_1101a__gen_shaOnly(___VshaInput)
	___Vsha.A1 = binary.BigEndian.Uint64(___Vsha.b256[0:8])
	___Vsha.A2 = binary.BigEndian.Uint64(___Vsha.b256[8:16])
	___Vsha.A3 = binary.BigEndian.Uint64(___Vsha.b256[16:24])
	___Vsha.A4 = binary.BigEndian.Uint64(___Vsha.b256[24:32])
} // _Fbase_1101__gen_shaT

func (___Vmd5 *_Tb128) _Fbase_1101__gen_md5T(___Vmd5Input *[]byte) {
	___Vmd5.b128 = _Fbase_1101b__gen_md5Only(___Vmd5Input)
	//___Vmd5 . A1 , _    =   strconv.ParseUint( string(___Vmd5 . b128[0:7])      , 16, 64 )
	//___Vmd5 . A2 , _    =   strconv.ParseUint( string(___Vmd5 . b128[8:15])     , 16, 64 )
	//encoding/binary
	//for __Vi:=0 ; __Vi < 8   ; __Vi ++ { ___Vmd5 . A1 +=   (uint8) (___Vmd5 . b128[__Vi]) }
	//for __Vi:=8 ; __Vi < 16  ; __Vi ++ { ___Vmd5 . A1 +=   (uint8) (___Vmd5 . b128[__Vi]) }
	___Vmd5.A1 = binary.BigEndian.Uint64(___Vmd5.b128[0:8])
	___Vmd5.A2 = binary.BigEndian.Uint64(___Vmd5.b128[8:16])

	_FpfN(" 289893 _Fbase_1101b__gen_md5Only [%x] (%x %x)", ___Vmd5.b128, ___Vmd5.A1, ___Vmd5.A2)
} // _Fbase_1101__gen_md5T

var _VgenMd5_last16 [16]byte

func _Fmd5__5s(___VbIn *[]byte) string {
	if nil == ___VbIn {
		return "12345"
	}
	__Vb := _Fmd5__5x(___VbIn)
	return _Spf("%x", __Vb)
}
func _Fmd5__5x(___VbIn *[]byte) []byte {
	if nil == ___VbIn {
		return []byte("12345")
	}
	__Vo := _FgenMd5___(___VbIn)
	return __Vo[:5]
}

func _FgenMd5___(___VbIn *[]byte) [16]byte {
	return md5.Sum(*___VbIn)
}
func _FgenMd5__(___VbIn *[]byte) []byte {
	__Vb16 := md5.Sum(*___VbIn)
	__Vo16 := make([]byte, 16)
	copy(__Vo16, __Vb16[:16])
	return __Vo16
}

func _FgenMd5_byString__(___Vstr string) []byte {
	_VgenMd5_last16 := md5.Sum([]byte(___Vstr))
	__VtimeMd5 := make([]byte, 16)
	copy(__VtimeMd5, _VgenMd5_last16[:16])
	return __VtimeMd5
} // _FgenMd5_byString__

func _FgenMd5_now16___() []byte {
	__VtimeMd5 := make([]byte, 16)
	__VtimeStr := []byte(time.Now().String())
	_VgenMd5_last16 := md5.Sum(__VtimeStr)
	copy(__VtimeMd5, _VgenMd5_last16[:16])
	return __VtimeMd5
} // _FgenMd5_now16___
func _FgenMd5_now32___() []byte {
	__VtimeMd5 := make([]byte, 32)
	copy(__VtimeMd5[16:], _VgenMd5_last16[:16])
	__VtimeStr := []byte(time.Now().String())
	_VgenMd5_last16 := md5.Sum(__VtimeStr)
	copy(__VtimeMd5, _VgenMd5_last16[:16])
	return __VtimeMd5
} // _FgenMd5_now32___

func _FgenMd5_nowB___(___Viv *[]byte) []byte {
	__VtimeStr := time.Now().String()
	_VgenMd5_last16 := md5.Sum([]byte(__VtimeStr + string(*___Viv)))
	__VtimeMd5 := make([]byte, 16)
	copy(__VtimeMd5, _VgenMd5_last16[:16])
	return __VtimeMd5
} // _FgenMd5_nowB___
