package main

import (
	"bytes"
	"encoding/hex"
	"io/ioutil"
	"os"
)

var (
	_VjsonConfig_need_save bool
	//_VC       _Tconfig
	//_VjsonConfig_bytes []byte
)

func _Fbase_104c__try_to_get_env_id128() {

	__Vstr := os.Getenv("id128")
	_FpfN(" 823813 01 read env id128 is (%d)[%s]", len(__Vstr), __Vstr)
	if "" == __Vstr || len(__Vstr) < 32 { // id env error , try use old id
		if 16 != len(_VC.MyId128) { // old id error -> I don't know what happens. exit.
			_FpfN(" 823813 03 read env id128 is NULL or too short, and the json.config file id128 error: \n (%d)[%0x] \n",
				len(_VC.MyId128), _VC.MyId128)
			_Fex1(" 823813 05 : Exit now ")
		}
		_FpfN(" 823813 07 env id128 is NULL , and so old id128 saved in config file ok: \n (%d)[%0x] \n",
			len(_VC.MyId128), _VC.MyId128)

		return
	}

	// func hex.DecodeString(s string) ([]byte, error)
	//if ( __Vstr[0:2] == "0x" ) || ( __Vstr[0:2] == "0X" ) {
	//__Vstr = _FmakeByte( []byte(__Vstr[2:]) )
	//}
	__Vstr = __Vstr[2:]
	__Vbyte, __Verr := hex.DecodeString(__Vstr)
	if nil != __Verr {
		_FpfN(" 823815 01 read env id128 is error , check what happens : \n (%d)[%s] \n %v\n",
			len(__Vbyte), string(__Vbyte), __Verr)
		_Fex1(" Exit now ")
	}
	if 16 != len(__Vbyte) {
		_FpfN(" 823815 02 read env id128 is len error , check what happens \n (%d)[%0x] \n",
			len(__Vbyte), string(__Vbyte))
		_Fex1(" Exit now ")
	}

	if bytes.Equal(__Vbyte, _VC.MyId128) {
		_FpfN(" 823815 03 read env id128 equals to json's id128\n (%d)[%0x] \n",
			len(__Vbyte), string(__Vbyte))
		return
	}

	_FpfN(" 823815 04 read env id128 NOT equals to json's id128\n env : (%d)[%0x] \n json: (%d)[%0x] \n",
		len(__Vbyte), string(__Vbyte),
		len(_VC.MyId128), string(_VC.MyId128))

	_VC.MyId128 = __Vbyte
	_VjsonConfig_need_save = true
} // _Fbase_104c__try_to_get_env_id128

func _Fbase_104b__try_to_check_json_config() {
	if _VC.Name != _Vself.ProjName {
		_VC.Name = _Vself.ProjName
		_VjsonConfig_need_save = true
	}
} // _Fbase_104b__try_to_check_json_config

// _Fbase_1203__gen_self_md5_sha
func _Fbase_104a__try_to_read_json_config_file() {
	__Vfname := _Vself.progPath + ".json"
	__VjsonConfig_bytes, __Verr := ioutil.ReadFile(__Vfname)
	if nil != __Verr {
		_FpfN(" 823816 01 read config file <"+__Vfname+"> error, try to gen it... %v", __Verr)
		_VjsonConfig_need_save = true
		return
	}

	_FdecJson___(" 823816 02 ", &__VjsonConfig_bytes, &_VC)

} // _Fbase_104a__try_to_read_json_config_file

func _Fbase_104d__try_to_save_json_config_to_file() {
	__Vbyte := _FencJsonExit("823818 01 : jsonConf encoding ", _VC)

	__Vfname := _Vself.progPath + ".json"
	_FwriteFileExit("823818 03 jsonconf writing ", __Vfname, &__Vbyte)

	__Vbyte2 := _FreadFileExit(" 823815 config file re-reading ", __Vfname)
	if bytes.Equal(__Vbyte, __Vbyte2) {
		_FpfN("823818 05 jsonconf writing succeed :<%s>", __Vbyte2)
		return
	}
	_Fex1(" 823818 08 : I don't know what happens.")

} // _Fbase_104d__try_to_save_json_config_to_file

func _Fbase_104e__try_to_reread_json_config_and_recheck_the_result() {
} // _Fbase_104e__try_to_reread_json_config_and_recheck_the_result

func _Fbase_104z__try_to_read_json_config_top() {

	_Fbase_104a__try_to_read_json_config_file()
	_Fbase_104b__try_to_check_json_config()
	_Fbase_104c__try_to_get_env_id128()

	_FpfN(" 823819 01 :%t ", _VjsonConfig_need_save)
	if true == _VjsonConfig_need_save {
		_Fbase_104d__try_to_save_json_config_to_file()
	}

	//_Vself.progMd5.b128 = _VC.MyId128
	//_Vself.ProjName = _VC.Name

	//_Fex1( " 381919 :Debug Stop here. " )
} // _Fbase_104z__try_to_read_json_config_top
