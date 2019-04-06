package main

import "bytes"

// note : all debug log begin pushed into tnClog will try to redirect to TCP debug monitorS.
func (___VtcpNode4 *_TtcpNodE) _FtcpNode__200301x_send__default() {
	_FpfN("283822 01 send start")
	for {
		select {
		case __Vds := <-___VtcpNode4.tnCHsendToAllClientI: // _TtcpNodeDataSend
			//_FpfN("283822 02 get from send-chain :{%s}", __Vds.String())

			if 0 == len(__Vds.tnsId128) {
				for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
					if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled {
						//_FpfN("283822 03 get from send-chain :{%s}", __Vds.String())
						___VtcpNode4.tnAcceptTCPs[__Vi]._FtcpNodeAccept_send(&__Vds)
					}
				}
			} else {
				for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
					if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled && (bytes.Equal(__Vds.tnsId128, ___VtcpNode4.tnAcceptTCPs[__Vi].taId128)) {
						_FpfN("283822 05 get from send-chain :{%s}", __Vds.String())
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

	if nil == ___Vs {
		_FpfN(" 838187 02 : why Null ? ")
		return
	}

	if 0 == ___Vs.tnsLen {
		_FpfN(" 838187 03 : why ZERO len ? ")
		return
	}

	if ___Vs.tnsLen != len(___Vs.tnsBuf) {
		_FpfN(" 838187 04 : why diff len ? ")
		return
	}

	if nil == ___Vacc.taConnTCP {
		_FpfN(" 838187 05 : why conn NULL ? ")
		return
	}

	_FpfN(" 838187 06 : <%v>sending %s", ___Vacc.taConnTCP, ___Vacc.String())

	//func (c *TCPConn) Write(b []byte) (int, error)
	___Vacc.taConnTCP.Write(___Vs.tnsBuf)

}
