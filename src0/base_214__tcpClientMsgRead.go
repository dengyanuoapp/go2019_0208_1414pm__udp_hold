// _TserviceTCP
package main

import (
    //"time"
)



func _FhandleTcpReceiveMsg01(___VconnTCP *_TconnTCP ) {
    for {
        _Fsleep_1s()
        _Ppn( " 183891 : under constructing " )
    }
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
