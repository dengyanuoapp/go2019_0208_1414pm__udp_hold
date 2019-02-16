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

    /*
    (*___VacceptTCP).Vlen,
    (*___VacceptTCP).VremoteAddr,
    (*___VacceptTCP).err =
    (*___VacceptTCP).tcpConn.ReadFromTCP((*___VacceptTCP).Vbuf)

    _FerrExit( (*___VacceptTCP).err )

    _FnullExit( " 183813 : why ___Vconn.ReadFromTCP addr error ?" , (*___VacceptTCP).VremoteAddr )

    _FnotNullRunTcp01( (*___VacceptTCP).callbackR , ___VacceptTCP )
    */
} // _FhandleTcpReceiveMsg01

// _TacceptTCP 
func _FhandleTcpReceiveMsg01_loop(___VacceptTCP *_TacceptTCP ) {
        _Ppn( " 183891 : under constructing " , (*___VacceptTCP).r64 )
        (*___VacceptTCP).r64 ++ 

        _Ppn( " 183892 : " , (*___VacceptTCP).serverTCP )
        /*
        _Ppf( " 183893 : " )
        for __Vi , __VacceptTcp := range (*(*___VacceptTCP).serverTCP).acceptTCPs {
            //_Ppf( " %d,%d,%d" , __Vi , __VacceptTcp.idx , __VacceptTcp.r64  )
            _Ppf( " %d" , __Vi )
        }
        */
        _Ppn()
} // _FhandleTcpReceiveMsg01_loop
