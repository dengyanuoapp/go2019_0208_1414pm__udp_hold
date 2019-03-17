package main

import (
	"time"
)

const (
	_T1s    = 1 * time.Second
	_T5s    = 5 * time.Second
	_T10s   = 10 * time.Second
	_T15s   = 15 * time.Second
	_T20s   = 20 * time.Second
	_T25s   = 25 * time.Second
	_T50s   = 50 * time.Second
	_T80s   = 80 * time.Second
	_T100s  = 100 * time.Second
	_T110s  = 110 * time.Second
	_T120s  = 120 * time.Second
	_T140s  = 140 * time.Second
	_T1200s = 1200 * time.Second
	_T2400s = 2400 * time.Second
	_T3600s = 3600 * time.Second
	_T7200s = 7200 * time.Second
	_T1ms   = 1 * time.Millisecond
	_T5ms   = 5 * time.Millisecond
	_T10ms  = 10 * time.Millisecond
	_T15ms  = 15 * time.Millisecond
	_T20ms  = 20 * time.Millisecond
	_T25ms  = 25 * time.Millisecond
	_T50ms  = 50 * time.Millisecond
	_T80ms  = 80 * time.Millisecond
	_T100ms = 100 * time.Millisecond
	_T110ms = 110 * time.Millisecond
	_T120ms = 120 * time.Millisecond
	_T140ms = 140 * time.Millisecond
	_T240ms = 240 * time.Millisecond
	_T360ms = 360 * time.Millisecond
	_T720ms = 720 * time.Millisecond
)

//_FtimeNow
//func _Fsleep_100s()     { time.Sleep(100 * time.Second) }      //

// _TudpNodeSt
type _TgapTimer struct {
	uTmGapX        time.Duration
	uTmCHudpReceLX *_TudpNodeSt
	//uTmCHudpSend02    *_TudpNodeSt
	//uTmCHbyteRece02   []byte
	//uTmCHbyteSend02   []byte
	uTmGapNewSession2 *_TgapNewSession // input , try to connect
	//uTmCHtryConn2     chan _TsrvInfo
	uTmCB150101init   func(*_TgapTimer) // if nil , use the default init procedure
	uTmCB150201filter func(*_TgapTimer) // if nil , use the default filter procedure to deal with receive
}

//_TUreqNewSession
type _TgapNewSession struct {
	name         string
	connected    bool
	tryCnt       int
	skipCnt      int
	updateUri    string
	updatePasswd []byte
	srvInfo      *_TsrvInfo
}
