package main

// _TacceptTCP 
// _TserviceTCP 
func _FcallbackForDebugMo(___VacceptTcp *_TacceptTCP ) {
    _Fpf( "188111 tcpRdump|l:%s|r:%s|" , (*___VacceptTcp).VlocalAddr , (*___VacceptTcp).VremoteAddr )
    _PpdN( (*___VacceptTcp).Vlen , &(*___VacceptTcp).Vbuf )
} // _FcallbackForDebugMo
