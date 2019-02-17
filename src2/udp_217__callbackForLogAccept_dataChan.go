package main

// _TacceptTCP 
// _TserviceTCP 
// _FhandleTcp_accept_dataReceiveMsg01
func _FcallbackForDebugLog_accept_dataChan(___VacceptTcp *_TacceptTCP ) {

    var __Vmsg      []byte
    var __Vstr      string

    for {
        _Fsleep_1ms()
        select {
        case __Vstr = <- (*___VacceptTcp).CreceiveErr :
            _FpfN( " %d : 181181 accept_dataChan receERR : %d , %s"     , (*___VacceptTcp).idx , len(__Vstr) , __Vstr )
        case __Vmsg = <- (*___VacceptTcp).CreceiveMsg :
            _FpfN( " %d : 181183 accept_dataChan receMsg : %d %d , %s"  , (*___VacceptTcp).idx , len(__Vmsg) , cap(__Vmsg) , __Vmsg )
        case __Vmsg = <- (*___VacceptTcp).CchanMsg :
            _FpfN( " %d : 181185 accept_dataChan chanMsg : %d %d , %s"  , (*___VacceptTcp).idx , len(__Vmsg) , cap(__Vmsg) , __Vmsg )
        }

        //_FhandleTcp_accept_dataReceiveMsg01__loop( ___VacceptTCP )
    }

} // _FcallbackForDebugLog_accept_dataChan
