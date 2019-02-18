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
//func _FhandleWaitForClientMsgUdp_main_top(___VserviceUdp *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
func (___VserviceUdp *_TserviceUDP ) _FhandleWaitForClientMsgUdp__read_main_top (){

    ___VserviceUdp.Vbuf         = make(      []byte , 2048 )   // silice : with var len
    ___VserviceUdp.CuIn01       = make( chan []byte , 5    )   // silice : with var len

    //    //func (c *UDPConn) LocalAddr() Addr
    ___VserviceUdp.VlocalAddr = ___VserviceUdp.udpConn.LocalAddr()

    for ; ; {
        _FhandleWaitForClientMsgUdp__read_main_loop( ___VserviceUdp )
    }
    (*___VserviceUdp.Cexit) <- "Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _FhandleWaitForClientMsgUdp__read_main_top

func _FhandleWaitForClientMsgUdp__read_main_loop(___VserviceUdp *_TserviceUDP ) {
    ___VserviceUdp.Vlen,
    ___VserviceUdp.VremoteAddr,
    ___VserviceUdp.err =
    ___VserviceUdp.udpConn.ReadFromUDP(___VserviceUdp.Vbuf)

    _FerrExit( "err 183911 : udp reading : " , ___VserviceUdp.err )

    _FnullExit( " err 183813 : why ___Vconn.ReadFromUDP addr error ?" , ___VserviceUdp.VremoteAddr )

    _FnotNullRunUdp01( ___VserviceUdp.UcallbackR , ___VserviceUdp )
} // _FhandleWaitForClientMsgUdp__read_main_loop
