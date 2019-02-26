package main

import (
)

// block is 128 bit , 16 byte
// key   is 256 bit , 32 byte : fixed use 256 aes , not 128 , 196 
type _Taes      struct {
    name            string
    key             []byte
} // _Taes      


func _FencAesCbc__only(___Vkey *[]byte , ___Viv *[]byte , ___VbyteIn *[]byte) ([]byte, error){
    return nil,nil
} // _FencAesCbc__only

func _FencAesCbcExit(___Vkey *[]byte , ___Viv *[]byte , ___VbyteIn *[]byte) ([]byte){
    __Vbyte , __Verr := _FencAesCbc__only(___Vkey , ___Viv , ___VbyteIn )
    _FerrExit( " 182811 " ,  __Verr )
    return __Vbyte
} // _FencAesCbcExit
