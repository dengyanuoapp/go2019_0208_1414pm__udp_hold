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

	_VsrvInfo_Dn    _TsrvInfo
	_VreqNewSession _TreqNewSession = _TreqNewSession{
		updateUri:    "https://github.com/jasas78/jsonOnly/raw/master/json/Dn2Fn.json.rand",
		updatePasswd: &_VpasswdDown_Dn,
		srvInfo:      &_VsrvInfo_Dn,
		UcallbackNS:  _FuserCallback_u03TM__connect_Dn2Fn,
	}
)

func _FuserCallback_u03TM__timer_Dn(___Vsvr *_TserviceUDP) {
	//_VuExtTimer_Dn.idx++
	_VreqNewSession._Fconnect_to_server_01__req_new_sessionID__default()
} // _FuserCallback_u03TM__timer_Dn

// being call in _Fconnect_to_server_03__real
func _FuserCallback_u03TM__connect_Dn2Fn(___VreqNewSession *_TreqNewSession) bool {
	//_FpfN(" 311916 01 ")
	return false // false -> call default func ; true --> already deal with , no call to default func call
} // _FuserCallback_u03TM__connect_Dn2Fn
