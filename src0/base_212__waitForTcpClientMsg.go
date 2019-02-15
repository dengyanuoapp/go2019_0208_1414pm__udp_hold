// _TserviceTCP
package main

import (
    //"encoding/json"
    //"flag"
    //"fmt"
    //"log"
    //"net"
    "time"
)


// https://golang.org/pkg/net/#TCPConn.ReadFromTCP
// func (c *TCPConn) ReadFromTCP(b []byte) (int, *TCPAddr, error)
//func _FhandleWaitForClientMsgTcpTop(___VserviceTCP *_TserviceTCP, ___Cexit chan string , ___Clog chan string ) {
func _FhandleWaitForClientMsgTcpTop(___VserviceTCP *_TserviceTCP ) {
    if (*___VserviceTCP).cAmount < 1        { (*___VserviceTCP).cAmount = 1     }
    if (*___VserviceTCP).cAmount > 100      { (*___VserviceTCP).cAmount = 100   }

    (*___VserviceTCP).clientConn = make([]_TconnTCP , (*___VserviceTCP).cAmount )

    //(*___VserviceTCP).Vbuf = make( []byte , 2048 )   // silice : with var len
    for ; ; {
        time.Sleep(100 * time.Millisecond)
        //_FhandleWaitForClientMsgTcpLoop01( ___VserviceTCP )
    }
    (*(*___VserviceTCP).Cexit) <- "Error : (" + (*___VserviceTCP).hostPortStr + ")"
} // _FhandleWaitForClientMsgTcpTop

/*
func _FhandleWaitForClientMsgTcpLoop01(___VserviceTCP *_TserviceTCP ) {
    (*___VserviceTCP).Vlen,
    (*___VserviceTCP).VremoteAddr,
    (*___VserviceTCP).err =
    (*___VserviceTCP).tcpConn.ReadFromTCP((*___VserviceTCP).Vbuf)

    _FerrExit( (*___VserviceTCP).err )

    _FnullExit( " 183813 : why ___Vconn.ReadFromTCP addr error ?" , (*___VserviceTCP).VremoteAddr )

    _FnotNullRunTcp01( (*___VserviceTCP).callbackR , ___VserviceTCP )
} // _FhandleWaitForClientMsgTcpLoop01
*/
