package main

type _TreqNewSession struct {
	Enabled bool

	skipCnt   int
	remainCnt int

	updateUri    string
	updatePasswd *[]byte

	srvLen  int
	srvIdx  int
	srvInfo *_TsrvInfo
} //    _TreqNewSession

// you can
func (___VreqNewSession *_TreqNewSession) _Fconnect_to_server_01__Default() {
	if ___VreqNewSession.Enabled {
		_FpfN(" 311911 01 ")
		_Fsleep_100s()
	} else {
		//_FpfN(" 311911 03 ")
		___VreqNewSession._Fconnect_to_server_01__req_new_sessionID__main_top()
		_FsleepRand_12_to_14s()
	}
} // _Fconnect_to_server_01__Default

func (___VreqNewSession *_TreqNewSession) _Fconnect_to_server_01__req_new_sessionID__main_top() {
	_FpfN(" 311912 01 ")
	_FpfN(" 311912 02 %v", ___VreqNewSession)
	//_FpfN(" 311912 03 %v", _Vself)
	_FpfN(" 311912 04 %v", _VS)
	//_FpfN(" 311912 05 %v", _Vconfig)
	_FpfN(" 311912 06 %v", _VC)
	_FpfN(" 311912 07 %v", ___VreqNewSession)

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

		___VreqNewSession.srvLen = len(___VreqNewSession.srvInfo.UriDn2Fn) // try to use the U[:] slice
		_FpfN(" 311913 04 : ==== ==== ==== srvLen %d", ___VreqNewSession.srvLen)
		if 0 == ___VreqNewSession.srvLen {
			_FpfN(" 311913 05 : Error : why Uri slice err ? : %d , %s ", ___VreqNewSession.srvLen, __nowUri)
			___VreqNewSession.skipCnt = 8 // skip 8 time , about 80 second , before recheck.
			return
		}
		___VreqNewSession.remainCnt = ___VreqNewSession.srvLen * 9 // try to use the U[:] slice , loop 9 time , then refresh URI

	}

	if 0 == ___VreqNewSession.remainCnt {
		_FpfN(" 311913 07 : why reach here ? [%v]", ___VreqNewSession)
	}

	_FpfN(" 311914 01 : try connect to idx  %d of %d , remain %d , [%v]", ___VreqNewSession.srvIdx, ___VreqNewSession.srvLen,
		___VreqNewSession.remainCnt, ___VreqNewSession)

	_Fex1(" 311914 99 ")
} // _Fconnect_to_server_01__req_new_sessionID__main_top
