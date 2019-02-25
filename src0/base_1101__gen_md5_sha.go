package main

import (
    //"bytes"
    //"os"
    //"io/ioutil"
    //"crypto/md5"
    "crypto/sha256"
    //"log"
    //"encoding/hex"
    //"time"
    "strconv"
)

type _Tsha          struct {
    b256                []byte
    A1                  uint64
    A2                  uint64
    A3                  uint64
    A4                  uint64
}

func _Fbase_1101a__gen_shaOnly( ___VbyteIn []byte ) ([]byte) {
    __Vb    :=  _FmakeByte32( sha256.Sum256( ___VbyteIn ) )
    return __Vb
} // _Fbase_1101a__gen_shaOnly

func ( ___Vsha _Tsha ) _Fbase_1101__gen_shaT( ___VshaInput []byte ) {
    ___Vsha . b256      =   _Fbase_1101a__gen_shaOnly( ___VshaInput )
    ___Vsha . A1 , _    =   strconv.ParseUint( string(___Vsha . b256[0:7])     , 16, 64 )
} // _Fbase_1101__gen_shaT

