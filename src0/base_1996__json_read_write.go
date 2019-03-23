package main

import (
	"fmt"
)

// _FencAesRand_only
// _FencAesRandExit
func _Ftry_gen_json01(___Vfname string, ___Vkey *[]byte, ___VoriginObj interface{}) {
	_FtestER__write_json_and_rand_Exit(" 192181 01", ___Vkey, "json/"+___Vfname+".json", ___VoriginObj)
} // _Ftry_gen_json01

func _Ftry_gen_gob01(___Vfname string, ___Vkey *[]byte, ___VoriginObj interface{}) {
	_FtestER__write_gob_and_rand_Exit(" 192181 02", ___Vkey, "json/"+___Vfname+".gob", ___VoriginObj)
} // _Ftry_gen_gob01

func _Ftry_download_rand_json01(___VdownUri string, ___Vkey *[]byte, ___VrecoverObj interface{}) ([]byte, error) {

	__Vtmp1, __Verr := _Fhttp_getLimit(true, ___VdownUri, 1024*64)
	//__Vtmp1, __Verr := _Fhttp_getAll(true, ___VdownUri)
	if nil != __Verr {
		return nil, fmt.Errorf(" 192182 01 : %v", __Verr)
	}

	//_FpfNhex(&__Vtmp1, 40, " 638192 %x :", _Fbase_1101b__gen_md5Only(&__Vtmp1))

	//__Vtmp2 := _FdecAesRandExit(" 638194 ", ___Vkey, &__Vtmp1)
	__Vtmp2, __Verr := _FdecAesRand__only(___Vkey, &__Vtmp1)
	//_FpfNhex(&__Vtmp2, 40, " 638195 %s :", __Vtmp2)
	if nil != __Verr {
		//_FpfN(" 638197 %v :", __Verr)
		//_FpfNhex(&__Vtmp1, 40, " 638198 receive error data :")
		return nil, fmt.Errorf(" 192182 03 : %v", __Verr)
	}

	if nil != ___VrecoverObj {
		_FdecJson___(" 192182 05 ", &__Vtmp2, ___VrecoverObj)
	}

	//_FpfNhex(&__Vtmp2, 40, " 192182 06 : %v ", ___VrecoverObj)

	return __Vtmp2, nil
} //  _Ftry_download_rand_json01

func _Ftry_download_rand_gob01(___VdownUri string, ___Vkey *[]byte, ___VrecoverObj interface{}) ([]byte, error) {

	__Vtmp1, __Verr := _Fhttp_getLimit(true, ___VdownUri, 1024*64)
	//__Vtmp1, __Verr := _Fhttp_getAll(true, ___VdownUri)
	if nil != __Verr {
		return nil, fmt.Errorf(" 192183 01 : %v", __Verr)
	}

	//_FpfNhex(&__Vtmp1, 40, " 638192 %x :", _Fbase_1101b__gen_md5Only(&__Vtmp1))

	//__Vtmp2 := _FdecAesRandExit(" 638194 ", ___Vkey, &__Vtmp1)
	__Vtmp2, __Verr := _FdecAesRand__only(___Vkey, &__Vtmp1)
	//_FpfNhex(&__Vtmp2, 40, " 638195 %s :", __Vtmp2)
	if nil != __Verr {
		//_FpfN(" 638197 %v :", __Verr)
		//_FpfNhex(&__Vtmp1, 40, " 638198 receive error data :")
		return nil, fmt.Errorf(" 192183 03 : %v", __Verr)
	}

	if nil != ___VrecoverObj {
		//_FdecJson___(" 192183 05 ", &__Vtmp2, ___VrecoverObj)
		_FdecGob___(" 192183 06 ", &__Vtmp2, ___VrecoverObj)
	}

	return __Vtmp2, nil
} //  _Ftry_download_rand_gob01
