package main

// _TacceptTCP 
// _TserviceTCP 
// _FhandleTcp_accept_dataReceiveMsg01
func _FcallbackForDebugLog_accept_dataChan(___VacceptTcp *_TacceptTCP ) {

    var __Vmsg      []byte
    var __Vstr      string

    for {
        select {
        case __Vstr = <- (*___VacceptTcp).CreceiveErr :
            _FpfN( " 181181 accept_dataChan receERR : %d , %s" , len(__Vstr) , __Vstr )
        case __Vmsg = <- (*___VacceptTcp).CreceiveMsg :
            _FpfN( " 181183 accept_dataChan receMsg : %d , %s" , len(__Vmsg) , __Vmsg )
        case __Vmsg = <- (*___VacceptTcp).CchanMsg :
            _FpfN( " 181185 accept_dataChan chanMsg : %d , %s" , len(__Vmsg) , __Vmsg )
        }

        //_FhandleTcp_accept_dataReceiveMsg01__loop( ___VacceptTCP )
    }

} // _FcallbackForDebugLog_accept_dataChan
