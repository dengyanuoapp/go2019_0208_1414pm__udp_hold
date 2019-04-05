package main

func _FudpDecode__700201x__receive__default(___Vutm *_TuDecode) {
	go _FudpDecode__700201x10__receive__default(___Vutm)
}

func _FudpDecode__700201x10__receive__default(___Vutm *_TuDecode) {
	//_FpfNdb(" 388195 01 : filte received start ")
	var __VunRece _TudpNodeDataRece
	var __Vdecode _Tdecode
	for {
		//_Fsleep_100s()
		//__VunRece = _TudpNodeDataRece{}
		//__Vdecode = _Tdecode{}
		select {
		case __VunRece = <-___Vutm.uTmCHunDataReceI: // _TudpNodeDataRece
			//_FpfNhex(&__VunRece.urrBuf, 38, " 388195 02 : %T , decode len <%d>: %11d", __VunRece.urrBuf, __VunRece.urrLen, _FtimeI64())
			__VunRece.
				_FdataPack__decode_from_udpNodeDataRece(&__Vdecode) // _Tdecode
			__Vdecode.remoteAddr = __VunRece.urrRemoteAddr
			//__Vdecode.receiveTime = _FtimeInt()
			_FpfNdb(" 388195 03 : decoder out:{%s} ", __Vdecode.String())

			switch __Vdecode.Type {
			case Cmd__loginS01genReplyTokenA, Cmd__loginS02genReplyTokenB,
				Cmd__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken: // 15540362231554036223
				//_FpfN(" 388195 05 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
				if nil == ___Vutm.uTmDecodeCmdLO {
					_FpfN(" 388195 06 : outChan null , ignore:%s", __Vdecode.String())
				} else {
					//_FpfN(" 388195 08 real outChain : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
					//_FpfN(" 388195 09 real outChain : %s", __Vdecode.String())
					(*___Vutm.uTmDecodeCmdLO) <- __Vdecode // 15540463611554046361
				}
			case Cmd__data:
				_FpfN(" 388197 01 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
				if nil == ___Vutm.uTmDecodeDataLO {
					_FpfN(" 388197 03 : outChan null , ignore ")
				} else {
					_FpfN(" 388197 05 real outChain : type %d", __Vdecode.Type)
					(*___Vutm.uTmDecodeDataLO) <- __Vdecode //
				}
			default:
				_FpfN(" 388197 09 : type %d : unknow how to deal with.", __Vdecode.Type)
			}
		}
	}
}
