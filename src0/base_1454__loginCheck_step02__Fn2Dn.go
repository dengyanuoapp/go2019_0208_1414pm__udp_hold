package main

import "reflect"

func (___Vlc *_TloginCheck) _FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn(___Vdecode *_Tdecode) {
	if _FcheckDecodeType(___Vdecode, Cmd__loginS01genReplyTokenA) {
		_FpfNdb(" 838393 01 type error , ignore ")
		return
	}

	////_FpfNdb(" 838393 02 : ...... ")
	//_FpfNdb(" 838393 04 : %s", ___Vdecode.String()) // 15540463611554046361
	////___Vlc.ulCmd.M        map[[16]byte]_TconnInfo // _TloginReq MeIdx128

	__VlenArr := ___Vdecode.Count128()
	if false == reflect.DeepEqual(__VlenArr, __Vstep02_LoginLenArr) {
		_FpfN(" 838393 05 _TloginCheck :%s", ___Vlc.String())
		_FpfN(" 838393 06 _Tdecode :%s", ___Vdecode.String())
		_FpfNdb(" 838393 07 len error , ignore %d ", __VlenArr)
		return
	}

	__Vk := _FgenB16(&___Vdecode.Dlogin.MeIdx128)

	if len(___Vlc.ulCmd.M) > 300 {
		_FdeleteOld_cmdStack(&___Vlc.ulCmd.M)
	}

	//_FpfNdb(" 838393 08 : key is <%x> ", __Vk)
	___Vdecode.Dlogin.TokenR = _FgenRand_nByte__(16)

	___Vlc.ulCmd.M[__Vk] = *___Vdecode

	//_FpfNdb(" 838393 09 : [decode:<%s>]", ___Vdecode.String())

	___Vlc._FloginCheck_step102y__sReply_tokenB(___Vdecode)
}

// _TloginReq _Tdecode
func (___Vlc *_TloginCheck) _FloginCheck_step102y__sReply_tokenB(___Vdecode *_Tdecode) {
	var __VunSend _TudpNodeDataSend

	__Vreq := _TloginReq{
		MeTime:   _FtimeInt(),                 // int64
		ReqStr:   " step102y__sReply_tokenB ", // string
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _VS.MySeq128,
		ToIdx128: ___Vdecode.Dlogin.MeIdx128, // []byte
		ToSeq128: ___Vdecode.Dlogin.MeSeq128, // []byte
		TokenL:   ___Vdecode.Dlogin.TokenR,   // []byte
		TokenR:   ___Vdecode.Dlogin.TokenL,   // []byte
	}

	//_FpfNdb(" 838394 01 start [req:<%s>]", __Vreq.String())
	if nil == ___Vlc.ulCHSendLO {
		_FpfN(" 838394 02 , why output-Chan nil ? ")
	} else {
		// _FpfN(" 838394 03 , fake Chan {%s}", __Vreq.String())
		__Vreq._FdataPack__100__loginReq(Cmd__loginS02genReplyTokenB, &__VunSend.usOutBuf)
		__VunSend.usToAddr = _TudpConnPort{
			DstAddr: ___Vdecode.remoteAddr,    // net.UDPAddr
			K256:    ___Vdecode.remotePortKey, // []byte
		}
		__VrandCnt := 4 + (_FgenRand_int() % 7) // gen 0 -- 6 , + 4 : so total 5 - 10
		//_FpfNdb(" 838394 04 (%d)[pre-unSend:<%s>]{%s}", __VrandCnt, __VunSend.String(), __Vreq.String())
		_CpfN(" 838394 05 (%d)[pre-unSend:<%s>]{%s}", __VrandCnt, __VunSend.String(), __Vreq.String())
		for __VrandCnt > 0 {
			__VrandCnt--
			(*___Vlc.ulCHSendLO) <- __VunSend // _TudpNodeDataSend
		}
	}
}
