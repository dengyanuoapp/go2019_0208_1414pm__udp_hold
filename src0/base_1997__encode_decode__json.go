package main

import (
    //"bytes"
    //"encoding/binary"
    //"encoding/gob"

    "encoding/json"
    //"io/ioutil"

    //"log"
    //"fmt"
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
func _FencJson(___V interface{}) ([]byte , error) {
    __VbByte , __Verr := json.Marshal( ___V )
    if __Verr != nil {
        _Perr( __Verr , "1831917 : gob.NewDecoder failed:" )
        return nil , __Verr
    }
    return __VbByte , __Verr
} // _FencJson

func _FencJsonExit( ___VeMsg string , ___V interface{} ) ([]byte) {
    //__Vbyte , __Verr := _FencJson( ___V )
    //_FerrExit( " 1831918 : json error : " + ___VeMsg , __Verr )
    //_FerrExit( " 1831918 : json error : " , __Verr )

    //__Vbyte , _:= _FencJson( ___V )
    __Vbyte , __Verr := _FencJson( ___V )
    _FerrExit( ___VeMsg + " 1831918 : json error : " , __Verr )

    return __Vbyte
} // _FencJsonExit

//    func Unmarshal(data []byte, v interface{}) error
func _FdecJson(___Vbyte []byte, ___Vout interface{}) {
    __Verr := json.Unmarshal( ___Vbyte , ___Vout )
    if __Verr != nil {
        _Perr( __Verr , "1831919 : gob.NewDecoder failed:" )
    }
} // _FdecJson




func _Fwrite_json_only_Exit(___VeMsg string , ___Vfname string, ___Vdst interface{}) {
	var __Vb []byte
	__Vb = _FencJsonExit(___VeMsg + " 381911 ", ___Vdst)
	_FwriteFileExit(___VeMsg + " 381912 ", ___Vfname, &__Vb)
} // _Fwrite_json_only_Exit

func _Fwrite_json_rand_only_Exit(___VeMsg string , ___Vkey *[]byte , ___Vfname string, ___Vdst interface{}) {
	var __Vb []byte
	__Vb = _FencJsonExit(___VeMsg + " 381915 ", ___Vdst)
	__Vb = _FencAesRandExit(___VeMsg + " 381916 ", ___Vkey , __Vb)
	_FwriteFileExit(___VeMsg + " 381917 ", ___Vfname, &__Vb)
} // _Fwrite_json_rand_only_Exit

func _Fwrite_json_and_rand_Exit(___VeMsg string , ___Vkey *[]byte , ___Vfname string, ___Vdst interface{}) {
    _Fwrite_json_only_Exit( ___VeMsg + " 381918 " , ___Vfname , ___Vdst )
    _Fwrite_json_rand_only_Exit( ___VeMsg + " 381919 " , ___Vkey , ___Vfname + ".rand" , ___Vdst )
} // _Fwrite_json_and_rand_Exit

