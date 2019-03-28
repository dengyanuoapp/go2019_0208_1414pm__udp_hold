package main

type _TloginGenerator struct {
	uTmSrvDownInfoLX   *_TsrvDownInfo      // server Info seting , update by Github , try to connect
	uTmCHugConnPortLO  *chan _TudpConnPort // every gap , gen tryConn package and push into this chan , usually a udpGroupNodeArr
	uTmCHconnectSuccI  chan bool           // connect/login succeed true/false
	uTmDecodeOutLX     *chan _Tdecode
	uTmCB850101init    func(*_TgapTimer) // if nil , use the default init procedure
	uTmCB850201chRece  func(*_TgapTimer) // if nil , use the default receive
	uTmCB850301ConnGen func(*_TgapTimer) // if nil , use the default receive
}
