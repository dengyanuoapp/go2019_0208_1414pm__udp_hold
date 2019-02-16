package main

// _TacceptTCP 
// _TserviceTCP 
func _FcallbackForDebugLog_service(___VacceptTcp *_TserviceTCP ) {
    //_Fpf( "283822 service" ); _Pn( )
    __VprStr := <- (*(*___VacceptTcp).Clog)
    _FpfN( "283823 service:%s" , __VprStr )
} // _FcallbackForDebugLog_service
