package main

func (___Vtn0 *_TtcpNodE) IRun(___Vidx int) {
	switch ___Vidx {
	case 200101:
		if nil == ___Vtn0.tnCBinit {
			___Vtn0.
				_FtcpNode__200101x__init_default()
		} else {
			___Vtn0.tnCBinit(___Vtn0)
		}
	case 200301:
		go ___Vtn0.
			_FtcpNode__200301x_send__default()
	case 200401:
		go ___Vtn0.
			_FtcpNode__200401x_accept_default() // each acc : _FtcpNodeAccept__200401x4__dataReceiveMsg01
	case 200801:
		go ___Vtn0.
			_FtcpNode__200801x_send__tester()
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

	___Vtn3.tnAcceptTCPs = make([]_TacceptTCP, ___Vtn3.tnAmount)
	___Vtn3.tnCHsendToAllClientI = make(chan _TtcpNodeDataSend, 10)
	go _Frun(___Vtn3, 200401) // _FtcpNode__200401x_accept_default

	_Fsleep(_T1s)

	go _Frun(___Vtn3, 200301) // _FtcpNode__200301x_send__default
	//go _Frun(___Vtn3, 200801) // _FtcpNode__200801x_send__tester

} // _FtcpNode__200101x__init_default
