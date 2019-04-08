package main

func _FuEncode__1100201x__packSend__default(___Vuen *_TuEncode) {

	var __Vue _Tencode
	var __Vus _TudpNodeDataSend // _TudpNodeDataSendX
	for {
		__Vus = _TudpNodeDataSend{}
		select {
		case __Vue = <-___Vuen.enCHencodeCmdI: // chan _Tencode
			_FpfNdb(" 849192 01 : try Encode CMD {%s}", __Vue.String())
			//___Vuen.
			//	_FdataMachin__1000201x11__connMap_insertId(&__Vue)
			__Vue.enLogin.
				_FdataPack__100__loginReq(__Vue.enType, &__Vus.usOutBuf)
		case __Vue = <-___Vuen.enCHencodeDataI: // chan _Tencode
			_FpfNdb(" 849192 05 : try Encode Data{%s}", __Vue.String())
			_FpfNdb(" 849192 06 : under constructing ")

		}
		___Vuen._FuEncode__1100201x2__fillAddr_sending(&__Vus)
	}
}
func (___Vuen *_TuEncode) _FuEncode__1100201x2__fillAddr_sending(___Vus *_TudpNodeDataSend) { // _TudpNodeDataSendX
	if nil == ___Vus {
		_Fex(" 849193 01 : why NULL ?")
	}
	if 0 == len(___Vus.usOutBuf) {
		_FpfN(" 849193 02 : len ZERO, ignore ")
	}

	_FpfN(" 849193 04 : try filling addr :%s", ___Vus.String())
}
