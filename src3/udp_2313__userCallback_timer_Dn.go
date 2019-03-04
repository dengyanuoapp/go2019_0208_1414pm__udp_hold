// _TuExtMRead
// _TserviceUDP
package main

import (
	//"net"
	"os"
	"strconv"
	"time"
)

var (
	_VuExtTimer_Dn _TuExtTimer = _TuExtTimer{
		name:     "timer_Dn",
		timgGap1: 25 * time.Second, // every try must more than 2 gap(10*2==20S)
		pw2:      _Taes{"password__used_to_connect_to_Fn", []byte(_FnPasswd)},
		// hex.DecodeString
	}

	_VsrvInfo_Dn    _TsrvInfo
	_VreqNewSession _TreqNewSession = _TreqNewSession{
		updateUri:    "https://github.com/jasas78/jsonOnly/raw/master/json/Dn2Fn.json.rand",
		updatePasswd: &_VpasswdDown_Dn,
		srvInfo:      &_VsrvInfo_Dn,
	}
)

func _FuserCallback_u03TM__timer_Dn(___Vsvr *_TserviceUDP) {
	_VuExtTimer_Dn.idx++

	if 2 == 3 {
		_FpfN(" 839111 : %d : trying to Connect to Fn using key (%s) \n                 %d : <%0x> ",
			_VuExtTimer_Dn.idx,
			_VuExtTimer_Dn.pw2.name,
			len(_VuExtTimer_Dn.pw2.key),
			_VuExtTimer_Dn.pw2.key)

		// func strconv.ParseUint(s string, base int, bitSize int) (uint64, error)
		//__Vu64 , __Verr := strconv.ParseUint( os.Getenv("id128") , 16, 64 ) // : invalid syntax : when 0x prefix
		__Vu64, __Verr := strconv.ParseUint(os.Getenv("id128"), 0, 0)
		if __Verr == nil {
			_FpfN(" 839113 : var id128 : %x , %s , %x", os.Getenv("id128"), os.Getenv("id128"), __Vu64)
		} else {
			_FpfN(" 839115 : var id128 : %x , %s , convert to u64 error : %v", os.Getenv("id128"), os.Getenv("id128"), __Verr)
		}
	}

	_VreqNewSession._Fconnect_to_server_01__Default()
} // _FuserCallback_u03TM__timer_Dn
