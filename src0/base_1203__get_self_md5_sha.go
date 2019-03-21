package main

import (
	"bytes"
	"io/ioutil"
	"os"
	//"crypto/md5"
	//"crypto/sha256"
	"encoding/hex"
	"log"
	"time"
	//"strconv"
)

type _Tself struct {
	ProjName string
	//MyId256  _Tb256

	progPath string
	progMd5  _Tb128 // md5sum    : 16 byte : 727bf338cf523b90baccd24cca30b919
	progSha  _Tb256 // sha256sum : 32 byte : 2c6e3b458d5c482bc52a1d7d4f5a7d7766381c9f07d9b32ca605ae45b4e473f5

	startTime    time.Time
	startTimEsha _Tb256
	meSeq128     []byte

	debugEnabled bool
} //    _Tself

func _Fbase_1203__detect_debug_enabled() {
	// prog : 1
	// prog x1 x2 x3 : 4
	__VparaLen := len(os.Args)
	if 1 != __VparaLen {
		// func hex.DecodeString(s string) ([]byte, error)
		__Vhex, __Verr := hex.DecodeString(os.Args[__VparaLen-1])
		if nil == __Verr {
			//if fmt.Sprintf("%x" , __Vhex) == fmt.Sprintf("%x" , _VS.progSha ) {
			if bytes.Equal(__Vhex, _VS.progSha.b256) {
				_VS.debugEnabled = true
			}
		}
	}
	log.Printf(" _VS.debugEnabled : %t\n", _VS.debugEnabled)
} // _Fbase_1203__detect_debug_enabled

func _Fbase_1203__gen_self_md5_sha() {

	_VS.progPath = os.Args[0]
	__Vcontent, __Verr := ioutil.ReadFile(_VS.progPath)
	if __Verr != nil {
		log.Fatalf("err138191", __Verr)
	}

	_VS.progSha._Fbase_1101__gen_shaT(&__Vcontent)
	_VS.progMd5._Fbase_1101__gen_md5T(&__Vcontent)
	_FpfN(" 838191 _Fbase_1101b__gen_md5Only return : [%x] (%x %x)", _VS.progMd5.b128, _VS.progMd5.A1, _VS.progMd5.A2)

	_Ppn("from:", _VS.progPath)
	_FpfN(" 8381191 File md5: [ %x ]", _VS.progMd5.b128)
	_FpfN(" 8381192 File sha: [ %x ] %x %x %x %x ", _VS.progSha.b256, _VS.progSha.A1, _VS.progSha.A2, _VS.progSha.A3, _VS.progSha.A4)

} // _Fbase_1203__gen_self_md5_sha

func _Fbase_1203__gen_rand_seed() {
	_VS.startTime = time.Now()
	__Vb := []byte(_Pspf("%x", _VS.startTime))
	_VS.startTimEsha._Fbase_1101__gen_shaT(&__Vb)
	_VS.meSeq128 = _FgenRand_nByte__(16)
} // _Fbase_1203__gen_rand_seed()

func _Fbase_1203__init_self_All() {
	_Fbase_1203__detect_debug_enabled()
	_Fbase_1203__gen_self_md5_sha()
	_Fbase_1203__gen_rand_seed()
} //    _Fbase_1203__init_self_All()
