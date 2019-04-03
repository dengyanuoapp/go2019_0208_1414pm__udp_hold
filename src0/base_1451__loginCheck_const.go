package main

import (
	"sync"
)

// _TloginReq
type _TconnInfo struct {
	ciId128 []byte
}

type _TcmdMap struct {
	M   map[[16]byte]_Tdecode // key by 16byte : MeIdx128
	mux sync.Mutex
}
type _TdataMap struct {
	M   map[[16]byte]_Tdecode // key by 16byte : MeIdx128
	mux sync.Mutex
}

type _TloginCheck struct {
	ucDecodeI         chan _Tdecode
	ulCHconnPortI     chan _TudpConnPort      // all data need to be sent by nodeS send here , then  will distribute to one of node
	ucCHSendLO        *chan _TudpNodeDataSend // ugCHSendI
	ucCB900101init    func(*_TloginCheck)     // if nil , use the default init procedure
	ucCB900201stCheck func(*_TloginCheck)     // if nil , use the default receive
	ucCmd             _TcmdMap
	ucData            _TdataMap
	ucTokenA          []byte
}

func _FcheckDecodeType(___Vdecode *_Tdecode, ___VwantType byte) bool { // match --> return false , others -> return true
	if nil == ___Vdecode {
		return true
	}

	if ___VwantType != ___Vdecode.Type {
		return true
	}
	return false
}
