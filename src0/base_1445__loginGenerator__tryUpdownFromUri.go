package main

// _TsrvDownInfo _TsrvInfo
func (___VnewSession *_TsrvDownInfo) _FudpDecode__800501x__tryGetSrvInfoFromUri() {
	_FpfNonce("238191 01 : %v t:%d", ___VnewSession, _FtimeInt())

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
		__nowUri3 = _Spf("%s.%x", __Vsi2.refreshUri, _VC.MyId128)

		//_, __Verr2 = _Ftry_download_rand_json01(__nowUri3, &__Vsi2.refreshPwd, &__VtmpSi2)
		__VtmpSi2 = _TsrvInfo{}
		_, __Verr2 = _Ftry_download_rand_gob01(__nowUri3, &__Vsi2.refreshPwd, &__VtmpSi2)
		if nil != __Verr2 {
			//_FpfN(" 311913 02 : Error : update Uri slice failed.: %s , %v ", __nowUri3, __Verr2)
			__nowUri3 = __Vsi2.refreshUri
			__VtmpSi2 = _TsrvInfo{}
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
