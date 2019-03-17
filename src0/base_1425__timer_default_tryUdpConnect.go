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
			if 0 == __VnewSession.tryCnt {
				if 0 == __VnewSession.skipCnt {
					__VsrvInfo := __VnewSession._FudpTimer__750102y__tryGetSrvInfoFromUri()
					if nil == __VsrvInfo {
						// do nothing , except set timeoutto nex 30s
						__VnewSession.skipCnt = 3
					} else {
					}
				} else {
				}
			} else { // 0 != __VnewSession.tryCnt
				if 0 == __VnewSession.skipCnt {
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

func (___Vsrv *_TsrvInfo) _FudpTimer__750102y__tryGetSrvInfoFromUri() *_TsrvInfo {
	return true
}
