package main

type _TdataMachinEid struct {
	diConnPort _TudpConnPort
}

type _TdataMachine struct {
	dmCHdataMachineIdI chan _TdataMachinEid
	dmCBinit           func(*_TdataMachine) // if nil , use the default init procedure // _FudpDecode__700101x__init__default
}

func (___Vdi *_TdataMachinEid) String() string {
	return _Pspf(
		" to{%s} ",
		___Vdi.diConnPort.String())
}
