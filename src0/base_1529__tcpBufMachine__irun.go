package main

func (___Vtbm *_TtcpBufMachine) IRun(___Vidx int) {
	switch ___Vidx {
	case 1500101:
		if nil == ___Vtbm.tbmCBinit {
			go _FtcpBufMachine__1500101x__init(___Vtbm)
		} else {
			go ___Vtbm.
				tbmCBinit(___Vtbm)
		}
	case 1500201:
		go ___Vtbm.
			_FtcpBufMachine__1500201x__chan_rece__default()
	default:
		_FpfNex(" 834821 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FtcpBufMachine__1500101x__init(___Vtbm *_TtcpBufMachine) {
	___Vtbm.tbmCHtcpReceBI = make(chan []byte, 50)

	switch _VS.RoleName {
	case "Cn":
		___Vtbm.tbmBufArr.tbaCntMax = 100 // Cn : 100 tunnel
	case "Dn": // _TtcpBufferArrX
		___Vtbm.tbmBufArr.tbaCntMax = 1000 // Cn : 1000 tunnel
	case "Fn":
		// do nothing.
	default:
		_FpfNex(" 834821 03 : unknown Role ")
	}

	___Vtbm.tbmBufArr.tbaCntFree = ___Vtbm.tbmBufArr.tbaCntMax
	___Vtbm.tbmBufArr.tbaMbuftunnel = make([]_TtcpBuftunnel, ___Vtbm.tbmBufArr.tbaCntMax)
	___Vtbm.tbmBufArr.tbaMtid = make(map[[16]byte]int)
	for __Vi := 0; __Vi < ___Vtbm.tbmBufArr.tbaCntMax; __Vi++ { // _TtcpBuFx
		___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi].tbtL2R.tbFreeCnt = 4096 - 3
		___Vtbm.tbmBufArr.tbaMbuftunnel[__Vi].tbtR2L.tbFreeCnt = 4096 - 3
	}

	_Fsleep(_T1s)

	go _Frun(___Vtbm, 1500201) // _FtcpBufMachine__1500201x__chan_rece__default
}
