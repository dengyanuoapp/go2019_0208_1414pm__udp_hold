package main

func _FudpDecode__800301x__connGen__default(___Vlg *_TloginGenerator) {
	__VnewSession := ___Vlg.lgSrvDownInfoLX
	_FpfNdb(" 138181 01 %#v ", ___Vlg.lgSrvDownInfoLX)

	for {
		//_Fsleep(__Vgap) // mini Gap , at least
		_FsleepRand_12_to_14s()
		if true == ___Vlg.lgConnected {
			__VnewSession.tryCnt = 0
			__VnewSession.skipCnt = 6 // set timeout to 60s , not try re-connect
		} else { // false == ___Vlg.lgConnected
			// tryCnt , skipCnt
			if 0 != __VnewSession.skipCnt {
				__VnewSession.skipCnt-- // wait only , do nothing.
			} else { // 0 == __VnewSession.skipCnt
				if 0 == __VnewSession.tryCnt { // 0,0 : re-download
					__VnewSession._FudpDecode__800501x__tryGetSrvInfoFromUri()
					if false == __VnewSession.srvInfo.ok {
						_FpfN(" 138181 04 : download failed. ")
						__VnewSession.skipCnt = 6 // srvInfo get error , wait 60s before retry
					} else {
						//_FpfN(" 138181 05 : download succeed. ")
						__VnewSession.tryCnt = 20 * len(__VnewSession.srvInfo.UriArrs)
						__VnewSession.skipCnt = 3
						if 0 == __VnewSession.tryCnt {
							_FpfN(" 138181 06 : download error , why zero ?. : %d ", __VnewSession.tryCnt)
						}
						// try 20 times of UriArrs to connect
					}

				} else { // xTry,0skip
					__VucPort := __VnewSession._FudpDecode__800401x__tryfillSendChan() // _TudpConnPort
					if nil == __VucPort {
						_FpfNdb(" 138181 09: why nil ?")
					} else {
						if nil == ___Vlg.lgCHuConnPortLO {
							_FpfNdb(" 138181 07: tryCnt %d , skipCnt %d, %v", __VnewSession.tryCnt, __VnewSession.skipCnt, __VucPort)
						} else {
							if 2 == 2 {
								_Fn()
								_FpfNdb(
									" 138181 08: tryToLogin, push port address to connPort-Chan , tryCnt %d , skipCnt %d",
									__VnewSession.tryCnt, __VnewSession.skipCnt)
							}
							(*___Vlg.lgCHuConnPortLO) <- (*__VucPort)
						}
						__VnewSession.tryCnt--
						__VnewSession.skipCnt = 2
					}
				}
			}
		}
	}
}
