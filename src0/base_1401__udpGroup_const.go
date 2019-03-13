package main

import (
	"net"
)

type _TudpNodeDataRece struct {
	unRemoteAddr net.Addr
	unInLen      int
	unInBuf      []byte
}
type _TudpNodeDataSend struct {
	unDstAddr net.Addr
	unOutBuf  []byte
}

// _TserviceUDP
// a udp node: input accept AES256 rand only , can output anything
type _TudpNodeSt struct {
	// the input password , all data try AES decode, if faild , then drop it.
	// if nil , then ignore(no filter , directly receive)
	unPwIn256LP  *[]byte
	unPwIn256    [32]byte
	unAddr       *net.UDPAddr
	unConn       *net.UDPConn
	unLocalAddr  net.Addr
	unErr        error
	unRerr       error
	unSerr       error
	unBuf        []byte
	unLen        int
	unRece       *chan _TudpNodeDataRece // if nil , drop it ; not-nil , put the received data into this chan
	unSend       chan _TudpNodeDataSend  // try get data from chan, then send it out.
	unCBrece     func(*_TudpNodeSt)      // if nil , use the default procedure to deal with receive
	unCBsend     func(*_TudpNodeSt)      // if nil , use the default procedure to deal with receive
	unRemoteAddr *net.UDPAddr
	name         string
	hostPortStr  string
}
