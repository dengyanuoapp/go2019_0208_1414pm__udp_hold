package main

import (
	"net"
	"sync"
	"time"
)

type _TudpNodeDataRece struct {
	urrRemoteAddr net.UDPAddr
	urrLen        int
	urrBuf        []byte
}
type _TudpNodeDataSend struct {
	usToAddr _TudpConnPort
	usOutBuf []byte
}

type _TuNodeDataRmap struct {
	mux        sync.Mutex
	unrMapNow  map[[16]byte]_TudpNodeDataRece
	unrMapLast map[[16]byte]_TudpNodeDataRece
}

// _TserviceUDP
// a udp node: input accept AES256 rand only , can output anything
type _TudpNodeSt struct {
	unRKeyLP      *[]byte
	unRkeyX       _Tkey256
	unAddr        *net.UDPAddr
	unConn        *net.UDPConn
	unLocalAddr   net.Addr
	unLoopGap     time.Duration
	unLocalPort   int
	unErr         error
	unRerr        error
	unSerr        error
	unRbuf        []byte // receive-buf
	unRlen        int
	unCHreceLX    *chan _TudpNodeDataRece // if nil , drop it ; not-nil , put the received data into this chan
	unCHsendX     chan _TudpNodeDataSend  // try get data from chan, then send it out.
	unCBinit      func(*_TudpNodeSt)      //
	unCBrece      func(*_TudpNodeSt)      // if nil , use the default procedure to deal with receive
	unCBsend      func(*_TudpNodeSt)      // if nil , use the default procedure to deal with send
	unCBgap       func(*_TudpNodeSt)      // if unLoopGap is not ZERO , call this.
	unRemoteAddr  net.UDPAddr
	unName        string
	unHostPortStr string
	unIdx         int
}
