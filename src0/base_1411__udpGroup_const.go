package main

type _TudpGroupSt struct {
	ugName         string
	ugHostPortStr  []string
	ugAmount       int
	ugRkeyLP       [](*[]byte)
	ugNodeS        []_TudpNodeSt
	ugCHuConnPortX chan _TudpConnPort
	ugCBinit       func(*_TudpGroupSt) //
	ugCBrece       func(*_TudpGroupSt) // if nil , use the default procedure to deal with receive
	ugCBsend       func(*_TudpGroupSt) // if nil , use the default procedure to deal with send
	ugCBportConn   func(*_TudpGroupSt) // if nil , use the default procedure to deal with send
	//ugCHreceLX     *chan _TudpNodeDataRece // if nil , drop it ; not-nil , put the received data into this chan
	//ugCHsendX      chan _TudpNodeDataSend  // try get data from chan, then send it out.
}
