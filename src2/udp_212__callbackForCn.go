package main

func _FcallbackInFnForCn(___VserviceUDP *_TserviceUDP ) {
    _Fpf( "341011 udpRdump |l:%s|r:%s|" , (*___VserviceUDP).VlocalAddr , (*___VserviceUDP).VremoteAddr )
    _PpdN( (*___VserviceUDP).Vlen , &(*___VserviceUDP).Vbuf )
} // _FcallbackInFnForCn
