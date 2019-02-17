package main

import (
    //"encoding/json"
    //"flag"
    //"fmt"
    //"log"
    "net"
    "sync"
)

type _TacceptTCP struct {
    enabled             bool
    idx                 int
    serverTCP           *_TserviceTCP
    connTCP             *net.TCPConn
    r64try              int64
    r64ok               int64
    err                 error

    Vbuf                []byte
    Vlen                int
    Verr                error
    VremoteAddr         net.Addr
    VlocalAddr          net.Addr

    Cstart              chan string
    CreceiveMsg         chan string
    CreceiveErr         chan string

    Cexit               *chan string
    Clog                *chan string
} // _TacceptTCP 

type _TserviceTCP struct {
    name                string
    hostPortStr         string
    cAmount             int

    tcpAddr             *net.TCPAddr
    tcpLisn             *net.TCPListener
    err                 error

    acceptTCPs          []_TacceptTCP
    clientMux           sync.Mutex
    clientCnt           int

    TcallbackS          func( *_TserviceTCP)    // _FcallbackForDebugLog_service
    TcallbackR          func( *_TacceptTCP)     // _FcallbackForDebugLog_accept_dataReceive
    TcallbackC          func( *_TacceptTCP)     // _FcallbackForDebugLog_accept_dataChan

    Cexit               *chan string
    Clog                *chan string
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

