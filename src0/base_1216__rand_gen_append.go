package main

import (
	//"encoding/binary"
	//"fmt"
	//"math/rand"
	"sync"
	"time"
)

const (
	_VsizeOfRandBuf_block = 1024 * 8 // * 16
	_VsizeOfRandBuf_byte  = _VsizeOfRandBuf_block * 16
)

type _TgenRand struct {
	buf    []byte
	remain uint32
	cnt64  uint64
	lock   sync.Mutex
}

var (
	_VgenRand _TgenRand
)

func _FreGenRandBuf___() {
	if 0 == len(_VgenRand.buf) {
		_FtrueExit(" 371911 buf size must large than 16bit ", _VsizeOfRandBuf_byte <= 0x10000)
		_VgenRand.buf = make([]byte, _VsizeOfRandBuf_byte)
	}
	__Vk := _FgenMd5_now32___()
	_Fsleep_1ms()
	__Viv := _FgenMd5_nowB___(&__Vk)

	__Vtmp, __Verr := _FencAesCbc__only___(&__Vk, &__Viv, &(_VgenRand.buf))
	_FerrExit(" 371913 ", __Verr)
	if 2 == 3 {
		_FpfN(" 371915 _FreGenRandBuf___ : len ( %d ) : %x %x %x ", len(__Vtmp), __Vtmp[:16], __Vtmp[16:32], __Vtmp[32:48])
		_FpfN(" 371916 _FreGenRandBuf___ : len ( %d ) : %x %x %x ", len(_VgenRand.buf), _VgenRand.buf[:16], _VgenRand.buf[16:32], _VgenRand.buf[32:48])
	}

	copy(_VgenRand.buf, __Vtmp[16:])
	_FnotEqExit(" 371917 ", _VsizeOfRandBuf_byte, len(_VgenRand.buf))
	//_FpfN(" 371915: len ( %d ) : %x %x %x ", len(_VgenRand.buf), _VgenRand.buf[:16], _VgenRand.buf[16:32], _VgenRand.buf[32:48])

	_VgenRand.remain = uint32(_VsizeOfRandBuf_byte)
	//_FpfhexN(&_VgenRand.buf, 40, " 371919: _FreGenRandBuf___ : bufSize[%d] ", _VsizeOfRandBuf_byte)
} // _FreGenRandBuf___

// gen rand byte slice , size is N
func _FgenRand_nByte__(___Vlen uint16) []byte {
	var __Vout []byte

	//_FpfN(" 938191 _FgenRand_nByte__ : _VgenRand.remain %d , need : %d ", _VgenRand.remain, ___Vlen)

	if ___Vlen == 0 {
		return __Vout
	}
	__VlenReq := uint32(___Vlen)

	_VgenRand.lock.Lock()
	for {
		if _VgenRand.remain == 0 {
			_FreGenRandBuf___()
		}
		//_FpfN(" 938192 _FgenRand_nByte__ : _VgenRand.remain %d , need : %d ", _VgenRand.remain, __VlenReq)

		if _VgenRand.remain >= __VlenReq {
			__Vnew := _VgenRand.remain - __VlenReq
			__Vout = make([]byte, __VlenReq)
			copy(__Vout, _VgenRand.buf[__Vnew:_VgenRand.remain])
			_VgenRand.remain = __Vnew
			//_FpfN(" 938194 _FgenRand_nByte__ : succeed gen : %d", __VlenReq)
			break
		} else {
			//_FpfN(" 938196 _FgenRand_nByte__ : skip the remain , regen ")
			_VgenRand.remain = 0
		}

		_Fsleep_1ms()
	}
	_VgenRand.lock.Unlock()

	_VgenRand.cnt64 += uint64(__VlenReq)

	//_FpfhexN(&__Vout, 24, " 938199 _FgenRand_nByte__ : (%d:%d) remain %d , Vout:0x%x ", _VgenRand.cnt64, __VlenReq, _VgenRand.remain, len(__Vout))

	return __Vout
} // _FgenRand_nByte__

func _FgenRand_nByte__testExit(___VloopAmount uint32) {
	var __Vu1, __Vu2, __Vu3 uint16
	var __Vcnt uint32

	__Vwant64 := uint64(___VloopAmount) * uint64(_VsizeOfRandBuf_byte)
	if 2 == 3 {
		_FpfN(" ___VloopAmount %d , %x ", ___VloopAmount, ___VloopAmount)
		_FpfN(" _VsizeOfRandBuf_byte %d , %x ", _VsizeOfRandBuf_byte, _VsizeOfRandBuf_byte)
		_FpfN(" __Vwant64 %d , %x ", __Vwant64, __Vwant64)
		_FpfN(" ___VloopAmount %d , %x ", ___VloopAmount, ___VloopAmount)
		_Fex1(" 389182 debug stop ")
	}

	for _VgenRand.cnt64 <= __Vwant64 {

		__Vb2 := _FgenRand_nByte__(2)
		__Vu1 = uint16(__Vb2[0])
		__Vu2 = uint16(__Vb2[1])
		__Vu3 = (__Vu1 << 8) + __Vu2

		//_FpfN(" 813911 _FgenRand_nByte__testExit : %x : %x %x %x , %d %d %d", __Vb2, __Vu1, __Vu2, __Vu3, __Vu1, __Vu2, __Vu3)
		__Vb3 := _FgenRand_nByte__(__Vu3)
		if 2 == 3 {
			_Fpf(" 813914 _FgenRand_nByte__testExit : %x , %d , %d : ", __Vb2, __Vu3, len(__Vb3))
			if len(__Vb3) <= 16 {
				_FpfN("%x", __Vb3)
			} else {
				_FpfN("%x (16 first ...)", __Vb3[:16])
			}
		}
		//_FpfN(" 813916 : %d : %d , %d ", __Vcnt, _VgenRand.cnt64, __Vwant64)
		__Vcnt++
	}
} // _FgenRand_nByte__testExit

func _FappendRandPAT0_15(___Vbyte *[]byte) []byte {

	if 0 == len(*___Vbyte) {
		return *___Vbyte
	}

	__Vtmp1 := _FgenRand_nByte__(1)

	__Vtmp2 := uint16(__Vtmp1[0])
	__Vtmp2 &= 0xF
	//_FpfN(" 381921 : add pat tail rand pat : %d ", __Vtmp2)

	__Vtmp3 := _FgenRand_nByte__(__Vtmp2)

	__Vout := append(*___Vbyte, __Vtmp3...)
	return __Vout
} // _FappendRandPAT0_15

func _FgenTimeRand16byte() []byte {
	var __Vb []byte
	__Vb = []byte(time.Now().String())
	return _Fbase_1101b__gen_md5Only(&__Vb)
} // _FgenTimeRand16byte
