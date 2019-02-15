package main

import (
    //"encoding/json"
    //"flag"
    //"fmt"
    //"log"
    "net"
)

type _TconnTCP struct {
} // connTCP 

type _TserviceTCP struct {
    name                string
    hostPortStr         string
    cAmount             int

    tcpAddr             *net.TCPAddr
    tcpLisn             *net.TCPListener
    err                 error

    clientConn          [] _TconnTCP

    callbackR   func( *_TserviceUDP)
    callbackW   func( *_TserviceUDP)
    Cexit       *chan string
    Clog        *chan string
} // _TserviceTCP 


func _FtryListenToTCP01( ___Vsvr *_TserviceTCP ) {
    // func ResolveTCPAddr(network, address string) (*TCPAddr, error)
    (*___Vsvr).tcpAddr , (*___Vsvr).err  = net.ResolveTCPAddr("tcp4", (*___Vsvr).hostPortStr)
    if (*___Vsvr).err != nil {
        _Fex( "err13813" , (*___Vsvr).err)
    }

    // func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
    (*___Vsvr).tcpLisn , (*___Vsvr).err  = net.ListenTCP("tcp4", (*___Vsvr).tcpAddr )
    if (*___Vsvr).err != nil {
        _Fex( "err13814" , (*___Vsvr).err)
    }
} // _FtryListenToTCP01
