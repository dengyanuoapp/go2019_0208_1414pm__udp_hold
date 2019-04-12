package main

import (
	"net"
)

func (___Vun *_TudpNodeSt) _FudpNode__500101x__listen() {
	___Vun.unCHsendI = make(chan _TudpNodeDataSend, 5) // silice : with var len

	//_FtryListenToUDP01()
	___Vun.unAddr, ___Vun.unErr = net.ResolveUDPAddr("udp4", ___Vun.unHostPortStr)
	if ___Vun.unErr != nil {
		_Fex1("918381 03", ___Vun.unErr)
	}

	// func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
	___Vun.unConn, ___Vun.unErr = net.ListenUDP("udp4", ___Vun.unAddr)
	if ___Vun.unErr != nil {
		_Fex1("918381 05", ___Vun.unErr)
	}

	//    //func (c *UDPConn) LocalAddr() Addr
	___Vun.unLocalAddr = ___Vun.unConn.LocalAddr()
	___Vun.unLocalPort = ___Vun.unLocalAddr.(*net.UDPAddr).Port

	_FpfN(" 918381 07 listen on : %s, %s , key %x", ___Vun.unLocalAddr.String(), ___Vun.unName, ___Vun.unRkeyX.B32)
	_CpfN(" 918381 08 listen on : %s, %s , key %x", ___Vun.unLocalAddr.String(), ___Vun.unName, ___Vun.unRkeyX.B32)

}
