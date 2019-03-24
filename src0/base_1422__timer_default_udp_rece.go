package main

func _FudpTimer__750201x__gap_receive__default(___Vutm *_TgapTimer) {
	//_FpfNdb(" 388193 01 : filte received start ")
	for {
		//_Fsleep_100s()
		select {
		case __Vrece := <-(*___Vutm.uTmCHudpReceLX.unCHreceLX): // _TudpNodeDataRece
			_FpfNhex(&__Vrece.urrBuf, 38, " 388193 06 : %T , decode len <%d>: %11d", __Vrece.urrBuf, __Vrece.urrLen, _FtimeI64())
		}
	}
}
