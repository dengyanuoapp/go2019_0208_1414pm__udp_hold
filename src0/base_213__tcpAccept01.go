// _TserviceTCP
package main

import (
    //"time"
)

func _FtcpAccept01(___VserviceTCP *_TserviceTCP ) {

    for __Vidx , __VclientConn := range (*___VserviceTCP).clientConn {
        __VclientConn.Vbuf      = make( []byte , 2048 )   // silice : with var len
        __VclientConn.idx       = __Vidx
        __VclientConn.enabled   = false
        __VclientConn.serverTCP = ___VserviceTCP
    }

    defer (*___VserviceTCP).tcpLisn.Close() //_FtryListenToTCP01
    for ; ; {
        _Fsleep_10ms()
        _FtcpAccept01_loop( ___VserviceTCP )
    }
    (*(*___VserviceTCP).Cexit) <- "Error 381911: (" + (*___VserviceTCP).hostPortStr + ")"

} // _FtcpAccept01

func _FtcpAccept01_loop(___VserviceTCP *_TserviceTCP ) {

    // func (l *TCPListener) AcceptTCP() (*TCPConn, error)
    __Vconn, __Verr := (*___VserviceTCP).tcpLisn.AcceptTCP()
    _FerrExit( " 188111 : tcp accept error " , __Verr )

    _FpfN( "accepting : max %d , now %d" , (*___VserviceTCP).cAmount , (*___VserviceTCP).clientCnt )
    if ( (*___VserviceTCP).cAmount > (*___VserviceTCP).clientCnt ) {
        __Vcnt := (*___VserviceTCP).clientCnt
        for _ , __VclientConn := range (*___VserviceTCP).clientConn {
            if ( __VclientConn.enabled == false ) {
                (*___VserviceTCP).clientMux.Lock()

                (*___VserviceTCP).clientCnt ++
                __VclientConn.enabled       = true
                __VclientConn.conn          = __Vconn

                (*___VserviceTCP).clientMux.Unlock()
                _FpfN( "accept succeed : old %d , now %d" , __Vcnt , (*___VserviceTCP).clientCnt )

                go _FhandleTcpReceiveMsg01( __VclientConn )
                // _TserviceTCP 
                break ;
            }
        }

        _FpfN( "after accept : old %d , now %d" , __Vcnt , (*___VserviceTCP).clientCnt )
        _FeqExit( "err 138191 : why not inc ? " , __Vcnt , (*___VserviceTCP).clientCnt )

    } else {
        __Vconn.Close()
        _FpfN( "refuse accept : max %d , now %d" , (*___VserviceTCP).cAmount , (*___VserviceTCP).clientCnt )
    }
} // _FtcpAccept01_loop
