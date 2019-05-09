package main

type _TtcpBufMachine struct {
	tbmCHtcpReceBI   chan []byte   // _TtcpNodE tnCHtcpReceBLO *chan []byte // byte of _TtcpNodeDataRece
	tbmCHtcpReceCmdI chan [17]byte // command of tunnel : byte 0:15 -> channelID, byte [16] -> cmd : // TcpNodeCmd__NULL tnCHtcpReceCmdLO
	tbmCHoutLO       *chan []byte
	tbmCBinit        func(*_TtcpBufMachine)
	tbmBufArr        _TtcpBufferArr // _TtcpBufferArrX

}

type _TtcpBufCellX struct{}
type _TtcpBufCell struct {
	tbcOffset int64
	tbcBuf    [1024]byte
	tbcLen    int
}

func (___Vtbc *_TtcpBufCell) String() string {
	__Vs := _Spf("(%d:%d:%d)[%s]", ___Vtbc.tbcOffset, ___Vtbc.tbcLen, ___Vtbc.tbcOffset+int64(___Vtbc.tbcLen), ___Vtbc.tbcBuf[:7])
	return __Vs
}

type _TtcpBuFx struct{}
type _TtcpBuF struct {
	tbCellArr     [4096]_TtcpBufCell
	tbStart       int
	tbEnd         int
	tbCellNow     _TtcpBufCell
	tbOffsetStart int64
	tbOffsetEnd   int64
	tbOffsetNow   int64
	tbFreeCnt     int
}

func (___Vtbf *_TtcpBuF) String() string {
	__Vs := _Spf("<now{%s} st/end:%d/%d free:%d (%d/%d/%d)>",
		___Vtbf.tbCellNow.String(), ___Vtbf.tbStart, ___Vtbf.tbEnd, ___Vtbf.tbFreeCnt,
		___Vtbf.tbOffsetStart, ___Vtbf.tbOffsetEnd, ___Vtbf.tbOffsetNow)
	return __Vs
}

type _TtcpBuftunnelX struct{}
type _TtcpBuftunnel struct {
	tbtEna      bool
	tbtType     int      //  Tunnel Type : 0 -> normal , 1 ... 255 -> special channel
	tbtTunnelNo int      //
	tbtLastL2R  int      //
	tbtLastR2L  int      //
	tbtLastALL  int      //
	tbtTidx     [16]byte // tunnel ID
	tbtL2R      _TtcpBuF // from local to remote // _TtcpBuFx
	tbtR2L      _TtcpBuF // from remote to local // _TtcpBuFx
}

func (___Vtbt *_TtcpBuftunnel) String() string {
	__Vs := _Spf("T:%d %d:%x\n:l2r[%s]\n:r2l[%s]\n",
		___Vtbt.tbtLastALL, ___Vtbt.tbtTunnelNo, ___Vtbt.tbtTidx[:2],
		___Vtbt.tbtL2R.String(), ___Vtbt.tbtR2L.String())
	return __Vs
}

type _TtcpBufferArrX struct{}
type _TtcpBufferArr struct {
	tbaCntMax     int
	tbaCntUsed    int
	tbaCntFree    int
	tbaMtid       map[[16]byte]int // tunnel ID
	tbaMbuftunnel []_TtcpBuftunnel // _TtcpBuftunnelX
}

func (___Vtbfa *_TtcpBufferArr) String() string {
	__Vs := _Spf("BF{%d/%d/%d,(%d)[",
		___Vtbfa.tbaCntUsed,
		___Vtbfa.tbaCntFree,
		___Vtbfa.tbaCntMax,
		len(___Vtbfa.tbaMtid))

	__FpfN(" 381817 01 ")

	__Vi2 := 0
	for __Vk2, __Vv2 := range ___Vtbfa.tbaMtid {
		__FpfN(" 381817 02 ")
		if 0 == __Vi2 {
			__FpfN(" 381817 03 ")
			__Vs += _Spf("%x %d", __Vk2[:2], __Vv2)
		} else {
			__FpfN(" 381817 04 ")
			__Vs += _Spf(",%x %d", __Vk2[:2], __Vv2)
		}
	}
	__FpfN(" 381817 05 ")

	__Vs += "]"
	for _, __Vv3 := range ___Vtbfa.tbaMtid {
		__FpfN(" 381817 06 ")
		__Vs += _Spf("(%d:%s)", __Vv3, ___Vtbfa.tbaMbuftunnel[__Vv3].String())
	}
	__FpfN(" 381817 07 ")
	return __Vs
}
