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
		_FdataMachin__1000501x__time_gap_dataChanLock(___Vdm)
	case 1000502:
		_FdataMachin__1000502x__time_gap_dataSendIdle(___Vdm)
	case 1000503:
		_FdataMachin__1000503x__time_gap_dataResend(___Vdm)
	default:
		_FpfNex(" 839191 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FdataMachin__1000101__main_init__default(___Vdm *_TdataMachine) {

	___Vdm.dmCHdataMachineIdI = make(chan _TdataMachinEid)

	___Vdm.dmMconn.M = make(map[[16]byte]_TdataMachinEconnMap)
	___Vdm.dmMconn.LockNow = make(map[[16]byte]int)
	___Vdm.dmMconn.LockLast = make(map[[16]byte]int)

	_Fsleep(_T1s)

	go _Frun(___Vdm, 1000201) // _FdataMachin__1000201x__receive__default
	go _Frun(___Vdm, 1000501) // _FdataMachin__1000501x__time_gap_dataChanLock
	go _Frun(___Vdm, 1000502) // _FdataMachin__1000502x__time_gap_dataSendIdle
	go _Frun(___Vdm, 1000503) // _FdataMachin__1000503x__time_gap_dataResend
}
