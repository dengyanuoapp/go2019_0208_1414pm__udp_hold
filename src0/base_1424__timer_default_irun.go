package main

func (___Vutm *_TgapTimer) IRun(___Vidx int) {
	switch ___Vidx {
	case 150001:
		if nil == ___Vutm.uTmCB150101init {
			_FudpTimer__150001x__init__default(___Vutm)
		} else {
			___Vutm.uTmCB150101init(___Vutm)
		}
	}
	//uTmCB150201filter
}

func _FudpTimer__150001x__init__default(___Vutm *_TgapTimer) {
	_FpfN(" 838918 01 : udpTimer 150001 ")

	if nil == ___Vutm.uTmCHudpRece01 {
		_Fex1(" 838918 02 : you must set the Receive _TudpNodeSt ")
	}

	__VuRece := make(chan _TudpNodeDataRece, 10)
	___Vutm.uTmCHudpRece01.unCHrece = &__VuRece

	if nil == ___Vutm.uTmCB150201filter {
		_FudpTimer__150201x__filter_receive__default(___Vutm)
	} else {
		___Vutm.uTmCB150201filter(___Vutm)
	}
}
