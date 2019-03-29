package main

func (___Vlc *_TloginCheck) IRun(___Vidx int) {
	switch ___Vidx {
	case 900101:
		if nil == ___Vlc.ucCB900101init {
			___Vlc._FloginCheck__900101x__init__default()
		} else {
			___Vlc.ucCB900101init(___Vlc)
		}
	case 900201:
		if nil == ___Vlc.ucCB900201stCheck {
			___Vlc._FloginCheck__900201x__standardCheck()
		} else {
			___Vlc.ucCB900201stCheck(___Vlc)
		}
	default:
		_FpfNex(" 838391 99 unknow :idx %d", ___Vidx)
	}
}

func (___Vlc *_TloginCheck) _FloginCheck__900101x__init__default() {

	___Vlc.ucDecodeI = make(chan _Tdecode, 8)

	_FpfNdb(" 838391 01 start ")
	go _Frun(___Vlc, 900201) // _FloginCheck__900201x__standardCheck
}

func (___Vlc *_TloginCheck) _FloginCheck__900201x__standardCheck() {
	_Fsleep(_T1s)
	if nil == ___Vlc.ucCHSendLO {
		_FpfNdb(" 838392 01 why no Chan output ? exit loop")
		return
	}

	_FpfN(" 838392 02 start")

	var __Vdecode _Tdecode
	for {
		//_Fsleep_100s()
		select {
		case __Vdecode = <-___Vlc.ucDecodeI: // _Tdecode
			_FpfNdb(" 838392 08 : %v", __Vdecode)
		}
	}
}
