// _TserviceTCP
package main

import (
    //"time"
)

func _FtcpAccept01(___VserviceTCP *_TserviceTCP ) {

    for _ , __VclientConn := range (*___VserviceTCP).clientConn {
        __VclientConn.Vbuf = make( []byte , 2048 )   // silice : with var len
    }

    defer (*___VserviceTCP).tcpLisn.Close() //_FtryListenToTCP01
    for ; ; {
        _Fsleep_10ms()
        _FtcpAccept01_loop( ___VserviceTCP )
    }
    (*(*___VserviceTCP).Cexit) <- "Error 381911: (" + (*___VserviceTCP).hostPortStr + ")"

} // _FtcpAccept01

func _FtcpAccept01_loop(___VserviceTCP *_TserviceTCP ) {

    //_, __Verr := (*___VserviceTCP).tcpLisn.Accept()
    __Vconn, __Verr := (*___VserviceTCP).tcpLisn.Accept()
    _FerrExit( " 188111 : tcp accept error " , __Verr )

    _FpfN( "accepting : max %d , now %d" , (*___VserviceTCP).cAmount , (*___VserviceTCP).clientCnt )
    if ( (*___VserviceTCP).cAmount > (*___VserviceTCP).clientCnt ) {
        _FpfN( "ok accept : max %d , now %d" , (*___VserviceTCP).cAmount , (*___VserviceTCP).clientCnt )
    } else {
        __Vconn.Close()
        _FpfN( "refuse accept : max %d , now %d" , (*___VserviceTCP).cAmount , (*___VserviceTCP).clientCnt )
    }
} // _FtcpAccept01_loop
