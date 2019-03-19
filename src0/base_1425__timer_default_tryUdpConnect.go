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
			if 0 != __VnewSession.skipCnt {
				__VnewSession.skipCnt-- // wait only , do nothing.
			} else { // 0 == __VnewSession.skipCnt
				if 0 == __VnewSession.tryCnt { // 0,0 : re-download
					__VnewSession._FudpTimer__750102y__tryGetSrvInfoFromUri()
					if nil == __VnewSession.srvInfo {
						__VnewSession.skipCnt = 6 // srvInfo get error , wait 60s before retry
					} else {
					}
				} else { // xTry,0
				}
			}
		}
		_Fsleep(__Vgap) // mini Gap , at least
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

func (___VnewSession *_TgapNewSession) _FudpTimer__750102y__tryGetSrvInfoFromUri() {
	___VnewSession.srvInfo = nil
}
