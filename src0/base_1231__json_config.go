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
    _FpfN( " 893871 read env id128 is (%d)[%s]" , len(__Vstr) , __Vstr )
    if ( "" == __Vstr ) {
        if ( 16 != len( _VjsonConfig_Now . Id128 ) ) {
            _FpfN( " 893873 read env id128 is NULL , and the json.config file id128 error: \n (%d)[%0x] \n" , 
            len(_VjsonConfig_Now . Id128 ) , _VjsonConfig_Now . Id128  )
            _Fex1( " 893874 : Exit now ")
        }
        _FpfN( " 893875 env id128 is NULL , and the json.config file id128 ok: \n (%d)[%0x] \n" , 
        len(_VjsonConfig_Now . Id128 ) , _VjsonConfig_Now . Id128  )

        return
    }

    // func hex.DecodeString(s string) ([]byte, error)
    __Vbyte , __Verr := hex.DecodeString( __Vstr )
    if ( nil != __Verr ) {
        _FpfN( " 893876 read env id128 is error , check what happens : \n (%d)[%s] \n" , 
        len( __Vbyte ) , string(__Vbyte)  )
        _Fex1( " Exit now ")
    }
    if ( 16 != len( __Vbyte ) ) {
        _FpfN( " 893877 read env id128 is len error , check what happens \n (%d)[%0x] \n" , 
        len( __Vbyte ) , string(__Vbyte)  )
        _Fex1( " Exit now ")
    }

    if ( __Vbyte == _VjsonConfig_Now . Id128 ) {
        _FpfN( " 893878 read env id128 equals to json's id128\n (%d)[%0x] \n" , 
        len( __Vbyte ) , string(__Vbyte)  )
        return
    }

    _VjsonConfig_Now . Id128 = __Vbyte 
    _FpfN( " 893879 read env id128 NOT equals to json's id128\n env : (%d)[%0x] \n json: (%d)[%0x] \n" , 
    len( __Vbyte ) , string(__Vbyte)  ,
    len( _VjsonConfig_Now . Id128 ) , string(_VjsonConfig_Now . Id128)  )

    _VjsonConfig_need_save = true
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
        _FpfN( " 389191 read config file <" + __Vfname + "> error, try to gen it... %v" , __Verr)
        _VjsonConfig_need_save = true
        return
    }

    _FdecJson( _VjsonConfig_bytes , &_VjsonConfig_Now )

} // _Fbase_104a__try_to_read_json_config_file

func _Fbase_104d__try_to_save_json_config_to_file() {
    _Fbase_104e__try_to_reread_json_config_and_recheck_the_result()
} // _Fbase_104d__try_to_save_json_config_to_file

func _Fbase_104e__try_to_reread_json_config_and_recheck_the_result() {
} // _Fbase_104e__try_to_reread_json_config_and_recheck_the_result

func _Fbase_104z__try_to_read_json_config_top() {

    _Fbase_104a__try_to_read_json_config_file()
    _Fbase_104b__try_to_check_json_config()
    _Fbase_104c__try_to_get_env_id128()
    if ( true == _VjsonConfig_need_save ) {
        _Fbase_104d__try_to_save_json_config_to_file() 
    }
    _Fex( " 381919 :Debug Stop here." )
} // _Fbase_104z__try_to_read_json_config_top

