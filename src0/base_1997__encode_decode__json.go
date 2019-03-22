package main

import (
	"bytes"
	//"encoding/binary"
	//"encoding/gob"

	"encoding/json"
	//"io/ioutil"
	//"log"
	"fmt"
	//"math"
)

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
func _FencJson(___V interface{}) ([]byte, error) {
	__VbByte, __Verr := json.Marshal(___V)
	if __Verr != nil {
		_Perr(__Verr, "1831917 : gob.NewDecoder failed:")
		return nil, __Verr
	}
	return __VbByte, __Verr
} // _FencJson

func _FencJsonExit(___VeMsg string, ___V interface{}) []byte {
	//__Vbyte , __Verr := _FencJson( ___V )
	//_FerrExit( " 1831918 : json error : " + ___VeMsg , __Verr )
	//_FerrExit( " 1831918 : json error : " , __Verr )

	//__Vbyte , _:= _FencJson( ___V )
	__Vbyte, __Verr := _FencJson(___V)
	_FerrExit(___VeMsg+" 1831918 : json error : ", __Verr)

	return __Vbyte
} // _FencJsonExit

//    func Unmarshal(data []byte, v interface{}) error
func _FdecJson___(___VeMsg string, ___Vbyte *[]byte, ___Vout interface{}) error {
	__Verr := json.Unmarshal(*___Vbyte, ___Vout)
	if __Verr != nil {
		return fmt.Errorf(___VeMsg+"1831919 : gob.NewDecoder failed: %v", __Verr)
	}
	return nil
} // _FdecJson___

func _Fwrite_json_only_Exit(___VeMsg string, ___Vfname string, ___Vobj interface{}) []byte {
	var __Vb []byte
	__Vb = _FencJsonExit(___VeMsg+" 381911 ", ___Vobj)
	_FwriteFileExit(___VeMsg+" 381912 ", ___Vfname, &__Vb)
	return __Vb
} // _Fwrite_json_only_Exit

func _Fwrite_json_rand_only_Exit(___VeMsg string, ___Vkey *[]byte, ___Vfname string, ___Vobj interface{}) {
	var __VbufTmp1, __VbufTmp2 []byte
	__VbufTmp1 = _FencJsonExit(___VeMsg+" 381913 ", ___Vobj)
	__VbufTmp2 = _FencAesRandExit(___VeMsg+" 381914 ", ___Vkey, &__VbufTmp1)
	_FwriteFileExit(___VeMsg+" 381915 ", ___Vfname, &__VbufTmp2)
} // _Fwrite_json_rand_only_Exit

func _Fread_json_rand_only_Exit(___VeMsg string, ___Vkey *[]byte, ___Vfname string, ___Vobj interface{}) []byte {
	var __VbufTmp1, __VbufTmp2 []byte
	__VbufTmp1 = _FreadFileExit(___VeMsg+" 819183 ", ___Vfname)
	__VbufTmp2 = _FdecAesRandExit(___VeMsg+" 819184 ", ___Vkey, &__VbufTmp1)
	//_FpfNhex(&__VbufTmp2, 80, " 819186 ")
	if nil == ___Vobj {
		_FdecJson___(___VeMsg+" 819187 ", &__VbufTmp2, ___Vobj)
	}
	return __VbufTmp2
} // _Fread_json_rand_only_Exit

func _FtestER__write_json_and_rand_Exit(___VeMsg string, ___Vkey *[]byte, ___Vfname string, ___Vobj interface{}) {
	__VbufText1 := _Fwrite_json_only_Exit(___VeMsg+" 381916 01 ", ___Vfname, ___Vobj)

	_Fwrite_json_rand_only_Exit(___VeMsg+" 381916 02 ", ___Vkey, ___Vfname+".rand", ___Vobj)

	__VbufText2 := _Fread_json_rand_only_Exit(___VeMsg+" 381916 03 ", ___Vkey, ___Vfname+".rand", ___Vobj)

	_FfalseExit(" 381916 04 ", bytes.Equal(__VbufText1, __VbufText2))

	_FpfNdb(" 381916 05 : gob rand test ok. ")
	//_FpfNhex(&__VbufText1, 80, " 381916 06 ")
	//_FpfNhex(&__VbufText2, 80, " 381916 07 ")
	//_FpfN(" 381916 08 : <%s>", __VbufText2)
} // _FtestER__write_json_and_rand_Exit
