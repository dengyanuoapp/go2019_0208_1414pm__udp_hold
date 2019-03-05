package main

import (
	"net"
)

type _TserviceUDP struct {
	name        string
	hostPortStr string

	uExtMR *_TuExtMRead
	uExtCI *_TuExtChanI
	uExtTM *_TuExtTimer

	udpAddr *net.UDPAddr
	udpConn *net.UDPConn
	Vuerr   error

	Vubuf        []byte
	Vulen        int
	VuremoteAddr *net.UDPAddr
	VulocalAddr  net.Addr

	UcallbackMR func(*_TserviceUDP) // _Fhandle_u01x__udpListen_Udp__read_main_top : deal with udp main loop data rece :
	UcallbackCI func(*_TserviceUDP) // _Fhandle_u02x__udpListen__chanIn_main_top : deal with outside chan data in for udp using another handle
	//UcallbackTM func(*_TserviceUDP) // _Fhandle_u03x__udpListen__timer__main_top : deal with timer ARRAY in udp

	UreqNewSessionMR *_TreqNewSession
	UreqNewSessionCI *_TreqNewSession
	UreqNewSessionTM *_TreqNewSession

	Cexit *chan string
	Clog  *chan string

	CuIn01  chan []byte
	CuOut01 *chan []byte
} // _TserviceUDP

func (___Vsvr *_TserviceUDP) _FtryListenToUDP01() {
	// func ResolveUDPAddr(network, address string) (*UDPAddr, error)
	___Vsvr.udpAddr, ___Vsvr.Vuerr = net.ResolveUDPAddr("udp4", ___Vsvr.hostPortStr)
	if ___Vsvr.Vuerr != nil {
		_Fex("err13811", ___Vsvr.Vuerr)
	}

	// func ListenUDP(network string, laddr *UDPAddr) (*UDPConn, error)
	___Vsvr.udpConn, ___Vsvr.Vuerr = net.ListenUDP("udp4", ___Vsvr.udpAddr)
	if ___Vsvr.Vuerr != nil {
		_Fex("err13812", ___Vsvr.Vuerr)
	}

	//    //func (c *UDPConn) LocalAddr() Addr
	___Vsvr.VulocalAddr = ___Vsvr.udpConn.LocalAddr()

	_Fpf("ok13813 : udp listen on: %28v , %28s , %28s \n", ___Vsvr.VulocalAddr, _FgetFuncName3(), ___Vsvr.name)
} // _FtryListenToUDP01
