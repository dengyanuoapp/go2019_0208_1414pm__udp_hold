package main

import "sync"

var ___VuEncode__1100201x__chanIn_mux sync.Mutex
var ___VuEncode__1100201x__send_mux sync.Mutex

func _FuEncode__1100201x__chanRece__default(___Vuen *_TuEncode) {

	var __Vue _Tencode
	var __Vus _TudpNodeDataSend // _TudpNodeDataSendX
	for {
		__Vus = _TudpNodeDataSend{}
		__Vue = _Tencode{}
		select {
		case __Vue = <-___Vuen.enCHencodeCkI: // chan _TencodeX
			___VuEncode__1100201x__chanIn_mux.Lock()
			_FpfNonce(" 849192 01 : try Encode CMD {%s}", __Vue.String())
			//___Vuen.
			//	_FdataMachin__1000201x11__rece_machineId_check_and_insert(&__Vue)
			//__Vue.enLogin.
		case __Vue = <-___Vuen.enCHencodeDataCommI: // chan _TencodeX
			___VuEncode__1100201x__chanIn_mux.Lock()
			//_CFpfN(" 849192 05 : try Encode idle Data{%s}", __Vue.String())
		}
		__Vue.
			_FdataPack__100__loginReq(__Vue.enType, &__Vus.usOutBuf)

		if 0 == __Vue.Ti {
			__Vue.Ti = _FgenRand_int()
		}
		___Vuen.
			_FuEncode__1100201x2__fillAddr_sending(__Vus, __Vue)

		___VuEncode__1100201x__chanIn_mux.Unlock()
	}
}

func (___Vuen *_TuEncode) _FuEncode__1100201x2__fillAddr_sending(___Vus _TudpNodeDataSend, ___Vue _Tencode) { // _TudpNodeDataSendX _TencodeX
	if nil == ___Vuen.enCHuDataSendLO { // *chan _TudpNodeDataSend // _TudpNodeDataSendX
		_FpfN(" 849193 01 : why NULL chan out? ignore :%v:%v", ___Vus, ___Vue)
		return
	}

	//	if nil == ___Vus || nil == ___Vue {
	//		_FpfNex(" 849193 01 : why NULL ? %v:%v", ___Vus, ___Vue)
	//	}
	if 0 == len(___Vus.usOutBuf) {
		_FpfN(" 849193 02 : len ZERO, ignore ")
		return
	}

	if ((0 == len(___Vue.enToId128)) || (16 == len(___Vue.enToId128))) && 32 == len(___Vue.enToConnPort.K256) {
		___Vus.usToAddr = ___Vue.enToConnPort // _TudpConnPort
	} else {
		_FpfN(" 849193 03 : under constructing addr :%s", ___Vue.String())
		return
	}

	___Vus.Ti = ___Vue.Ti

	_FpfNonce(" 849193 04 : filling addr ok:%s", ___Vus.String())
	(*___Vuen.enCHuDataSendLO) <- ___Vus

	if 0 != ___Vue.enDelay {
		go _FudpNodeDataSend__delaySend(___Vue.enDelay, ___Vus, ___Vuen.enCHuDataSendLO)
	}

	__Vi := ___Vue.enMultiSend
	for __Vi != 0 {
		___VuEncode__1100201x__send_mux.Lock()
		(*___Vuen.enCHuDataSendLO) <- ___Vus
		___VuEncode__1100201x__send_mux.Unlock()
		__Vi--
	}
}

func _FudpNodeDataSend__delaySend(___Vdelay int, ___Vus _TudpNodeDataSend, ___VchUnds *chan _TudpNodeDataSend) { // _TudpNodeDataSendX
	if 0 == ___Vdelay {
		return
	}
	if nil == ___VchUnds {
		return
	}

	_Fsleep_1sX(___Vdelay)

	_CpfN("849193 08 delay_send (%d) {%s}", ___Vdelay, ___Vus.String()) // keykey

	___VuEncode__1100201x__send_mux.Lock()
	(*___VchUnds) <- ___Vus
	___VuEncode__1100201x__send_mux.Unlock()
}
