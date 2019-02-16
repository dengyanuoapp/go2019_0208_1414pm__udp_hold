// _TconnTCP
// _TserviceTCP
// _FtcpAccept01
package main

import (
    //"time"
)



func _FhandleTcpReceiveMsg01(___VconnTCP *_TconnTCP ) {
    for {
        _Fsleep_1s()
        _Ppn( " 183891 : under constructing " , (*___VconnTCP).r64 )
        (*___VconnTCP).r64 ++ 
    }

    _Ppf( " 183892 : " )
    for __Vi , __VconnTcp := range (*(*___VconnTCP).serverTCP).clientTCPs {
        _Ppf( " %d,%d,%d" , __Vi , __VconnTcp.idx , __VconnTcp.r64  )
    }
    _Ppn()

    /*
    (*___VconnTCP).Vlen,
    (*___VconnTCP).VremoteAddr,
    (*___VconnTCP).err =
    (*___VconnTCP).tcpConn.ReadFromTCP((*___VconnTCP).Vbuf)

    _FerrExit( (*___VconnTCP).err )

    _FnullExit( " 183813 : why ___Vconn.ReadFromTCP addr error ?" , (*___VconnTCP).VremoteAddr )

    _FnotNullRunTcp01( (*___VconnTCP).callbackR , ___VconnTCP )
    */
} // _FhandleTcpReceiveMsg01
