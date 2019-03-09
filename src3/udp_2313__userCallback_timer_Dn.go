// _TuExtMRead
// _TserviceUDP
package main

import (
	//"net"
	//"os"
	//"strconv"
	"time"
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
		updateUri:    "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/FnWaitDn.json.rand",
		updatePasswd: &_Vpasswd_udp_FnWaitDn_download_config,
		srvInfo:      &_VsrvInfo_Dn,
		//UnewSessionCall304: _FuserCallback_u04__,
		//UnewSessionCall301: _FuserCallback_u01__reqNewSession_in_Dn_to_fn,
		UnewSessionCall308: _FuserCallback_u08_packageData_saveTo_tmpBuf,
	}
)

// _Tconfig
type _TreqIneedToLogin struct {
	MeName   string
	MeIdx128 []byte
	MeSeq128 []byte
	MeTime   int64
}

var _VreqIneedToLogin__Dn _TreqIneedToLogin

func _FuserCallback_u08_packageData_saveTo_tmpBuf(___VreqNS *_TUreqNewSession) {

	_VreqIneedToLogin__Dn = _TreqIneedToLogin{
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _FgenRand_nByte__(16),
		MeTime:   time.Now().Unix(),
	}
	//_FpfN(" 838981 01 : %v ", _VreqIneedToLogin__Dn)

	//__Vb := []byte(" 838981 02 ")
	//___VreqNS.sendBuf081 = &__Vb

	__VbufTmp1 := _FencJsonExit(" 838981 03 ", &_VreqIneedToLogin__Dn)

	//	__VbufTmp2 := _FencAesRandExit(" 838981 04 ", &(___VreqNS.srvInfo.K256), &__VbufTmp1)
	//	_FpfNhex(&(___VreqNS.srvInfo.K256), 40, " 838981 05 pack data using key  : ")
	//	_FpfN(" 838981 06 pack data to (len origin : %d / packed : %d ) ", len(__VbufTmp1), len(__VbufTmp2))
	//
	//	__VbufTmp3 := _FencAesRandExit(" 838981 07 ", &(___VreqNS.srvInfo.K256), &__VbufTmp1)
	//	//_FpfNhex(&(___VreqNS.srvInfo.K256), 40, " 838981 08 pack data using key  : ")
	//	_FpfN(" 838981 09 pack data to (len origin : %d / packed : %d ) ", len(__VbufTmp1), len(__VbufTmp3))

	__VbufTmp2, __VbufTmp3 := _FencAesRandExit2(" 838981 00 ", &(___VreqNS.srvInfo.K256), &__VbufTmp1)

	___VreqNS.sendBuf081 = &__VbufTmp2
	___VreqNS.sendBuf082 = &__VbufTmp3

} // _FuserCallback_u08_packageData_saveTo_tmpBuf
