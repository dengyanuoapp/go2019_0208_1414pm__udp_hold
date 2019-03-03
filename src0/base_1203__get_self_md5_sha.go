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
	MyId256  _Tb256

	progPath string
	progMd5  _Tb128 // md5sum    : 16 byte : 727bf338cf523b90baccd24cca30b919
	progSha  _Tb256 // sha256sum : 32 byte : 2c6e3b458d5c482bc52a1d7d4f5a7d7766381c9f07d9b32ca605ae45b4e473f5

	startTime    time.Time
	startTimEsha _Tb256

	debugEnabled bool
} //    _Tself

func _Fbase_1203__gen_self_md5_sha() {

	_Vself.progPath = os.Args[0]
	__Vcontent, __Verr := ioutil.ReadFile(_Vself.progPath)
	if __Verr != nil {
		log.Fatalf("err138191", __Verr)
	}

	_Vself.progSha._Fbase_1101__gen_shaT(&__Vcontent)
	_Vself.progMd5._Fbase_1101__gen_md5T(&__Vcontent)
	_FpfN(" 838191 _Fbase_1101b__gen_md5Only return : [%x] (%x %x)", _Vself.progMd5.b128, _Vself.progMd5.A1, _Vself.progMd5.A2)

	// prog : 1
	// prog x1 x2 x3 : 4
	__VparaLen := len(os.Args)
	if 1 != __VparaLen {
		// func hex.DecodeString(s string) ([]byte, error)
		__Vhex, __Verr := hex.DecodeString(os.Args[__VparaLen-1])
		if nil == __Verr {
			//if fmt.Sprintf("%x" , __Vhex) == fmt.Sprintf("%x" , _Vself.progSha ) {
			if bytes.Equal(__Vhex, _Vself.progSha.b256) {
				_Vself.debugEnabled = true
			}
		}
	}
	log.Printf(" _Vself.debugEnabled : %t\n", _Vself.debugEnabled)

	_Ppn("from:", _Vself.progPath)
	_FpfN(" 8381191 File md5: [ %x ]", _Vself.progMd5.b128)
	_FpfN(" 8381192 File sha: [ %x ] %x %x %x %x ", _Vself.progSha.b256, _Vself.progSha.A1, _Vself.progSha.A2, _Vself.progSha.A3, _Vself.progSha.A4)

} // _Fbase_1203__gen_self_md5_sha

func _Fbase_103__gen_rand_seed() {
	_Vself.startTime = time.Now()
	__Vb := []byte(_Pspf("%x", _Vself.startTime))
	_Vself.startTimEsha._Fbase_1101__gen_shaT(&__Vb)
} // _Fbase_103__gen_rand_seed()
