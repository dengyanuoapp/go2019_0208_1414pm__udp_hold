package main

import (
	"sync"
)

var (
	_Vgap_nothingToLost  = 39 // 3*(12+1) == 39
	_Vgap_skip_idle_send = 7
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

func (___Vdmcc *_TdataMachinEconnectClient) String() string {
	__Vlen := len(___Vdmcc.dccConnPortArr)
	__Vs := _Spf("eid:{%s} (%d)[", ___Vdmcc.dccID.String(), __Vlen)
	for __Vi := 0; __Vi < __Vlen; __Vi++ {
		if 0 == __Vi {
			__Vs += _Spf("%s", ___Vdmcc.dccConnPortArr[__Vi].String())
		} else {
			__Vs += _Spf(" | %s", ___Vdmcc.dccConnPortArr[__Vi].String())
		}
	}
	__Vs += _Spf("] amount:%d lastRead:%d lastPr:%d cntLast:%d cntNow:%d lastSendIdx:%d ",
		___Vdmcc.dccConnPortAmount,
		___Vdmcc.dccLastReceTime,
		___Vdmcc.dccLastPrTime,
		___Vdmcc.dccLockCntLast,
		___Vdmcc.dccLockCntLast,
		___Vdmcc.dccLastSendIdx,
	)
	return __Vs
}
func (___Vdmdc *_TdataMachinEdataClient) String() string {
	__Vlen := len(___Vdmdc.ddcConnPortArr)
	__Vs := _Spf("eid:{%s} (%d)[", ___Vdmdc.ddcID.String(), __Vlen)
	for __Vi := 0; __Vi < __Vlen; __Vi++ {
		if 0 == __Vi {
			__Vs += _Spf("%s", ___Vdmdc.ddcConnPortArr[__Vi].String())
		} else {
			__Vs += _Spf(" | %s", ___Vdmdc.ddcConnPortArr[__Vi].String())
		}
	}
	__Vs += _Spf("] amount:%d lastRead:%d lastSend:%d lastSendIdx:%d ",
		___Vdmdc.ddcConnPortAmount,
		___Vdmdc.ddcLastReceTime,
		___Vdmdc.ddcLastSendTime,
		___Vdmdc.ddcLastSendIdx,
	)
	return __Vs
}

func (___Vdmdc *_TdataMachinEdataClient) Clear() {
	(*___Vdmdc) = _TdataMachinEdataClient{}
	___Vdmdc.ddcConnPortStrMap = make(map[string]byte)
}
func (___Vdmcc *_TdataMachinEconnectClient) Clear() {
	(*___Vdmcc) = _TdataMachinEconnectClient{}
	___Vdmcc.dccConnPortStrMap = make(map[string]byte)
}

type _TdataMachinEconnectClient struct {
	dccID             _TdataMachinEid
	dccConnPortArr    []_TudpConnPort
	dccConnPortStrMap map[string]byte
	dccConnPortAmount int
	dccLastReceTime   int
	dccLastPrTime     int
	dccLockCntLast    int
	dccLockCntNow     int
	dccLastSendIdx    int
}
type _TdataMachinEdataClient struct {
	ddcID             _TdataMachinEid
	ddcConnPortArr    []_TudpConnPort
	ddcConnPortStrMap map[string]byte
	ddcConnPortAmount int
	ddcLastReceTime   int
	ddcLastSendTime   int
	ddcLastSendIdx    int
}

const _VallowClientMax = 1000
const _VallowTunnelPerClientMax = 150

type _TdataMachinEconnectSt struct {
	dcsMidx        map[[16]byte]int
	dcsMm          [_VallowClientMax]_TdataMachinEconnectClient
	dcsMusedAmount int
	dcsMfreeAmount int
	dcsMlastInsIdx int // the last insert place , the next insert can be start here
	dcsMux         sync.Mutex
}
type _TdataMachinEdataSt struct {
	ddsMidx        map[[16]byte]int
	ddsMm          [_VallowClientMax]_TdataMachinEdataClient
	ddsMusedAmount int
	ddsMfreeAmount int
	ddsMlastInsIdx int // the last insert place , the next insert can be start here
	ddsMux         sync.Mutex
}

type _TdataMachine struct {
	dmChSendIdleNoteInternalUSE chan byte            // a random timer , send idle note to main receive loop. internal use only.
	dmChSwapLoginCkInfoForLock  chan byte            // a 5s timer , send swap note to main receive loop. internal use only.
	dmCBinit                    func(*_TdataMachine) // _FdataMachin__1000101__main_init__default
	dmCBrece                    func(*_TdataMachine) //
	dmCBprReceKey               func(*_TdataMachine) //
	dmCBprSendKey               func(*_TdataMachine) //
	dmMconn                     _TdataMachinEconnectSt
	dmMdata                     _TdataMachinEdataSt
	dmCHdataMachineIdI          chan _TdataMachinEid  // loginChecker will fill this chan when check-token ok.
	dmCHloginGenMachineIdLO     *chan _TdataMachinEid // _VloginGenerator_Dn.lgCHdataMachineIdI, fill this chan to told the loginGen to stop
	dmCHdecodeDataI             chan _Tdecode         // from uDeCHdecodeDataLO  *chan _Tdecode of decoder
	dmCHencodeLO                *chan _Tencode        // _TencodeX
	//dmCH  unCHreceByteLO    *chan _TudpNodeDataRece // if nil , drop it ; not-nil , put the received data into this chan
	//dmCH  unCHsendI     chan _TudpNodeDataSend  // try get data from chan, then send it out.
}
