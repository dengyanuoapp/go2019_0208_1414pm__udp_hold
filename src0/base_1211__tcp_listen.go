package main

import (
    "net"
    "sync"
)

type _TacceptTCP struct {
    enabled             bool
    idx                 int
    serverTCP           *_TserviceTCP
    connTCP             *net.TCPConn
    err                 error

    r64try              int64
    r64ok               int64
    r64eof              int64
    w64try              int64
    w64ok               int64
    w64eof              int64

    Vbuf                []byte
    Vbuf2               []byte
    Vlen                int
    Verr                error
    VremoteAddr         net.Addr
    VlocalAddr          net.Addr

    Cstart              chan string
    CreceiveMsg         chan []byte
    CchanMsg            chan []byte
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

    TcallbackSvrDataChan          func( *_TserviceTCP)    // _FcallbackForDebugLog_service_dataChan
    TcallbackAccDataRece          func( *_TacceptTCP)     // _FuserCallback__LogAccept_dataReceive__Fn
    TcallbackAccDataChan          func( *_TacceptTCP)     // _FcallbackForDebugLog_accept_dataChan

    Cexit               *chan string
    Clog                *chan string
} // _TserviceTCP 

func ( ___Vsvr *_TserviceTCP ) _FtryListenToTCP01()  {
    // func ResolveTCPAddr(network, address string) (*TCPAddr, error)
    ___Vsvr.tcpAddr , ___Vsvr.err  = net.ResolveTCPAddr("tcp4", ___Vsvr.hostPortStr)
    if ___Vsvr.err != nil {
        _Fex( "err13813" , ___Vsvr.err)
    }

    // func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
    ___Vsvr.tcpLisn , ___Vsvr.err  = net.ListenTCP("tcp4", ___Vsvr.tcpAddr )
    if ___Vsvr.err != nil {
        _Fex( "err13814" , ___Vsvr.err)
    }
} // _FtryListenToTCP01

