package main

import (
	"net"
)

// _TreqIneedToLogin
// _TsrvInfo
// _TgapTimer
//uTmGapNewSession2 *_TgapNewSession
func _FudpTimer__750102x__init__tryUdpConn__default(___Vgtm *_TgapTimer) {
	__VnewSession := ___Vgtm.uTmGapNewSession2
	_FpfNdb(" 138181 01 %#v ", ___Vgtm.uTmGapNewSession2)

	__Vgap := ___Vgtm.uTmGapX
	if 0 == __Vgap {
		_FpfNdb(" 138181 02 %#v ", ___Vgtm.uTmGapNewSession2)
		_FpfN(" 138181 03 : sorry , gap is ZERO , skip gap loop. ")
		return
	}

	for {
		//_Fsleep(__Vgap) // mini Gap , at least
		_FsleepRand_12_to_14s()
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
					__VucPort := __VnewSession._FudpTimer__750102z__tryfillSendChan()
					if nil == ___Vgtm.uTmUconnPortLX {
						_FpfNdb(" 138181 07: tryCnt %d , skipCnt %d, %v", __VnewSession.tryCnt, __VnewSession.skipCnt, __VucPort)
					} else {
						*___Vgtm.uTmUconnPortLX <- *__VucPort
						_FsleepRand_12_to_14s()
						*___Vgtm.uTmUconnPortLX <- *__VucPort
					}
					__VnewSession.tryCnt--
					__VnewSession.skipCnt = 2
				}
			}
		}
	}
}

// _TgapNewSession _TsrvInfo _TudpConnPort
func (___VnewSession *_TgapNewSession) _FudpTimer__750102z__tryfillSendChan() *_TudpConnPort {
	__Vlen2 := len(___VnewSession.srvInfo.UriArrs)
	__Vlen3 := len(___VnewSession.srvInfo.K256)
	__Vidx2 := ___VnewSession.tryCnt % __Vlen2

	//Uri: ___VnewSession.srvInfo.UriArrs[__Vidx2],
	// func ResolveUDPAddr(network, address string) (*UDPAddr, error)
	//__Verr2 error
	__VuAddr, __Verr2 := net.ResolveUDPAddr("udp4", ___VnewSession.srvInfo.UriArrs[__Vidx2])
	if nil != __Verr2 {
		_FpfN(" 388181 01 : %v", __Verr2)
		return nil
	}

	__VuConn := _TudpConnPort{}
	__VuConn.DstAddr = *__VuAddr

	if __Vidx2 >= __Vlen3 {
		__VuConn.K256 = ___VnewSession.srvInfo.K256[0]
	} else {
		__VuConn.K256 = ___VnewSession.srvInfo.K256[__Vidx2]
	}
	return &__VuConn
}

// _TgapNewSession _TsrvInfo
func (___VnewSession *_TgapNewSession) _FudpTimer__750102y__tryGetSrvInfoFromUri() {
	_FpfNdb("238191 01 : %v", ___VnewSession)

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
			//_FpfN(" 311913 02 : Error : update Uri slice failed.: %s , %v ", __nowUri3, __Verr2)
			__nowUri3 = __Vsi2.refreshUri
			_, __Verr2 = _Ftry_download_rand_json01(__nowUri3, &__Vsi2.refreshPwd, &__VtmpSi2)
			if nil != __Verr2 {
				_FpfN(" 311913 03 : Error : update Uri slice failed.: %s , %v ", __nowUri3, __Verr2)
				return
			} else {
				//_FpfN(" 311913 04 : ok : %s , %v ", __nowUri3, __VtmpSi2)
			}
		} else {
			_FpfN(" 311913 05 : ok : %s , %v ", __nowUri3, __VtmpSi2)
		}

		if nil == __VtmpSi2.refreshPwd {
			__VtmpSi2.refreshPwd = ___VnewSession.updatePasswd // use default passwd if nil
		}

		___VnewSession.srvInfo = __VtmpSi2

		if "" == ___VnewSession.srvInfo.refreshUri {
			//_FpfN(" 311913 06 : %v", ___VnewSession.srvInfo)
			break
		}

		if ___VnewSession.srvInfo.refreshUri == __nowUri3 { //_VsrvInfo_Dn     _TsrvInfo
			_FpfN(" 311913 07 : %s , %s , %s", __VtmpSi2.refreshUri, ___VnewSession.srvInfo.refreshUri, __nowUri3)
			break
		}

		_FpfN(" 311913 08 : download loop ing")

	}

	//_FpfNdb(" 311913 09 : all ok : %s , %v ", __nowUri3, __VtmpSi2)

	___VnewSession.srvInfo.ok = true
}
