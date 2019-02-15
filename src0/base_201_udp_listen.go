package main

import (
    //"encoding/json"
    //"flag"
    //"fmt"
    //"log"
    "net"
)

type _TserviceUDP struct {
    name        string
    port        string
    laddr       string
    udpAddr     *net.UDPAddr
    udpConn     *net.UDPConn
    err         error
}


func _FtryListenToUDP01( ___Vsvr *_TserviceUDP ) {
    // func ResolveUDPAddr(network, address string) (*UDPAddr, error)
    (*___Vsvr).udpAddr , (*___Vsvr).err  = net.ResolveUDPAddr("udp4", (*___Vsvr).port)
    if (*___Vsvr).err != nil {
        _Fex( "err13811" , (*___Vsvr).err)
    }

    // func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
    (*___Vsvr).udpConn , (*___Vsvr).err  = net.ListenUDP("udp", (*___Vsvr).udpAddr )
    if (*___Vsvr).err != nil {
        _Fex( "err13812" , (*___Vsvr).err)
    }
} // _FtryListenToUDP01

