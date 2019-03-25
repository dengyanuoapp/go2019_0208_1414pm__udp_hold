package main

func _FudpTimer__750201x__gap_receive__default(___Vutm *_TgapTimer) {
	//_FpfNdb(" 388193 01 : filte received start ")
	if nil != ___Vutm.uTmCHudpReceLX.unCHreceLX {
		go _FudpTimer__750201x10__gap_receive__default(___Vutm)
	} else {
		_FpfNdb(" 388193 09 : why no input channel ?")
	}
}

func _FudpTimer__750201x10__gap_receive__default(___Vutm *_TgapTimer) {
	//_FpfNdb(" 388195 01 : filte received start ")
	for {
		//_Fsleep_100s()
		select {
		case __Vrece := <-(*___Vutm.uTmCHudpReceLX.unCHreceLX): // _TudpNodeDataRece
			_FpfNhex(&__Vrece.urrBuf, 38, " 388195 06 : %T , decode len <%d>: %11d", __Vrece.urrBuf, __Vrece.urrLen, _FtimeI64())
		}
	}
}
