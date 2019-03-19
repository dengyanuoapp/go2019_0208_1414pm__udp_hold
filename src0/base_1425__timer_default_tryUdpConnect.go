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
		if true == __VnewSession.connected {
			__VnewSession.tryCnt = 0
			__VnewSession.skipCnt = 6 // set timeout to 60s , not try re-connect
		} else { // true != __VnewSession.connected
			// tryCnt , skipCnt
			if 0 == __VnewSession.skipCnt {
				if 0 == __VnewSession.tryCnt { // 0,0 : re-download
					__VsrvInfo := __VnewSession._FudpTimer__750102y__tryGetSrvInfoFromUri()
					if nil == __VsrvInfo {
						// get srvInfo error, set timeout to 80s
						__VnewSession.skipCnt = 3
					} else {
						var __Vreq _TreqIneedToLogin
						___Vgtm.uTmReqIneedToLogin <- __Vreq
						_Fsleep(__Vgap)
						___Vgtm.uTmReqIneedToLogin <- __Vreq
						// add another gap before retry
						__VnewSession.skipCnt = 2
					}
				} else { // xTry,0
					__VnewSession.skipCnt--
				}
			} else { // 0 == __VnewSession.skipCnt
				if 0 != __VnewSession.tryCnt {
				} else {
				}
			}
		}
		//		__Vsrv3 := __VnewSession._Fanalyze_srvInfo__whether_trying_to_connect()
		//		if nil != __Vsrv3 {
		//			//___Vgtm.uTmGapNewSession2
		//		}
		_Fsleep(__Vgap) // mini Gap , at least
	}
}

func (___VnewSession *_TgapNewSession) _FudpTimer__750102y__tryGetSrvInfoFromUri() *_TsrvInfo {
	return nil
}
