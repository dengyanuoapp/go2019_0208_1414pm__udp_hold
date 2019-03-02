package main

// _FencAesRand_only
// _FencAesRandExit
func _Ftry_gen_json01(___Vfname string, ___Vkey *[]byte, ___VoriginObj interface{}) {

	//_FgenRand_nByte__testExit(8)
	//_FaesRand_test__en_de_Exit("asa90as90sa90as90as9as82391")

	_FtestER__write_json_and_rand_Exit(" 192181 ", ___Vkey, "json/"+___Vfname+".json", ___VoriginObj)
} // _Ftry_gen_json01

func _Ftry_download_rand_json01(___VdownUri string, ___Vkey *[]byte, ___VrecoverObj interface{}) {

	__Vtmp1, __Verr := _Fhttp_getLimit(___VdownUri, 1024*64)
	_FerrExit(" 638191 ", __Verr)

	//_FpfhexN(&__Vtmp1, 40, " 638192 %x :", _Fbase_1101b__gen_md5Only(&__Vtmp1))

	__Vtmp2 := _FdecAesRandExit(" 638196 ", ___Vkey, &__Vtmp1)
	//_FpfhexN(&__Vtmp2, 40, " 638197 %s :", __Vtmp2)

	if nil != ___VrecoverObj {
		_FdecJson___(" 638198 ", &__Vtmp2, ___VrecoverObj)
	}

} //  _Ftry_download_rand_json01
