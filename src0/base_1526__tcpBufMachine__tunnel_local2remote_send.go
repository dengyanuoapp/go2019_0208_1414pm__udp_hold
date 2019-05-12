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
	__FpfNdb(" ###### 398385 02 : %x ", ___Vtunnel.tbtTidx) // _TtcpBuftunnelX
	___VtunnelL2R := &(___Vtunnel.tbtL2R)                   // _TtcpBuFx
	if 0 != ___VtunnelL2R.tbCellNow.tbcLen {                // _TtcpBufCell
		___VtunnelL2R.tbCellArr[___VtunnelL2R.tbOffsetEnd].tbcOffset = ___VtunnelL2R.tbCellNow.tbcOffset
		___VtunnelL2R.tbCellArr[___VtunnelL2R.tbOffsetEnd].tbcLen = ___VtunnelL2R.tbCellNow.tbcLen

		copy(___VtunnelL2R.tbCellArr[___VtunnelL2R.tbOffsetEnd].tbcBuf[:],
			___VtunnelL2R.tbCellNow.tbcBuf[:___VtunnelL2R.tbCellNow.tbcLen])

		___VtunnelL2R.tbOffsetEnd++
		// _TtcpBufCell
		if 4096 == ___VtunnelL2R.tbOffsetEnd {
			___VtunnelL2R.tbOffsetEnd = 0
		}
		___VtunnelL2R.tbCellNow.tbcOffset += int64(___VtunnelL2R.tbCellNow.tbcLen)
		___VtunnelL2R.tbCellNow.tbcLen = 0 // _TtcpBufCell

		_CFpfNdb(" ###### 398385 05 : %x ", ___Vtunnel.tbtTidx) // _TtcpBuftunnelX
	}

}
