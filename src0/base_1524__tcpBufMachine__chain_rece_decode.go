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

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__findOrCreate_local2remote_tunnel(___VtnRece *_TtcpNodeDataRece) *_TtcpBuftunnel {
	if nil == ___Vtbm || ___Vtbm.tbmBufArr.tbaCntMax <= 0 {
		_CFpfN(" 381812 11 : no socket avaiable free:%d max:%d , {%s} ", ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
		return nil
	}

	if nil == ___VtnRece || 16 != len(___VtnRece.TnrId128) || 0 == len(___VtnRece.TnrBuf) {
		_CFpfN(" 381812 12 : rece _TtcpNodeDataRece error {%s} ", ___VtnRece.String())
		return nil
	}

	if ___Vtbm.tbmBufArr.tbaCntFree <= 3 {
		_CFpfN(" 381812 13 : no free Buf socket avaiable free:%d max:%d , {%s} ",
			___Vtbm.tbmBufArr.tbaCntFree, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
		return nil
	}

	__FpfN(" 381812 14 : free:%d max:%d , {%s} ",
		___Vtbm.tbmBufArr.tbaCntFree, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())

	var __Vk16 [16]byte
	copy(__Vk16[:], ___VtnRece.TnrId128[:16])
	//_FcopyByte(&__Vk16, &___VtnRece.TnrId128, 16)

	__Vi3, __Vok3 := ___Vtbm.tbmBufArr.tbaMtid[__Vk16]
	if false == __Vok3 {
		if ___Vtbm.tbmBufArr.tbaCntFree <= 4 { // not 3 , if insert , must 4 avaiable
			_CFpfN(" 381813 21 : no free Buf socket avaiable free:%d max:%d , {%s} ",
				___Vtbm.tbmBufArr.tbaCntFree, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
			return nil
		}

		for __Vi3 = 0; __Vi3 < ___Vtbm.tbmBufArr.tbaCntMax; __Vi3++ {
			_CFpfN(" 381813 23 vi:%d max:%d ", __Vi3, ___Vtbm.tbmBufArr.tbaCntMax)
			if false == ___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3].tbtEna {
				_CFpfN(" 381813 25 vi:%d max:%d ", __Vi3, ___Vtbm.tbmBufArr.tbaCntMax)
				break
			}
		}
		if __Vi3 >= ___Vtbm.tbmBufArr.tbaCntMax {
			_FpfNex(" 381813 27 : why not found free buffer socket ? ")
			return nil
		}
		_CFpfN(" 381813 29 : new create ")
		___Vtbm.tbmBufArr.tbaMtid[__Vk16] = __Vi3
		___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3] = _TtcpBuftunnel{ // _TtcpBuftunnelX
			tbtEna:  true,
			tbtTidx: __Vk16, // tunnel ID
			tbtL2R: _TtcpBuF{ // from local to remote // _TtcpBuFx
				tbFreeCnt: 4096 - 3,
			},
			tbtR2L: _TtcpBuF{
				tbFreeCnt: 4096 - 3,
			},
		}
	} else {
		_CFpfN(" 381814 31 : alread exist.")
	}

	__FpfN(" 381814 33 ")
	if 0 > __Vi3 || ___Vtbm.tbmBufArr.tbaCntMax <= __Vi3 {
		_CFpfN(" 381814 35 : idx error {%d/%d} ",
			__Vi3, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
		return nil
	}

	__FpfN(" 381814 37 ")
	___Vtunnel := &(___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3])
	return ___Vtunnel
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
