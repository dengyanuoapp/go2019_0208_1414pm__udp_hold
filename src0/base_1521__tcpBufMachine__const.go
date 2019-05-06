package main

type _TtcpBufMachine struct {
	tbmCHtcpReceBI chan []byte // _TtcpNodE tnCHtcpReceBLO *chan []byte // byte of _TtcpNodeDataRece
	tbmCHoutLO     *chan []byte
	tbmCBinit      func(*_TtcpBufMachine)
	tbmBufArr      _TtcpBufferArr // _TtcpBufferArrX

}

type _TtcpBufCellX struct{}
type _TtcpBufCell struct {
	tbcBuf [1024]byte
	tbcLen int
}

func (___Vtbc *_TtcpBufCell) String() string {
	__Vs := _Spf("(%d)[%s]", ___Vtbc.tbcLen, ___Vtbc.tbcBuf[:3])
	return __Vs
}

type _TtcpBuFx struct{}
type _TtcpBuF struct {
	tbCellArr [4096]_TtcpBufCell
	tbStart   int
	tbEnd     int
	tbCellNow _TtcpBufCell
}

func (___Vtbf *_TtcpBuF) String() string {
	__Vlen := 0
	if ___Vtbf.tbStart != ___Vtbf.tbEnd {
		__Vst := ___Vtbf.tbStart
		for __Vst != ___Vtbf.tbEnd {
			__Vlen += ___Vtbf.tbCellArr[__Vst].tbcLen
			__Vst++
			if __Vst >= 4096 {
				__Vst = 0
			}
		}
	}
	__Vs := _Spf("now{%s} st/end:%d/%d len:%d",
		___Vtbf.tbCellNow.String(), ___Vtbf.tbStart, ___Vtbf.tbEnd, __Vlen)
	return __Vs
}

type _TtcpBuftunnelX struct{}
type _TtcpBuftunnel struct {
	tbtEna      bool
	tbtTunnelNo byte     //  Tunnel Number : 0 -> normal , 1 ... 255 -> special channel
	tbtTidx     [16]byte // tunnel ID
	tbtL2R      _TtcpBuF // from local to remote
	tbtR2L      _TtcpBuF // from remote to local
}

func (___Vtbt *_TtcpBuftunnel) String() string {
	__Vs := _Spf("%d:%x\n:l2r[%s]\n:r2l[%s]\n", ___Vtbt.tbtTunnelNo, ___Vtbt.tbtTidx[:2],
		___Vtbt.tbtL2R.String(), ___Vtbt.tbtR2L.String())
	return __Vs
}

type _TtcpBufferArrX struct{}
type _TtcpBufferArr struct {
	tbaCntMax     int
	tbaCntUsed    int
	tbaCntFree    int
	tbaMtid       map[[16]byte]int // tunnel ID
	tbaMbuftunnel []_TtcpBuftunnel
}

func (___Vtbfa *_TtcpBufferArr) String() string {
	__Vs := _Spf("BF{%d/%d/%d,(%d)[",
		___Vtbfa.tbaCntUsed,
		___Vtbfa.tbaCntFree,
		___Vtbfa.tbaCntMax,
		len(___Vtbfa.tbaMtid))

	_FpfN(" 381817 01 ")

	__Vi2 := 0
	for __Vk2, __Vv2 := range ___Vtbfa.tbaMtid {
		_FpfN(" 381817 02 ")
		if 0 == __Vi2 {
			_FpfN(" 381817 03 ")
			__Vs += _Spf("%x %d", __Vk2[:2], __Vv2)
		} else {
			_FpfN(" 381817 04 ")
			__Vs += _Spf(",%x %d", __Vk2[:2], __Vv2)
		}
	}
	_FpfN(" 381817 05 ")

	__Vs += "]"
	for _, __Vv3 := range ___Vtbfa.tbaMtid {
		_FpfN(" 381817 06 ")
		__Vs += _Spf("(%d:%s)", __Vv3, ___Vtbfa.tbaMbuftunnel[__Vv3])
	}
	_FpfN(" 381817 07 ")
	return __Vs
}
