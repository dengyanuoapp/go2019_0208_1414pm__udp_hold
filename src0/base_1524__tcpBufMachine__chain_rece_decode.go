package main

// _TtcpBufCellX
// tbmCHtcpLocal2RemoteBI chan []byte // _TtcpNodE tnCHtcpReceBLO *chan []byte // byte of _TtcpNodeDataRece
func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__1500201y1__chan_rece__Local2Remote(___VbIN *[]byte) {

	__VtnRece := _TtcpNodeDataRece{}
	// _FdecGob___(___VeMsg string, ___VbyteIn []byte, ___VoutObjLp interface{}) error {
	__Verr2 := _FdecGob___(" 381811 01 ", *___VbIN, &__VtnRece)
	_FerrExit(" 381811 02 : ", __Verr2)

	_CFpfN(" 381811 03 : tcpBuf rece from local {%s} ", __VtnRece.String()) // _TtcpNodeDataRece
	if false == ___Vtbm._Finsert_local2remote_buf(&__VtnRece) {
		return
	}
}

// _TtcpBufferArrX _TtcpNodeDataRece
func (___Vtbm *_TtcpBufMachine) _Finsert_local2remote_buf(___VtnRece *_TtcpNodeDataRece) bool {
	__Vtunnel := ___Vtbm._FtcpBufMachine__findOrCreate_local2remote_tunnel(___VtnRece)
	if nil == __Vtunnel { // _TtcpBuftunnelX
		_CFpfN(" 381815 41 : rece _TtcpNodeDataRece error {%s} ", ___VtnRece.String())
		return false
	}

	__VtcNow := &(__Vtunnel.tbtL2R.tbCellNow) // _TtcpBufCell _TtcpBuFx from local to remote

	__FpfN(" 381815 43 ")
	__VremainLen := len(___VtnRece.TnrBuf)
	if __VremainLen > 1024 {
		_CFpfN(" 381815 45 : Len ERROR (%d){%s} ", __VremainLen, ___VtnRece.String())
		return false
	}

	__FpfN(" 381815 47 ")
	__VremainBuf := ___VtnRece.TnrBuf // _TtcpBuftunnelX

	if 0 == __VtcNow.tbcLen { // _TtcpBufCell
		_CFpfN(" 381816 51 : len:%d ", __VremainLen)
		__VtcNow.tbcLen = __VremainLen
		_CFpfN(" 381816 52 %d/%d", len(__VtcNow.tbcBuf), len(__VremainBuf))
		copy(__VtcNow.tbcBuf[:], __VremainBuf)
		_CFpfN(" 381816 53 : len:%d ", __VremainLen)
	} else {
		_CFpfN(" 381816 54 : len:%d ", __VremainLen)
		if __VtcNow.tbcLen+__VremainLen > 1024 { // _TtcpBufCell
			_CFpfN(" 381816 55 , extra-len , need one-more bufcell: len:%d / %d", __VremainLen, __VtcNow.tbcLen)
			__Vt := 1024 - __VtcNow.tbcLen
			copy(__VtcNow.tbcBuf[__VtcNow.tbcLen:], __VremainBuf[:__Vt]) // ok , a full bufCell is ready. put into queue

			__Vtunnel.tbtL2R.tbCellArr[__Vtunnel.tbtL2R.tbEnd] = *__VtcNow // _TtcpBufCell _TtcpBuFx from local to remote
			__Vtunnel.tbtL2R.tbEnd++
			__Vtunnel.tbtL2R.tbOffsetEnd = __Vtunnel.tbtL2R.tbOffsetNow
			__Vtunnel.tbtL2R.tbOffsetNow += 1024
			if __Vtunnel.tbtL2R.tbEnd >= ___Vtbm.tbmBufArr.tbaCntMax {
				__Vtunnel.tbtL2R.tbEnd = 0
			}
			__Vtunnel.tbtL2R.tbFreeCnt--

			__VtcNow.tbcLen = __VremainLen - __Vt
			copy(__VtcNow.tbcBuf[:], __VremainBuf[__Vt:]) // put the remain into the nowBufCell
		} else {
			_CFpfN(" 381816 57 , ok-len , no need bufcell: len:%d / %d", __VremainLen, __VtcNow.tbcLen)
			copy(__VtcNow.tbcBuf[__VtcNow.tbcLen:], __VremainBuf)
			__VtcNow.tbcLen += __VremainLen

		}
	}

	// _TtcpBufferArrX
	_CFpfN(" 381816 59 : after insert Buf (%d/%d) {%s}",
		__Vtunnel.tbtTunnelNo, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
	return true
}
