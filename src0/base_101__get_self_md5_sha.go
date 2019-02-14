package main

//import "fmt"
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
    _P.pn("from ", os.Args[0])
    _P.pfN("from:%s", os.Args[0])

    __Vcontent, __Verr := ioutil.ReadFile( os.Args[0] )
    if __Verr != nil {
        _P.EX(__Verr)
    }

    _self_md5 := md5.Sum(__Vcontent)
    _self_sha := sha256.Sum256(__Vcontent)
    _P.pf("File md5: [ %x ]\n", _self_md5)
    _P.pf("File sha: [ %x ]\n", _self_sha)

} // _Fbase_101__get_self_md5_sha
