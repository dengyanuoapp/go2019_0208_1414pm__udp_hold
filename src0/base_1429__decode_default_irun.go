package main

func (___Vgtm *_TuDecode) IRun(___Vidx int) {
	switch ___Vidx {
	case 700101:
		if nil == ___Vgtm.uTmCB700101init {
			_FudpDecode__700101x__init__default(___Vgtm)
		} else {
			___Vgtm.uTmCB700101init(___Vgtm)
		}
	case 700201:
		if nil == ___Vgtm.uTmCB700201rece2 {
			_FudpDecode__700201x__receive__default(___Vgtm)
		} else {
			___Vgtm.uTmCB700201rece2(___Vgtm)
		}
	case 700301:
		if nil == ___Vgtm.uTmCB700301gap2 {
			_FudpDecode__700301x__loop__default(___Vgtm)
		} else {
			___Vgtm.uTmCB700301gap2(___Vgtm)
		}
	default:
		_FpfNex(" 839182 99 unknow :idx %d", ___Vidx)
	}
}

// try received from uTmCHunDataReceLI
func _FudpDecode__700101x__init__default(___Vgtm *_TuDecode) {
	_FpfN(" 838918 01 : udpTimer 700101 ")

	if nil == ___Vgtm.uTmCHunDataReceLI {
		_Fex1(" 838918 02 : you must set the Receive _TudpNodeSt ")
	}

	__VuRece := make(chan _TudpNodeDataRece, 10)
	___Vgtm.uTmCHunDataReceLI.unCHreceLO = &__VuRece

	_Fsleep(_T1s)

	// IRun _FudpDecode__700201x__receive__default
	go _Frun(___Vgtm, 700201)

	// _FudpDecode__700301x__loop__default
	go _Frun(___Vgtm, 700301)
}
