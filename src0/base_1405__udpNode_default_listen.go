package main

import (
	"net"
)

func (___Vun *_TudpNodeSt) _FudpNode__140201x__listen() {
	___Vun.unCHsend = make(chan _TudpNodeDataSend, 5) // silice : with var len

	//_FtryListenToUDP01()
	___Vun.unAddr, ___Vun.unErr = net.ResolveUDPAddr("udp4", ___Vun.hostPortStr)
	if ___Vun.unErr != nil {
		_Fex("918381 03", ___Vun.unErr)
	}

	// func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
	___Vun.unConn, ___Vun.unErr = net.ListenUDP("udp4", ___Vun.unAddr)
	if ___Vun.unErr != nil {
		_Fex("918381 05", ___Vun.unErr)
	}

	//    //func (c *UDPConn) LocalAddr() Addr
	___Vun.unLocalAddr = ___Vun.unConn.LocalAddr()
	___Vun.unLocalPort = ___Vun.unLocalAddr.(*net.UDPAddr).Port

	//_FpfN(" 918381 061 %v", ___Vun)
	//_FpfN(" 918381 062 %+v", ___Vun)
	//_FpfN(" 918381 063 %#v", ___Vun)
	//_FpfN(" 918381 065 ___Vun.unConn       %v", ___Vun.unConn)
	//_FpfN(" 918381 066 ___Vun.unAddr       %v", ___Vun.unAddr)
	_FpfN(" 918381 067 ___Vun.unLocalAddr  %v", ___Vun.unLocalAddr)
	//_FpfN(" 918381 068 ___Vun.unRemoteAddr %v", ___Vun.unRemoteAddr)

}
