package main

//import "time"

// _TudpNodeSt
type _TuEncode struct {
	enCHuDataSendLO     *chan _TudpNodeDataSend // _TudpNodeDataSendX
	enCHencodeCmdI      chan _Tencode
	enCHencodeDataI     chan _Tencode
	enCB1100101init     func(*_TuDecode) // if nil , use the default init procedure // _FudpDecode__700101x__init__default
	enCB1100201packSend func(*_TuDecode) // if nil , use the default receive        // _FudpDecode__700201x__receive__default
	//enCHunDataReceLI *_TudpNodeSt // an udpNode pointer , if not nil , read from it's unCHreceLO
}
