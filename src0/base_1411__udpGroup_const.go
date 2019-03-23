package main

type _TudpGroupSt struct {
	ugName         string
	ugHostPortStr  []string
	ugAmount       int
	ugRkeyLP       [](*[]byte)
	ugNodeS        []_TudpNodeSt
	ugCHuConnPortX chan _TudpConnPort
	ugCBinit       func(*_TudpGroupSt)           //
	ugCBchInSend   func(*_TudpGroupSt)           // if nil , use the default procedure to deal with send
	ugCHtmpReceX   chan _TudpNodeDataRece        // unCHreceLX    *chan _TudpNodeDataRece //
	ugCHtmpSendLX  [](*(chan _TudpNodeDataSend)) // unCHsendX     chan _TudpNodeDataSend  //
	ugLastSendIdx  int
	//ugCBrece       func(*_TudpGroupSt) // if nil , use the default procedure to deal with receive
	//ugCBsend       func(*_TudpGroupSt) // if nil , use the default procedure to deal with send
}
