package main

import "time"

// _TencodeX
// _TdataTran
type _TrecePackThenEncodeAsLoad struct {
	pelGap                 time.Duration // _T1s
	pelCHudpNodeDataReceBI chan []byte
	pelCHencodeBLO         *chan _TdataTran
	pelCBinit              func(*_TrecePackThenEncodeAsLoad)
	pelChOutputGenGap      chan byte // every 1s gen output
}
