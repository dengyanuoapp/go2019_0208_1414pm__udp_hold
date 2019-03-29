package main

type _TudpGroupSt struct {
	ugName         string
	ugHostPortStr  []string
	ugAmount       int
	ugRkeyLP       [](*[]byte)
	ugNodeS        []_TudpNodeSt
	ugCBinit       func(*_TudpGroupSt) //
	ugCBchInSend   func(*_TudpGroupSt) // if nil , use the default procedure to deal with send
	ugCHuConnPortI chan _TudpConnPort
	ugCHtmpReceO   chan _TudpNodeDataRece        // unCHreceLO    *chan _TudpNodeDataRece //
	ugCHtmpSendLI  [](*(chan _TudpNodeDataSend)) // unCHsendI     chan _TudpNodeDataSend  //
	ugLastSendIdx  int
	//ugDecodeI      chan _Tdecode
	//ugCBrece       func(*_TudpGroupSt) // if nil , use the default procedure to deal with receive
	//ugCBsend       func(*_TudpGroupSt) // if nil , use the default procedure to deal with send
}
