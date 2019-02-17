package main

// _TacceptTCP 
// _TserviceTCP 
// note : all debug log begin pushed into Clog will try to redirect to TCP debug monitorS.
func _FcallbackForDebugLog_service_dataChan(___VacceptTcp *_TserviceTCP ) {
    //_Fpf( "283822 service" ); _Pn( )
    __VprStr := <- (*(*___VacceptTcp).Clog)
    _FpfN( "283823 service:%s" , __VprStr )
} // _FcallbackForDebugLog_service_dataChan
