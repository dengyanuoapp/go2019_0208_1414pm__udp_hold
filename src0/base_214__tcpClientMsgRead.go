// _TacceptTCP
// _TserviceTCP
// _FtcpAccept01
package main

import (
    //"time"
)



func _FhandleTcpReceiveMsg01(___VacceptTCP *_TacceptTCP ) {
    for {
        _Fsleep_1s()
        _FhandleTcpReceiveMsg01_loop( ___VacceptTCP ) 

    }

} // _FhandleTcpReceiveMsg01

// func (c *TCPConn) Read(b []byte) (int, error)
// _TacceptTCP 
func _FhandleTcpReceiveMsg01_loop(___VacceptTCP *_TacceptTCP ) {
    //_FhandleTcpReceiveMsg01_Debug(___VacceptTCP )

    (*___VacceptTCP).r64try ++ 

    (*___VacceptTCP).Vlen,
    (*___VacceptTCP).Verr =
    (*___VacceptTCP).connTCP.Read((*___VacceptTCP).Vbuf)

    _FerrExit( " reading from tcp 31911 " , (*___VacceptTCP).Verr )

    _FnullExit( " 183813 : why ___Vconn.ReadFromTCP addr error ?" , (*___VacceptTCP).VremoteAddr )

    (*___VacceptTCP).r64ok ++ 

    _Fpf( " 183814 : " )
    _PpdN( (*___VacceptTCP).Vlen , &((*___VacceptTCP).Vbuf) )
    /*

    _FnotNullRunTcp01( (*___VacceptTCP).callbackR , ___VacceptTCP )
    */

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
