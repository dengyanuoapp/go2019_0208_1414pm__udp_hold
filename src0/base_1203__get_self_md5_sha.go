package main

import (
    "bytes"
    "os"
    "io/ioutil"
    //"crypto/md5"
    //"crypto/sha256"
    "log"
    "encoding/hex"
    "time"
    //"strconv"
)

var (
    _self_prog          string

    _self_startTime     time.Time
    _self_startTimeSha  _Tb256


    _self_md5           _Tb128 // md5sum    : 16 byte : 727bf338cf523b90baccd24cca30b919
    _self_sha           _Tb256 // sha256sum : 32 byte : 2c6e3b458d5c482bc52a1d7d4f5a7d7766381c9f07d9b32ca605ae45b4e473f5

    _self_id64          uint64

    _self_rand          _Tb256

    _debugEnabled       bool
)

func _Fbase_1203__gen_self_md5_sha() {

    _self_prog = os.Args[0]
    __Vcontent, __Verr := ioutil.ReadFile( _self_prog )
    if __Verr != nil {
        log.Fatalf( "err138191" , __Verr )
    }

    _self_sha . _Fbase_1101__gen_shaT( __Vcontent )
    _self_md5 . _Fbase_1101__gen_md5T( __Vcontent )

    // prog : 1
    // prog x1 x2 x3 : 4
    __VparaLen := len( os.Args )
    if 1 != __VparaLen {
        // func hex.DecodeString(s string) ([]byte, error)
        __Vhex , __Verr := hex.DecodeString( os.Args[ __VparaLen - 1 ] )
        if nil == __Verr {
            //if fmt.Sprintf("%x" , __Vhex) == fmt.Sprintf("%x" , _self_sha ) {
            if bytes.Equal( __Vhex , _self_sha.b256 ) {
                _debugEnabled = true
            }
        }
    }
    log.Printf( " _debugEnabled : %t\n" , _debugEnabled )


    _Ppn("from:", _self_prog )
    _FpfN("8381191 File md5: [ %x ]", _self_md5.b128)
    _FpfN("8381192 File sha: [ %x ] %x %x %x %x ", _self_sha.b256 , _self_sha.A1 , _self_sha.A2 , _self_sha.A3 , _self_sha.A4 )

} // _Fbase_1203__gen_self_md5_sha

func _Fbase_103__gen_rand_seed() {
    _self_startTime             = time.Now()
    _self_startTimeSha          . _Fbase_1101__gen_shaT( []byte(_Pspf( "%x" , _self_startTime )) )
} // _Fbase_103__gen_rand_seed() 

func _Fbase_105__get_or_gen_id64() {

    _self_id64                  =   _Fhexstr2uint64( os.Getenv("idhex") )
    if ( 0 == _self_id64 ) {
        _Fex1( " err23818911 " )
    }
} // _Fbase_105__get_or_gen_id64

