package main

const _VallowClientMax = 1000
const _VallowTunnelPerClientMax = 150
const _Vgap_nothingToLost  = 39 // 3*(12+1) == 39
const _Vgap_skip_idle_send = 7

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

type _TdataMachine struct {
	dmChSendIdleNoteInternalUSE chan byte            // a random timer , send idle note to main receive loop. internal use only.
	dmChSwapLoginCkInfoForLock  chan byte            // a 5s timer , send swap note to main receive loop. internal use only.
	dmChCheckTimeOutDieClient   chan byte            // a 5s timer , send swap note to main receive loop. internal use only.
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
