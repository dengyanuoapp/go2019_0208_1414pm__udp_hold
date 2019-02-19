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
	var __VbBuf bytes.Buffer        // Stand-in for a __VbBuf connection
	__Venc := gob.NewEncoder(&__VbBuf) // Will write to __VbBuf.
	__Verr := __Venc.Encode( ___V )
	if __Verr != nil {
        _Perr( __Verr , "1831915 : gob.NewEncoder failed:" )
        return nil
	}
	return __VbBuf.Bytes()
} // _FencGob
