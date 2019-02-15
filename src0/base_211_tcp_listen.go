package main

import (
    //"encoding/json"
    //"flag"
    //"fmt"
    //"log"
    "net"
)

type _TserviceTCP struct {
    name        string
    port        string
    laddr       string
    tcpAddr     *net.TCPAddr
    tcpLisn     *net.TCPListener
    err         error
}


func _FtryListenToTCP01( ___Vsvr *_TserviceTCP ) {
    // func ResolveTCPAddr(network, address string) (*TCPAddr, error)
    (*___Vsvr).tcpAddr , (*___Vsvr).err  = net.ResolveTCPAddr("tcp4", (*___Vsvr).port)
    if (*___Vsvr).err != nil {
        _Fex( "err13811" , (*___Vsvr).err)
    }

    // func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
    (*___Vsvr).tcpLisn , (*___Vsvr).err  = net.ListenTCP("tcp", (*___Vsvr).tcpAddr )
    if (*___Vsvr).err != nil {
        _Fex( "err13812" , (*___Vsvr).err)
    }
} // _FtryListenToTCP01

