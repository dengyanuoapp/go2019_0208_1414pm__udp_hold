// _TserviceTCP
package main

import (
    "time"
)

func _FtcpAccept01(___VserviceTCP *_TserviceTCP ) {
    for _ , __VclientConn := range (*___VserviceTCP).clientConn {
        __VclientConn.Vbuf = make( []byte , 2048 )   // silice : with var len
    }

    for ; ; {
        time.Sleep(100 * time.Millisecond)
        _FtcpAccept01_loop( ___VserviceTCP )
    }
    (*(*___VserviceTCP).Cexit) <- "Error 381911: (" + (*___VserviceTCP).hostPortStr + ")"
} // _FtcpAccept01

func _FtcpAccept01_loop(___VserviceTCP *_TserviceTCP ) {
} // _FtcpAccept01_loop
