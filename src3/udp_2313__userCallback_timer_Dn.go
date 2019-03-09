// _TuExtMRead
// _TserviceUDP
package main

import (
//"net"
//"os"
//"strconv"
//"time"
)

var (
	//	_VuExtTimer_Dn _TuExtTimer = _TuExtTimer{
	//		name:     "timer_Dn",
	//		timgGap1: 25 * time.Second, // every try must more than 2 gap(10*2==20S)
	//		pw2:      _Taes{"password__used_to_connect_to_Fn", []byte(_FnPasswd)},
	//		// hex.DecodeString
	//	}

	_VsrvInfo_Dn        _TsrvInfo
	_VUreqNewSession_Dn _TUreqNewSession = _TUreqNewSession{
		updateUri:    "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/Dn2Fn.json.rand",
		updatePasswd: &_VpasswdDown_Dn,
		srvInfo:      &_VsrvInfo_Dn,
		//UnewSessionCall04: _FuserCallback_u04__,
		//UnewSessionCall01: _FuserCallback_u01__reqNewSession_in_Dn_to_fn,
		UnewSessionCall08: _FuserCallback_u08_packageData_saveTo_tmpBuf,
	}
)

// _Tconfig
type _TreqIneedToLogin struct {
	MeName   string
	MeIdx128 []byte
	MeSeq16  []byte
}

var _VreqIneedToLogin__Dn _TreqIneedToLogin

func _FuserCallback_u08_packageData_saveTo_tmpBuf(___VreqNS *_TUreqNewSession) {

	_VreqIneedToLogin__Dn = _TreqIneedToLogin{
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq16:  _FgenRand_nByte__(16),
	}
	//_FpfN(" 311919 01 : %v ", _VreqIneedToLogin__Dn)

	//__Vb := []byte(" 311919 02 ")
	//___VreqNS.sendBuf08 = &__Vb

	__VbufTmp1 := _FencJsonExit(" 311919 05 ", &_VreqIneedToLogin__Dn)
	__VbufTmp2 := _FencAesRandExit(" 311919 06 ", &(___VreqNS.srvInfo.K256), &__VbufTmp1)
	_FpfhexN(&(___VreqNS.srvInfo.K256), 80, " 311919 07 pack data using key  : ")
	_FpfN(" 311919 08 pack data to (len origin : %d / packed : %d ) ", len(__VbufTmp1), len(__VbufTmp2))
	___VreqNS.sendBuf08 = &__VbufTmp2

} // _FuserCallback_u08_packageData_saveTo_tmpBuf
