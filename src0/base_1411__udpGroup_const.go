package main

type _TudpGroupSt struct {
	ugChTmpReceO   chan _TudpNodeDataRece        // unCHreceLO    *chan _TudpNodeDataRece // interal used
	ugChTmpSendLI  [](*(chan _TudpNodeDataSend)) // unCHsendI     chan _TudpNodeDataSend  // interal used
	ugName         string
	ugHostPortStr  []string
	ugAmount       int
	ugRkeyLP       [](*[]byte)
	ugNodeS        []_TudpNodeSt
	ugCBinit       func(*_TudpGroupSt) //
	ugCBchInSend   func(*_TudpGroupSt) // if nil , use the default procedure to deal with send
	ugCHuConnPortI chan _TudpConnPort
	ugCHSendLI     chan _TudpNodeDataSend // all data need to be sent by nodeS send here , then  will distribute to one of node
	ugLastSendIdx  int
	//ugDecodeI      chan _Tdecode
	//ugCBrece       func(*_TudpGroupSt) // if nil , use the default procedure to deal with receive
	//ugCBsend       func(*_TudpGroupSt) // if nil , use the default procedure to deal with send
}
