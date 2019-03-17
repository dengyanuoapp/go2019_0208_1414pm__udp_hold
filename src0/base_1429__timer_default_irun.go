package main

func (___Vgtm *_TgapTimer) IRun(___Vidx int) {
	switch ___Vidx {
	case 150101:
		if nil == ___Vgtm.uTmCB150101init {
			_FudpTimer__150101x__init__default(___Vgtm)
		} else {
			___Vgtm.uTmCB150101init(___Vgtm)
		}
	case 150102:
		_FudpTimer__150102x__init__tryUdpConn__default(___Vgtm)
	case 150201:
		if nil == ___Vgtm.uTmCB150201filter {
			_FudpTimer__150201x__filter_receive__default(___Vgtm)
		} else {
			___Vgtm.uTmCB150201filter(___Vgtm)
		}
	}
	//uTmCB150201filter
}

// try received from uTmCHudpReceLX
func _FudpTimer__150101x__init__default(___Vgtm *_TgapTimer) {
	_FpfN(" 838918 01 : udpTimer 150101 ")

	if nil == ___Vgtm.uTmCHudpReceLX {
		_Fex1(" 838918 02 : you must set the Receive _TudpNodeSt ")
	}

	__VuRece := make(chan _TudpNodeDataRece, 10)
	___Vgtm.uTmCHudpReceLX.unCHreceLX = &__VuRece

	go _Frun(___Vgtm, 150201)
}
