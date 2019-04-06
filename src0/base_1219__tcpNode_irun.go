package main

func (___VsrvT *_TtcpNodE) IRun(___Vidx int) {
	switch ___Vidx {
	case 200101:
		if nil == ___VsrvT.tnCBinit {
			___VsrvT._FserviceTcp__200101x__init_default()
		} else {
			___VsrvT.tnCBinit(___VsrvT)
		}
	default:
		_FpfNex(" 739181 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func (___VserviceTcp *_TtcpNodE) _FserviceTcp__200101x__init_default() {

	___VserviceTcp._FtryListenToTCP01()

	if ___VserviceTcp.cAmount < 1 {
		___VserviceTcp.cAmount = 1
	}
	if ___VserviceTcp.cAmount > 100 {
		___VserviceTcp.cAmount = 100
	}
	___VserviceTcp.acceptTCPs = make([]_TacceptTCP, ___VserviceTcp.cAmount)
	go _Fhandle_tcpAccept01(___VserviceTcp)

	for {
		_Fsleep_1ms()

		_FnotNullRun011_tcp_service_chan(___VserviceTcp.TcallbackSvrDataChan, ___VserviceTcp)

	}
	*___VserviceTcp.Cexit <- "Error 183818: (" + ___VserviceTcp.hostPortStr + ")"
} // _FserviceTcp__200101x__init_default
