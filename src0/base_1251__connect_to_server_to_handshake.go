package main

import (
	"net"
)

func (___VUreqNewSession *_TUreqNewSession) IRun(___Vidx int) {
	switch ___Vidx {
	case 1:
		if nil == ___VUreqNewSession.UnsCallback01 {
			___VUreqNewSession._Fconnect_to_server_01x__req_new_sessionID__default()
		} else {
			___VUreqNewSession.UnsCallback01(___VUreqNewSession) //_FuserCallback_u03TM__timer_Dn(___VUreqNewSession)
		}
	case 4:
		//_Fex1(" 381991 08 ")
		if nil == ___VUreqNewSession.UnsCallback02 {
			___VUreqNewSession.UnsCallback02(___VUreqNewSession)
		} else {
			___VUreqNewSession._Fconnect_to_server_04x__real_default()
		}
	default:
		_FpfNex(" 381991 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

// you can
func (___VreqNewSession *_TUreqNewSession) _Fconnect_to_server_01x__req_new_sessionID__default() {
	if ___VreqNewSession.Enabled {
		_FpfN(" 311911 01 ")
		_Fsleep_100s()
	} else {
		//_FpfN(" 311911 03 ")
		___VreqNewSession._Fconnect_to_server_01y__req_new_sessionID__main_top()

		_FsleepRand_12_to_14s()
		//_Fsleep_1s() // speed up , to test mem leak.
	}
} // _Fconnect_to_server_01x__req_new_sessionID__default

func (___VreqNewSession *_TUreqNewSession) _Fconnect_to_server_01y__req_new_sessionID__main_top() {
	//_FpfN(" 311912 01 ")
	//_FpfN(" 311912 02 %v", ___VreqNewSession)
	//	//_FpfN(" 311912 03 %v", _Vself)
	//	_FpfN(" 311912 04 %v", _VS)
	//	//_FpfN(" 311912 05 %v", _Vconfig)
	//	_FpfN(" 311912 06 %v", _VC)
	//	_FpfN(" 311912 07 %v", ___VreqNewSession)

	if 0 != ___VreqNewSession.skipCnt {
		___VreqNewSession.skipCnt--
		return // skip then exit
	}
	if 0 == ___VreqNewSession.remainCnt {
		__nowUri := ___VreqNewSession.updateUri
		var __Verr error
		for {
			__nowUri2 := _Pspf("%s.%x", __nowUri, _VC.MyId128)
			_, __Verr = _Ftry_download_rand_json01(__nowUri2, ___VreqNewSession.updatePasswd, ___VreqNewSession.srvInfo)
			if nil != __Verr {
				_FpfN(" 311913 01 : Error : update Uri slice failed.: %s , %v ", __nowUri2, __Verr)
				_, __Verr = _Ftry_download_rand_json01(__nowUri, ___VreqNewSession.updatePasswd, ___VreqNewSession.srvInfo)
			}
			if nil != __Verr {
				_FpfN(" 311913 02 : Error : update Uri slice failed.: %s , %v ", __nowUri, __Verr)
				___VreqNewSession.skipCnt = 4 // skip 4 time , about 40 second
				return
			}
			_FpfN(" 311913 03 : ok : %s , %v ", __nowUri, ___VreqNewSession.srvInfo)
			if ___VreqNewSession.srvInfo.Guri == __nowUri { //_VsrvInfo_Dn     _TsrvInfo
				break
			}

			// or , the new uri need to be used. try tu use it.
			__nowUri = ___VreqNewSession.srvInfo.Guri
		}

		___VreqNewSession.srvLen = len(___VreqNewSession.srvInfo.UriArrs) // try to use the U[:] slice
		_FpfN(" 311913 04 : ==== ==== ==== srvLen %d", ___VreqNewSession.srvLen)
		if 0 == ___VreqNewSession.srvLen {
			_FpfN(" 311913 05 : Error : why Uri slice err ? : %d , %s ", ___VreqNewSession.srvLen, __nowUri)
			___VreqNewSession.skipCnt = 8 // skip 8 time , about 80 second , before recheck.
			return
		}
		___VreqNewSession.remainCnt = ___VreqNewSession.srvLen * 9 // try to use the U[:] slice , loop 9 time , then refresh URI
	}

	if 0 == ___VreqNewSession.remainCnt {
		_FpfNex(" 311913 07 : why reach here ? [%v]", ___VreqNewSession)
	}

	//_FpfN(" 311914 01 : try connect to idx  %d of %d , remain %d , [%v]", ___VreqNewSession.srvIdx, ___VreqNewSession.srvLen,
	//	___VreqNewSession.remainCnt, ___VreqNewSession)

	if 2 == 3 {
		_FpfN(" 311914 02 : try connect to idx  %d of %d , remain %d ,[%v]", ___VreqNewSession.srvIdx, ___VreqNewSession.srvLen,
			___VreqNewSession.remainCnt, ___VreqNewSession.srvInfo.UriArrs[___VreqNewSession.srvIdx])
	}

	_Frun(___VreqNewSession, 4) // _Fconnect_to_server_04x__real_default

	___VreqNewSession.srvIdx++
	if ___VreqNewSession.srvIdx >= ___VreqNewSession.srvLen {
		___VreqNewSession.srvIdx = 0
	}

	___VreqNewSession.remainCnt--

	//_Fex1(" 311914 99 ")
} // _Fconnect_to_server_01y__req_new_sessionID__main_top

func (___VreqNewSession *_TUreqNewSession) _Fconnect_to_server_04x__real_default() {
	__VdstUaddrStr := ___VreqNewSession.srvInfo.UriArrs[___VreqNewSession.srvIdx]
	_FpfN(" 311917 01 : try connect to idx  %d of %d , remain %d ,[%v]", ___VreqNewSession.srvIdx, ___VreqNewSession.srvLen,
		___VreqNewSession.remainCnt, __VdstUaddrStr)

	___Vsvr := ___VreqNewSession.serviceUdP
	__VudpConn := ___Vsvr.udpConn
	//_FpfN(" 311917 02 : %v " , ___Vsvr)
	_FpfN(" 311917 03 : %v ", __VudpConn)

	// func ResolveUDPAddr(network, address string) (*UDPAddr, error)
	__VuAddr, __Verr1 := net.ResolveUDPAddr("udp4", __VdstUaddrStr)
	if __Verr1 != nil {
		_FpfN("311917 04 : udpDst mistake<%s>[%v]", __VdstUaddrStr, __Verr1)
		return
	}

	// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
	_, __Verr2 := __VudpConn.WriteToUDP([]byte(" 311917 05 "), __VuAddr)
	if __Verr2 != nil {
		_FpfN("311917 06 : udp send error <%s>[%v]", __VdstUaddrStr, __Verr2)
		return
	}
	_FpfN("311917 07 : udp send succeed 01 dst<%s>, local<%v>, listen<%v>", __VdstUaddrStr, ___Vsvr.VulocalAddr, ___Vsvr.udpAddr)

	_FsleepRand_12_to_14s()
} // _Fconnect_to_server_04__real_default
