package main

var (
    _VencAesRand_iv []byte
)

func _FencAesRand__gen_iv(){
}// _FencAesRand__gen_iv
func _FencAesRand_only(___Vkey *[]byte , ___VbyteIn *[]byte) ([]byte, error){
	_FpfN( " 192391 key (%d) %x , byte (%d) %x , %s ", len(*___Vkey) , *___Vkey , len(*___VbyteIn) , *___VbyteIn , string(*___VbyteIn) )

    _FencAesRand__gen_iv()

    __Vbyte , __Verr    :=  _FencAesCbc__only( ___Vkey , &_VencAesRand_iv , ___VbyteIn )

    return __Vbyte , __Verr
} // _FencAesRand_only

func _FencAesRandExit(___VeMsg string , ___Vkey *[]byte , ___VbyteIn *[]byte) ([]byte){
    __Vb , __Verr := _FencAesRand_only( ___Vkey , ___VbyteIn )
	_FerrExit( ___VeMsg + " 192395 ", __Verr )
    return __Vb
} // _FencAesRandExit

