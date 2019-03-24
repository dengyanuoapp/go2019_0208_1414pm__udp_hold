package main

func _FudpTimer__750201x__filter_receive__default(___Vutm *_TgapTimer) {
	_FpfN(" 388193 01 : filte received 150201 ")
	for {
		//_Fsleep_100s()
		select {
		case __Vrece := <-(*___Vutm.uTmCHudpReceLX.unCHreceLX): // _TudpNodeDataRece
			//_FpfN(" 388193 02 : filte received 150201 %T ", __Vrece)
			//_FpfNhex(&__Vrece.urrBuf, 32, " 388193 03 : %T %d", __Vrece.urrBuf, __Vrece.urrLen)
			//_FpfN(" 388193 04 : %T len <%d>:\n '%s'\n==------end======= ", __Vrece.urrBuf, __Vrece.urrLen, __Vrece.urrBuf)
			//_Ppt(" 388193 05 : ", __Vrece.urrBuf)
			_FpfNhex(&__Vrece.urrBuf, 40, " 388193 06 : %T , decode len <%d>: ", __Vrece.urrBuf, __Vrece.urrLen)
		}
		_Ppf("                                     %d\n", _FtimeI64())
	}
}
