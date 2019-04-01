package main

import "reflect"

var __VfirstLoginLenArr []int = []int{0, 0, 16, 16, 16, 0}

func (___Vlc *_TloginCheck) _FloginCheck_step900201y__s2Reply_tokenB(___Vdecode *_Tdecode) {
	if _FcheckDecodeType(___Vdecode, Cmd__loginS1ReqTryNoToken) {
		_FpfNdb(" 838393 01 type error , ignore ")
		return
	}

	////_FpfNdb(" 838393 01 : ...... ")
	//_FpfNdb(" 838393 02 : %s", ___Vdecode.String()) // 15540463611554046361
	////___Vlc.ucCmd.M        map[[16]byte]_TconnInfo // _TloginReq MeIdx128

	__VlenArr := ___Vdecode.Count128()
	if false == reflect.DeepEqual(__VlenArr, __VfirstLoginLenArr) {
		_FpfNdb(" 838393 03 len error , ignore %d ", __VlenArr)
		return
	}

	__Vk := _FgenB16(&___Vdecode.Dlogin.MeIdx128)

	if len(___Vlc.ucCmd.M) > 300 {
		_FdeleteOld_cmdStack(&___Vlc.ucCmd.M)
	}

	//_FpfNdb(" 838393 04 : key is <%x> ", __Vk)
	___Vdecode.Dlogin.TokenR = _FgenRand_nByte__(16)

	___Vlc.ucCmd.M[__Vk] = *___Vdecode

	_FpfNdb(" 838393 06 : %s", ___Vdecode.String())

	___Vlc._FloginCheck_step102y__sReply_tokenB(___Vdecode)
}

var __VmaxCmdPerid int = 100

func _FdeleteOld_cmdStack(___Vm *map[[16]byte]_Tdecode) {
	var __Vdel [][16]byte
	__Vnow := _FtimeInt()
	for __k, __v := range *___Vm {
		if __Vnow-__v.receiveTime > __VmaxCmdPerid {
			__Vdel = append(__Vdel, __k)
		}
	}
	for _, __v := range __Vdel {
		delete(*___Vm, __v)
	}
}

// _TloginReq _Tdecode
func (___Vlc *_TloginCheck) _FloginCheck_step102y__sReply_tokenB(___Vdecode *_Tdecode) {
	//var ___VnewUnSend _TudpNodeDataSend

	___Vreq := _TloginReq{
		MeTime:   _FtimeInt(),                 // int64
		ReqStr:   " step102y__sReply_tokenB ", // string
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _VS.meSeq128,
		ToIdx128: ___Vdecode.Dlogin.MeIdx128, // []byte
		ToSeq128: ___Vdecode.Dlogin.MeSeq128, // []byte
		TokenL:   ___Vdecode.Dlogin.TokenR,   // []byte
		TokenR:   ___Vdecode.Dlogin.TokenL,   // []byte
	}

	_FpfNdb(" 838394 01 start %s", ___Vreq.String())
	if nil == ___Vlc.ucCHSendLO {
		_FpfN(" 838394 02 , why output-Chan nil ? ")
	} else {
		_FpfN(" 838394 03 , fake Chan ")
		//__VucPort._FdataPack__101__udpConnPort(&__VusData.usOutBuf)
		// (*___Vlc.ucCHSendLO) <- ___VnewUnSend // _TudpNodeDataSend
	}
}
