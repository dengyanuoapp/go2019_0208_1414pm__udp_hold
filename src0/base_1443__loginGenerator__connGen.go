package main

func _FloginGen__800301x__connGen__default(___Vlg *_TloginGenerator) {
	__VnewSession := ___Vlg.lgSrvDownInfoLX // *_TsrvDownInfo
	//_FpfNdb(" 138171 01 %#v ", ___Vlg.lgSrvDownInfoLX)
	_FpfNdb(" 138171 02 %s ", ___Vlg.lgSrvDownInfoLX.String())

	for {
		//_Fsleep(__Vgap) // mini Gap , at least
		_FsleepRand_12_to_14s()

		if true == ___Vlg.lgConnected {
			_CFpfN(" 138171 02 : don't need to connect server ")
		} else {
			if _FtimeInt()-__VnewSession.lastDownTime > 3600 { // 0,0 : re-download
				_CFpfN(" 138171 03 : trying to update server info. ")
				___Vlg.
					_FloginGen__800301y1__connGen__update_serverInfo(__VnewSession)
			} else {
				_CFpfN(" 138171 05 : trying to connect server. ")
				___Vlg.
					_FloginGen__800301y2__connGen__tryConnect(__VnewSession)
			}
		}
	}
}

func (___Vlg *_TloginGenerator) _FloginGen__800301y1__connGen__update_serverInfo(___VnewSession *_TsrvDownInfo) {

	_CFpfN(" 138172 02 : don't need to connect server ")

	___VnewSession.
		_FudpDecode__800501x__tryGetSrvInfoFromUri()
	if false == ___VnewSession.srvInfo.ok { // _TsrvInfo
		_FpfN(" 138172 03 : download failed. ")
		_Fsleep(_T60s)
	} else {
		_Fsleep(_T5s)
		_FpfN(" 138172 04 : download ok. %s", ___VnewSession.String())
		_CpfN(" 138172 05 : download ok. %s", ___VnewSession.String())
		___VnewSession.lastDownTime = _FtimeInt()
	}
}

func (___Vlg *_TloginGenerator) _FloginGen__800301y2__connGen__tryConnect(___VnewSession *_TsrvDownInfo) {
	___VnewSession.tryCnt++
	__VucPort := ___VnewSession.
		_FudpDecode__800401x__tryfillSendChan() // _TudpConnPort
	if nil == __VucPort {
		_FpfNdb(" 138173 06: why nil ?")
		_Fsleep(_T120s)
		___VnewSession.lastDownTime = 0 // req-redownload
	} else {
		if nil == ___Vlg.lgCHuConnPortLO {
			_FpfNdb(" 138173 07: Error ! tryCnt %d , skipCnt %d, %v", ___VnewSession.tryCnt, ___VnewSession.skipCnt, __VucPort)
		} else {
			_CpfN(" 138173 08: tryToLogin, push port address to connPort-Chan , %11d, tryCnt %d , skipCnt %d", // check-important keykey
				_FtimeInt(), ___VnewSession.tryCnt, ___VnewSession.skipCnt)
			(*___Vlg.lgCHuConnPortLO) <- (*__VucPort)

		}
	}
	_Fsleep(_T60s)
}
