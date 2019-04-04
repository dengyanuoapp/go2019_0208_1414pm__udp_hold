package main

import (
	"bytes"
	"reflect"
)

func (___Vlc *_TloginCheck) _FloginCheck_step900201y__s4accept_tokenB_resetData_Fn(___Vdecode *_Tdecode) {
	if _FcheckDecodeType(___Vdecode, Cmd__loginS03acceptWithToken) {
		_FpfNdb(" 838383 01 type error , ignore ")
		return
	}

	__VlenArr := ___Vdecode.Count128()
	if false == reflect.DeepEqual(__VlenArr, __Vstep03_LoginLenArr) {
		_FpfNdb(" 838383 03 len error , ignore %d ", __VlenArr)
		return
	}

	if //false == bytes.Equal(___Vdecode.Dlogin.TokenR, ___Vlc.ulTokenA) || // the Dn's id
	false == bytes.Equal(___Vdecode.Dlogin.ToIdx128, _VC.MyId128) ||
		false == bytes.Equal(___Vdecode.Dlogin.ToSeq128, _VS.MySeq128) {
		_FpfN(" 838383 05 : error : no equal, ignore. {tk:%x,id:%x,seq:%x}{%s} ",
			___Vlc.ulTokenA,
			_VC.MyId128[:5],
			_VS.MySeq128[:5],
			___Vdecode.String())
		return
	}

	if (_FtimeInt() - ___Vdecode.receiveTime) > __VmaxCmdPerid {
		_FpfN(" 838383 06 : error : timeOut. %s ", ___Vdecode.String())
		return
	}

	__Vk128 := _FgenB16(&___Vdecode.Dlogin.MeIdx128) // Dn's id
	___Vlc.ulCmd.mux.Lock()                          //               _TcmdMap
	__Vold, __Vok4 := ___Vlc.ulCmd.M[__Vk128]        //               _TcmdMap
	___Vlc.ulCmd.mux.Unlock()                        //               _TcmdMap

	if false == __Vok4 {
		_FpfN(" 838383 07 : error : not found(key:%x), %s ", String5(&___Vdecode.Dlogin.MeIdx128), ___Vdecode.String())
		//_FpfN(" 838383 08 : %#v", ___Vlc.ulCmd.M)
		//_Fex1(" 838383 09 ")
		return
	}
	if (_FtimeInt() - __Vold.receiveTime) > __VmaxCmdPerid {
		_FpfN(" 838383 11 : error : timeOut. %s ", __Vold.String())
		_FpfN(" 838383 12 : error : timeOut. %s ", ___Vdecode.String())
		return
	}

	if false == bytes.Equal(___Vdecode.Dlogin.MeIdx128, __Vold.Dlogin.MeIdx128) ||
		false == bytes.Equal(___Vdecode.Dlogin.MeSeq128, __Vold.Dlogin.MeSeq128) ||
		false == bytes.Equal(___Vdecode.Dlogin.ToIdx128, __Vold.Dlogin.ToIdx128) ||
		false == bytes.Equal(___Vdecode.Dlogin.ToSeq128, __Vold.Dlogin.ToSeq128) ||
		false == bytes.Equal(___Vdecode.Dlogin.TokenL, __Vold.Dlogin.TokenL) ||
		false == bytes.Equal(___Vdecode.Dlogin.TokenR, __Vold.Dlogin.TokenR) {
		_FpfN("   838383 18 %s ", __Vold.String())
		_FpfNex(" 838383 19 %s ", ___Vdecode.String())
	}

	_FpfN("   838383 20 %s ", __Vold.String())
	_FpfNex(" 838383 21 %s ", ___Vdecode.String())

	___Vlc._FloginCheck_step04__accept_tokenB_Fn(___Vdecode)
}

// _TloginReq _Tdecode
func (___Vlc *_TloginCheck) _FloginCheck_step04__accept_tokenB_Fn(___Vdecode *_Tdecode) {

	if nil == ___Vlc.ulCHdataMachineIdLO {
		_FpfN(" 838384 06 , why output-Chan nil ? ulCHdataMachineIdLO")
	} else {
		__Vid := _TdataMachinEid{
			diConnPort: _TudpConnPort{
				___Vdecode.remoteAddr,     // net.UDPAddr
				___Vdecode.remotePortKey}, // []byte
			diIdx128: ___Vdecode.Dlogin.MeIdx128, // []byte
			diSeq128: ___Vdecode.Dlogin.MeSeq128, // []byte
			diToken:  ___Vdecode.Dlogin.TokenL,   // []byte
		}
		_FpfNdb(" 838384 07 [insertId:<%s>]", __Vid.String())
		(*___Vlc.ulCHdataMachineIdLO) <- __Vid
	}
}