package main

import "reflect"

func (___Vlc *_TloginCheck) _FloginCheck_step900201y__s3accept_tokenA_file03send(___Vdecode *_Tdecode) {
	if _FcheckDecodeType(___Vdecode, Cmd__loginS02genReplyTokenB) {
		_FpfNdb(" 838381 01 type error , ignore ")
		return
	}

	////_FpfNdb(" 838381 01 : ...... ")
	//_FpfNdb(" 838381 02 : %s", ___Vdecode.String()) // 15540463611554046361
	////___Vlc.ucCmd.M        map[[16]byte]_TconnInfo // _TloginReq MeIdx128

	__VlenArr := ___Vdecode.Count128()
	if false == reflect.DeepEqual(__VlenArr, __Vstep03_LoginLenArr) {
		_FpfNdb(" 838381 03 len error , ignore %d ", __VlenArr)
		return
	}

	__Vk := _FgenB16(&___Vdecode.Dlogin.MeIdx128)

	if len(___Vlc.ucCmd.M) > 300 {
		_FdeleteOld_cmdStack(&___Vlc.ucCmd.M)
	}

	//_FpfNdb(" 838381 04 : key is <%x> ", __Vk)
	___Vdecode.Dlogin.TokenR = _FgenRand_nByte__(16)

	___Vlc.ucCmd.M[__Vk] = *___Vdecode

	//_FpfNdb(" 838381 06 : [decode:<%s>]", ___Vdecode.String())

	___Vlc._FloginCheck_step03__accept_tokenA(___Vdecode)
}

// _TloginReq _Tdecode
func (___Vlc *_TloginCheck) _FloginCheck_step03__accept_tokenA(___Vdecode *_Tdecode) {
	var __VunSend _TudpNodeDataSend

	__Vreq := _TloginReq{
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

	_FpfNdb(" 838382 01 start [req:<%s>]", __Vreq.String())
	if nil == ___Vlc.ucCHSendLO {
		_FpfN(" 838382 02 , why output-Chan nil ? ")
	} else {
		_FpfN(" 838382 03 , fake Chan ")
		__Vreq._FdataPack__100__loginReq(Cmd__loginS03acceptWithToken, &__VunSend.usOutBuf)
		__VunSend.usToAddr = _TudpConnPort{
			DstAddr: ___Vdecode.remoteAddr,    // net.UDPAddr
			K256:    ___Vdecode.remotePortKey, // []byte
		}
		_FpfNdb(" 838382 04 [pre-unSend:<%s>]{%s}", __VunSend.String(), __Vreq.String())
		(*___Vlc.ucCHSendLO) <- __VunSend // _TudpNodeDataSend
	}
}
