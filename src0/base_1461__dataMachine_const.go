package main

type _TdataMachinEid struct {
	diConnPort _TudpConnPort
}

type _TdataMachine struct {
	dmCHdataMachineIdI chan _TdataMachinEid
}

func (___Vdi *_TdataMachinEid) String() string {
	return _Pspf(
		" to{%s} ",
		___Vdi.diConnPort.String())
}
