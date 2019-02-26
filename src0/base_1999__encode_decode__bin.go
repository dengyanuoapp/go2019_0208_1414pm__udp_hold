package main

import (
    "bytes"
    "encoding/binary"
    //"encoding/gob"

    //"encoding/json"
    //"io/ioutil"

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

