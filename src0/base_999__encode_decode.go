package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"

    "encoding/json"

	//"log"
	//"fmt"
    //"math"
)

func _FencBin(___V interface{}) []byte {
    __VbBbuf := new(bytes.Buffer)
	__Verr := binary.Write(__VbBbuf, binary.LittleEndian, ___V)
	if __Verr != nil {
        _Perr( __Verr , "1831913 : binary.Write failed:" )
	}
	return __VbBbuf.Bytes()
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

//  func Value.Indirect(v Value) Value
func _FencRF(___V interface{}) []byte {

    /*
    var __Vbyte []byte
    __Vbyte = *((*[]byte)(&___V))

	size := dataSize(v)
    */

	//buf := make([]byte, size)
    return nil
} // _FencRF

//    func Marshal(v interface{}) ([]byte, error)
func _FencJson(___V interface{}) ([]byte , error) {
	__VbByte , __Verr := json.Marshal( ___V )
	if __Verr != nil {
        _Perr( __Verr , "1831918 : gob.NewDecoder failed:" )
        return nil , __Verr
	}
	return __VbByte , __Verr
} // _FencJson

//    func Unmarshal(data []byte, v interface{}) error
func _FdecJson(___Vbyte []byte, ___Vout interface{}) {
	__Verr := json.Unmarshal( ___Vbyte , ___Vout )
	if __Verr != nil {
        _Perr( __Verr , "1831919 : gob.NewDecoder failed:" )
	}
} // _FdecJson
