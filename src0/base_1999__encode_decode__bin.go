package main

import (
	"bytes"
	"encoding/binary"
	//"encoding/gob"
	//"encoding/json"
	//"io/ioutil"
	//"log"
	//"fmt"
	//"math"
)

func _FencBin(___V interface{}) []byte {
	__VbBbuf := new(bytes.Buffer)
	__Verr := binary.Write(__VbBbuf, binary.LittleEndian, ___V)
	if __Verr != nil {
		_FpfN("1831913 01 : binary.Write failed:%v", __Verr)
	}
	return __VbBbuf.Bytes()
} // _FencBin

func _FencBin1(___V _TtestLeng01) []byte {
	__VbBbuf := new(bytes.Buffer)
	__Verr := binary.Write(__VbBbuf, binary.LittleEndian, ___V)
	if __Verr != nil {
		_FpfN("1831913 02: binary.Write failed:%v", __Verr)
	}
	return __VbBbuf.Bytes()
} // _FencBin1

func _FencBin2(___V _TtestLeng02) []byte {
	__VbBbuf := new(bytes.Buffer)
	__Verr := binary.Write(__VbBbuf, binary.LittleEndian, ___V)
	if __Verr != nil {
		_FpfN("1831913 03 : binary.Write failed:%v", __Verr)
	}
	return __VbBbuf.Bytes()
} // _FencBin2
