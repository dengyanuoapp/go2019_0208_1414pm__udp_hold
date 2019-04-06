package main

import (
	"net"
	"sync"
)

type _TtcpNodE struct {
	tnName               string
	tnHostPortStr        string
	tnAmount             int
	tnAddr               *net.TCPAddr
	tnListener           *net.TCPListener
	tnLisnAddr           net.Addr
	tnErr                error
	tnCBinit             func(*_TtcpNodE) // _FtcpNode__200101x__init_default
	tnCHsendToAllClientI chan []byte
	tnAcceptTCPs         []_TacceptTCP
	tnClientMux          sync.Mutex
	tnClientCnt          int
	tnCBaccDataChan      func(*_TacceptTCP) // _FuserCallback__accept_dataChan__Log_Fn
	tnCBsvrDataChan      func(*_TtcpNodE)   // _FuserCallback__service_dataChan__Log_Fn
} // _TtcpNodE

type _TacceptTCP struct {
	taChanId128   []byte
	taEnabled     bool
	taIdx         int
	taServerTCP   *_TtcpNodE
	taConnTCP     *net.TCPConn
	taErr         error
	tnCHsendI     chan []byte
	tnCHreceLO    *chan _TtcpNodeDataRece
	taR           _Tcount
	taW           _Tcount
	taBuf         []byte
	taRdata       _TtcpNodeDataRece
	taLen         int
	taErr2        error
	taRemoteAddr  net.Addr
	taLocalAddr   net.Addr
	taCreceiveMsg chan []byte
	taCchanMsg    chan []byte
	taCreceiveErr chan string
} // _TacceptTCP
