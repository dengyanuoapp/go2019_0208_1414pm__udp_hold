package main

// _TreqIneedToLogin
// _TsrvInfo
// _TgapTimer
func _FudpTimer__150102x__init__tryUdpConn__default(___Vgtm *_TgapTimer) {
	__Vsrv2 := ___Vgtm.uTmSetSrvInfo02
	_Fdebug1(" 138181 01 %#v ", ___Vgtm.uTmSetSrvInfo02)

	__Vgap := ___Vgtm.uTmGap02
	if 0 == __Vgap {
		_Fdebug1(" 138181 02 %#v ", ___Vgtm.uTmSetSrvInfo02)
		_FpfN(" 138181 03 : sorry , gap is ZERO , skip gap loop. ")
		return
	}

	for {
		if true == __Vsrv2.runing {
			__Vsrv2.tryCnt2 = 0
			__Vsrv2.tryCnt3 = 3 // if within 3 * gap , not try re-connect
			_Fsleep(__Vgap)
		} else {
			if 0 == __Vsrv2.tryCnt2 {
				if 0 == __Vsrv2.tryCnt3 {
				} else {
				}
			} else {
			}
		}
		//		__Vsrv3 := __Vsrv2._Fanalyze_srvInfo__whether_trying_to_connect()
		//		if nil != __Vsrv3 {
		//			//___Vgtm.uTmSetSrvInfo02
		//		}
		_Fsleep(__Vgap)
	}
}

func (___Vsrv *_TsrvInfo) _Fanalyze_srvInfo__whether_trying_to_connect() *_TsrvInfo {
	return nil
}
