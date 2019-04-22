package main

// _TencodeX
// _TdataTran
type _TrecePackThenEncodeAsLoad struct {
	pelCHinI   chan _TdataTran
	pelCHoutLO *chan _TdataTran
	pelCBinit  func(*_TrecePackThenEncodeAsLoad)
}
