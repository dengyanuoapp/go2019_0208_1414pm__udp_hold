package main

import (
	"sync"
)

type _TdataMachinEid struct {
	diConnPort _TudpConnPort
	diIdx128   []byte
	diSeq128   []byte
	diToken    []byte
}

func (___Vdi *_TdataMachinEid) String() string {
	return _Spf(
		" to{%s} id:%x,%x tk:%x",
		___Vdi.diConnPort.String(),
		String5(&___Vdi.diIdx128),
		String5(&___Vdi.diSeq128),
		String5(&___Vdi.diToken),
	)
}

type _TdataMachinEconnMap struct {
	dmmID          _TdataMachinEid
	dmmConnPortArr []_TudpConnPort
	dmmLastAccTime int
}

type _TdataMachineConnSt struct {
	M        map[[16]byte]_TdataMachinEconnMap
	LockLast map[[16]byte]int
	LockNow  map[[16]byte]int
	mux      sync.Mutex
}

type _TdataMachine struct {
	dmCHdataMachineIdI      chan _TdataMachinEid
	dmCHloginGenMachineIdLO *chan _TdataMachinEid // _VloginGenerator_Dn.lgCHdataMachineIdI,
	//dmCHloginGenMachineIdLO *chan [16]byte       // _VloginGenerator_Dn.lgCHdataMachineIdI,
	dmCBinit func(*_TdataMachine) // _FdataMachin__1000101__main_init__default
	dmCBrece func(*_TdataMachine) //
	dmMconn  _TdataMachineConnSt
}
