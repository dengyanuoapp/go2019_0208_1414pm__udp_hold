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
	VuSrvInfo    *_TsrvInfo

	UuserLoopCall11 func(*_TserviceUDP) // UcallbackMR: _FuserCallback_u01M__dataRece_Fn__main_top,
	// _FuserCallback_u01M__dataRece_Fn__main_top _FuserCallback_u01M__dataRece_Sn _FuserCallback_u01M__dataRece_Cn

	UuserLoopCall21 func(*_TserviceUDP) // UcallbackCI: _FuserCallback_chanIn_Fn__main_top,
	// _FuserCallback_chanIn_Fn__main_top

	UsrvGoCall201 func(*_TserviceUDP) // _Fhandle_u01x__udpListen_Udp__read_main_top__default : deal with udp main loop data rece :
	UsrvGoCall202 func(*_TserviceUDP) // _Fhandle_u02x__udpListen__chanIn_main_top__default
	UsrvGoCall203 func(*_TserviceUDP) // _Fhandle_u03x__udpListen__timer__main_top__default

	// _Fconnect_to_server_01x__req_new_sessionID__default
	UreqNewSessionTM *_TUreqNewSession // called by UsrvGoCall203 , if nil , then UsrvGoCall203 stop

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
