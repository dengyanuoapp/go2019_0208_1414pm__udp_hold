// _TserviceTCP
package main

import (
    //"time"
)

func _Fhandle_tcpAccept01(___VserviceTcp *_TserviceTCP ) {

    for __Vi:=0; __Vi < ___VserviceTcp.cAmount ; __Vi ++ {
        ___VserviceTcp.acceptTCPs[__Vi].Vbuf         = make( []byte , 2048 )   // silice : with var len
        ___VserviceTcp.acceptTCPs[__Vi].idx          = __Vi
        ___VserviceTcp.acceptTCPs[__Vi].enabled      = false
        ___VserviceTcp.acceptTCPs[__Vi].serverTCP    = ___VserviceTcp

        ___VserviceTcp.acceptTCPs[__Vi].Clog         = ___VserviceTcp.Clog
        ___VserviceTcp.acceptTCPs[__Vi].Cexit        = ___VserviceTcp.Cexit

        ___VserviceTcp.acceptTCPs[__Vi].Cstart       = make (chan string, 1  )
        ___VserviceTcp.acceptTCPs[__Vi].CreceiveMsg  = make (chan []byte, 10 )
        ___VserviceTcp.acceptTCPs[__Vi].CchanMsg     = make (chan []byte, 10 )
        ___VserviceTcp.acceptTCPs[__Vi].CreceiveErr  = make (chan string, 1  )

        go ___VserviceTcp.acceptTCPs[__Vi] . _FhandleTcp_accept_dataReceiveMsg01( )
        go ___VserviceTcp.acceptTCPs[__Vi] . _FhandleTcp__accept_dataChan__main_top( )
    }

    defer ___VserviceTcp.tcpListener.Close() //_FtryListenToTCP01
    for ; ; {
        _Fsleep_1ms()
        _FtcpAccept01_loop( ___VserviceTcp )
    }
    (*___VserviceTcp.Cexit) <- "Error 381911: (" + ___VserviceTcp.hostPortStr + ")"

} // _Fhandle_tcpAccept01

func _FtcpAccept01_loop(___VserviceTcp *_TserviceTCP ) {

    // func (l *TCPListener) AcceptTCP() (*TCPConn, error)
    __Vconn, __Verr := ___VserviceTcp.tcpListener.AcceptTCP()
    _FerrExit( " 381810 : tcp accept error " , __Verr )

    //func (c *TCPConn) Write(b []byte) (int, error)
    //_FpfN( "381812 accepting 1 :%x\n" , _self_sha )
    __Vconn.Write( []byte(_Pspf("%x\n",_self_sha)) ) 
    //_FpfN( "381814 accepting 2 :%s\n" , _self_sha )

    //_FpfN( "381816 accepting : max %d , now %d" , ___VserviceTcp.cAmount , ___VserviceTcp.clientCnt )
    if ( ___VserviceTcp.cAmount > ___VserviceTcp.clientCnt ) {
        __Vcnt := ___VserviceTcp.clientCnt
        for __Vi:=0; __Vi < ___VserviceTcp.cAmount ; __Vi ++ {
            __VacceptTcp := &(___VserviceTcp.acceptTCPs[__Vi])
            if ( __VacceptTcp.enabled == false ) {
                // func (c *TCPConn) LocalAddr() Addr
                // func (c *TCPConn) RemoteAddr() Addr
                __VacceptTcp.VlocalAddr         = __Vconn.LocalAddr()
                __VacceptTcp.VremoteAddr        = __Vconn.RemoteAddr()
                //_FpfN( "381818 l:%s r:%s"     , __VacceptTcp.VlocalAddr , __VacceptTcp.VremoteAddr )

                // acceptTcpINC / acceptTcpDEC : begin
                ___VserviceTcp.clientMux.Lock()

                ___VserviceTcp.clientCnt ++
                __VacceptTcp.enabled             = true
                __VacceptTcp.connTCP             = __Vconn

                ___VserviceTcp.clientMux.Unlock()
                // acceptTcpINC / acceptTcpDEC : end

                //_FpfN( "381815 accept succeed : old %d , now %d" , __Vcnt , ___VserviceTcp.clientCnt )

                __VacceptTcp.Cstart              <- " 183191 start: " + _FtimeNow()
                //__VacceptTcp.Cstart              = _FtimeNow()
                // _TserviceTCP 
                break ;
            }
        }

        //_FpfN( "381817 after accept : old %d , now %d" , __Vcnt , ___VserviceTcp.clientCnt )
        _FeqExit( "381819 err : why not inc ? " , __Vcnt , ___VserviceTcp.clientCnt )

    } else {
        __Vconn.Close()
        _FpfN( "38181b refuse accept : max reach %d , now %d" , ___VserviceTcp.cAmount , ___VserviceTcp.clientCnt )
    }
} // _FtcpAccept01_loop
