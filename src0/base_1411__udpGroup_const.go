package main

type _TudpGroupSt struct {
	ugName        string
	ugHostPortStr []string
	ugAmount      int
	ugRkeyLP      [](*[]byte)
	ugNodeS       []_TudpNodeSt
	ugCHrece      *chan _TudpNodeDataRece // if nil , drop it ; not-nil , put the received data into this chan
	ugCHsend      chan _TudpNodeDataSend  // try get data from chan, then send it out.
	ugCBinit      func(*_TudpGroupSt)     //
	ugCBrece      func(*_TudpGroupSt)     // if nil , use the default procedure to deal with receive
	ugCBsend      func(*_TudpGroupSt)     // if nil , use the default procedure to deal with send
}
