package main

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__1500301x__timegap_timeout_delete() {
	for {
		_Fsleep(_T1s)
		___Vtbm.tbmChCheckTunnelTimeOut <- 1
	}
}
