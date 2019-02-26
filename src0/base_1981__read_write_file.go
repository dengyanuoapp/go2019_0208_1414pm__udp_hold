package main

import (
    //"bytes"
    //"encoding/binary"
    //"encoding/gob"

    //"encoding/json"
    "io/ioutil"

    //"log"
    //"fmt"
    //"math"
)

func _FwriteFile( ___Vfname string , ___Vcontent *[]byte ) error {
    __Verr := ioutil.WriteFile(___Vfname, *___Vcontent, 0666);
    return __Verr
} // _FwriteFile

func _FwriteFileExit( ___VerrMsg , ___Vfname string , ___Vcontent *[]byte ) {
    __Verr := _FwriteFile( ___Vfname , ___Vcontent )
    if ( nil == __Verr ) {
        return
    }
    _FpfN( " 388191 file write error : " + ___VerrMsg + "\n %v" , __Verr )
    _Fex1( "" )
} // _FwriteFileExit

//    _VjsonConfig_bytes, __Verr := ioutil.ReadFile( __Vfname )
func _FreadFile( ___Vfname string ) ([]byte , error) {
    __Vcontent , __Verr := ioutil.ReadFile(___Vfname)
    return __Vcontent , __Verr
} // _FreadFile

func _FreadFileExit( ___VerrMsg , ___Vfname string ) ( []byte ) {
    __Vcontent , __Verr := ioutil.ReadFile(___Vfname)
    if ( nil == __Verr ) {
        return __Vcontent
    }
    _FpfN( " 823181 file read error : " + ___VerrMsg + "\n %v" , __Verr )
    _Fex1( "" )
    return nil
} // _FreadFileExit
