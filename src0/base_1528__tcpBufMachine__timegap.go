package main

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__1500301x__timegap_timeout_delete() {
	for {
		_Fsleep(_T1s)
		___Vtbm.tbmChCheckTunnelTimeOut <- 1
	}
}
func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__1500302x__timegap_bufSendTunnelCheck() {
	___Vtbm.tbmChCheckLocal2RemoteGap = make(chan byte, 50)
}
