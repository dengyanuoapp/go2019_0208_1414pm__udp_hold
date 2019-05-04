package main

type _TtcpBufMachine struct {
	bnbCHinI    chan byte
	bnbCHoutLO1 *chan byte
	bnbCHoutLO2 *chan byte
	bnbCHoutLO3 *chan byte
	bnbCHoutLO4 *chan byte
	bnbCHoutLO5 *chan byte
	bnbCBinit   func(*_TtcpBufMachine)
}
