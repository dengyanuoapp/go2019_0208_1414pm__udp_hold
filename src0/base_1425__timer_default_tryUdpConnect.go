package main

// _TreqIneedToLogin
// _TsrvInfo
// _TgapTimer
func _FudpTimer__150102x__init__tryUdpConn__default(___Vgtm *_TgapTimer) {
	__Vns := ___Vgtm.uTmGapNewSession2
	_Fdebug1(" 138181 01 %#v ", ___Vgtm.uTmGapNewSession2)

	__Vgap := ___Vgtm.uTmGapX
	if 0 == __Vgap {
		_Fdebug1(" 138181 02 %#v ", ___Vgtm.uTmGapNewSession2)
		_FpfN(" 138181 03 : sorry , gap is ZERO , skip gap loop. ")
		return
	}

	for {
		if true == __Vns.connected {
			__Vns.tryCnt = 0
			__Vns.skipCnt = 3 // if within 3 * gap , not try re-connect
			_Fsleep(__Vgap)
		} else {
			if 0 == __Vns.tryCnt {
				if 0 == __Vns.skipCnt {
				} else {
				}
			} else {
			}
		}
		//		__Vsrv3 := __Vns._Fanalyze_srvInfo__whether_trying_to_connect()
		//		if nil != __Vsrv3 {
		//			//___Vgtm.uTmGapNewSession2
		//		}
		_Fsleep(__Vgap)
	}
}

func (___Vsrv *_TsrvInfo) _Fanalyze_srvInfo__whether_trying_to_connect() *_TsrvInfo {
	return nil
}
