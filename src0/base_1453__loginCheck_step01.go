package main

func (___Vlc *_TloginCheck) _FloginCheck_step102__sReply_tokenB(___Vdecode *_Tdecode) {
	if _FcheckDecodeType(___Vdecode, Cmd__loginS1ReqTryNoToken) {
		_FpfNdb(" 838393 01 type error , ignore ")
		return
	}

	_FpfNdb(" 838393 01 : ...... ")
	//___Vlc.ucMapConnA        map[[16]byte]_TconnInfo // _TreqIneedToLogin MeIdx128

	__Vk := _FgenB16(&___Vdecode.D__loginS1ReqTryNoToken.MeIdx128)

	if len(___Vlc.ucMapConnA) > 300 {
		_FdeleteOld_conA(&___Vlc.ucMapConnA)
	}

	_FpfNdb(" 838393 03 : key is <%x> ", __Vk)
	___Vdecode.D__loginS1ReqTryNoToken.TokenA = _FgenRand_nByte__(16)

	___Vlc.ucMapConnA[__Vk] = *___Vdecode
}

func _FdeleteOld_conA(___Vm *map[[16]byte]_Tdecode) {
	var __Vdel [][16]byte
	__Vnow := _FtimeInt()
	for __k, __v := range *___Vm {
		if __Vnow-__v.receiveTime > 100 {
			__Vdel = append(__Vdel, __k)
		}
	}
	for _, __v := range __Vdel {
		delete(*___Vm, __v)
	}
}
