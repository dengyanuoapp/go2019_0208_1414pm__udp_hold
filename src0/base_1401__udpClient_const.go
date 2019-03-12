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
	unPwIn256    [32]byte // the input password
	unAddr       *net.UDPAddr
	unConn       *net.UDPConn
	unRemoteAddr *net.UDPAddr
	unLocalAddr  net.Addr
	unErr        error
	unBuf        []byte
	unLen        int
	unRece       _TudpNodeDataRece
	unSend       _TudpNodeDataSend
}
