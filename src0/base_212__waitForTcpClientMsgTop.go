// _TserviceTCP
package main

import (
    //"time"
)


func ( ___VserviceTcp *_TserviceTCP ) _FhandleWaitForClientMsgTcp__main_top() {
    if ___VserviceTcp.cAmount < 1        { ___VserviceTcp.cAmount = 1     }
    if ___VserviceTcp.cAmount > 100      { ___VserviceTcp.cAmount = 100   }
    ___VserviceTcp.acceptTCPs = make([]_TacceptTCP , ___VserviceTcp.cAmount )
    go _FtcpAccept01( ___VserviceTcp )

    for ; ; {
        _Fsleep_1ms()

        _FnotNullRun011_tcp_service_chan( ___VserviceTcp .TcallbackS , ___VserviceTcp )

    }
    *___VserviceTcp.Cexit <- "Error 183818: (" + ___VserviceTcp.hostPortStr + ")"
} // _FhandleWaitForClientMsgTcp__main_top

