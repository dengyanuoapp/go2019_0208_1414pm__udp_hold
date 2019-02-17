package main

// _TacceptTCP 
// _TserviceTCP 
func _FcallbackForDebugLog_accept_dataReceive(___VacceptTcp *_TacceptTCP ) {

    /*
    _Fpf( "188111 tcpRdump|l:%s|r:%s|" , (*___VacceptTcp).VlocalAddr , (*___VacceptTcp).VremoteAddr )
    _PpdN( (*___VacceptTcp).Vlen , &(*___VacceptTcp).Vbuf )
    */
    /*
    (*(*___VacceptTcp).Clog) <- _Pspf( "188112 tcpRdump |l:%s|r:%s|(%d)" , 
    (*___VacceptTcp).VlocalAddr , (*___VacceptTcp).VremoteAddr , (*___VacceptTcp).Vlen )
    */
    //(*(*___VacceptTcp).CreceiveMsg) <- make([]byte , (*___VacceptTcp).Vbuf[0:((*___VacceptTcp).Vlen-1)] )
    //(*___VacceptTcp).CreceiveMsg <- (*___VacceptTcp).Vbuf
    //(*___VacceptTcp).CreceiveMsg <- (*___VacceptTcp).Vbuf[0:((*___VacceptTcp).Vlen-1)] 
    //(*___VacceptTcp).CreceiveMsg <- make([]byte , (*___VacceptTcp).Vbuf[0:((*___VacceptTcp).Vlen-1)] , (*___VacceptTcp).Vlen )
    __Vbyte := make([]byte , (*___VacceptTcp).Vlen )
    copy(__Vbyte , (*___VacceptTcp).Vbuf[0:((*___VacceptTcp).Vlen-1)] )
    (*___VacceptTcp).CreceiveMsg <- __Vbyte

} // _FcallbackForDebugLog_accept_dataReceive
