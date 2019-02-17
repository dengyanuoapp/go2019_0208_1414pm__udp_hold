// _TacceptTCP
// _TserviceTCP
// _FtcpAccept01
package main

import (
    "io"
)



func _FhandleTcpReceiveMsg01(___VacceptTCP *_TacceptTCP ) {
    //var __VcanReceiveMsg bool = true
    for {
        if ( 3 == 2 ) {
            <-(*___VacceptTCP).Cstart
        } else {
            _FpfN( " 188118 rece run start at " + <-(*___VacceptTCP).Cstart )
        }

        //if ( false == (*___VacceptTCP).enabled ) { continue }

        for {
            _Fsleep_1ms()

            //__VcanReceiveMsg = (*___VacceptTCP).enabled
            if false == (*___VacceptTCP).enabled        { break }

            _FhandleTcpReceiveMsg01_loop( ___VacceptTCP )
        }
    }

} // _FhandleTcpReceiveMsg01

// func (c *TCPConn) Read(b []byte) (int, error)
// _TacceptTCP
func _FhandleTcpReceiveMsg01_loop(___VacceptTCP *_TacceptTCP ) bool {
    if ( 4 == 2 ) { _FhandleTcpReceiveMsg01_Debug(___VacceptTCP ) }

    (*___VacceptTCP).r64try ++

    (*___VacceptTCP).Vlen,
    (*___VacceptTCP).Verr =
    (*___VacceptTCP).connTCP.Read((*___VacceptTCP).Vbuf)

    // _FtcpAccept01
    if ( (*___VacceptTCP).Verr == io.EOF ) { // lost the connect.
        // acceptTcpINC / acceptTcpDEC : begin
        (*(*___VacceptTCP).serverTCP)   .clientMux.Lock()

        (*(*___VacceptTCP).serverTCP)   .clientCnt --
        (*___VacceptTCP)                .enabled       = false
        (*___VacceptTCP)                .connTCP.Close()

        (*(*___VacceptTCP).serverTCP)   .clientMux.Unlock()
        // acceptTcpINC / acceptTcpDEC : end
        return false
    }

    _FerrExit( " reading from tcp 831911 " , (*___VacceptTCP).Verr )

    _FnullExit( " 183813 : why ___Vconn.ReadFromTCP addr error ?" , (*___VacceptTCP).VremoteAddr )

    (*___VacceptTCP).r64ok ++

    /*
    _Fpf( " 183814 | l:%s | r:%s | " , (*___VacceptTCP).VlocalAddr , (*___VacceptTCP).VremoteAddr )
    _PpdN( (*___VacceptTCP).Vlen , &((*___VacceptTCP).Vbuf) )
    */

    // _FcallbackForDebugLog_accept
    _FnotNullRunTcp02_accept( (*(*___VacceptTCP).serverTCP) .callbackR , ___VacceptTCP )

    return true
} // _FhandleTcpReceiveMsg01_loop

func _FhandleTcpReceiveMsg01_Debug(___VacceptTCP *_TacceptTCP ) {

    _Ppn( " 183891 : under constructing " , (*___VacceptTCP).r64try , (*___VacceptTCP).r64ok )

    if ( nil == (*___VacceptTCP).serverTCP ) {
        _Ppn( " 183892 : (*___VacceptTCP).serverTCP == nil " )
    } else {

        _Ppf( " 183893 : " )
        for __Vi:=0; __Vi < (*(*___VacceptTCP).serverTCP) .cAmount ; __Vi ++ {
            __VacceptTcp := &((*(*___VacceptTCP).serverTCP).acceptTCPs[__Vi])
            _Ppf( " %d,%d,%d,%d" , __Vi , (*__VacceptTcp).idx , (*__VacceptTcp).r64try  , (*__VacceptTcp).r64ok  )
        }
        _Ppn()
    }
} // _FhandleTcpReceiveMsg01_Debug
