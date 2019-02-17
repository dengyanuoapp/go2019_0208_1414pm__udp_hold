package main

//import "fmt"
import "os"
import "io/ioutil"
import "crypto/md5"
import "crypto/sha256"
import "log"
//import "std"

var (
    _self_md5 [16]byte // md5sum    : 727bf338cf523b90baccd24cca30b919
    _self_sha [32]byte // sha256sum :  2c6e3b458d5c482bc52a1d7d4f5a7d7766381c9f07d9b32ca605ae45b4e473f5
)

func _Fbase_101__get_self_md5_sha() {

    __Vcontent, __Verr := ioutil.ReadFile( os.Args[0] )
    if __Verr != nil {
        log.Fatalf( "err138191" , __Verr )
    }

    _self_md5 = md5.Sum(__Vcontent)
    _self_sha = sha256.Sum256(__Vcontent)

    _Ppn("from:", os.Args[0])
    _FpfN("File md5: [ %x ]", _self_md5)
    _FpfN("File sha: [ %x ]", _self_sha)

} // _Fbase_101__get_self_md5_sha
