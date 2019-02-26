package main

func _FencAesRand_only(___Vdst interface{}) ([]byte, error){
    // 	var __Vb []byte
    // 	__Vb = _FencJsonExit(" 381915 ", ___Vdst)
    // 	//__Vb = _FencAesRandExit(" 381916 ", ___Vb)
    // 	_FwriteFileExit(" 381917 ", ___Vfname, &__Vb)
    return nil, nil
} // _FencAesRand_only

func _FencAesRandExit(___VeMsg string , ___Vkey *[]byte , ___Vdst interface{}) ([]byte){
    __Vb , __Verr := _FencAesRand_only( ___Vdst )
	_FerrExit( ___VeMsg + " 192395 ", __Verr )
    return __Vb
} // _FencAesRandExit

