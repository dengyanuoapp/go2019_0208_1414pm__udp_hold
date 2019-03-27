package main

func _FudpTimer__750201x__gap_receive__default(___Vutm *_TgapTimer) {
	//_FpfNdb(" 388193 01 : filte received start ")
	if (nil != ___Vutm.uTmCHudpReceLX) && (nil != ___Vutm.uTmCHudpReceLX.unCHreceLX) {
		go _FudpTimer__750201x10__gap_receive__default(___Vutm)
	} else {
		_FpfNdb(" 388193 09 : why no input channel ?")
	}
}

func _FudpTimer__750201x10__gap_receive__default(___Vutm *_TgapTimer) {
	//_FpfNdb(" 388195 01 : filte received start ")
	var __Vundr _TudpNodeDataRece
	var __Vdecode _Tdecode
	for {
		//_Fsleep_100s()
		select {
		case __Vundr = <-(*___Vutm.uTmCHudpReceLX.unCHreceLX): // _TudpNodeDataRece
			//_FpfNhex(&__Vundr.urrBuf, 38, " 388195 06 : %T , decode len <%d>: %11d", __Vundr.urrBuf, __Vundr.urrLen, _FtimeI64())
			__Vundr._FdataPack__301__dataDecode_loginS1ReqTryNoToken(&__Vdecode) // _Tdecode
			__Vdecode.remoteAddr = __Vundr.urrRemoteAddr
			__Vdecode.receiveTime = _FtimeInt()

			//_FpfNdb(" 388195 07 : %#v, key %x", __Vundr, __Vdecode.remotePortKey)
			_FpfNdb(" 388195 08 : %#v, key %x", __Vdecode, __Vdecode.remotePortKey)
		}
	}
}
