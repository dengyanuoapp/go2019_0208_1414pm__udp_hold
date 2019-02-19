package main

import (
	"bytes"
	"encoding/binary"
	//"fmt"
	"encoding/gob"
	//"log"
)

func _FencBin(___V interface{}) []byte {
	__Vbuf := new(bytes.Buffer)
	__Verr := binary.Write(__Vbuf, binary.LittleEndian, ___V)
	if __Verr != nil {
        _Perr( __Verr , "1831911 : binary.Write failed:" )
        return nil
	}
	return __Vbuf.Bytes()
} // _FencBin

func _FencGob(___V interface{}) []byte {
	//__VbBuf := new(bytes.Buffer)
	var __VbBuf bytes.Buffer
	__Venc := gob.NewEncoder(&__VbBuf) // Will write to __VbBuf.
	__Verr := __Venc.Encode( ___V )
	if __Verr != nil {
        _Perr( __Verr , "1831915 : gob.NewEncoder failed:" )
        return nil
	}
	return __VbBuf.Bytes()
} // _FencGob

func _FdecGob(___Vbyte []byte, ___Vout interface{}) {
	__Vdec := gob.NewDecoder(bytes.NewReader(___Vbyte)) // Will write to __VbBuf.
	__Verr := __Vdec.Decode( ___Vout )
	if __Verr != nil {
        _Perr( __Verr , "1831917 : gob.NewDecoder failed:" )
	}
} // _FdecGob
