package main

type _TreqNewSession struct {
	Enabled  bool
	RetryCNT uint16
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
	_FpfN(" 311912 03 %v", _VS)
	_FpfN(" 311912 04 %v", _VC)

	_Fex1(" 311912 09 ")
} // _Fconnect_to_server_01__req_new_sessionID__main_top
