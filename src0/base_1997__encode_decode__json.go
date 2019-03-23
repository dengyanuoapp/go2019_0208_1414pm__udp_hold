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
func _FencJson___(___V interface{}) ([]byte, error) {
	_FinterfaceObjIsPointerOrExit("819181 01", ___V)
	__VbByte, __Verr := json.Marshal(___V)
	if __Verr != nil {
		_Perr(__Verr, "819181 05 : gob.NewDecoder failed:")
		return nil, __Verr
	}
	return __VbByte, __Verr
} // _FencJson___

func _FencJsonExit(___VeMsg string, ___V interface{}) []byte {
	//__Vbyte , __Verr := _FencJson___( ___V )
	//_FerrExit( " 1831918 : json error : " + ___VeMsg , __Verr )
	//_FerrExit( " 1831918 : json error : " , __Verr )

	//__Vbyte , _:= _FencJson___( ___V )
	__Vbyte, __Verr := _FencJson___(___V)
	_FerrExit(___VeMsg+" 819181 11 : json error : ", __Verr)

	return __Vbyte
} // _FencJsonExit

//    func Unmarshal(data []byte, v interface{}) error
func _FdecJson___(___VeMsg string, ___Vbyte *[]byte, ___VoutObj interface{}) error {
	_FinterfaceObjIsPointerOrExit(" 819181 20 ", ___VoutObj)
	//_FpfN("819181 21 : %v", ___VoutObj)
	//_FpfNhex(___Vbyte, 800, "819181 22 :")
	__Verr := json.Unmarshal(*___Vbyte, ___VoutObj)
	if __Verr != nil {
		_FpfN("819181 24 : %v , %T , %T", __Verr, ___VoutObj, &___VoutObj)

		if 2 == 2 {
			var __Vsi _TsrvInfo
			__Verr2 := json.Unmarshal(*___Vbyte, &__Vsi)
			_FpfN("819181 25 : %v , %v : %T , %T ", __Verr2, __Vsi, __Vsi, &__Vsi)
		}

		//__Vo4 := *___VoutObj
		//_FpfN("819181 26 : %T", __Vo4)

		return fmt.Errorf(___VeMsg+"819181 28 : gob.NewDecoder failed: %v", __Verr)
	}
	//_FpfN("819181 29 : %v", ___VoutObj)
	return nil
} // _FdecJson___

func _Fwrite_json_only_Exit(___VeMsg string, ___Vfname string, ___Vobj interface{}) []byte {
	var __Vb []byte
	__Vb = _FencJsonExit(___VeMsg+" 819181 31", ___Vobj)
	_FwriteFileExit(___VeMsg+" 819181 32 ", ___Vfname, &__Vb)
	return __Vb
} // _Fwrite_json_only_Exit

func _Fwrite_json_rand_only_Exit(___VeMsg string, ___Vkey *[]byte, ___Vfname string, ___Vobj interface{}) {
	var __VbufTmp1, __VbufTmp2 []byte
	//_FpfN("819184 01 : %T , %x , %#v", ___Vobj, ___Vkey, ___Vobj)
	__VbufTmp1 = _FencJsonExit(___VeMsg+" 819184 03", ___Vobj)
	__VbufTmp2 = _FencAesRandExit(___VeMsg+" 819184 05", ___Vkey, &__VbufTmp1)
	_FwriteFileExit(___VeMsg+" 819184 07", ___Vfname, &__VbufTmp2)
} // _Fwrite_json_rand_only_Exit

func _Fread_json_rand_only_Exit(___VeMsg string, ___Vkey *[]byte, ___Vfname string, ___Vobj interface{}) []byte {
	var __VbufTmp1, __VbufTmp2 []byte
	//_FpfN("819183 01 : %T , %x , %#v", ___Vobj, ___Vkey, ___Vobj)
	__VbufTmp1 = _FreadFileExit(___VeMsg+" 819183 03 ", ___Vfname)
	__VbufTmp2 = _FdecAesRandExit(___VeMsg+" 819183 05", ___Vkey, &__VbufTmp1)
	//_FpfNhex(&__VbufTmp2, 80, " 819183 07")
	if nil == ___Vobj {
		_FpfNex(" 819183 11 : why nil")
	} else {
		//_FpfNhex(&__VbufTmp2, 80, " 819183 13")
		_FdecJson___(___VeMsg+" 819183 15 ", &__VbufTmp2, ___Vobj)
	}
	//_FpfN(" 819183 17")
	return __VbufTmp2
} // _Fread_json_rand_only_Exit

func _FtestER__write_json_and_rand_Exit(___VeMsg string, ___Vkey *[]byte, ___Vfname string, ___Vobj interface{}) {
	__VbufText1 := _Fwrite_json_only_Exit(___VeMsg+" 381916 01 ", ___Vfname, ___Vobj)

	//_FpfN("381916 11 : %T , %x , %#v", ___Vobj, ___Vkey, ___Vobj)

	_Fwrite_json_rand_only_Exit(___VeMsg+" 381916 02 ", ___Vkey, ___Vfname+".rand", ___Vobj)

	__VbufText2 := _Fread_json_rand_only_Exit(___VeMsg+" 381916 03 ", ___Vkey, ___Vfname+".rand", ___Vobj)

	//_FpfN("381916 12 : %x , %#v", ___Vkey, ___Vobj)
	_FfalseExit(" 381916 04 ", bytes.Equal(__VbufText1, __VbufText2))

	if 3 == 3 {
		var __Vobj3 _TsrvInfo
		_Fread_json_rand_only_Exit(___VeMsg+" 381916 13 ", ___Vkey, ___Vfname+".rand", &__Vobj3)
		__VbufText3 := []byte(_Pspf("%#v", __Vobj3))
		_FwriteFileExit(___VeMsg+" 381916 15 ", ___Vfname+".rand.ex.printf.txt", &__VbufText3)
		//_FpfN("381916 14 : %x , %#v", ___Vkey, __Vobj3)
	}

	_FpfNdb(" 381916 06 : gob rand test ok. ")
	//_FpfNhex(&__VbufText1, 80, " 381916 07 ")
	//_FpfNhex(&__VbufText2, 80, " 381916 08 ")
	//_FpfN(" 381916 09 : <%s>", __VbufText2)
} // _FtestER__write_json_and_rand_Exit
