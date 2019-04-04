package main

import (
	"bytes"
	"reflect"
)

func (___Vlc *_TloginCheck) _FloginCheck_step900201y__s3accept_tokenA_file03send_Dn(___Vdecode *_Tdecode) {
	if _FcheckDecodeType(___Vdecode, Cmd__loginS02genReplyTokenB) {
		_FpfNdb(" 838381 01 type error , ignore ")
		return
	}

	////_FpfNdb(" 838381 01 : ...... ")
	//_FpfNdb(" 838381 02 : %s", ___Vdecode.String()) // 15540463611554046361
	////___Vlc.ulCmd.M        map[[16]byte]_TconnInfo // _TloginReq MeIdx128

	__VlenArr := ___Vdecode.Count128()
	if false == reflect.DeepEqual(__VlenArr, __Vstep03_LoginLenArr) {
		_FpfNdb(" 838381 03 len error , ignore %d ", __VlenArr)
		return
	}

	//__Vk16 := _FgenB16(&___Vdecode.Dlogin.MeIdx128)
	//__Vdc2, __Vok2 := ___Vlc.ulCmd.M[__Vk16] // _Tdecode , bool
	//	if false == __Vok2 {
	//		_FpfN(" 838381 04 : error : no recorde<%x>, ignore. %s ", __Vk16[:8], ___Vdecode.String())
	//		return
	//	}

	if false == bytes.Equal(___Vdecode.Dlogin.TokenR, ___Vlc.ulTokenA) || // the Dn's id
		false == bytes.Equal(___Vdecode.Dlogin.ToIdx128, _VC.MyId128) ||
		false == bytes.Equal(___Vdecode.Dlogin.ToSeq128, _VS.MySeq128) {
		_FpfN(" 838381 05 : error : no equal, ignore. %s ", ___Vdecode.String())
		return
	}

	//_FpfNex(" 838381 07 %s ", ___Vdecode.String())

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
		MeSeq128: _VS.MySeq128,
		ToIdx128: ___Vdecode.Dlogin.MeIdx128, // []byte
		ToSeq128: ___Vdecode.Dlogin.MeSeq128, // []byte
		TokenL:   ___Vdecode.Dlogin.TokenR,   // []byte
		TokenR:   ___Vdecode.Dlogin.TokenL,   // []byte
	}

	_FpfN(" 838382 08 %s ", ___Vdecode.String())
	_FpfNex(" 838382 09 %s ", __Vreq.String())

	_FpfNdb(" 838382 01 start [req:<%s>]", __Vreq.String())
	if nil == ___Vlc.ulCHSendLO {
		_FpfN(" 838382 02 , why output-Chan nil ? ")
	} else {
		_FpfN(" 838382 03 , fake Chan ")
		__Vreq._FdataPack__100__loginReq(Cmd__loginS03acceptWithToken, &__VunSend.usOutBuf)
		__VunSend.usToAddr = _TudpConnPort{
			DstAddr: ___Vdecode.remoteAddr,    // net.UDPAddr
			K256:    ___Vdecode.remotePortKey, // []byte
		}
		_FpfNdb(" 838382 04 [pre-unSend:<%s>]{%s}", __VunSend.String(), __Vreq.String())
		(*___Vlc.ulCHSendLO) <- __VunSend // _TudpNodeDataSend
	}
}
