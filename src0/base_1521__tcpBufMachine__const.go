package main

type _TtcpBufMachine struct {
	tbmCHinI   chan []byte
	tbmCHoutLO *chan []byte
	tbmCBinit  func(*_TtcpBufMachine)
}
