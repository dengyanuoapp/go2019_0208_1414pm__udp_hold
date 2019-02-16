// _TserviceTCP
package main

import (
    //"time"
)

func _FtcpAccept01(___VserviceTCP *_TserviceTCP ) {

    for __Vi:=0; __Vi < (*___VserviceTCP).cAmount ; __Vi ++ {
        (*___VserviceTCP).acceptTCPs[__Vi].Vbuf         = make( []byte , 2048 )   // silice : with var len
        (*___VserviceTCP).acceptTCPs[__Vi].idx          = __Vi
        (*___VserviceTCP).acceptTCPs[__Vi].enabled      = false
        (*___VserviceTCP).acceptTCPs[__Vi].serverTCP    = ___VserviceTCP

        (*___VserviceTCP).acceptTCPs[__Vi].Cin01        = make (chan string, 10 )
        (*___VserviceTCP).acceptTCPs[__Vi].Cout01       = make (chan string, 10 )
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
        for __Vi:=0; __Vi < (*___VserviceTCP).cAmount ; __Vi ++ {
            __VclientConn := &((*___VserviceTCP).acceptTCPs[__Vi])
            if ( (*__VclientConn).enabled == false ) {
                // acceptTcpINC / acceptTcpDEC : begin
                (*___VserviceTCP).clientMux.Lock()

                (*___VserviceTCP).clientCnt ++
                (*__VclientConn).enabled       = true
                (*__VclientConn).connTCP       = __Vconn

                (*___VserviceTCP).clientMux.Unlock()
                // acceptTcpINC / acceptTcpDEC : end

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
