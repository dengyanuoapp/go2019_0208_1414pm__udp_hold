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

type _Tb256          struct {
    b256                []byte
    A1                  uint64
    A2                  uint64
    A3                  uint64
    A4                  uint64
}
type _Tb128          struct {
    b128                []byte
    A1                  uint64
    A2                  uint64
}

func _Fbase_1101a__gen_shaOnly( ___VbyteIn []byte ) ([]byte) {
    __Vb    :=  _FmakeByte32( sha256.Sum256( ___VbyteIn ) )
    return __Vb
} // _Fbase_1101a__gen_shaOnly

func ( ___Vsha _Tb256 ) _Fbase_1101__gen_shaT( ___VshaInput []byte ) {
    ___Vsha . b256      =   _Fbase_1101a__gen_shaOnly( ___VshaInput )
    ___Vsha . A1 , _    =   strconv.ParseUint( string(___Vsha . b256[0:7])      , 16, 64 )
    ___Vsha . A2 , _    =   strconv.ParseUint( string(___Vsha . b256[8:15])     , 16, 64 )
    ___Vsha . A3 , _    =   strconv.ParseUint( string(___Vsha . b256[16:23])    , 16, 64 )
    ___Vsha . A4 , _    =   strconv.ParseUint( string(___Vsha . b256[24:31])    , 16, 64 )
} // _Fbase_1101__gen_shaT

func ( ___Vmd5 _Tb128 ) _Fbase_1101__gen_md5T( ___Vmd5Input []byte ) {
    ___Vmd5 . b128      =   _Fbase_1101a__gen_shaOnly( ___Vmd5Input )
    ___Vmd5 . A1 , _    =   strconv.ParseUint( string(___Vmd5 . b128[0:7])      , 16, 64 )
    ___Vmd5 . A2 , _    =   strconv.ParseUint( string(___Vmd5 . b128[8:15])     , 16, 64 )
} // _Fbase_1101__gen_md5T

