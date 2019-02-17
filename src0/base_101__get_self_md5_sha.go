package main

import "fmt"
import "os"
import "io/ioutil"
import "crypto/md5"
import "crypto/sha256"
import "log"
import "encoding/hex"

var (
    _self_md5           []byte // md5sum    : 16 byte : 727bf338cf523b90baccd24cca30b919
    _self_sha           []byte // sha256sum : 32 byte : 2c6e3b458d5c482bc52a1d7d4f5a7d7766381c9f07d9b32ca605ae45b4e473f5
    _debugEnabled       bool
)

func _Fbase_101__get_self_md5_sha() {

    __Vcontent, __Verr := ioutil.ReadFile( os.Args[0] )
    if __Verr != nil {
        log.Fatalf( "err138191" , __Verr )
    }

    _self_md5 = _FmakeByte( 16 , md5.Sum(__Vcontent) )
    _self_sha = _FmakeByte( 32 , sha256.Sum256(__Vcontent) )

    // prog : 1
    // prog x1 x2 x3 : 4
    __VparaLen := len( os.Args )
    if 1 != __VparaLen {
        // func hex.DecodeString(s string) ([]byte, error)
        __Vhex , __Verr := hex.DecodeString( os.Args[ __VparaLen - 1 ] )
        if nil == __Verr {
            if fmt.Sprintf("%x" , __Vhex) == fmt.Sprintf("%x" , _self_sha ) {
                _debugEnabled = true
            }
        }
    }
    log.Printf( " _debugEnabled : %t\n" , _debugEnabled )


    _Ppn("from:", os.Args[0])
    _FpfN("File md5: [ %x ]", _self_md5)
    _FpfN("File sha: [ %x ]", _self_sha)

} // _Fbase_101__get_self_md5_sha
