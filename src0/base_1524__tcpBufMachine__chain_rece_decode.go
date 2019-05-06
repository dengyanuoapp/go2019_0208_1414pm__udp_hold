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

	if ___Vtbm.tbmBufArr.tbaCntFree <= 0 || ___Vtbm.tbmBufArr.tbaCntMax <= 0 {
		_CFpfN(" 381812 02 : no free Buf socket avaiable free:%d max:%d , {%s} ",
			___Vtbm.tbmBufArr.tbaCntFree, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
		return false
	}

	_CFpfN(" 381812 03 : free:%d max:%d , {%s} ",
		___Vtbm.tbmBufArr.tbaCntFree, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())

	var __Vk16 [16]byte
	//_FcopyByte(&__Vk16, &___VtnRece.TnrId128, 16)
	copy(__Vk16[:], ___VtnRece.TnrId128[:16])

	__Vi3, __Vok3 := ___Vtbm.tbmBufArr.tbaMtid[__Vk16]
	if false == __Vok3 {
		for __Vi3 = 0; __Vi3 < ___Vtbm.tbmBufArr.tbaCntMax; __Vi3++ {
			_CFpfN(" 381813 05 vi:%d max:%d ", __Vi3, ___Vtbm.tbmBufArr.tbaCntMax)
			if false == ___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3].tbtEna {
				_CFpfN(" 381813 06 vi:%d max:%d ", __Vi3, ___Vtbm.tbmBufArr.tbaCntMax)
				break
			}
		}
		if __Vi3 >= ___Vtbm.tbmBufArr.tbaCntMax {
			_FpfNex(" 381813 07 : why not found free buffer socket ? ")
			return false
		}
		_CFpfN(" 381813 08 : new create ")
		___Vtbm.tbmBufArr.tbaMtid[__Vk16] = __Vi3
		___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3] = _TtcpBuftunnel{ // _TtcpBuftunnelX
			tbtEna:  true,
			tbtTidx: __Vk16, // tunnel ID
		}
	} else {
		_CFpfN(" 381813 09 : alread exist.")
	}

	__FpfN(" 381814 01 ")
	if 0 > __Vi3 || ___Vtbm.tbmBufArr.tbaCntMax <= __Vi3 {
		_CFpfN(" 381814 02 : idx error {%d/%d} ",
			__Vi3, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
		return false
	}

	__FpfN(" 381814 03 ")
	__Vtunnel := &(___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3])

	__VtcNow := &(__Vtunnel.tbtL2R.tbCellNow) // _TtcpBufCell _TtcpBuFx from local to remote

	_FpfN(" 381814 04 ")
	__VremainLen := len(___VtnRece.TnrBuf)
	if __VremainLen > 1024 {
		_CFpfN(" 381814 05 : Len ERROR (%d){%s} ", __VremainLen, ___VtnRece.String())
		return false
	}

	__FpfN(" 381814 07 ")
	__VremainBuf := ___VtnRece.TnrBuf

	if 0 == __VtcNow.tbcLen { // _TtcpBufCell
		__FpfN(" 381814 08 ")
		__VtcNow.tbcLen = __VremainLen
		__FpfN(" 381814 09 %d/%d", len(__VtcNow.tbcBuf), len(__VremainBuf))
		copy(__VtcNow.tbcBuf[:], __VremainBuf)
		__FpfN(" 381814 10 ")
	} else {
		__FpfN(" 381814 11 ")
	}

	// _TtcpBufferArrX
	_CFpfN(" 381814 29 : after insert Buf (%d/%d) {%s}",
		__Vi3, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
	return true
}
