package main

import (
	"net"
)

// _TserviceUDP
type _TudpClientSt struct {
	ucPw128 [16]byte

	ucAddr *net.UDPAddr
	ucConn *net.UDPConn
	ucErr  error

	VucBuf        []byte
	VucLen        int
	VucRemoteAddr *net.UDPAddr
	VucLocalAddr  net.Addr
	VucSrvInfo    *_TsrvInfo
}
