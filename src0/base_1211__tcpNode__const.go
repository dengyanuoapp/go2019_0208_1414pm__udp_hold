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
	tnName        string
	tnHostPortStr string
	tnAmount      int

	tnAddr     *net.TCPAddr
	tnListener *net.TCPListener
	tnLisnAddr net.Addr
	tnErr      error

	tnAcceptTCPs []_TacceptTCP
	tnClientMux  sync.Mutex
	tnClientCnt  int

	tnCBaccDataRece func(*_TacceptTCP) // _FuserCallback__Accept_dataReceive__Log_Fn
	tnCBaccDataChan func(*_TacceptTCP) // _FuserCallback__accept_dataChan__Log_Fn
	tnCBsvrDataChan func(*_TtcpNodE)   // _FuserCallback__service_dataChan__Log_Fn
	tnCBinit        func(*_TtcpNodE)   // _FtcpNode__200101x__init_default

	tnCexit *chan string
	tnClog  *chan string
} // _TtcpNodE
