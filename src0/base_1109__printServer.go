package main

import (
)

func _FhandleDebugServer(___VserviceUdp *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
    __Vbuf := make( []byte , 2048 )   // silice : with var len

    __Vlen, __Vaddr, __Verr := ___VserviceUdp.udpConn.ReadFromUDP(__Vbuf)

    if __Verr != nil {
        return
    }
    if __Vaddr == nil {
        _Fex( " 183811 : why ___Vconn.ReadFromUDP addr error ?" , nil )
    }

    //_FpdN( __Vlen , &__Vbuf )

    _Fpf( "|%s|" , __Vaddr )
    _PpdLN( __Vlen , &__Vbuf )

} // _FhandleFnClientCn
