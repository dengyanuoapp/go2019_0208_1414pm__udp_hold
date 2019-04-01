package main

import (
	"time"
)

//_FtimeNow
//func _Fsleep_100s()     { time.Sleep(100 * time.Second) }      //

// _TudpNodeSt
type _TuDecode struct {
	uTmCHunDataReceLI *_TudpNodeSt // an udpNode pointer , if not nil , read from it's unCHreceLO
	uTmDecodeLO       *chan _Tdecode
	uTmCB700101init   func(*_TuDecode) // if nil , use the default init procedure // _FudpDecode__700101x__init__default
	uTmCB700201rece2  func(*_TuDecode) // if nil , use the default receive        // _FudpDecode__700201x__gap_receive__default
	uTmCB700301gap2   func(*_TuDecode) // if nil , use the default gap loop       // _FudpDecode__700301x__gap_loop__default
	uTmGapX           time.Duration
	//uTmCHugConnPortLO *chan _TudpConnPort // nor-nil ==> every gap , gen tryConn package and push into this chan , usually a udpGroupNodeArr
	//uTmSrvDownInfoLX  *_TsrvDownInfo      // input , try to connect
	//uTmToken          []byte
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
}
