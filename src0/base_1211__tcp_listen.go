package main

import (
	"net"
	"sync"
)

type _TacceptTCP struct {
	enabled   bool
	idx       int
	serverTCP *_TserviceTCP
	connTCP   *net.TCPConn
	err       error

	r64try int64
	r64ok  int64
	r64eof int64
	w64try int64
	w64ok  int64
	w64eof int64

	Vbuf        []byte
	Vbuf2       []byte
	Vlen        int
	Verr        error
	VremoteAddr net.Addr
	VlocalAddr  net.Addr

	Cstart      chan string
	CreceiveMsg chan []byte
	CchanMsg    chan []byte
	CreceiveErr chan string

	Cexit *chan string
	Clog  *chan string
} // _TacceptTCP

type _TserviceTCP struct {
	name        string
	hostPortStr string
	cAmount     int

	tcpAddr     *net.TCPAddr
	tcpListener *net.TCPListener
	tcpLisnAddr net.Addr
	err         error

	acceptTCPs []_TacceptTCP
	clientMux  sync.Mutex
	clientCnt  int

	TcallbackSvrDataChan func(*_TserviceTCP) // _FuserCallback__service_dataChan__Log_Fn
	TcallbackAccDataRece func(*_TacceptTCP)  // _FuserCallback__Accept_dataReceive__Log_Fn
	TcallbackAccDataChan func(*_TacceptTCP)  // _FuserCallback__accept_dataChan__Log_Fn

	TsrvGoCallback10101 func(*_TserviceTCP)

	Cexit *chan string
	Clog  *chan string
} // _TserviceTCP

func (___Vsvr *_TserviceTCP) _FtryListenToTCP01() {
	// func ResolveTCPAddr(network, address string) (*TCPAddr, error)
	___Vsvr.tcpAddr, ___Vsvr.err = net.ResolveTCPAddr("tcp4", ___Vsvr.hostPortStr)
	if ___Vsvr.err != nil {
		_Fex("err13815", ___Vsvr.err)
	}

	// func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
	___Vsvr.tcpListener, ___Vsvr.err = net.ListenTCP("tcp4", ___Vsvr.tcpAddr)
	if ___Vsvr.err != nil {
		_Fex("err13816", ___Vsvr.err)
	}

	// func (l *TCPListener) Addr() Addr ;
	//type Addr interface { Network() string (for example, "tcp", "udp") ; String() string  (for example, "192.0.2.1:25", "[2001:db8::1]:80") }
	___Vsvr.tcpLisnAddr = ___Vsvr.tcpListener.Addr()
	_Fpf("ok13817 : tcp listen on: %28v , %28s , %28s \n", ___Vsvr.tcpLisnAddr, _FgetFuncName3(), ___Vsvr.name)

} // _FtryListenToTCP01
