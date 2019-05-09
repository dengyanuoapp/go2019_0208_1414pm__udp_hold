package main

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__local2remote_remove_an_tunnel(__Vb *[17]byte) {

	switch __Vb[16] {
	case TcpNodeCmd__Eof:
		// ok
	default: // TcpNodeCmd__NULL
		_CFpfN(" 381819 01 : no such a TcpNodeCmd ")
		return
	}

	var __Vk16 [16]byte
	copy(__Vk16[:], (*__Vb)[:16])

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
