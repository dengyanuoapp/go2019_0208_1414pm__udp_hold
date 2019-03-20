package main

// _TreqIneedToLogin
// _TsrvInfo
// _TgapTimer
//uTmGapNewSession2 *_TgapNewSession
func _FudpTimer__750102x__init__tryUdpConn__default(___Vgtm *_TgapTimer) {
	__VnewSession := ___Vgtm.uTmGapNewSession2
	_Fdebug1(" 138181 01 %#v ", ___Vgtm.uTmGapNewSession2)

	__Vgap := ___Vgtm.uTmGapX
	if 0 == __Vgap {
		_Fdebug1(" 138181 02 %#v ", ___Vgtm.uTmGapNewSession2)
		_FpfN(" 138181 03 : sorry , gap is ZERO , skip gap loop. ")
		return
	}

	for {
		_Fsleep(__Vgap) // mini Gap , at least
		if true == __VnewSession.connected {
			__VnewSession.tryCnt = 0
			__VnewSession.skipCnt = 6 // set timeout to 60s , not try re-connect
		} else { // true != __VnewSession.connected
			// tryCnt , skipCnt
			if 0 != __VnewSession.skipCnt {
				__VnewSession.skipCnt-- // wait only , do nothing.
			} else { // 0 == __VnewSession.skipCnt
				if 0 == __VnewSession.tryCnt { // 0,0 : re-download
					__VnewSession._FudpTimer__750102y__tryGetSrvInfoFromUri()
					if false == __VnewSession.srvInfo.ok {
						__VnewSession.skipCnt = 6 // srvInfo get error , wait 60s before retry
					} else {
						__VnewSession.tryCnt = 20 * len(__VnewSession.srvInfo.UriArrs)
						__VnewSession.skipCnt = 2
						// try 20 times of UriArrs to connect
					}
				} else { // xTry,0skip
					__VnewSession._FudpTimer__750102z__tryfillSendChan()
					__VnewSession.tryCnt--
					__VnewSession.skipCnt = 2
				}
			}
		}
	}
}

func (___VnewSession *_TgapNewSession) _FudpTimer__750102z__tryfillSendChan() {
	//						var __Vreq _TreqIneedToLogin
	//						___Vgtm.uTmReqIneedToLogin <- __Vreq
	//						_Fsleep(__Vgap)
	//						___Vgtm.uTmReqIneedToLogin <- __Vreq
	//						// add another gap before retry
	//						__VnewSession.skipCnt = 2
}

// _TgapNewSession _TsrvInfo
func (___VnewSession *_TgapNewSession) _FudpTimer__750102y__tryGetSrvInfoFromUri() {
	_Fdebug1("238191 01 : %v", ___VnewSession)

	___VnewSession.srvInfo.ok = false

	__Vsi2 := _TsrvInfo{
		name:       ___VnewSession.name,
		refreshUri: ___VnewSession.updateUri,
		refreshPwd: ___VnewSession.updatePasswd,
	}

	var (
		__VtmpSi2 _TsrvInfo
		__Verr2   error
		__nowUri3 string
	)
	//for __VtmpSi2.refreshUri != __Vsi2.refreshUri {
	// _Fconnect_to_server_01y__req_new_sessionID__main_top
	for {
		_Fsleep(_T1s) // mini Gap , at least
		__nowUri3 = _Pspf("%s.%x", __Vsi2.refreshUri, _VC.MyId128)

		_, __Verr2 = _Ftry_download_rand_json01(__nowUri3, &__Vsi2.refreshPwd, &__VtmpSi2)
		if nil != __Verr2 {
			_FpfN(" 311913 04 : Error : update Uri slice failed.: %s , %v ", __nowUri3, __Verr2)
			__nowUri3 = __Vsi2.refreshUri
			_, __Verr2 = _Ftry_download_rand_json01(__nowUri3, &__Vsi2.refreshPwd, &__VtmpSi2)
			if nil != __Verr2 {
				_FpfN(" 311913 05 : Error : update Uri slice failed.: %s , %v ", __nowUri3, __Verr2)
				return
			} else {
				_FpfN(" 311913 06 : ok : %s , %v ", __nowUri3, __VtmpSi2)
			}
		} else {
			_FpfN(" 311913 07 : ok : %s , %v ", __nowUri3, __VtmpSi2)
		}

		if nil == __VtmpSi2.refreshPwd {
			__VtmpSi2.refreshPwd = ___VnewSession.updatePasswd // use default passwd if nil
		}

		___VnewSession.srvInfo = __VtmpSi2

		_FpfN(" 311913 08 : %s , %s", ___VnewSession.srvInfo.refreshUri, __nowUri3)

		if ___VnewSession.srvInfo.refreshUri == __nowUri3 { //_VsrvInfo_Dn     _TsrvInfo
			break
		}

		_FpfN(" 311913 08 : download loop ing")

	}

	_FpfN(" 311913 09 : all ok : %s , %v ", __nowUri3, __VtmpSi2)

	___VnewSession.srvInfo.ok = false
}
