package main

import (
	"bytes"
	"reflect"
)

func (___Vlc *_TloginCheck) _FloginCheck_step900201y__s3accept_tokenA_fill03send_Dn(___Vdecode *_Tdecode) {
	if _FcheckDecodeType(___Vdecode, Cmd__loginS02genReplyTokenB) {
		_FpfNdb(" 838381 01 type error , ignore ")
		return
	}

	__VlenArr := ___Vdecode.Count128()
	if false == reflect.DeepEqual(__VlenArr, __Vstep03_LoginLenArr) {
		_FpfNdb(" 838381 03 len error , ignore %d ", __VlenArr)
		return
	}

	if false == bytes.Equal(___Vdecode.Dlogin.TokenR, ___Vlc.ulTokenA) || // the Dn's id
		false == bytes.Equal(___Vdecode.Dlogin.ToIdx128, _VC.MyId128) ||
		false == bytes.Equal(___Vdecode.Dlogin.ToSeq128, _VS.MySeq128) {
		_FpfN(" 838381 05 : error : no equal, ignore. %s ", ___Vdecode.String())
		return
	}

	if (_FtimeInt() - ___Vdecode.receiveTime) > __VmaxCmdPerid {
		_FpfN(" 838381 06 : error : timeOut. %s ", ___Vdecode.String())
		return
	}

	if (_FtimeInt() - ___Vlc.ulGenTime) > __VmaxCmdPerid {
		_FpfN(" 838381 07 : error : timeOut. %s ", ___Vdecode.String())
		return
	}

	//_FpfNex(" 838381 08 %s ", ___Vdecode.String())

	___Vlc._FloginCheck_step03__accept_tokenA_Dn(___Vdecode)
}

// _TloginReq _Tdecode
func (___Vlc *_TloginCheck) _FloginCheck_step03__accept_tokenA_Dn(___Vdecode *_Tdecode) {
	var __VunSend _TudpNodeDataSend

	__Vreq := _TloginReq{
		MeTime:   _FtimeInt(),               // int64
		ReqStr:   " step03__accept_tokenA ", // string
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _VS.MySeq128,
		ToIdx128: ___Vdecode.Dlogin.MeIdx128, // []byte
		ToSeq128: ___Vdecode.Dlogin.MeSeq128, // []byte
		TokenL:   ___Vdecode.Dlogin.TokenR,   // []byte
		TokenR:   ___Vdecode.Dlogin.TokenL,   // []byte
	}

	//_FpfN("   838382 10 %s ", ___Vdecode.String())
	//_FpfNex(" 838382 11 %s ", __Vreq.String())

	//_FpfNdb(" 838382 01 start [req:<%s>]", __Vreq.String())
	if nil == ___Vlc.ulCHSendLO {
		_FpfN(" 838382 02 , why output-Chan nil ? ulCHSendLO")
	} else {
		//_FpfN(" 838382 03 , fake Chan ")
		__Vreq._FdataPack__100__loginReq(Cmd__loginS03acceptWithToken, &__VunSend.usOutBuf)
		__VunSend.usToAddr = _TudpConnPort{
			DstAddr: ___Vdecode.remoteAddr,    // net.UDPAddr
			K256:    ___Vdecode.remotePortKey, // []byte
		}
		_FpfNdb(" 838382 04 [pre-unSend:<%s>]{%s}", __VunSend.String(), __Vreq.String())
		(*___Vlc.ulCHSendLO) <- __VunSend // _TudpNodeDataSend
	}

	if nil == ___Vlc.ulCHdataMachineIdLO {
		_FpfN(" 838382 06 , why output-Chan nil ? ulCHdataMachineIdLO")
	} else {
		__Vid := _TdataMachinEid{
			diConnPort: _TudpConnPort{
				___Vdecode.remoteAddr,     // net.UDPAddr
				___Vdecode.remotePortKey}, // []byte
			diIdx128: ___Vdecode.Dlogin.MeIdx128, // []byte
			diSeq128: ___Vdecode.Dlogin.MeSeq128, // []byte
			diToken:  ___Vdecode.Dlogin.TokenL,   // []byte
		}
		_FpfNdb(" 838382 07 [insertId:<%s>]", __Vid.String())
		(*___Vlc.ulCHdataMachineIdLO) <- __Vid
	}
}
