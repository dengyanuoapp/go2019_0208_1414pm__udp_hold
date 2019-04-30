package main

import (
	"bytes"
	"reflect"
)

func (___Vlc *_TloginCheck) _FloginCheck_step900201y__s3accept_tokenA_fill03send_Dn(___Vdecode *_Tdecode, ___Venc *_Tencode) {
	if _FcheckDecodeType(___Vdecode, LoadT__loginS02genReplyTokenB) {
		_FpfNdb(" 838381 01 type error , ignore ")
		return
	}

	__VlenArr := ___Vdecode.Count128()
	if false == reflect.DeepEqual(__VlenArr, __Vstep03_LoginLenArr) {
		_FpfNdb(" 838381 03 len error , ignore %d ", __VlenArr)
		return
	}

	if false == bytes.Equal(___Vdecode.DElogin.LgTokenR, ___Vlc.ulTokenA) || // the Dn's id
		false == bytes.Equal(___Vdecode.DElogin.LgToIdx128, _VC.MyId128) ||
		false == bytes.Equal(___Vdecode.DElogin.LgToSeq128, _VS.MySeq128) {
		_FpfN(" 838381 05 : error : no equal, ignore. %s : ulTokenA:%s MyId128:%s MySeq128:%s ",
			___Vdecode.String(), String5s(&___Vlc.ulTokenA), String5s(&_VC.MyId128), String5s(&_VS.MySeq128))
		return
	}

	if (_FtimeInt() - ___Vdecode.DEreceiveTime) > __VmaxCmdPerid {
		_FpfN(" 838381 06 : error : timeOut. %s ", ___Vdecode.String())
		return
	}

	if (_FtimeInt() - ___Vlc.ulGenTime) > __VmaxCmdPerid {
		_FpfN(" 838381 07 : error : timeOut. %s ", ___Vdecode.String())
		return
	}

	//_FpfNex(" 838381 08 %s ", ___Vdecode.String())
	*___Venc = _Tencode{
		EnToConnPort: _TudpConnPort{
			DstAddr: ___Vdecode.DEremoteAddr,    // net.UDPAddr
			K256:    ___Vdecode.DEremotePortKey, // []byte
		},
		EnLoadType: LoadT__loginS03acceptWithToken,
		EnLogin: _TloginReq{
			LgMeRand5:  _FgenRand_nByte__(5),
			LgMeTime:   _FtimeInt(),
			LgReqStr:   " step03__accept_tokenA ",
			LgMeName:   _VC.Name,
			LgMeIdx128: _VC.MyId128,
			LgMeSeq128: _VS.MySeq128,
			LgToIdx128: ___Vdecode.DElogin.LgMeIdx128, // []byte
			LgToSeq128: ___Vdecode.DElogin.LgMeSeq128, // []byte
			LgTokenL:   ___Vdecode.DElogin.LgTokenR,   // []byte
			LgTokenR:   ___Vdecode.DElogin.LgTokenL,   // []byte
		},
	}

	___Vlc._FloginCheck_step03__accept_tokenA_Dn(___Vdecode)
}

// _TloginReq _Tdecode
func (___Vlc *_TloginCheck) _FloginCheck_step03__accept_tokenA_Dn(___Vdecode *_Tdecode) {
	if nil == ___Vlc.ulCHdataMachineIdLO {
		_FpfN(" 838382 06 , why output-Chan nil ? ulCHdataMachineIdLO")
	} else {
		__Vid := _TdataMachinEid{
			diRole: ___Vdecode.DErole, // _TdecodeX
			diConnPort: _TudpConnPort{
				___Vdecode.DEremoteAddr,     // net.UDPAddr
				___Vdecode.DEremotePortKey}, // []byte
			diIdx128: ___Vdecode.DElogin.LgMeIdx128, // []byte
			diSeq128: ___Vdecode.DElogin.LgMeSeq128, // []byte
			diToken:  ___Vdecode.DElogin.LgTokenL,   // []byte
		}
		//_FpfNdb(" 838382 07 [reset-dataMachineID:<%s>]", __Vid.String())
		(*___Vlc.ulCHdataMachineIdLO) <- __Vid
	}
}
