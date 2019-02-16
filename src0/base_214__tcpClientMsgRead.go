// _TserviceTCP
package main

import (
    "time"
)



func _FhandleTcpReceiveMsg01(___VserviceTCP *_TserviceTCP ) {
    (*___VserviceTCP).Vlen,
    (*___VserviceTCP).VremoteAddr,
    (*___VserviceTCP).err =
    (*___VserviceTCP).tcpConn.ReadFromTCP((*___VserviceTCP).Vbuf)

    _FerrExit( (*___VserviceTCP).err )

    _FnullExit( " 183813 : why ___Vconn.ReadFromTCP addr error ?" , (*___VserviceTCP).VremoteAddr )

    _FnotNullRunTcp01( (*___VserviceTCP).callbackR , ___VserviceTCP )
} // _FhandleTcpReceiveMsg01
