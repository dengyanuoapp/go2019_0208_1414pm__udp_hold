package main

func (___Vgtm *_TgapTimer) IRun(___Vidx int) {
	switch ___Vidx {
	case 150101:
		if nil == ___Vgtm.uTmCB150101init {
			_FudpTimer__150101x__init__default(___Vgtm)
		} else {
			___Vgtm.uTmCB150101init(___Vgtm)
		}
	case 150201:
		if nil == ___Vgtm.uTmCB150201filter {
			_FudpTimer__150201x__filter_receive__default(___Vgtm)
		} else {
			___Vgtm.uTmCB150201filter(___Vgtm)
		}
	}
	//uTmCB150201filter
}

// try received from uTmCHudpRece01
func _FudpTimer__150101x__init__default(___Vgtm *_TgapTimer) {
	_FpfN(" 838918 01 : udpTimer 150101 ")

	if nil == ___Vgtm.uTmCHudpRece01 {
		_Fex1(" 838918 02 : you must set the Receive _TudpNodeSt ")
	}

	__VuRece := make(chan _TudpNodeDataRece, 10)
	___Vgtm.uTmCHudpRece01.unCHrece = &__VuRece

	go _Frun(___Vgtm, 150201)
}
