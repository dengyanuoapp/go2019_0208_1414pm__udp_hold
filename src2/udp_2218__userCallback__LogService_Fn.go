// _TacceptTCP
// _TtcpNodE
package main

// note : all debug log begin pushed into tnClog will try to redirect to TCP debug monitorS.
func _FuserCallback__service_dataChan__Log_Fn(___VtcpNode4 *_TtcpNodE) {
	//_Fpf( "283822 service" ); _Pn( )
	__VprStr := <-*___VtcpNode4.tnClog
	//_FpfN( "283823 service:%s" , __VprStr )

	for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
		if ___VtcpNode4.tnAcceptTCPs[__Vi].enabled {
			// the tnClog -> CchanMsg
			___VtcpNode4.tnClientMux.Lock()

			if ___VtcpNode4.tnAcceptTCPs[__Vi].enabled {
				___VtcpNode4.tnAcceptTCPs[__Vi].CchanMsg <- []byte(__VprStr)
			}

			___VtcpNode4.tnClientMux.Unlock()
		}
	}
} // _FuserCallback__service_dataChan__Log_Fn
