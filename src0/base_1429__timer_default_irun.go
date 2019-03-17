package main

func (___Vgtm *_TgapTimer) IRun(___Vidx int) {
	switch ___Vidx {
	case 750101:
		if nil == ___Vgtm.uTmCB750101init {
			_FudpTimer__750101x__init__default(___Vgtm)
		} else {
			___Vgtm.uTmCB750101init(___Vgtm)
		}
	case 750102:
		_FudpTimer__750102x__init__tryUdpConn__default(___Vgtm)
	case 750201:
		if nil == ___Vgtm.uTmCB750201filter {
			_FudpTimer__750201x__filter_receive__default(___Vgtm)
		} else {
			___Vgtm.uTmCB750201filter(___Vgtm)
		}
	default:
		_FpfNex(" 839182 99 unknow :idx %d", ___Vidx)
	}
	//uTmCB750201filter
}

// try received from uTmCHudpReceLX
func _FudpTimer__750101x__init__default(___Vgtm *_TgapTimer) {
	_FpfN(" 838918 01 : udpTimer 750101 ")

	if nil == ___Vgtm.uTmCHudpReceLX {
		_Fex1(" 838918 02 : you must set the Receive _TudpNodeSt ")
	}

	__VuRece := make(chan _TudpNodeDataRece, 10)
	___Vgtm.uTmCHudpReceLX.unCHreceLX = &__VuRece

	go _Frun(___Vgtm, 750201) // IRun
}
