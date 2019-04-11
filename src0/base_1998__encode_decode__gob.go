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

func _FencGob__only(___VobjIn interface{}) ([]byte, error) {
	_FinterfaceObjIsPointerOrExit("1831915 01", ___VobjIn)
	//__VbBuf := new(bytes.Buffer)
	var __VbBuf bytes.Buffer
	__Venc := gob.NewEncoder(&__VbBuf) // Will write to __VbBuf.
	__Verr := __Venc.Encode(___VobjIn)
	if __Verr != nil {
		return nil,
			fmt.Errorf("1831915 02: gob.NewEncoder failed:%v", __Verr)
	}
	return __VbBuf.Bytes(), nil
} // _FencGob__only

func _FdecGob__only(___VbyteIn []byte, ___VoutObjLp interface{}) {
	_FinterfaceObjIsPointerOrExit("1831915 03", ___VoutObjLp)
	__Vdec := gob.NewDecoder(bytes.NewReader(___VbyteIn)) // Will write to __VbBuf.
	__Verr := __Vdec.Decode(___VoutObjLp)
	if __Verr != nil {
		_FpfN("1831915 04: gob.NewDecoder failed:%v", __Verr)
	}
} // _FdecGob__only

func _FdecGob___(___VeMsg string, ___VbyteIn []byte, ___VoutObjLp interface{}) error {
	_FinterfaceObjIsPointerOrExit("1831915 05", ___VoutObjLp)
	__Vdec := gob.NewDecoder(bytes.NewReader(___VbyteIn)) // Will write to __VbBuf.
	__Verr := __Vdec.Decode(___VoutObjLp)
	if __Verr != nil {
		return fmt.Errorf(___VeMsg+"1831915 06 : gob.NewDecoder failed: %v", __Verr)
	}
	return nil
} // _FdecGob___

func _Fwrite_gob_only_Exit(___VeMsg string, ___Vfname string, ___Vobj interface{}) []byte {
	var __Vb []byte
	__Vb = _FencGobExit(___VeMsg+" 1831915 06 ", ___Vobj)
	_FwriteFileExit(___VeMsg+" 1831915 07 ", ___Vfname, &__Vb)
	return __Vb
} // _Fwrite_gob_only_Exit

func _FencGobExit(___VeMsg string, ___V interface{}) []byte {
	__Vbyte, __Verr := _FencGob__only(___V)
	_FerrExit(___VeMsg+" 1831915 08 : gob error : ", __Verr)

	return __Vbyte
} // _FencGobExit

func _Fwrite_gob_rand_only_Exit(___VeMsg string, ___Vkey []byte, ___Vfname string, ___Vobj interface{}) {
	var __VbufTmp1, __VbufTmp2 []byte
	__VbufTmp1 = _FencGobExit(___VeMsg+" 1831916 01 ", ___Vobj)
	__VbufTmp2 = _FencAesRandExit(___VeMsg+" 1831916 02 ", ___Vkey, __VbufTmp1)
	_FwriteFileExit(___VeMsg+" 1831916 03 ", ___Vfname, &__VbufTmp2)
} // _Fwrite_gob_rand_only_Exit

func _Fread_gob_rand_only_Exit(___VeMsg string, ___Vkey []byte, ___Vfname string, ___Vobj interface{}) []byte {
	var __VbufTmp1, __VbufTmp2 []byte
	__VbufTmp1 = _FreadFileExit(___VeMsg+" 1831916 05 ", ___Vfname)
	__VbufTmp2 = _FdecAesRandExit(___VeMsg+" 1831916 06 ", ___Vkey, __VbufTmp1)
	//_FpfNhex(&__VbufTmp2, 80, " 1831916 07 ")
	if nil == ___Vobj {
		_FpfNex(" 1831916 08 , why nil ?")
	} else {
		_FdecGob___(___VeMsg+" 1831916 09 ", __VbufTmp2, ___Vobj)
	}
	return __VbufTmp2
} // _Fread_gob_rand_only_Exit

func _FtestER__write_gob_and_rand_Exit(___VeMsg string, ___Vkey []byte, ___Vfname string, ___Vobj interface{}) {
	__VbufText1 := _Fwrite_gob_only_Exit(___VeMsg+" 1831911 00 ", ___Vfname, ___Vobj)
	_Fwrite_json_only_Exit(___VeMsg+" 1831911 01 ", ___Vfname+".json", ___Vobj)

	_Fwrite_gob_rand_only_Exit(___VeMsg+" 1831911 02 ", ___Vkey, ___Vfname+".rand", ___Vobj)

	__VbufText2 := _Fread_gob_rand_only_Exit(___VeMsg+" 1831911 03 ", ___Vkey, ___Vfname+".rand", ___Vobj)

	_FfalseExit(" 1831911 05 ", bytes.Equal(__VbufText1, __VbufText2))

	if 3 == 3 {
		var __Vobj3 _TsrvInfo
		_Fread_gob_rand_only_Exit(___VeMsg+" 1831911 13 ", ___Vkey, ___Vfname+".rand", &__Vobj3)
		__VbufText3 := []byte(_Spf("%#v", __Vobj3))
		_FwriteFileExit(___VeMsg+" 1831911 14 ", ___Vfname+".rand.ex.printf.txt", &__VbufText3)
	}

	_FpfNdb(" 1831911 06 : gob rand test ok. ")
	//_FpfNhex(&__VbufText1, 80, " 1831911 07 ")
	//_FpfNhex(&__VbufText2, 80, " 1831911 08 ")
	//_FpfN(" 1831911 09 : <%s>", __VbufText2)
} // _FtestER__write_gob_and_rand_Exit
