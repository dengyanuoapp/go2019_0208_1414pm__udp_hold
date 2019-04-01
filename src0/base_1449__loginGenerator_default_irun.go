package main

import (
	"net"
)

func (___Vlg *_TloginGenerator) IRun(___Vidx int) {
	switch ___Vidx {
	case 800101:
		_FudpDecode__800101x__init__tryUdpLogin__default(___Vlg)
	default:
		_FpfNex(" 739182 99 unknow :idx %d", ___Vidx)
	}
}

// _TloginReq
// _TsrvInfo
// _TuDecode
// _TloginGenerator
// ulSrvDownInfoLX *_TsrvDownInfo
func _FudpDecode__800101x__init__tryUdpLogin__default(___Vlg *_TloginGenerator) {
	__VnewSession := ___Vlg.ulSrvDownInfoLX
	_FpfNdb(" 138181 01 %#v ", ___Vlg.ulSrvDownInfoLX)

	// ___Vlg.ulSrvDownInfoLX

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
					__VnewSession._FudpDecode__750102y__tryGetSrvInfoFromUri()
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
					___Vlg.ulTmpToken = _FgenRand_nByte__(16)
					__VucPort := __VnewSession._FudpDecode__750102z__tryfillSendChan() // _TudpConnPort
					if nil != __VucPort {
						__VucPort.TKme = ___Vlg.ulTmpToken
						if nil == ___Vlg.ulCHunSendLO {
							_FpfNdb(" 138181 07: tryCnt %d , skipCnt %d, %v", __VnewSession.tryCnt, __VnewSession.skipCnt, __VucPort)
						} else {
							_Pn()
							_FpfNdb(" 138181 08: tryCnt %d , skipCnt %d, Token %x", __VnewSession.tryCnt, __VnewSession.skipCnt, __VucPort.TKme)
							var __VusData _TudpNodeDataSend // _TudpConnPort
							__VucPort._FdataPack__101__udpConnPort(&__VusData.usOutBuf)
							__VusData.usToAddr = *__VucPort   // _TudpConnPort
							*___Vlg.ulCHunSendLO <- __VusData //
							_FsleepRand_12_to_14s()
							*___Vlg.ulCHunSendLO <- __VusData // 15540362231554036223
						}
						__VnewSession.tryCnt--
						__VnewSession.skipCnt = 2
					}
				}
			}
		}
	}
}

// _TsrvDownInfo _TsrvInfo _TudpConnPort
func (___VnewSession *_TsrvDownInfo) _FudpDecode__750102z__tryfillSendChan() *_TudpConnPort {
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

// _TsrvDownInfo _TsrvInfo
func (___VnewSession *_TsrvDownInfo) _FudpDecode__750102y__tryGetSrvInfoFromUri() {
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

		//_, __Verr2 = _Ftry_download_rand_json01(__nowUri3, &__Vsi2.refreshPwd, &__VtmpSi2)
		_, __Verr2 = _Ftry_download_rand_gob01(__nowUri3, &__Vsi2.refreshPwd, &__VtmpSi2)
		if nil != __Verr2 {
			//_FpfN(" 311913 02 : Error : update Uri slice failed.: %s , %v ", __nowUri3, __Verr2)
			__nowUri3 = __Vsi2.refreshUri
			//_, __Verr2 = _Ftry_download_rand_json01(__nowUri3, &__Vsi2.refreshPwd, &__VtmpSi2)
			_, __Verr2 = _Ftry_download_rand_gob01(__nowUri3, &__Vsi2.refreshPwd, &__VtmpSi2)
			if nil != __Verr2 {
				_FpfN(" 311913 03 : Error : update Uri slice failed.: %s , %v ", __nowUri3, __Verr2)
				return
			} else {
				//_FpfN(" 311913 05 : ok : %s , %v ", __nowUri3, __VtmpSi2)
			}
		} else {
			_FpfN(" 311913 06 : ok : %s , %v ", __nowUri3, __VtmpSi2)
		}

		if 0 == len(__VtmpSi2.UriArrs) {
			_FpfN(" 311913 07 : error decode error : %s , %v ", __nowUri3, __VtmpSi2)
			return
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
