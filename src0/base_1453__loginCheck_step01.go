package main

func (___Vlc *_TloginCheck) _FloginCheck_step900201y__sReply_tokenB(___Vdecode *_Tdecode) {
	if _FcheckDecodeType(___Vdecode, Cmd__loginS1ReqTryNoToken) {
		_FpfNdb(" 838393 01 type error , ignore ")
		return
	}

	_FpfNdb(" 838393 01 : ...... ")
	//___Vlc.ucM.cmAnow        map[[16]byte]_TconnInfo // _TloginReq MeIdx128

	__Vk := _FgenB16(&___Vdecode.D__loginS1ReqTryNoToken.MeIdx128)

	if len(___Vlc.ucM.cmAnow) > 300 {
		_FdeleteOld_conA(&___Vlc.ucM.cmAnow)
	}

	_FpfNdb(" 838393 03 : key is <%x> ", __Vk)
	___Vdecode.D__loginS1ReqTryNoToken.TokenA = _FgenRand_nByte__(16)

	___Vlc.ucM.cmAnow[__Vk] = *___Vdecode

	___Vlc._FloginCheck_step102y__sReply_tokenB(___Vdecode)
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

// _TloginReq
func (___Vlc *_TloginCheck) _FloginCheck_step102y__sReply_tokenB(___Vdecode *_Tdecode) {
	//var ___VnewUnSend _TudpNodeDataSend

	___Vreq := _TloginReq{
		MeTime:   _FtimeInt(),                 // int64
		ReqStr:   " step102y__sReply_tokenB ", // string
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _VS.meSeq128,
		ToIdx128: ___Vdecode.D__loginS1ReqTryNoToken.MeIdx128, // []byte
		ToSeq128: ___Vdecode.D__loginS1ReqTryNoToken.MeSeq128, // []byte
		// TokenA   : , // []byte
		// TokenB   : , // []byte
	}

	_FpfNdb(" 838394 01 start %s", ___Vreq.String())
	if nil == ___Vlc.ucCHSendLO {
		_FpfN(" 838394 02 , why output-Chan nil ? ")
	} else {
		_FpfN(" 838394 03 , fake Chan ")
		// (*___Vlc.ucCHSendLO) <- ___VnewUnSend // _TudpNodeDataSend
	}
}
