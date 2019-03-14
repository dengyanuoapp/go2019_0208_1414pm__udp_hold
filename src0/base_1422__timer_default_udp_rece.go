package main

func _FudpTimer__150201x__filter_receive__default(___Vutm *_TudpTimer) {
	__VchRece := ___Vutm.uTmCHudpRece.unCHrece
	_FpfN(" 388193 01 : filte received 150201 ")
	for {
		select {
		case __Vrece := <-(*__VchRece):
			_FpfN(" 388193 02 : filte received 150201 %T : %+v", __Vrece, __Vrece)
		}
	}
}
