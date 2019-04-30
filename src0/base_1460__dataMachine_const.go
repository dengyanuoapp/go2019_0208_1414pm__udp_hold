package main

const _VallowClientMax = 1000
const _VallowTunnelPerClientMax = 150

// send idle == 8-12, so , check == 12*4==36 == every 39 second check per time , if 2 check lost , connect lost (25)
//const _Vgap_connectLostTimeOut = 39 // 3*(12+1) == 39
//const _Vgap_connectLostTimeOut = 52 // 4*(12+1) == 52
//const _Vgap_connectLostTimeOut = 78 // 4*(12+1) == 52
//const _Vgap_connectLostCheckDealy = _T25s
const _Vgap_connectLostCheckDealy = _T15s
const _Vgap_connectLostTimeOut = 65 // 4*(12+1) == 52
const _Vgap_skip_idle_send = 7

type _TdataMachinEid struct {
	diRole     string
	diConnPort _TudpConnPort
	diIdx128   []byte
	diSeq128   []byte
	diToken    []byte
}

func (___Vdi *_TdataMachinEid) String() string {
	return _Spf(
		"W:%s to:{%s} id:%s,%s tk:%s",
		___Vdi.diRole,
		___Vdi.diConnPort.String(),
		String5s(&___Vdi.diIdx128),
		String5s(&___Vdi.diSeq128),
		String5s(&___Vdi.diToken),
	)
}

type _TdataMachine struct {
	dmChSendIdleNoteInternalUSE_sendIdleKeep  chan byte            // a random timer , send idle note to main receive loop. internal use only.
	dmChSwapLoginCkInfoForLock   chan byte            // a 5s timer , send swap note to main receive loop. internal use only.
	dmChCheckTimeOutDieClient    chan byte            // a 5s timer , send swap note to main receive loop. internal use only.
	dmCBinit                     func(*_TdataMachine) // _FdataMachin__1000101__main_init__default
	dmCBrece                     func(*_TdataMachine) //
	dmCBprReceKey                func(*_TdataMachine) //
	dmCBprSendKey                func(*_TdataMachine) //
	dmCBdebugInfo                func(*_TdataMachine) //
	dmMconn                      _TdataMachinEconnectSt
	dmMdata                      _TdataMachinEdataSt
	dmCHdataMachineIdI           chan _TdataMachinEid  // loginChecker will fill this chan when check-token ok.
	dmCHloginGenMachineIdLO      *chan _TdataMachinEid // _VloginGenerator_Dn.lgCHdataMachineIdI, fill this chan to told the loginGen to stop
	dmCHdecodeDataI              chan _Tdecode         // from uDeCHdecodeDataLO  *chan _Tdecode of decoder
	dmCHrepackDecodeC2sLO        *chan _Tdecode        // _TencodeX , get repacked-c2s decode as spec-data , then send to loginChecker
	dmCHencodeIdleLO             *chan _Tencode        // _TencodeX , get encode  repack to idle package ,then send
	dmCHencodeDataSpecBLO        *chan _TdataTran      // _TencodeX , used for Fn-Dn, and other special use
	dmCHencodeData9999BLO        *chan _TdataTran      // _TencodeX , used for normal data tunnel
	dmCHencodeDataSpecFnWaitCnBI chan []byte           // byte of _TencodeX , used for Fn-Dn, and other special use
	dmCHencodeData9999BI         chan []byte           // byte of _TencodeX , used for normal data tunnel
	dmCHdebugInfoI               chan byte             // when received , output the debug info
	//dmCH  unCHreceByteLO    *chan _TudpNodeDataRece // if nil , drop it ; not-nil , put the received data into this chan
	//dmCH  unCHsendI     chan _TudpNodeDataSend  // try get data from chan, then send it out.
}
