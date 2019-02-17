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

        (*___VserviceTCP).acceptTCPs[__Vi].Clog         = ___VserviceTCP.Clog
        (*___VserviceTCP).acceptTCPs[__Vi].Cexit        = ___VserviceTCP.Cexit

        (*___VserviceTCP).acceptTCPs[__Vi].Cstart       = make (chan string, 1  )
        (*___VserviceTCP).acceptTCPs[__Vi].CreceiveMsg  = make (chan []byte, 10 )
        (*___VserviceTCP).acceptTCPs[__Vi].CchanMsg     = make (chan []byte, 10 )
        (*___VserviceTCP).acceptTCPs[__Vi].CreceiveErr  = make (chan string, 1  )

        go _FhandleTcp_accept_dataReceiveMsg01( &((*___VserviceTCP).acceptTCPs[__Vi]) )
        go _FcallbackForDebugLog_accept_dataChan( &((*___VserviceTCP).acceptTCPs[__Vi]) )
    }

    defer (*___VserviceTCP).tcpLisn.Close() //_FtryListenToTCP01
    for ; ; {
        _Fsleep_1ms()
        _FtcpAccept01_loop( ___VserviceTCP )
    }
    (*(*___VserviceTCP).Cexit) <- "Error 381911: (" + (*___VserviceTCP).hostPortStr + ")"

} // _FtcpAccept01

func _FtcpAccept01_loop(___VserviceTCP *_TserviceTCP ) {

    // func (l *TCPListener) AcceptTCP() (*TCPConn, error)
    __Vconn, __Verr := (*___VserviceTCP).tcpLisn.AcceptTCP()
    _FerrExit( " 381810 : tcp accept error " , __Verr )

    //func (c *TCPConn) Write(b []byte) (int, error)
    _FpfN( "accepting 1 :%s\n" , _self_sha )
    __Vconn.Write( _self_sha ) ; __Vconn.Write( []byte("\n") )
    _FpfN( "accepting 2 :%s\n" , _self_sha )

    //_FpfN( "381811 accepting : max %d , now %d" , (*___VserviceTCP).cAmount , (*___VserviceTCP).clientCnt )
    if ( (*___VserviceTCP).cAmount > (*___VserviceTCP).clientCnt ) {
        __Vcnt := (*___VserviceTCP).clientCnt
        for __Vi:=0; __Vi < (*___VserviceTCP).cAmount ; __Vi ++ {
            __VacceptTcp := &((*___VserviceTCP).acceptTCPs[__Vi])
            if ( (*__VacceptTcp).enabled == false ) {
                // func (c *TCPConn) LocalAddr() Addr
                // func (c *TCPConn) RemoteAddr() Addr
                (*__VacceptTcp).VlocalAddr         = (*__Vconn).LocalAddr()
                (*__VacceptTcp).VremoteAddr        = (*__Vconn).RemoteAddr()
                //_FpfN( "381813 l:%s r:%s"     , (*__VacceptTcp).VlocalAddr , (*__VacceptTcp).VremoteAddr )

                // acceptTcpINC / acceptTcpDEC : begin
                (*___VserviceTCP).clientMux.Lock()

                (*___VserviceTCP).clientCnt ++
                (*__VacceptTcp).enabled             = true
                (*__VacceptTcp).connTCP             = __Vconn

                (*___VserviceTCP).clientMux.Unlock()
                // acceptTcpINC / acceptTcpDEC : end

                //_FpfN( "381815 accept succeed : old %d , now %d" , __Vcnt , (*___VserviceTCP).clientCnt )

                (*__VacceptTcp).Cstart              <- " 183191 start: " + _FtimeNow()
                //(*__VacceptTcp).Cstart              = _FtimeNow()
                // _TserviceTCP 
                break ;
            }
        }

        //_FpfN( "381817 after accept : old %d , now %d" , __Vcnt , (*___VserviceTCP).clientCnt )
        _FeqExit( "381819 err : why not inc ? " , __Vcnt , (*___VserviceTCP).clientCnt )

    } else {
        __Vconn.Close()
        _FpfN( "38181b refuse accept : max reach %d , now %d" , (*___VserviceTCP).cAmount , (*___VserviceTCP).clientCnt )
    }
} // _FtcpAccept01_loop
