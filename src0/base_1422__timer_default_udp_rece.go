package main

func _FudpTimer__150201x__filter_receive__default(___Vutm *_TgapTimer) {
	_FpfN(" 388193 01 : filte received 150201 ")
	for {
		//_Fsleep_100s()
		select {
		case __Vrece := <-(*___Vutm.uTmCHudpReceLX.unCHreceLX): // _TudpNodeDataRece
			//_FpfN(" 388193 02 : filte received 150201 %T ", __Vrece)
			//_FpfNhex(&__Vrece.unInBuf, 32, " 388193 03 : %T %d", __Vrece.unInBuf, __Vrece.unInLen)
			_FpfN(" 388193 04 : %T %d: '%s'", __Vrece.unInBuf, __Vrece.unInLen, __Vrece.unInBuf)
			//_Ppt(" 388193 05 : ", __Vrece.unInBuf)
		}
		_Ppf("                                     %d\n", _FtimeI64())
	}
}
