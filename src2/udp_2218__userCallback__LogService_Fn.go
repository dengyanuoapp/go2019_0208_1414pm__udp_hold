// _TacceptTCP
// _TtcpNodE
package main

// note : all debug log begin pushed into Clog will try to redirect to TCP debug monitorS.
func _FuserCallback__service_dataChan__Log_Fn(___VtcpNode4 *_TtcpNodE) {
	//_Fpf( "283822 service" ); _Pn( )
	__VprStr := <-*___VtcpNode4.Clog
	//_FpfN( "283823 service:%s" , __VprStr )

	for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
		if ___VtcpNode4.acceptTCPs[__Vi].enabled {
			// the Clog -> CchanMsg
			___VtcpNode4.clientMux.Lock()

			if ___VtcpNode4.acceptTCPs[__Vi].enabled {
				___VtcpNode4.acceptTCPs[__Vi].CchanMsg <- []byte(__VprStr)
			}

			___VtcpNode4.clientMux.Unlock()
		}
	}
} // _FuserCallback__service_dataChan__Log_Fn
