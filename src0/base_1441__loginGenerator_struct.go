package main

type _TloginGenerator struct {
	lgSrvDownInfoLX    *_TsrvDownInfo          // server Info seting , update by Github , try to connect
	lgCB850101init     func(*_TloginGenerator) // if nil , use the default init procedure
	lgCB850201chRece   func(*_TloginGenerator) // if nil , use the default receive
	lgCB850301ConnGen  func(*_TloginGenerator) // if nil , use the default receive
	lgCHuConnPortLO    *chan _TudpConnPort     // all data need to be sent by nodeS send here , then  will distribute to one of node
	lgCHdataMachineIdI chan _TdataMachinEid    // lgCHconnectSuccI   chan bool  // connect/login succeed true/false
}
