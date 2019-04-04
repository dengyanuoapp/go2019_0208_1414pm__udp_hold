package main

type _TdataMachinEid struct {
	diConnPort _TudpConnPort
	diIdx128   []byte
	diSeq128   []byte
	diToken    []byte
}

type _TdataMachine struct {
	dmCHdataMachineIdI chan _TdataMachinEid
	dmCBinit           func(*_TdataMachine) // _FdataMachin__1000101__main_init__default
	dmCBrece           func(*_TdataMachine) //
}

func (___Vdi *_TdataMachinEid) String() string {
	return _Pspf(
		" to{%s} id:%x,%x tk:%x",
		___Vdi.diConnPort.String(),
		String5(&___Vdi.diIdx128),
		String5(&___Vdi.diSeq128),
		String5(&___Vdi.diToken),
	)
}
