package main

import (
	"bytes"
	"encoding/binary"
	//"fmt"
)

func _FencBin(___V interface{}) []byte {
	__Vbuf := new(bytes.Buffer)
	err := binary.Write(__Vbuf, binary.LittleEndian, ___V)
	if err != nil {
        _Perr( err , "1831911 : binary.Write failed:" )
        return nil
	}
	return __Vbuf.Bytes()
} // _FencBin

