package main

// _TudpNodeSt
type _TudpGroupSt struct {
	ugCHSendI        chan _TudpNodeDataSend        // all data need to be sent by nodeS send here , then  will distribute to one of node
	ugChTmpSendLI    [](*(chan _TudpNodeDataSend)) // unCHsendI     chan _TudpNodeDataSend  // interal used
	ugLastSendIdx    int
	ugName           string
	ugHostPortStr    []string
	ugAmount         int
	ugRkeyLP         [](*[]byte)
	ugNodeS          []_TudpNodeSt
	ugCBinit         func(*_TudpGroupSt)     //
	ugCBchInSend     func(*_TudpGroupSt)     // if nil , use the default procedure to deal with send
	ugCHreceByteLO   *chan []byte            // unCHreceByteLO    *chan _TudpNodeDataRece // interal used
	ugCHreceStructLO *chan _TudpNodeDataRece // unCHreceLO    *chan _TudpNodeDataRece // interal used
}
