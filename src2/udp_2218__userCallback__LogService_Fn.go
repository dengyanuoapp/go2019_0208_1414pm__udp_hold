// _TacceptTCP 
// _TserviceTCP 
package main

// note : all debug log begin pushed into Clog will try to redirect to TCP debug monitorS.
func _FuserCallback__Log_service_dataChan__Fn(___VserviceTcp *_TserviceTCP ) {
    //_Fpf( "283822 service" ); _Pn( )
    __VprStr := <- *___VserviceTcp.Clog
    //_FpfN( "283823 service:%s" , __VprStr )

    for __Vi:=0 ; __Vi < ___VserviceTcp.cAmount ; __Vi ++ {
        if ___VserviceTcp.acceptTCPs[__Vi].enabled {
            // the Clog -> CchanMsg
            ___VserviceTcp.clientMux.Lock()

            if ___VserviceTcp.acceptTCPs[__Vi].enabled {
                ___VserviceTcp.acceptTCPs[__Vi].CchanMsg <- []byte(__VprStr)
            }

            ___VserviceTcp.clientMux.Unlock()
        }
    }
} // _FuserCallback__Log_service_dataChan__Fn
