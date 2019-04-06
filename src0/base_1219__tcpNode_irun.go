package main

func (___Vtn0 *_TtcpNodE) IRun(___Vidx int) {
	switch ___Vidx {
	case 200101:
		if nil == ___Vtn0.tnCBinit {
			___Vtn0._FtcpNode__200101x__init_default()
		} else {
			___Vtn0.tnCBinit(___Vtn0)
		}
	default:
		_FpfNex(" 739181 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func (___Vtn3 *_TtcpNodE) _FtcpNode__200101x__init_default() {

	___Vtn3._FtcpNode__tryListen01()

	if ___Vtn3.tnAmount < 1 {
		___Vtn3.tnAmount = 1
	}
	if ___Vtn3.tnAmount > 100 {
		___Vtn3.tnAmount = 100
	}

	___Vtn3.acceptTCPs = make([]_TacceptTCP, ___Vtn3.tnAmount)
	go _Fhandle_tcpAccept01(___Vtn3)

	for {
		_Fsleep_1ms()

		_FnotNullRun011_tcp_service_chan(___Vtn3.tnCBsvrDataChan, ___Vtn3)

	}
	*___Vtn3.Cexit <- "Error 183818: (" + ___Vtn3.tnHostPortStr + ")"
} // _FtcpNode__200101x__init_default
