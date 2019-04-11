package main

import (
	"time"
)

//_FtimeNow
//func _Fsleep_100s()     { time.Sleep(100 * time.Second) }      //

// _TudpNodeSt
type _TuDecode struct {
	//uTmCHunDataReceI  chan _TudpNodeDataRece // an udpNode pointer , if not nil , read from it's unCHreceLO
	uTmCHunDataReceI  chan []byte    // an udpNode pointer , if not nil , read from it's unCHreceLO
	uTmCHdecodeCmdLO  *chan _Tdecode //_TdecodeX
	uTmCHdecodeDataLO *chan _Tdecode
	uTmCB700101init   func(*_TuDecode) // if nil , use the default init procedure // _FudpDecode__700101x__init__default
	uTmCB700201rece2  func(*_TuDecode) // if nil , use the default receive        // _FudpDecode__700201x__receive__default
	uTmCB700301gap2   func(*_TuDecode) // if nil , use the default gap loop       // _FudpDecode__700301x__loop__default
	uTmGapX           time.Duration
	//uTmCHunDataReceLI *_TudpNodeSt // an udpNode pointer , if not nil , read from it's unCHreceLO
}

//_TUreqNewSession
type _TsrvDownInfo struct {
	name         string
	connected    bool
	tryCnt       int
	skipCnt      int
	updateUri    string
	updatePasswd []byte
	srvInfo      _TsrvInfo
	lastDownTime int
}

func (___Vsi *_TsrvDownInfo) String() string {
	return _Spf("n:%s cOK:%t try:%d skip:%d uUri:%s uPwd:%s srv{%s} last:%d",
		___Vsi.name,
		___Vsi.connected,
		___Vsi.tryCnt,
		___Vsi.skipCnt,
		___Vsi.updateUri,
		String5(&___Vsi.updatePasswd),
		___Vsi.srvInfo.String(),
		___Vsi.lastDownTime,
	)
}
