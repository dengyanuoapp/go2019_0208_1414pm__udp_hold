package main

// _TacceptTCP 
// _TserviceTCP 
func _FcallbackForDebugMo(___VacceptTcp *_TacceptTCP ) {
    _Fpf( "|l:%s|r:%s|" , (*___VacceptTcp).VlocalAddr , (*___VacceptTcp).VremoteAddr )
    _PpdN( (*___VacceptTcp).Vlen , &(*___VacceptTcp).Vbuf )
} // _FcallbackForDebugMo
