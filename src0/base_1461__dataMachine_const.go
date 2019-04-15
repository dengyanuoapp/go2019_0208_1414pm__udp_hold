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
		"to:{%s} id:%s,%s tk:%s",
		___Vdi.diConnPort.String(),
		String5(&___Vdi.diIdx128),
		String5(&___Vdi.diSeq128),
		String5(&___Vdi.diToken),
	)
}

type _TdataMachinEconnMap struct {
	dmmID             _TdataMachinEid
	dmmConnPortArr    []_TudpConnPort
	dmmConnPortAmount int
	dmmLastReadTime   int
	dmmLastSendTime   int
	dmmLastPrTime     int
	dmmLastSendIdx    int
}

func (___Vdmem *_TdataMachinEconnMap) String() string {
	__Vlen := len(___Vdmem.dmmConnPortArr)
	__Vs := _Spf("eid:{%s} (%d)[", ___Vdmem.dmmID.String(), __Vlen)
	for __Vi := 0; __Vi < __Vlen; __Vi++ {
		if 0 == __Vi {
			__Vs += _Spf("%s", ___Vdmem.dmmConnPortArr[__Vi].String())
		} else {
			__Vs += _Spf(" | %s", ___Vdmem.dmmConnPortArr[__Vi].String())
		}
	}
	__Vs += _Spf("] amount:%d lastRead:%d lastSend:%d lastPr:%d",
		___Vdmem.dmmConnPortAmount,
		___Vdmem.dmmLastReadTime,
		___Vdmem.dmmLastSendTime,
		___Vdmem.dmmLastPrTime)
	return __Vs
}

type _TdataMachineConnSt struct {
	M        map[[16]byte]_TdataMachinEconnMap
	LockLast map[[16]byte]int
	LockNow  map[[16]byte]int
	mux      sync.Mutex
}

type _TdataMachine struct {
	dmCBinit                func(*_TdataMachine) // _FdataMachin__1000101__main_init__default
	dmCBrece                func(*_TdataMachine) //
	dmCBprReceKey           func(*_TdataMachine) //
	dmCBprSendKey           func(*_TdataMachine) //
	dmMconn                 _TdataMachineConnSt
	dmCHdataMachineIdI      chan _TdataMachinEid  // loginChecker will fill this chan when check-token ok.
	dmCHloginGenMachineIdLO *chan _TdataMachinEid // _VloginGenerator_Dn.lgCHdataMachineIdI, fill this chan to told the loginGen to stop
	dmCHdecodeDataI         chan _Tdecode         // from uTmCHdecodeDataLO  *chan _Tdecode of decoder
	dmCHencodeLO            *chan _Tencode        // _TencodeX
	//dmCH  unCHreceByteLO    *chan _TudpNodeDataRece // if nil , drop it ; not-nil , put the received data into this chan
	//dmCH  unCHsendI     chan _TudpNodeDataSend  // try get data from chan, then send it out.
}
