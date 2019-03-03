package main

type _TreqNewSession struct {
	Enabled      bool
	RetryCNT     uint16
	UpdateUri    string
	UpdatePasswd *[]byte
	SrvInfo      interface{}
} //    _TreqNewSession

// you can
func (___VreqNewSession *_TreqNewSession) _Fconnect_to_server_01__Default() {
	if ___VreqNewSession.Enabled {
		_FpfN(" 311911 01 ")
		_Fsleep_100s()
	} else {
		//_FpfN(" 311911 03 ")
		___VreqNewSession._Fconnect_to_server_01__req_new_sessionID__main_top()
		_FsleepRand_12_to_14s()
	}
} // _Fconnect_to_server_01__Default

func (___VreqNewSession *_TreqNewSession) _Fconnect_to_server_01__req_new_sessionID__main_top() {
	_FpfN(" 311912 01 ")
	_FpfN(" 311912 02 %v", ___VreqNewSession)
	//_FpfN(" 311912 03 %v", _Vself)
	_FpfN(" 311912 04 %v", _VS)
	//_FpfN(" 311912 05 %v", _Vconfig)
	_FpfN(" 311912 06 %v", _VC)
	_FpfN(" 311912 07 %v", ___VreqNewSession)

	if 0 == ___VreqNewSession.RetryCNT {
		_, __Verr := _Ftry_download_rand_json01(___VreqNewSession.UpdateUri, ___VreqNewSession.UpdatePasswd, ___VreqNewSession.SrvInfo)
		if nil == __Verr {
			_FpfN(" 311912 08 : ok : %s , %v ", ___VreqNewSession.UpdateUri, ___VreqNewSession.SrvInfo)
		} else {
			_FpfN(" 311912 09 : Error : %s , %v ", ___VreqNewSession.UpdateUri, __Verr)
		}

	}

	_Fex1(" 311912 99 ")
} // _Fconnect_to_server_01__req_new_sessionID__main_top
