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

func (___Vdm *_TdataMachine) _FdataMachin__1000601y__encodeData_sendEnc_only(___Vdmdc *_TdataMachinEdataClient, ___Venc *_Tencode) {
	if ___Venc == nil || ___Vdmdc == nil {
		_FpfN(" 381922 00 while nil ? <%v> <%v> ", ___Vdmdc, ___Venc)
	}

	__VtoAddrAmount := len(___Vdmdc.ddcConnPortArr) // _TdataMachinEdataClient
	if __VtoAddrAmount < 1 {
		_FpfN(" 381922 01 len error :%d ", __VtoAddrAmount)
		return
	}

	__Vidx := _FgenRand_int() % __VtoAddrAmount
	if ___Vdmdc.ddcLastSendIdx == __Vidx {
		__Vidx = _FgenRand_int() % __VtoAddrAmount
	}
	___Vdmdc.ddcLastSendIdx = __Vidx

	// _TdataMachinEdataClient
	__Vconn := ___Vdmdc.ddcConnPortArr[__Vidx]

	___Venc.EnToConnPort = __Vconn // _TudpConnPort
	___Venc.EnTi = _FtimeInt()

	//_FpfN(" 381922 02 {%#v} ", ___Vdmdc) // _TdataMachinEconnectClient
	//_FpfN(" 381922 03 {%s} ", ___Vdmdc.String())
	//_FpfN(" 381922 04 : myID:%x mySeq:%x id:%x", _VC.MyId128, _VS.MySeq128, ___Vid)

	//_FpfN(" 381922 05 _TdataMachine -> _Tencode {%s}", __Venc.String()) // _TencodeX
	//_CpfN(" 381922 06 _TdataMachine -> _Tencode {%s}", __Venc.String()) // _TencodeX

	//_Fex("381922 07")

	___Vdmdc.ddcLastSendTime = _FtimeInt() // _TdataMachinEconnectClient

	___Vdm.
		_FdataMachin__1000601x__encodeData_sendMux(___Venc)
}
