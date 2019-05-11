package main

func (___Vtbm *_TtcpBufMachine) tcpBufMachine__tunnel_local2remote_send_to_dataMachine_top() {
	if nil == ___Vtbm.tbmCHtoDataMachineBLO {
		return
	}
	if 0 == cap(*___Vtbm.tbmCHtoDataMachineBLO) {
		_FpfNex(" ###### 398384 01 : _TtcpBufMachine tcpBufMachine__tunnel_local2remote_send_to_dataMachine_top why chan ZERO len ?")
		return
	}

	if 0 == ___Vtbm.tbmBufArr.tbaCntUsed {
		__FpfN(" 398384 02 : _TtcpBufMachine tcpBufMachine__tunnel_local2remote_send_to_dataMachine_top no tunnel now .")
		return
	}

	for _, __Vidx := range ___Vtbm.tbmBufArr.tbaMtid { // _TtcpBufferArrX
		__Vtunnel := &(___Vtbm.tbmBufArr.tbaMbuftunnel[__Vidx])
		___Vtbm.tcpBufMachine__tunnel_local2remote_send_to_dataMachine_eachTunnel(__Vtunnel)
	}
}

func (___Vtbm *_TtcpBufMachine) tcpBufMachine__tunnel_local2remote_send_to_dataMachine_eachTunnel(___Vtunnel *_TtcpBuftunnel) {
	//_CFpfN(" ###### 398385 01 : _TtcpBufMachine _FtcpBufMachine__checkAndDeleteTimeoutTunnel %x ", ___Vtunnel.tbtTidx)
	_CFpfNdb(" ###### 398385 02 : %x ", ___Vtunnel.tbtTidx)

}
