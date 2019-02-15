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
//func _FhandleFnWaitForUdpTop(___VserviceUDP *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
func _FhandleFnWaitForUdpTop(___VserviceUDP *_TserviceUDP ) {

    (*___VserviceUDP).Vbuf = make( []byte , 2048 )   // silice : with var len
    for ; ; {
        _FhandleFnWaitForClientCnLoop( ___VserviceUDP )
    }
    (*(*___VserviceUDP).Cexit) <- "Error : (" + (*___VserviceUDP).hostPortStr + ")"
} // _FhandleFnWaitForUdpTop

func _FhandleFnWaitForClientCnLoop(___VserviceUDP *_TserviceUDP ) {
    (*___VserviceUDP).Vlen,
    (*___VserviceUDP).VremoteAddr,
    (*___VserviceUDP).err =
    (*___VserviceUDP).udpConn.ReadFromUDP((*___VserviceUDP).Vbuf)

    _FerrExit( (*___VserviceUDP).err )

    _FnullExit( " 183813 : why ___Vconn.ReadFromUDP addr error ?" , (*___VserviceUDP).VremoteAddr )

    //    //func (c *UDPConn) LocalAddr() Addr
    //    //(*___VserviceUDP).VlocalAddr = (*___VserviceUDP).udpConn.LocalAddr()
    //    //_Fpf( "|%s|r:%s|me:%s|" , (*___VserviceUDP).hostPortStr , (*___VserviceUDP).VremoteAddr.String() , (*___VserviceUDP).VlocalAddr.String() )
    //    _Fpf( "|l:%s|r:%s|" , (*___VserviceUDP).hostPortStr , (*___VserviceUDP).VremoteAddr.String() )
    //    _PpdN( (*___VserviceUDP).Vlen , &(*___VserviceUDP).Vbuf )

    _FnotNullRun( (*___VserviceUDP).callbackR , ___VserviceUDP )
} // _FhandleFnWaitForClientCnLoop
