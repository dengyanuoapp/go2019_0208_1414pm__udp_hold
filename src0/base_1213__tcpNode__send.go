package main

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
				_FtcpNode__200301x2_sendAll__default(&__Vds)

		case __Vb := <-___VtcpNode4.tnCHtcpSendBI: //        chan []byte        // byte of _TtcpNodeDataSend
			___TtcpNodE__mux.Lock()
			_CFpfN("283821 04 _TtcpNodE tnCHtcpSendBI :{%s}", String9s(&__Vb))

		case __VbCmd := <-___VtcpNode4.tnCHtcpSendCmdI: // chan [17]byte
			___TtcpNodE__mux.Lock() // command of tunnel : byte 0:15 -> channelID, byte [16] -> cmd : // TcpNodeCmd__NULL
			_CFpfN("283821 06 _TtcpNodE tnCHtcpSendCmdI :{%x %s}", __VbCmd[:16], StrTcpNodeCmd(__VbCmd[16]))

			___VtcpNode4.
				_FtcpNode__tryAcction_Cmd(__VbCmd)

		} // end Select
		___TtcpNodE__mux.Unlock()
	} // end for
} //
