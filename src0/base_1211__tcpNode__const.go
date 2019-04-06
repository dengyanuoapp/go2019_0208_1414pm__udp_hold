package main

import (
	"net"
	"sync"
)

type _TacceptTCP struct {
	enabled   bool
	idx       int
	serverTCP *_TtcpNodE
	connTCP   *net.TCPConn
	err       error

	//	r64try int64
	//	r64ok  int64
	//	r64eof int64
	//	w64try int64
	//	w64ok  int64
	//	w64eof int64
	cR _Tcount
	cW _Tcount

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

type _TtcpNodE struct {
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

	TcallbackAccDataRece func(*_TacceptTCP) // _FuserCallback__Accept_dataReceive__Log_Fn
	TcallbackAccDataChan func(*_TacceptTCP) // _FuserCallback__accept_dataChan__Log_Fn

	TcallbackSvrDataChan func(*_TtcpNodE) // _FuserCallback__service_dataChan__Log_Fn
	tnCBinit             func(*_TtcpNodE) // _FtcpNode__200101x__init_default

	Cexit *chan string
	Clog  *chan string
} // _TtcpNodE
