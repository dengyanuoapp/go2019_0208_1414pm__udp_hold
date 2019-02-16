// _TserviceTCP
package main

import (
    "time"
)


func _FhandleWaitForClientMsgTcpTop(___VserviceTCP *_TserviceTCP ) {
    if (*___VserviceTCP).cAmount < 1        { (*___VserviceTCP).cAmount = 1     }
    if (*___VserviceTCP).cAmount > 100      { (*___VserviceTCP).cAmount = 100   }
    (*___VserviceTCP).clientConn = make([]_TconnTCP , (*___VserviceTCP).cAmount )
    go _FtcpAccept01( ___VserviceTCP )

    for ; ; {
        time.Sleep(100 * time.Millisecond)
        //_FhandleWaitForClientMsgTcpLoop01( ___VserviceTCP )
    }
    (*(*___VserviceTCP).Cexit) <- "Error 183818: (" + (*___VserviceTCP).hostPortStr + ")"
} // _FhandleWaitForClientMsgTcpTop

