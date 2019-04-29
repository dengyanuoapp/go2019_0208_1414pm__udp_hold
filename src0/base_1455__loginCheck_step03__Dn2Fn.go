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

	if false == bytes.Equal(___Vdecode.DElogin.TokenR, ___Vlc.ulTokenA) || // the Dn's id
		false == bytes.Equal(___Vdecode.DElogin.ToIdx128, _VC.MyId128) ||
		false == bytes.Equal(___Vdecode.DElogin.ToSeq128, _VS.MySeq128) {
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
			MeRand5:  _FgenRand_nByte__(5),
			MeTime:   _FtimeInt(),
			ReqStr:   " step03__accept_tokenA ",
			MeName:   _VC.Name,
			MeIdx128: _VC.MyId128,
			MeSeq128: _VS.MySeq128,
			ToIdx128: ___Vdecode.DElogin.MeIdx128, // []byte
			ToSeq128: ___Vdecode.DElogin.MeSeq128, // []byte
			TokenL:   ___Vdecode.DElogin.TokenR,   // []byte
			TokenR:   ___Vdecode.DElogin.TokenL,   // []byte
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
			diConnPort: _TudpConnPort{
				___Vdecode.DEremoteAddr,     // net.UDPAddr
				___Vdecode.DEremotePortKey}, // []byte
			diIdx128: ___Vdecode.DElogin.MeIdx128, // []byte
			diSeq128: ___Vdecode.DElogin.MeSeq128, // []byte
			diToken:  ___Vdecode.DElogin.TokenL,   // []byte
		}
		//_FpfNdb(" 838382 07 [reset-dataMachineID:<%s>]", __Vid.String())
		(*___Vlc.ulCHdataMachineIdLO) <- __Vid
	}
}
