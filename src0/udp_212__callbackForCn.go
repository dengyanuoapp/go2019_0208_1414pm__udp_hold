package main

func _FcallbackInFnForCn(___VserviceUDP *_TserviceUDP ) {
    _Fpf( "|l:%s|r:%s|" , (*___VserviceUDP).hostPortStr , (*___VserviceUDP).VremoteAddr.String() )
    _PpdN( (*___VserviceUDP).Vlen , &(*___VserviceUDP).Vbuf )
} // _FcallbackInFnForCn