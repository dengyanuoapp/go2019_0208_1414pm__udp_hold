package main

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__sendCmdRemote2Local(___Vk16 [16]byte, ___Vcmd byte) {
	// tbmCHtcpRemote2LocalCmdLO
	if nil == ___Vtbm.tbmCHtcpRemote2LocalCmdLO {
		return
	}
	if 0 == cap(*___Vtbm.tbmCHtcpRemote2LocalCmdLO) {
		_FpfNex(" ###### 398387 01 : _TtcpBufMachine _FtcpBufMachine__sendCmdRemote2Local why chan ZERO len ?")
		return
	}

	__V17 := [17]byte{}
	copy(__V17[:], ___Vk16[:])
	__V17[16] = ___Vcmd

	(*(___Vtbm.tbmCHtcpRemote2LocalCmdLO)) <- __V17
}

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__checkAndDeleteTimeoutTunnel(__Vb byte) {
	__FpfN(" ###### 398388 01 : _TtcpBufMachine _FtcpBufMachine__checkAndDeleteTimeoutTunnel %d ", __Vb)

	__Vt := _FtimeInt()
	__VkArr := [][16]byte{}
	for __Vk16, __Vid := range ___Vtbm.tbmBufArr.tbaMtid {
		__Vtunel := &(___Vtbm.tbmBufArr.tbaMbuftunnel[__Vid]) // _TtcpBuftunnelX
		if __Vt-__Vtunel.tbtLastALL > 20 {
			__VkArr = append(__VkArr, __Vk16)
			___Vtbm.tbmBufArr.tbaMbuftunnel[__Vid] = _TtcpBuftunnel{}
		}
	}

	for _, __Vk16 := range __VkArr {
		delete(___Vtbm.tbmBufArr.tbaMtid, __Vk16)
		___Vtbm.tbmBufArr.tbaCntUsed--
		___Vtbm.tbmBufArr.tbaCntFree++
		___Vtbm.
			_FtcpBufMachine__sendCmdRemote2Local(__Vk16, TcpNodeCmd__ErrTimeout)
	}

	if 0 < len(__VkArr) {
		_CFpfN(" ###### 398388 09 : _TtcpBufMachine _FtcpBufMachine__checkAndDeleteTimeoutTunnel delete %d {%s}",
			len(__VkArr), ___Vtbm.tbmBufArr.String())
	}
}

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

	_CFpfN(" 381819 03 : TcpNodeCmd : EOF , try delete : %x", __Vk16[:3])

	__Vi3, __Vok3 := ___Vtbm.tbmBufArr.tbaMtid[__Vk16]

	if false == __Vok3 {
		_CFpfN(" 381819 05 : TcpNodeCmd : why not found in the tunnel arr ? : %x", __Vk16[:3])
		return
	}

	___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3] = _TtcpBuftunnel{} // _TtcpBuftunnelX
	___Vtbm.tbmBufArr.tbaCntUsed--
	___Vtbm.tbmBufArr.tbaCntFree++
	delete(___Vtbm.tbmBufArr.tbaMtid, __Vk16)
}

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__findOrCreate_local2remote_tunnel(___VtnRece *_TtcpNodeDataRece) *_TtcpBuftunnel {
	if nil == ___Vtbm || ___Vtbm.tbmBufArr.tbaCntMax <= 0 {
		_CFpfN(" 381812 11 : no socket avaiable free:%d max:%d , {%s} ", ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
		//___Vtbm.  _FtcpBufMachine__sendCmdRemote2Local(__Vk16, TcpNodeCmd__ErrNoTunnel)
		return nil
	}

	if nil == ___VtnRece || 16 != len(___VtnRece.TnrId128) || 0 == len(___VtnRece.TnrBuf) {
		_CFpfN(" 381812 12 : rece _TtcpNodeDataRece error {%s} ", ___VtnRece.String())
		//___Vtbm.  _FtcpBufMachine__sendCmdRemote2Local(__Vk16, TcpNodeCmd__ErrR2Ldata)
		return nil
	}

	var __Vk16 [16]byte
	copy(__Vk16[:], ___VtnRece.TnrId128[:16])
	//_FcopyByte(&__Vk16, &___VtnRece.TnrId128, 16)

	if ___Vtbm.tbmBufArr.tbaCntFree <= 3 {
		_CFpfN(" 381812 13 : no free Buf socket avaiable free:%d max:%d , {%s} ",
			___Vtbm.tbmBufArr.tbaCntFree, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
		___Vtbm.
			_FtcpBufMachine__sendCmdRemote2Local(__Vk16, TcpNodeCmd__ErrNoTunnel)
		return nil
	}

	__FpfN(" 381812 14 : free:%d max:%d , {%s} ",
		___Vtbm.tbmBufArr.tbaCntFree, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())

	__Vi3, __Vok3 := ___Vtbm.tbmBufArr.tbaMtid[__Vk16]
	if false == __Vok3 {

		if 0 != ___VtnRece.TnrOffset { // _TtcpNodeDataRece
			_CFpfN(" 381813 21 : why not ZERO offset (%d) ? close this tunnel. ",
				___VtnRece.TnrOffset)
			___Vtbm.
				_FtcpBufMachine__sendCmdRemote2Local(__Vk16, TcpNodeCmd__ErrOffset)
			return nil
		}

		if ___Vtbm.tbmBufArr.tbaCntFree <= 4 { // not 3 , if insert , must 4 avaiable
			_CFpfN(" 381813 22 : no free Buf socket avaiable free:%d max:%d , {%s} ",
				___Vtbm.tbmBufArr.tbaCntFree, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
			___Vtbm.
				_FtcpBufMachine__sendCmdRemote2Local(__Vk16, TcpNodeCmd__ErrNoTunnel)
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
			___Vtbm.
				_FtcpBufMachine__sendCmdRemote2Local(__Vk16, TcpNodeCmd__ErrNoTunnel)
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
		___Vtbm.tbmBufArr.tbaCntUsed++
		___Vtbm.tbmBufArr.tbaCntFree--
	} else {
		_CFpfN(" 381814 31 : alread exist.")
	}

	__FpfN(" 381814 33 ")
	if 0 > __Vi3 || ___Vtbm.tbmBufArr.tbaCntMax <= __Vi3 {
		_CFpfN(" 381814 35 : idx error {%d/%d} ",
			__Vi3, ___Vtbm.tbmBufArr.tbaCntMax, ___Vtbm.tbmBufArr.String())
		___Vtbm.
			_FtcpBufMachine__sendCmdRemote2Local(__Vk16, TcpNodeCmd__ErrNoTunnel)
		return nil
	}

	__FpfN(" 381814 37 ")
	___Vtunnel := &(___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi3])

	___Vtunnel.tbtLastL2R = _FtimeInt()
	___Vtunnel.tbtLastALL = ___Vtunnel.tbtLastL2R

	return ___Vtunnel
}
