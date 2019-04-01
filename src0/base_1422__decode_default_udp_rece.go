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

			switch __Vdecode.Type {
			case Cmd__loginS1ReqTryNoToken, Cmd__loginS2ReplyTmpToken, Cmd__loginS3ReqWithToken: // 15540362231554036223
				_FpfN(" 388195 05 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenA)
				if nil != ___Vutm.uTmDecodeCmdLO {
					_FpfN(" 388195 06 real outChain : type %d", __Vdecode.Type)
					(*___Vutm.uTmDecodeCmdLO) <- __Vdecode // 15540463611554046361
				}
			case Cmd__data:
				_FpfN(" 388195 07 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenA)
				if nil != ___Vutm.uTmDecodeDataLO {
					_FpfN(" 388195 08 real outChain : type %d", __Vdecode.Type)
					(*___Vutm.uTmDecodeDataLO) <- __Vdecode //
				}
			default:
				_FpfN(" 388195 08 : type %d", __Vdecode.Type)
			}
		}
	}
}
