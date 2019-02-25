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
    _VjsonConfig_bytes          []byte
)

func _Fbase_104c__try_to_get_env_id128() (___VrtB []byte) {

    __Vstr := os.Getenv("id128")
    _FpfN( " env id128 is (%d)[%s]" , len(__Vstr) , __Vstr )
    if ( "" == __Vstr ) {
        return
    }
    _Fex( " id128 now is [%s] " , __Vstr )
    return
    //_self_id128                  =   _Fhexstr2uint64( os.Getenv("id128") )
} // _Fbase_104c__try_to_get_env_id128

func _Fbase_104b__try_to_check_json_config() {
    if ( _VjsonConfig_Now . Name != _VprojectName ) {
        _VjsonConfig_Now . Name = _VprojectName
        _VjsonConfig_need_save = true
    }
} // _Fbase_104b__try_to_check_json_config

// _Fbase_1203__gen_self_md5_sha
func _Fbase_104a__try_to_read_json_config_file() {
    __Vfname := _self_prog + ".json"
    _VjsonConfig_bytes, __Verr := ioutil.ReadFile( __Vfname )
    if ( nil != __Verr ) {
        _FpfN( " 389191 read config file <" + __Vfname + "> error, try to gen it..." )
        _VjsonConfig_need_save = true
        return
    }

    _FdecJson( _VjsonConfig_bytes , &_VjsonConfig_Now )

} // _Fbase_104a__try_to_read_json_config_file

func _Fbase_104z__try_to_read_json_config_top() {

    _Fbase_104a__try_to_read_json_config_file()
    _Fbase_104b__try_to_check_json_config()
    _Fbase_104c__try_to_get_env_id128()

} // _Fbase_104z__try_to_read_json_config_top

