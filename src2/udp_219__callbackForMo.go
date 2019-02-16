package main

func _FcallbackForDebugMo(___VserviceTCP *_TserviceTCP ) {
    _Fpf( "|l:%s|r:%s|" , (*___VserviceTCP).VlocalAddr , (*___VserviceTCP).VremoteAddr )
    _PpdN( (*___VserviceTCP).Vlen , &(*___VserviceTCP).Vbuf )
} // _FcallbackForDebugMo
