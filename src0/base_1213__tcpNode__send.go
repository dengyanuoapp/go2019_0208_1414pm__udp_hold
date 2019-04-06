// _TacceptTCP
// _TtcpNodE
package main

// note : all debug log begin pushed into tnClog will try to redirect to TCP debug monitorS.
func (___VtcpNode4 *_TtcpNodE) _FtcpNode__200301x_send__default() {
	//_Fpf( "283822 01 service" ); _Pn( )
	__VprStr := <-*___VtcpNode4.tnClog
	//_FpfN( "283822 02 service:%s" , __VprStr )

	for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
		if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled {
			// the tnClog -> taCchanMsg
			___VtcpNode4.tnClientMux.Lock()

			if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled {
				___VtcpNode4.tnAcceptTCPs[__Vi].taCchanMsg <- []byte(__VprStr)
			}

			___VtcpNode4.tnClientMux.Unlock()
		}
	}
} //
