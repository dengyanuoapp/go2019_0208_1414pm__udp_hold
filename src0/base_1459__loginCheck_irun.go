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
	___Vlc.ucM.cmAnow = make(map[[16]byte]_Tdecode)
	___Vlc.ucM.cmAlast = make(map[[16]byte]_Tdecode)
	___Vlc.ucM.cmAlas2 = make(map[[16]byte]_Tdecode)

	_FpfNdb(" 838391 01 start ")
	go _Frun(___Vlc, 900201) // _FloginCheck__900201x__standardCheck
}
