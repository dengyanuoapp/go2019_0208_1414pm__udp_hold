package main

import "bytes"
import "sync"

var ___TtcpNodE__mux sync.Mutex

// note : all debug log begin pushed into tnClog will try to redirect to TCP debug monitorS.
func (___VtcpNode4 *_TtcpNodE) _FtcpNode__200301x_send__default() {
	//_FpfN("283821 01 send start")
	for {
		select {
		case __Vds := <-___VtcpNode4.tnCHsendToAllClientI: // _TtcpNodeDataSend
			//_FpfN("283821 02 get from send-chain :{%s}", __Vds.String())
			___TtcpNodE__mux.Lock()
			___VtcpNode4.
				_FtcpNode__200301x2_send__default(&__Vds)

		case __Vb := <-___VtcpNode4.tnCHtcpSendBI: //        chan []byte        // byte of _TtcpNodeDataSend
			___TtcpNodE__mux.Lock()
			_CFpfN("283821 04 _TtcpNodE tnCHtcpSendBI :{%s}", String9s(&__Vb))

		case __VbCmd := <-___VtcpNode4.tnCHtcpSendCmdI: // chan [17]byte
			___TtcpNodE__mux.Lock() // command of tunnel : byte 0:15 -> channelID, byte [16] -> cmd : // TcpNodeCmd__NULL
			_CFpfN("283821 06 _TtcpNodE tnCHtcpSendCmdI :{%x %x}", __VbCmd[:16], __VbCmd[16])

		} // end Select
		___TtcpNodE__mux.Unlock()
	} // end for
} //

func (___VtcpNode4 *_TtcpNodE) _FtcpNode__200301x2_send__default(___Vds *_TtcpNodeDataSend) {
	if 0 == len(___Vds.tnsId128) {
		for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
			if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled {
				//_FpfN("283822 03 get from send-chain :{%s}", ___Vds.String())
				___VtcpNode4.
					tnAcceptTCPs[__Vi].
					_FtcpNodeAccept_send(___Vds)
			}
		}
	} else {
		for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
			if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled && (bytes.Equal(___Vds.tnsId128, ___VtcpNode4.tnAcceptTCPs[__Vi].taId128)) {
				_FpfN("283822 05 get from send-chain :{%s}", ___Vds.String())
				___VtcpNode4.
					tnAcceptTCPs[__Vi].
					_FtcpNodeAccept_send(___Vds)
			}
		}
	}
}

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

	//_FpfN(" 838187 06 : <%v>sending %s", ___Vacc.taConnTCP, ___Vacc.String())

	//func (c *TCPConn) Write(b []byte) (int, error)
	___Vacc.
		taConnTCP.
		Write(___Vs.tnsBuf)

}
