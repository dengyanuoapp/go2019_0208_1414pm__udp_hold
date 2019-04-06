package main

func (___VtcpNode0 *_TtcpNodE) IRun(___Vidx int) {
	switch ___Vidx {
	case 200101:
		if nil == ___VtcpNode0.tnCBinit {
			___VtcpNode0._FtcpNode__200101x__init_default()
		} else {
			___VtcpNode0.tnCBinit(___VtcpNode0)
		}
	default:
		_FpfNex(" 739181 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func (___VtcpNode3 *_TtcpNodE) _FtcpNode__200101x__init_default() {

	___VtcpNode3._FtryListenToTCP01()

	if ___VtcpNode3.cAmount < 1 {
		___VtcpNode3.cAmount = 1
	}
	if ___VtcpNode3.cAmount > 100 {
		___VtcpNode3.cAmount = 100
	}
	___VtcpNode3.acceptTCPs = make([]_TacceptTCP, ___VtcpNode3.cAmount)
	go _Fhandle_tcpAccept01(___VtcpNode3)

	for {
		_Fsleep_1ms()

		_FnotNullRun011_tcp_service_chan(___VtcpNode3.TcallbackSvrDataChan, ___VtcpNode3)

	}
	*___VtcpNode3.Cexit <- "Error 183818: (" + ___VtcpNode3.hostPortStr + ")"
} // _FtcpNode__200101x__init_default
