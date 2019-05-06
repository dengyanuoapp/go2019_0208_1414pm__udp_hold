package main

// _TtcpBufCellX
// tbmCHtcpReceBI chan []byte // _TtcpNodE tnCHtcpReceBLO *chan []byte // byte of _TtcpNodeDataRece
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
	if nil == ___VtnRece || 16 != len(___VtnRece.TnrId128) || 0 == len(___VtnRece.TnrBuf) {
		_CFpfN(" 381812 01 : rece _TtcpNodeDataRece error {%s} ", ___VtnRece.String())
		return false
	}

	if ___Vtbm.tbmBufArr.tbaCntFree <= 0 {
		_CFpfN(" 381812 02 : no free Buf socket avaiable {%s} ", ___Vtbm.tbmBufArr.String())
		return false
	}

	var __Vk16 [16]byte
	//_FcopyByte(&__Vk16, &___VtnRece.TnrId128, 16)
	copy(__Vk16[:], ___VtnRece.TnrId128[:16])

	__Vi3, __Vok3 := ___Vtbm.tbmBufArr.tbaMtid[__Vk16]
	if false == __Vok3 {
		for __Vi3 = 0; __Vi3 < ___Vtbm.tbmBufArr.tbaCntMax; __Vi3++ {
			if false == ___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3].tbtEna {
				break
			}
			if __Vi3 >= ___Vtbm.tbmBufArr.tbaCntMax {
				_FpfNex(" 381812 03 : why not found free buffer socket ? ")
				return false
			}
			___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3] = _TtcpBuftunnel{ // _TtcpBuftunnelX
				tbtEna:  true,
				tbtTidx: __Vk16, // tunnel ID
			}
		}
	}

	if 0 > __Vi3 || ___Vtbm.tbmBufArr.tbaCntMax <= __Vi3 {
		_CFpfN(" 381812 04 : idx error {%d/%d} ",
			__Vi3, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
		return false
	}

	__Vtunnel := &(___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3])

	__VtcNow := &(__Vtunnel.tbtL2R.tbCellNow) // _TtcpBufCell _TtcpBuFx from local to remote

	__VremainLen := len(___VtnRece.TnrBuf)
	if __VremainLen > 1024 {
		_CFpfN(" 381812 05 : Len ERROR (%d){%s} ", __VremainLen, ___VtnRece.String())
		return false
	}

	__VremainBuf := ___VtnRece.TnrBuf

	if 0 == __VtcNow.tbcLen { // _TtcpBufCell
		__VtcNow.tbcLen = __VremainLen
		copy(__VtcNow.tbcBuf[:], __VremainBuf)
	} else {
	}

	_CFpfN(" 381812 09 : after insert Buf {%d/%d} ",
		__Vi3, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
	return true
}
