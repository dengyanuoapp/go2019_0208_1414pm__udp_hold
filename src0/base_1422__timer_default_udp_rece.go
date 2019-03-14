package main

func _FudpTimer__150201x__filter_receive__default(___Vutm *_TudpTimer) {
	_FpfN(" 388193 01 : filte received 150201 ")
	for {
		select {
		case __Vrece := <-(*___Vutm.uTmCHudpRece.unCHrece): // _TudpNodeDataRece
			_FpfN(" 388193 02 : filte received 150201 %T ", __Vrece)
			_FpfNhex(__Vrece.unInBuf, 32, " 388193 03 : %T %d", __Vrece.unInBuf, __Vrece.unInLen)
			//_Ppt(" 388193 04 : ", __Vrece.unInBuf)
		}
	}
}
