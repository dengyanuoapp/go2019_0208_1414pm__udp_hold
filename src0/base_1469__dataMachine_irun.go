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
	case 1000504:
		___Vdm.
			_FdataMachin__1000504x__checkTimeOutDieClient()
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
	___Vdm.dmCHdebugInfoI = make(chan byte, 50)
	___Vdm.dmCHencodeDataSpecI = make(chan _TdataTran, 50) // _TencodeX , used for Fn-Dn
	___Vdm.dmCHencodeData9999I = make(chan _TdataTran, 50) // _TencodeX , used for normal data tunnel

	___Vdm.dmChSendIdleNoteInternalUSE = make(chan byte, 1) // a random timer , send idle note to main receive loop. internal use only.
	___Vdm.dmChSwapLoginCkInfoForLock = make(chan byte, 1)  // a 5s timer , send swap note to main receive loop. internal use only.
	___Vdm.dmChCheckTimeOutDieClient = make(chan byte, 1)   // a 80s timer , send swap note to main receive loop. internal use only.

	___Vdm.dmMconn.dcsMidx = make(map[[16]byte]int)
	___Vdm.dmMdata.ddsMidx = make(map[[16]byte]int)

	for __Vi := 0; __Vi < _VallowClientMax; __Vi++ {
		___Vdm.dmMconn.dcsMm[__Vi].dccConnPortStrMap = make(map[string]byte) // _TdataMachinEconnectClient
		___Vdm.dmMdata.ddsMm[__Vi].ddcConnPortStrMap = make(map[string]byte) // _TdataMachinEdataClient
	}

	___Vdm.dmMconn.dcsMfreeAmount = _VallowClientMax // _TdataMachinEconnectSt
	___Vdm.dmMdata.ddsMfreeAmount = _VallowClientMax // _TdataMachinEdataSt

	_Fsleep(_T1s)

	go _Frun(___Vdm, 1000201) // _FdataMachin__1000201x__receive__default // <-___Vdm.dmCHdataMachineIdI // <-___Vdm.dmCHdecodeDataI
	go _Frun(___Vdm, 1000501) // _FdataMachin__1000501x__swapLoginCkInfoForLock__timeoutGen // (*___Vdm.dmCHloginGenMachineIdLO)<-
	go _Frun(___Vdm, 1000502) // _FdataMachin__1000502x__dataSendIdle__gen_time_gap // (*___Vdm.dmCHloginGenMachineIdLO) <-
	go _Frun(___Vdm, 1000503) // _FdataMachin__1000503x__time_gap_dataResend
	go _Frun(___Vdm, 1000504) // _FdataMachin__1000504x__checkTimeOutDieClient
}

func (___Vdm *_TdataMachine) _FdataMachin__1000501x__swapLoginCkInfoForLock__timeoutGen() {
	for {
		_Fsleep(_T5s)
		___Vdm.
			dmChSwapLoginCkInfoForLock <- 1 // a 5s timer , send swap note to main receive loop. internal use only.
	}
}
func (___Vdm *_TdataMachine) _FdataMachin__1000502x__dataSendIdle__gen_time_gap() {

	for {
		_FsleepRand_8_to_12s() // _Vgap_connectLostTimeOut
		___Vdm.dmChSendIdleNoteInternalUSE <- 1
	}
}
func (___Vdm *_TdataMachine) _FdataMachin__1000504x__checkTimeOutDieClient() {
	for {
		_Fsleep(_Vgap_connectLostCheckDealy) // _Vgap_connectLostTimeOut
		___Vdm.
			dmChCheckTimeOutDieClient <- 1 // a 5s timer , send swap note to main receive loop. internal use only.
	}
}
