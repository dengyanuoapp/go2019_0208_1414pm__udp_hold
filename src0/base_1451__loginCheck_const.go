package main

import (
	"sync"
)

// _TreqIneedToLogin
type _TconnInfo struct {
	ciId128 []byte
}

type _TconnMap struct {
	cmAnow  map[[16]byte]_Tdecode // _TreqIneedToLogin MeIdx128
	cmAlast map[[16]byte]_Tdecode // _TreqIneedToLogin MeIdx128
	cmAlas2 map[[16]byte]_Tdecode // _TreqIneedToLogin MeIdx128
	cmMux   sync.Mutex
}

type _TloginCheck struct {
	ucDecodeI         chan _Tdecode
	ucCHSendLO        *chan _TudpNodeDataSend // ugCHSendI
	ucCB900101init    func(*_TloginCheck)     // if nil , use the default init procedure
	ucCB900201stCheck func(*_TloginCheck)     // if nil , use the default receive
	ucM               _TconnMap
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
