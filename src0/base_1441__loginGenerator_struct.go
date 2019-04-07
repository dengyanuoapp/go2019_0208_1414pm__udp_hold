package main

type _TloginGenerator struct {
	lgSrvDownInfoLX    *_TsrvDownInfo          // server Info seting , update by Github , try to connect
	lgCB850101init     func(*_TloginGenerator) // if nil , use the default init procedure // _FudpDecode__800101x__init__tryUdpLogin__default
	lgCB850201chRece   func(*_TloginGenerator) // if nil , use the default receive // _FudpDecode__800201x__chan_in__default
	lgCB850301ConnGen  func(*_TloginGenerator) // if nil , use the default receive // _FudpDecode__800301x__connGen__default
	lgCHuConnPortLO    *chan _TudpConnPort     // all data need to be sent by nodeS send here , then  will distribute to one of node
	lgCHdataMachineIdI chan _TdataMachinEid    // lgCHconnectSuccI   chan bool  // connect/login succeed true/false
}
