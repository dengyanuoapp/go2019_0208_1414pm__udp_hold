package main

type _TloginGenerator struct {
	ulCHconnectSuccI  chan bool // connect/login succeed true/false
	ulTmpToken        []byte
	ulSrvDownInfoLX   *_TsrvDownInfo      // server Info seting , update by Github , try to connect
	ulCB850101init    func(*_TuDecode)    // if nil , use the default init procedure
	ulCB850201chRece  func(*_TuDecode)    // if nil , use the default receive
	ulCB850301ConnGen func(*_TuDecode)    // if nil , use the default receive
	ulCHuConnPortLO   *chan _TudpConnPort // all data need to be sent by nodeS send here , then  will distribute to one of node
	//ulCHunSendLO      *chan _TudpNodeDataSend // all data need to be sent by nodeS send here , then  will distribute to one of node
}
