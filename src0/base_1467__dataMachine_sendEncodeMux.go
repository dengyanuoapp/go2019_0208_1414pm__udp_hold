package main

import "sync"

var ___V_FdataMachin__1000601x__encodeData_sendMux__mux sync.Mutex

func (___Vdm *_TdataMachine) _FdataMachin__1000601x__encodeData_sendMux(___Venc *_Tencode) {
	//dmCHencodeIdleLO            chan _Tencode // _TencodeX
	defer ___V_FdataMachin__1000601x__encodeData_sendMux__mux.Unlock()
	___V_FdataMachin__1000601x__encodeData_sendMux__mux.Lock()

	if nil == ___Vdm.dmCHencodeIdleLO {
		_FpfNonce(" 481919 01 : error : chan-out NULL ")
		return
	}

	(*___Vdm.dmCHencodeIdleLO) <- (*___Venc)
}
