package main

func (___Vutm *_TudpTimer) IRun(___Vidx int) {
	switch ___Vidx {
	case 150001:
		if nil == ___Vutm.uTmCB150001init {
			_FudpTimer__150001x__init__default(___Vutm)
		} else {
			___Vutm.uTmCB150001init(___Vutm)
		}
	}
	//uTmCB150201filter
}

func _FudpTimer__150001x__init__default(___Vutm *_TudpTimer) {
	_FpfN(" 838918 01 : udpTimer 150001 ")
}
