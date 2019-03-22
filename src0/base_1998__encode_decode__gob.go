package main

import (
	"bytes"
	//"encoding/binary"
	"encoding/gob"
	//"encoding/json"
	//"io/ioutil"
	//"log"
	"fmt"
	//"math"
)

func _FencGob__only(___V interface{}) ([]byte, error) {
	//__VbBuf := new(bytes.Buffer)
	var __VbBuf bytes.Buffer
	__Venc := gob.NewEncoder(&__VbBuf) // Will write to __VbBuf.
	__Verr := __Venc.Encode(___V)
	if __Verr != nil {
		_Perr(__Verr, "1831915 01: gob.NewEncoder failed:")
		return nil, __Verr
	}
	return __VbBuf.Bytes(), nil
} // _FencGob__only

func _FdecGob__only(___Vbyte *[]byte, ___Vout interface{}) {
	__Vdec := gob.NewDecoder(bytes.NewReader(*___Vbyte)) // Will write to __VbBuf.
	__Verr := __Vdec.Decode(___Vout)
	if __Verr != nil {
		_Perr(__Verr, "1831915 03: gob.NewDecoder failed:")
	}
} // _FdecGob__only

func _FdecGob___(___VeMsg string, ___Vbyte *[]byte, ___Vout interface{}) error {
	__Vdec := gob.NewDecoder(bytes.NewReader(*___Vbyte)) // Will write to __VbBuf.
	__Verr := __Vdec.Decode(___Vout)
	if __Verr != nil {
		return fmt.Errorf(___VeMsg+"1831915 05 : gob.NewDecoder failed: %v", __Verr)
	}
	return nil
} // _FdecGob___
