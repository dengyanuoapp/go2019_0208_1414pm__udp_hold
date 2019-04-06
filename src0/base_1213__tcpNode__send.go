package main

import "bytes"

// note : all debug log begin pushed into tnClog will try to redirect to TCP debug monitorS.
func (___VtcpNode4 *_TtcpNodE) _FtcpNode__200301x_send__default() {
	_FpfN("283822 01 send start")
	for {
		select {
		case __Vds := <-___VtcpNode4.tnCHsendToAllClientI: // _TtcpNodeDataSend
			_FpfN("283822 02 send chain:%s", __Vds.String)

			if 0 == len(__Vds.tnsId128) {
				for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
					if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled {
						___VtcpNode4.tnAcceptTCPs[__Vi]._FtcpNodeAccept_send(&__Vds)
					}
				}
			} else {
				for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
					if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled && (bytes.Equal(__Vds.tnsId128, ___VtcpNode4.tnAcceptTCPs[__Vi].taId128)) {
						___VtcpNode4.tnAcceptTCPs[__Vi]._FtcpNodeAccept_send(&__Vds)
					}
				}
			}

		} // end select
	} // end for
} //

// _TacceptTCP
// _TtcpNodE
func (___Vacc *_TacceptTCP) _FtcpNodeAccept_send(___Vs *_TtcpNodeDataSend) {
	defer ___Vacc.taServerTCP.tnClientMux.Unlock()
	___Vacc.taServerTCP.tnClientMux.Lock()

	if false == ___Vacc.taEnabled {
		_FpfN(" 838187 01 : already lost connect:%s", ___Vacc.String())
		return
	}

	_FpfN(" 838187 02 : sending %s", ___Vacc.String())
}
