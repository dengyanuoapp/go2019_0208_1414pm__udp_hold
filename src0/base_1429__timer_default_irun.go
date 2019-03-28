package main

func (___Vgtm *_TgapTimer) IRun(___Vidx int) {
	switch ___Vidx {
	case 750101:
		if nil == ___Vgtm.uTmCB750101init {
			_FudpTimer__750101x__init__default(___Vgtm)
		} else {
			___Vgtm.uTmCB750101init(___Vgtm)
		}
	case 750201:
		if nil == ___Vgtm.uTmCB750201rece2 {
			_FudpTimer__750201x__gap_receive__default(___Vgtm)
		} else {
			___Vgtm.uTmCB750201rece2(___Vgtm)
		}
	case 750301:
		if nil == ___Vgtm.uTmCB750301gap2 {
			_FudpTimer__750301x__gap_loop__default(___Vgtm)
		} else {
			___Vgtm.uTmCB750301gap2(___Vgtm)
		}
	default:
		_FpfNex(" 839182 99 unknow :idx %d", ___Vidx)
	}
}

// try received from uTmCHunDataReceLI
func _FudpTimer__750101x__init__default(___Vgtm *_TgapTimer) {
	_FpfN(" 838918 01 : udpTimer 750101 ")

	if nil == ___Vgtm.uTmCHunDataReceLI {
		_Fex1(" 838918 02 : you must set the Receive _TudpNodeSt ")
	}

	__VuRece := make(chan _TudpNodeDataRece, 10)
	___Vgtm.uTmCHunDataReceLI.unCHreceLX = &__VuRece

	_Fsleep(_T1s)

	// IRun _FudpTimer__750201x__gap_receive__default
	go _Frun(___Vgtm, 750201)

	// _FudpTimer__750301x__gap_loop__default
	go _Frun(___Vgtm, 750301)
}
