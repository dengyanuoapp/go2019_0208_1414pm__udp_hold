package main

func _FencAesRand_only(___Vkey *[]byte , ___Vbyte *[]byte) ([]byte, error){
	_FpfN( " 192391 key %x , byte %x , %s ", ___Vkey , ___Vbyte , string(*___Vbyte) )
    // 	var __Vb []byte
    // 	__Vb = _FencJsonExit(" 381915 ", ___Vbyte)
    // 	//__Vb = _FencAesRandExit(" 381916 ", ___Vb)
    // 	_FwriteFileExit(" 381917 ", ___Vfname, &__Vb)
    return nil, nil
} // _FencAesRand_only

func _FencAesRandExit(___VeMsg string , ___Vkey *[]byte , ___Vbyte *[]byte) ([]byte){
    __Vb , __Verr := _FencAesRand_only( ___Vkey , ___Vbyte )
	_FerrExit( ___VeMsg + " 192395 ", __Verr )
    return __Vb
} // _FencAesRandExit

