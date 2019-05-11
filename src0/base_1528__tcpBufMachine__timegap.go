package main

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__1500301x__timegap_timeout_delete() {
	for {
		_Fsleep(_T1s)
		___Vtbm.
			tbmChCheckTunnelTimeOut <- 1
	}
}

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__1500302x__timegap_bufSendTunnelCheck() {
	for {
		_Fsleep(_T240ms)
		___Vtbm.
			tbmChCheckLocal2RemoteGap <- 1
		__FpfN(" 388348 01 : _FtcpBufMachine__1500302x__timegap_bufSendTunnelCheck tick , len %d , cap %d",
			len(___Vtbm.tbmChCheckLocal2RemoteGap), cap(___Vtbm.tbmChCheckLocal2RemoteGap))
	}
}
