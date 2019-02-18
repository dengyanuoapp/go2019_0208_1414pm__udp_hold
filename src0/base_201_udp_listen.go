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

    UcallbackR  func( *_TserviceUDP) // deal with data rece from udp
    UcallbackC  func( *_TserviceUDP) // deal with data rece from chan
    Cexit       *chan string
    Clog        *chan string

    CuIn01      chan []byte
    CuOut01     *chan []byte

}


func ( ___Vsvr *_TserviceUDP )_FtryListenToUDP01()  {
    // func ResolveUDPAddr(network, address string) (*UDPAddr, error)
    ___Vsvr.udpAddr , ___Vsvr.err  = net.ResolveUDPAddr("udp4", ___Vsvr.hostPortStr)
    if ___Vsvr.err != nil {
        _Fex( "err13811" , ___Vsvr.err)
    }

    // func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
    ___Vsvr.udpConn , ___Vsvr.err  = net.ListenUDP("udp4", ___Vsvr.udpAddr )
    if ___Vsvr.err != nil {
        _Fex( "err13812" , ___Vsvr.err)
    }
} // _FtryListenToUDP01

