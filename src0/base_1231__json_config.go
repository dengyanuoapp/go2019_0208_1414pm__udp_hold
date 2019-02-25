package main

import (
    "os"
    "io/ioutil"
)

type    _TjsonConfig    struct {
    Name            string
    Id128           []byte
} //    _TjsonConfig

var (
    _VjsonConfig_need_save      bool
    _VjsonConfig_Now            _TjsonConfig
)

func _Fbase_105__get_or_gen_id128() (___VrtB []byte) {

    __Vstr := os.Getenv("id128")
    _FpfN( " env id128 is (%d)[%s]" , len(__Vstr) , __Vstr )
    if ( "" == __Vstr ) {
        return
    }
    _Fex( " id128 now is [%s] " , __Vstr )
    return
    //_self_id128                  =   _Fhexstr2uint64( os.Getenv("id128") )
} // _Fbase_105__get_or_gen_id128


// _Fbase_1203__gen_self_md5_sha
func _Fbase_104a__try_to_read_json_config_file() {
    __Vfname := _self_prog + ".json"
    __Vcontent, __Verr := ioutil.ReadFile( __Vfname )
    if ( nil != __Verr ) {
        _FpfN( " 389191 read config file <" + __Vfname + "> error, try to gen it..." )
        _VjsonConfig_need_save = true
        return
    }
} // _Fbase_104a__try_to_read_json_config_file

func _Fbase_104__try_to_read_json_config_top() {

    _Fbase_104a__try_to_read_json_config_file()

    _Fbase_105__get_or_gen_id128()
} // _Fbase_104__try_to_read_json_config_top

