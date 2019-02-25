package main

import (
    "bytes"
    "os"
    "io/ioutil"
    "crypto/md5"
    "crypto/sha256"
    "log"
    "encoding/hex"
    "time"
    "strconv"
)

var (
    _self_startTime     time.Time
    _self_startTimeSha  _Tb256


    _self_md5           _Tb128 // md5sum    : 16 byte : 727bf338cf523b90baccd24cca30b919
    _self_sha           _Tb256 // sha256sum : 32 byte : 2c6e3b458d5c482bc52a1d7d4f5a7d7766381c9f07d9b32ca605ae45b4e473f5

    _self_id64          uint64

    _self_rand          _Tb256

    _debugEnabled       bool
)

func _Fbase_1203__gen_self_md5_sha() {

    __Vcontent, __Verr := ioutil.ReadFile( os.Args[0] )
    if __Verr != nil {
        log.Fatalf( "err138191" , __Verr )
    }

    __Vself_md5 := md5.Sum(__Vcontent)
    __Vself_sha := sha256.Sum256(__Vcontent)

    _self_md5   = _FmakeByte( __Vself_md5[:] )
    _self_sha   = _FmakeByte( __Vself_sha[:] )
    _self_shA1  , _ = strconv.ParseUint( string(_self_sha[0:7])     , 16, 64 ) 
    _self_shA2  , _ = strconv.ParseUint( string(_self_sha[8:15])    , 16, 64 ) 
    _self_shA3  , _ = strconv.ParseUint( string(_self_sha[16:23])   , 16, 64 ) 
    _self_shA4  , _ = strconv.ParseUint( string(_self_sha[24:31])   , 16, 64 ) 

    // prog : 1
    // prog x1 x2 x3 : 4
    __VparaLen := len( os.Args )
    if 1 != __VparaLen {
        // func hex.DecodeString(s string) ([]byte, error)
        __Vhex , __Verr := hex.DecodeString( os.Args[ __VparaLen - 1 ] )
        if nil == __Verr {
            //if fmt.Sprintf("%x" , __Vhex) == fmt.Sprintf("%x" , _self_sha ) {
            if bytes.Equal( __Vhex , _self_sha ) {
                _debugEnabled = true
            }
        }
    }
    log.Printf( " _debugEnabled : %t\n" , _debugEnabled )


    _Ppn("from:", os.Args[0])
    _FpfN("File md5: [ %x ]", _self_md5)
    _FpfN("File sha: [ %x ] %x %x %x %x ", _self_sha , _self_shA1 , _self_shA2 , _self_shA3 , _self_shA4 )

} // _Fbase_1203__gen_self_md5_sha

func _Fbase_103__gen_rand_seed() {
    _self_startTime             = time.Now()
    __Vsha  :=  sha256.Sum256( []byte(_Pspf( "%x" , _self_startTime )) )
    _self_startTimeSha          = _FmakeByte( __Vsha[:] )
    _self_startTimeShA1 , _     = strconv.ParseUint( _self_startTimeSha[0:7]     , 16, 64 ) 
    _self_startTimeShA2 , _     = strconv.ParseUint( _self_startTimeSha[8:15]    , 16, 64 ) 
    _self_startTimeShA3 , _     = strconv.ParseUint( _self_startTimeSha[16:23]   , 16, 64 ) 
    _self_startTimeShA4 , _     = strconv.ParseUint( _self_startTimeSha[24:31]   , 16, 64 ) 

} // _Fbase_103__gen_rand_seed() 

func _Fbase_105__get_or_gen_id64() {

    _self_id64                  =   _Fhexstr2uint64( os.Getenv("idhex") )
    if ( 0 == _self_id64 ) {
        _Fex1( " err23818911 " )
    }
} // _Fbase_105__get_or_gen_id64

