package main

type _TtcpBufMachine struct {
	tbmCHtcpReceBI chan []byte // _TtcpNodE tnCHtcpReceBLO *chan []byte // byte of _TtcpNodeDataRece
	tbmCHoutLO     *chan []byte
	tbmCBinit      func(*_TtcpBufMachine)
}
