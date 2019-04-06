package main

import (
	"net"
	"sync"
)

type _TacceptTCP struct {
	taEnabled   bool
	taIdx       int
	taServerTCP *_TtcpNodE
	taConnTCP   *net.TCPConn
	taErr       error

	//	r64try int64
	//	r64ok  int64
	//	r64eof int64
	//	w64try int64
	//	w64ok  int64
	//	w64eof int64
	taR _Tcount
	taW _Tcount

	taBuf        []byte
	taBuf2       []byte
	taLen        int
	taErr2       error
	taRemoteAddr net.Addr
	taLocalAddr  net.Addr

	taCstart      chan string
	taCreceiveMsg chan []byte
	taCchanMsg    chan []byte
	taCreceiveErr chan string

	taClog *chan string
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

	tnClog *chan string
} // _TtcpNodE
