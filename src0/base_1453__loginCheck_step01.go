package main

func (___Vlc *_TloginCheck) _FloginCheck_step102__sReply_tokenB(___Vdecode *_Tdecode) {
	if _FcheckDecodeType(___Vdecode, Cmd__loginS1ReqTryNoToken) {
		_FpfNdb(" 838393 01 type error , ignore ")
		return
	}

	_FpfNdb(" 838393 01 : ...... ")
	//___Vlc.ucMapConnA        map[[16]byte]_TconnInfo // _TreqIneedToLogin MeIdx128

	__Vk := _FgenB16(&___Vdecode.D__loginS1ReqTryNoToken.MeIdx128)

	_FpfNdb(" 838393 03 : key is <%x> ", __Vk)
	___Vlc.ucMapConnA[__Vk] = *___Vdecode
}
