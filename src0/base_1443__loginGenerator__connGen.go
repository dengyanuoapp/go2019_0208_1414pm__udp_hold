package main

func _FudpDecode__800301x__connGen__default(___Vlg *_TloginGenerator) {
	__VnewSession := ___Vlg.lgSrvDownInfoLX // *_TsrvDownInfo
	//_FpfNdb(" 138181 01 %#v ", ___Vlg.lgSrvDownInfoLX)
	_FpfNdb(" 138181 02 %s ", ___Vlg.lgSrvDownInfoLX.String())

	for {
		//_Fsleep(__Vgap) // mini Gap , at least
		_FsleepRand_12_to_14s()
		__Vnow3 := _FtimeInt()
		if __Vnow3-__VnewSession.lastDownTime > 3600 { // 0,0 : re-download
			__VnewSession.
				_FudpDecode__800501x__tryGetSrvInfoFromUri()
			if false == __VnewSession.srvInfo.ok { // _TsrvInfo
				_FpfN(" 138181 03 : download failed. ")
				_Fsleep(_T60s)
			} else {
				_Fsleep(_T5s)
				_FpfN(" 138181 04 : download ok. %s", __VnewSession.String())
				_CpfN(" 138181 05 : download ok. %s", __VnewSession.String())
				__VnewSession.lastDownTime = _FtimeInt()
			}
		} else { // xTry,0skip
			__VnewSession.tryCnt++
			__VucPort := __VnewSession.
				_FudpDecode__800401x__tryfillSendChan() // _TudpConnPort
			if nil == __VucPort {
				_FpfNdb(" 138181 06: why nil ?")
				_Fsleep(_T120s)
				__VnewSession.lastDownTime = 0 // req-redownload
			} else {
				if nil == ___Vlg.lgCHuConnPortLO {
					_FpfNdb(" 138181 07: Error ! tryCnt %d , skipCnt %d, %v", __VnewSession.tryCnt, __VnewSession.skipCnt, __VucPort)
				} else {
					_CpfN(
						" 138181 08: tryToLogin, push port address to connPort-Chan , %11d, tryCnt %d , skipCnt %d", // check-important
						_FtimeInt(), __VnewSession.tryCnt, __VnewSession.skipCnt)
					(*___Vlg.lgCHuConnPortLO) <- (*__VucPort)

				}
			}
			_Fsleep(_T60s)
		}
	}
}
