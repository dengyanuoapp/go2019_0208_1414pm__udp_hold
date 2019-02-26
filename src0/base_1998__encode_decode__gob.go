package main

import (
	"bytes"
	//"encoding/binary"
	"encoding/gob"
	//"encoding/json"
	//"io/ioutil"
	//"log"
	//"fmt"
	//"math"
)

func _FencGob(___V interface{}) []byte {
	//__VbBuf := new(bytes.Buffer)
	var __VbBuf bytes.Buffer
	__Venc := gob.NewEncoder(&__VbBuf) // Will write to __VbBuf.
	__Verr := __Venc.Encode(___V)
	if __Verr != nil {
		_Perr(__Verr, "1831915 : gob.NewEncoder failed:")
		return nil
	}
	return __VbBuf.Bytes()
} // _FencGob

func _FdecGob(___Vbyte []byte, ___Vout interface{}) {
	__Vdec := gob.NewDecoder(bytes.NewReader(___Vbyte)) // Will write to __VbBuf.
	__Verr := __Vdec.Decode(___Vout)
	if __Verr != nil {
		_Perr(__Verr, "1831917 : gob.NewDecoder failed:")
	}
} // _FdecGob
