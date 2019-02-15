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
    hostPortStr string

    udpAddr     *net.UDPAddr
    udpConn     *net.UDPConn
    err         error

    Vbuf        []byte
    Vlen        int
    VremoteAddr *net.UDPAddr
    VlocalAddr  net.Addr

    callbackR   func( *_TserviceUDP)
    callbackW   func( *_TserviceUDP)
    Cexit       *chan string
    Clog        *chan string
}


func _FtryListenToUDP01( ___Vsvr *_TserviceUDP ) {
    // func ResolveUDPAddr(network, address string) (*UDPAddr, error)
    (*___Vsvr).udpAddr , (*___Vsvr).err  = net.ResolveUDPAddr("udp4", (*___Vsvr).hostPortStr)
    if (*___Vsvr).err != nil {
        _Fex( "err13811" , (*___Vsvr).err)
    }

    // func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
    (*___Vsvr).udpConn , (*___Vsvr).err  = net.ListenUDP("udp4", (*___Vsvr).udpAddr )
    if (*___Vsvr).err != nil {
        _Fex( "err13812" , (*___Vsvr).err)
    }
} // _FtryListenToUDP01
