package main

type _TtcpBufMachine struct {
	tbmCHtcpReceBI chan []byte // _TtcpNodE tnCHtcpReceBLO *chan []byte // byte of _TtcpNodeDataRece
	tbmCHoutLO     *chan []byte
	tbmCBinit      func(*_TtcpBufMachine)
}

type _TtcpBuF struct {
	tbBuf   [4096][1024]byte
	tbLen   [4096]int
	tbStart int
	tbEnd   int
}
type _TtcpBuftunnel struct {
	tbtType byte
	tbtIdx  [16]byte
	tbtL2R  _TtcpBuF // from local to remote
	tbtR2L  _TtcpBuF // from remote to local
}

type _TtcpBufferArr struct {
	tbaCntMax  int
	tbaCntUsed int
	tbaCntFree int
	tbaMid     map[[16]byte]int
	tbaMbuf    []_TtcpBuftunnel
}
