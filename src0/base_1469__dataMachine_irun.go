package main

func (___Vdm *_TdataMachine) IRun(___Vidx int) {
	switch ___Vidx {
	case 1000101:
		if nil == ___Vdm.dmCBinit {
			_FdataMachin__1000101__main_init__default(___Vdm)
		} else {
			___Vdm.dmCBinit(___Vdm)
		}
	default:
		_FpfNex(" 839191 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FdataMachin__1000101__main_init__default(___Vdm *_TdataMachine) {

	___Vdm.dmCHdataMachineIdI = make(chan _TdataMachinEid)
}
