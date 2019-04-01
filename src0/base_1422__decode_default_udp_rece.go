package main

func _FudpDecode__700201x__receive__default(___Vutm *_TuDecode) {
	//_FpfNdb(" 388193 01 : filte received start ")
	if (nil != ___Vutm.uTmCHunDataReceLI) && (nil != ___Vutm.uTmCHunDataReceLI.unCHreceLO) {
		go _FudpDecode__700201x10__receive__default(___Vutm)
	} else {
		_FpfNdb(" 388193 09 : why no input channel ?")
	}
}

func _FudpDecode__700201x10__receive__default(___Vutm *_TuDecode) {
	//_FpfNdb(" 388195 01 : filte received start ")
	var __Vundr _TudpNodeDataRece
	var __Vdecode _Tdecode
	for {
		//_Fsleep_100s()
		select {
		case __Vundr = <-(*___Vutm.uTmCHunDataReceLI.unCHreceLO): // _TudpNodeDataRece
			//_FpfNhex(&__Vundr.urrBuf, 38, " 388195 03 : %T , decode len <%d>: %11d", __Vundr.urrBuf, __Vundr.urrLen, _FtimeI64())
			__Vundr._FdataPack__decode_from_udpNodeDataRece(&__Vdecode) // _Tdecode
			__Vdecode.remoteAddr = __Vundr.urrRemoteAddr
			__Vdecode.receiveTime = _FtimeInt()

			if nil == ___Vutm.uTmDecodeCmdLO {
				//_FpfNdb(" 388195 05 : %#v, key %x", __Vundr, __Vdecode.remotePortKey)
				_FpfNdb(" 388195 06 : %#v, key %x", __Vdecode, __Vdecode.remotePortKey)
			} else {
				__VoutT := false
				switch __Vdecode.Type {
				case Cmd__loginS1ReqTryNoToken: // 15540362231554036223
					_FpfN(" 388195 07 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenA)
					__VoutT = true
				default:
					_FpfN(" 388195 08 : type %d", __Vdecode.Type)
				}
				if __VoutT {
					if nil != ___Vutm.uTmDecodeCmdLO {
						_FpfN(" 388195 09 real outChain : type %d", __Vdecode.Type)
						(*___Vutm.uTmDecodeCmdLO) <- __Vdecode // 15540463611554046361
					}
				}
			}
		}
	}
}
