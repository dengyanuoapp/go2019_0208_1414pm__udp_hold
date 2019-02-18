// _TserviceUDP
package main

import (
    //"encoding/json"
    //"flag"
    //"fmt"
    //"log"
    //"net"
)


// https://golang.org/pkg/net/#UDPConn.ReadFromUDP
// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
//func _FhandleWaitForClientMsgUdpTop(___VserviceUdp *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
func _FhandleWaitForClientMsgUdpTop(___VserviceUdp *_TserviceUDP ) {

    ___VserviceUdp.Vbuf = make( []byte , 2048 )   // silice : with var len
    //    //func (c *UDPConn) LocalAddr() Addr
    ___VserviceUdp.VlocalAddr = ___VserviceUdp.udpConn.LocalAddr()

    for ; ; {
        _FhandleWaitForClientMsgUdpLoop01( ___VserviceUdp )
    }
    (*___VserviceUdp.Cexit) <- "Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _FhandleWaitForClientMsgUdpTop

func _FhandleWaitForClientMsgUdpLoop01(___VserviceUdp *_TserviceUDP ) {
    ___VserviceUdp.Vlen,
    ___VserviceUdp.VremoteAddr,
    ___VserviceUdp.err =
    ___VserviceUdp.udpConn.ReadFromUDP(___VserviceUdp.Vbuf)

    _FerrExit( "err 183911 : udp reading : " , ___VserviceUdp.err )

    _FnullExit( " err 183813 : why ___Vconn.ReadFromUDP addr error ?" , ___VserviceUdp.VremoteAddr )

    _FnotNullRunUdp01( ___VserviceUdp.UcallbackR , ___VserviceUdp )
} // _FhandleWaitForClientMsgUdpLoop01
