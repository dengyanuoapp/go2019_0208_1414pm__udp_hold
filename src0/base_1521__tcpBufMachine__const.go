package main

type _TtcpBufMachine struct {
	tbmCHtcpReceBI chan []byte
	tbmCHoutLO     *chan []byte
	tbmCBinit      func(*_TtcpBufMachine)
}
