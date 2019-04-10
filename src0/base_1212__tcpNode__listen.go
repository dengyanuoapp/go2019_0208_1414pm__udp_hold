// _TtcpNodE
package main

import (
	//"time"
	"net"
)

func (___Vtn2 *_TtcpNodE) _FtcpNode__tryListen01() {
	// func ResolveTCPAddr(network, address string) (*TCPAddr, error)
	___Vtn2.tnAddr, ___Vtn2.tnErr = net.ResolveTCPAddr("tcp4", ___Vtn2.tnHostPortStr)
	if ___Vtn2.tnErr != nil {
		_Fex1("728191 01", ___Vtn2.tnErr)
	}

	// func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
	___Vtn2.tnListener, ___Vtn2.tnErr = net.ListenTCP("tcp4", ___Vtn2.tnAddr)
	if ___Vtn2.tnErr != nil {
		_Fex1("728191 03", ___Vtn2.tnErr)
	}

	// func (l *TCPListener) Addr() Addr ;
	//type Addr interface { Network() string (for example, "tcp", "udp") ; String() string  (for example, "192.0.2.1:25", "[2001:db8::1]:80") }
	___Vtn2.tnLisnAddr = ___Vtn2.tnListener.Addr()
	_Fpf("728191 05 : tcp listen on: %28v , %28s , %28s \n", ___Vtn2.tnLisnAddr, _FgetFuncName3(), ___Vtn2.tnName)

} // _FtcpNode__tryListen01
