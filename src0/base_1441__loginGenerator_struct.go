package main

type _TloginGenerator struct {
	ulTmpToken       []byte
	ulSrvDownInfoLX  *_TsrvDownInfo      // server Info seting , update by Github , try to connect
	ulCHugConnPortLO *chan _TudpConnPort // every gap , gen tryConn package and push into this chan , usually a udpGroupNodeArr
	ulCHconnectSuccI chan bool           // connect/login succeed true/false
	//ulDecodeOutLX     *chan _Tdecode
	ulCB850101init    func(*_TgapTimer) // if nil , use the default init procedure
	ulCB850201chRece  func(*_TgapTimer) // if nil , use the default receive
	ulCB850301ConnGen func(*_TgapTimer) // if nil , use the default receive
}
