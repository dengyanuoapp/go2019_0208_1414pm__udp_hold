package main

import "fmt"
import "os"
import "io/ioutil"
import "crypto/md5"
import "crypto/sha256"
//import "log"

var (
    _self_md5 []byte // md5sum    : 727bf338cf523b90baccd24cca30b919
    _self_sha []byte // sha256sum :  2c6e3b458d5c482bc52a1d7d4f5a7d7766381c9f07d9b32ca605ae45b4e473f5
)

func _Fbase_101__get_self_md5_sha() {
    _Ppn("1:from:", os.Args[0])
    _Ppf("2:from:%s", os.Args[0])
    _Ppf("2:from:%s\n", os.Args[0])

    __Vcontent, __Verr := ioutil.ReadFile( os.Args[0] )
    if __Verr != nil {
        _Pex(__Verr)
    }

    _self_md5 := md5.Sum(__Vcontent)
    _self_sha := sha256.Sum256(__Vcontent)
    _Ppf("31:File md5: [ %x ]", _self_md5)
    _Ppf("32:File md5: [ %x ]", _self_md5[0:]) ; _Ppn()
    _Ppf("33:File md5: [ %x ]", _self_md5[0:]) ; _Ppn()
    fmt.Printf("34:File md5: [ %x ]", _self_md5) ; _Ppn()
    _Ppf("4:File sha: [ %x ]", _self_sha)

} // _Fbase_101__get_self_md5_sha
