package main

func (___Vdm *_TdataMachine) IRun(___Vidx int) {
	switch ___Vidx {
	case 1000101:
		if nil == ___Vdm.dmCBinit {
			_FdataMachin__1000101__main_init__default(___Vdm)
		} else {
			___Vdm.dmCBinit(___Vdm)
		}
	case 1000201:
		if nil == ___Vdm.dmCBrece {
			_FdataMachin__1000201x__receive__default(___Vdm)
		} else {
			___Vdm.dmCBrece(___Vdm)
		}
	case 1000501:
		___Vdm.
			_FdataMachin__1000501x__swapLoginCkInfoForLock__timeoutGen()
	case 1000502:
		___Vdm.
			_FdataMachin__1000502x__dataSendIdle__gen_time_gap()
	case 1000503:
		___Vdm.
			_FdataMachin__1000503x__time_gap_dataResend()
	//case 1000601:
	//	___Vdm.
	//		_FdataMachin__1000601x__encodeData_sendMux()
	default:
		_FpfNex(" 839191 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FdataMachin__1000101__main_init__default(___Vdm *_TdataMachine) {

	___Vdm.dmCHdataMachineIdI = make(chan _TdataMachinEid, 50)
	___Vdm.dmCHdecodeDataI = make(chan _Tdecode, 50)
	___Vdm.dmChSendIdleNoteInternalUSE = make(chan byte, 1) // a random timer , send idle note to main receive loop. internal use only.
	___Vdm.dmChSwapLoginCkInfoForLock = make(chan byte, 1)  // a 5s timer , send swap note to main receive loop. internal use only.

	___Vdm.dmMconn.ddcMidx = make(map[[16]byte]int)
	___Vdm.dmMdata.ddsMidx = make(map[[16]byte]int)
	for __Vi := 0; __Vi < _VallowClientMax; __Vi++ {
		___Vdm.dmMconn.ddcM[__Vi].dmmConnPortStrMap = make(map[string]byte) // _TdataMachinEconnectClient
		___Vdm.dmMdata.ddsM[__Vi].dmdConnPortStrMap = make(map[string]byte) // _TdataMachinEdataClient
	}
	___Vdm.dmMconn.ddcMfreeAmount = _VallowClientMax
	___Vdm.dmMdata.ddsMfreeAmount = _VallowClientMax

	_Fsleep(_T1s)

	go _Frun(___Vdm, 1000201) // _FdataMachin__1000201x__receive__default // <-___Vdm.dmCHdataMachineIdI // <-___Vdm.dmCHdecodeDataI
	go _Frun(___Vdm, 1000501) // _FdataMachin__1000501x__swapLoginCkInfoForLock__timeoutGen // (*___Vdm.dmCHloginGenMachineIdLO)<-
	go _Frun(___Vdm, 1000502) // _FdataMachin__1000502x__dataSendIdle__gen_time_gap // (*___Vdm.dmCHloginGenMachineIdLO) <-
	go _Frun(___Vdm, 1000503) // _FdataMachin__1000503x__time_gap_dataResend
}
