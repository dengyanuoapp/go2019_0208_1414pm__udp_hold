package main

import (
//"net"
)

// _TudpNodeSt
// _Fhandle_u01x__udpListen_Udp__read_main_top__default
func (___Vun *_TudpNodeSt) IRun(___Vidx int) {
	switch ___Vidx {
	case 500101:
		if nil == ___Vun.unCBinit {
			_FudpNode__500101__main_init__default(___Vun)
		} else {
			___Vun.unCBinit(___Vun)
		}
	default:
		_FpfNex(" 937191 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FudpNode__500101__main_init__default(___Vun *_TudpNodeSt) {

	___Vun.unRbuf = make([]byte, 1500)
	//_FpfN(" 918381 01 : nodeArr[ %d ] < %s , %s >using %x ", ___Vun.unIdx, ___Vun.unName, ___Vun.unHostPortStr, ___Vun.unRKeyLP)

	___Vun.unRmap.unrMapLas2 = make(map[string]_TuNodeDataRcnt)
	___Vun.unRmap.unrMapLast = make(map[string]_TuNodeDataRcnt)
	___Vun.unRmap.unrMapNow = make(map[string]_TuNodeDataRcnt)

	___Vun.unRkeyX.initKey(___Vun.unRKeyLP) // initKey

	___Vun._FudpNode__500101x__listen()

	go ___Vun._FudpNode__540211x__gap() // active only if unLoopGap is set to some delayGap ==>> pipe the unRmap
	// if gap set ==> replace unCBrece from default-direct to _FudpNode__540211z__receiveCallBack_withTimeGap

	_Fsleep(_T50ms)

	go ___Vun._FudpNode__500101y__receive()

	go ___Vun._FudpNode__500101z__send()

	//_Fex1(" 918381 09 ")
}
