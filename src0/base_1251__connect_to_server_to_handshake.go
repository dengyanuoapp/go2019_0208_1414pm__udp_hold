package main

func (___VUreqNewSession *_TUreqNewSession) IRun() {
	if nil == ___VUreqNewSession.UcallbackTM {
		___VUreqNewSession._Fconnect_to_server_01__req_new_sessionID__default()
	} else {
		___VUreqNewSession.UcallbackTM(___VUreqNewSession)
		//_FuserCallback_u03TM__timer_Dn(___VUreqNewSession)
	}
}

// you can
func (___VreqNewSession *_TUreqNewSession) _Fconnect_to_server_01__req_new_sessionID__default() {
	if ___VreqNewSession.Enabled {
		_FpfN(" 311911 01 ")
		_Fsleep_100s()
	} else {
		//_FpfN(" 311911 03 ")
		___VreqNewSession._Fconnect_to_server_02__req_new_sessionID__main_top()

		_FsleepRand_12_to_14s()
		//_Fsleep_1s() // speed up , to test mem leak.
	}
} // _Fconnect_to_server_01__req_new_sessionID__default

func (___VreqNewSession *_TUreqNewSession) _Fconnect_to_server_02__req_new_sessionID__main_top() {
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

	___VreqNewSession._Fconnect_to_server_03__real()
	//_Frun( ___VreqNewSession )

	___VreqNewSession.srvIdx++
	if ___VreqNewSession.srvIdx >= ___VreqNewSession.srvLen {
		___VreqNewSession.srvIdx = 0
	}

	___VreqNewSession.remainCnt--

	//_Fex1(" 311914 99 ")
} // _Fconnect_to_server_02__req_new_sessionID__main_top
func (___VreqNewSession *_TUreqNewSession) _Fconnect_to_server_03__real() {
	//_FpfN(" 311915 01 ")
	//_Ftry_req
	//UriArrs []string // try-Uris
	//K256    []byte   // passwd to connect the this server
	if ___VreqNewSession.UcallbackNS(___VreqNewSession) { // _FuserCallback_u03TM__connect_Dn2Fn
		return
	}
	___VreqNewSession._Fconnect_to_server_04__real_default()
} // _Fconnect_to_server_03__real

//func (___VreqNewSession *_TUreqNewSession) _Fconnect_to_server_04__real_default() {
func (___VreqNewSession *_TUreqNewSession) _Fconnect_to_server_04__real_default() {
	_FpfN(" 311917 01 : try connect to idx  %d of %d , remain %d ,[%v]", ___VreqNewSession.srvIdx, ___VreqNewSession.srvLen,
		___VreqNewSession.remainCnt, ___VreqNewSession.srvInfo.UriArrs[___VreqNewSession.srvIdx])
	//_FpfN(" 311917 02 : %v " , ___VreqNewSession
} // _Fconnect_to_server_04__real_default
