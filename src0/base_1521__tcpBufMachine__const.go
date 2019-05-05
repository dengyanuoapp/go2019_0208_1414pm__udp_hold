package main

type _TtcpBufMachine struct {
	tbmCHtcpReceBI chan []byte // _TtcpNodE tnCHtcpReceBLO *chan []byte // byte of _TtcpNodeDataRece
	tbmCHoutLO     *chan []byte
	tbmCBinit      func(*_TtcpBufMachine)
}

type _TtcpBufCellX struct{}
type _TtcpBufCell struct {
	tbcBuf [1024]byte
	tbcLen int
}

type _TtcpBuFx struct{}
type _TtcpBuF struct {
	tbCellArr [4096]_TtcpBufCell
	tbStart   int
	tbEnd     int
	tbCellNow _TtcpBufCell
}

type _TtcpBuftunnelX struct{}
type _TtcpBuftunnel struct {
	tbtTunnelNo byte     //  Tunnel Number : 0 -> normal , 1 ... 255 -> special channel
	tbtTidx     [16]byte // tunnel ID
	tbtL2R      _TtcpBuF // from local to remote
	tbtR2L      _TtcpBuF // from remote to local
}

type _TtcpBufferArrX struct{}
type _TtcpBufferArr struct {
	tbaCntMax  int
	tbaCntUsed int
	tbaCntFree int
	tbaMtid    map[[16]byte]int // tunnel ID
	tbaMbuf    []_TtcpBuftunnel
}
