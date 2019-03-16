package main

import (
	"net"
)

type _TudpNodeDataRece struct {
	unInRemoteAddr net.UDPAddr
	unInLen        int
	unInBuf        []byte
}
type _TudpNodeDataSend struct {
	unDstAddr net.UDPAddr
	unOutBuf  []byte
}

// _TserviceUDP
// a udp node: input accept AES256 rand only , can output anything
type _TudpNodeSt struct {
	unRKeyLP      *[]byte
	unRkeyX       _Tkey256
	unAddr        *net.UDPAddr
	unConn        *net.UDPConn
	unLocalAddr   net.Addr
	unLocalPort   int
	unErr         error
	unRerr        error
	unSerr        error
	unRbuf        []byte // receive-buf
	unLen         int
	unCHrece      *chan _TudpNodeDataRece // if nil , drop it ; not-nil , put the received data into this chan
	unCHsend      chan _TudpNodeDataSend  // try get data from chan, then send it out.
	unCBinit      func(*_TudpNodeSt)      //
	unCBrece      func(*_TudpNodeSt)      // if nil , use the default procedure to deal with receive
	unCBsend      func(*_TudpNodeSt)      // if nil , use the default procedure to deal with send
	unRemoteAddr  *net.UDPAddr
	unName        string
	unHostPortStr string
}
