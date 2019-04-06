// _TtcpNodE
package main

import (
	//"time"
	"net"
)

func (___Vtn2 *_TtcpNodE) _FtcpNode__tryListen01() {
	// func ResolveTCPAddr(network, address string) (*TCPAddr, error)
	___Vtn2.tcpAddr, ___Vtn2.err = net.ResolveTCPAddr("tcp4", ___Vtn2.hostPortStr)
	if ___Vtn2.err != nil {
		_Fex("err13815", ___Vtn2.err)
	}

	// func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
	___Vtn2.tcpListener, ___Vtn2.err = net.ListenTCP("tcp4", ___Vtn2.tcpAddr)
	if ___Vtn2.err != nil {
		_Fex("err13816", ___Vtn2.err)
	}

	// func (l *TCPListener) Addr() Addr ;
	//type Addr interface { Network() string (for example, "tcp", "udp") ; String() string  (for example, "192.0.2.1:25", "[2001:db8::1]:80") }
	___Vtn2.tcpLisnAddr = ___Vtn2.tcpListener.Addr()
	_Fpf("ok13817 : tcp listen on: %28v , %28s , %28s \n", ___Vtn2.tcpLisnAddr, _FgetFuncName3(), ___Vtn2.name)

} // _FtcpNode__tryListen01
