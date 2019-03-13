package main

import (
	"net"
)

// _TudpNodeSt
// _Fhandle_u01x__udpListen_Udp__read_main_top__default
func (___Vun *_TudpNodeSt) IRun(___Vidx int) {
	switch ___Vidx {
	case 140201:
		_FsrvGroup__140201__main_top__default(___Vun)
	default:
		_FpfNex(" 937191 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FsrvGroup__140201__main_top__default(___Vun *_TudpNodeSt) {
	if nil != ___Vun.unPwIn256LP {
		copy(___Vun.unPwIn256[:], *___Vun.unPwIn256LP)
	}
	_FpfN(" 918381 01 : %x ", ___Vun.unPwIn256)

	___Vun._FsrvGroup__140201x__listen()

	go ___Vun._FsrvGroup__140201y__receive()

	go ___Vun._FsrvGroup__140201z__send()

	//_Fex1(" 918381 09 ")
}

func (___Vun *_TudpNodeSt) _FsrvGroup__140201x__listen() {
	___Vun.unBuf = make([]byte, 2048)               // silice : with var len
	___Vun.unSend = make(chan _TudpNodeDataSend, 5) // silice : with var len

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

	//_FpfN(" 918381 061 %v", ___Vun)
	//_FpfN(" 918381 062 %+v", ___Vun)
	//_FpfN(" 918381 063 %#v", ___Vun)
	//_FpfN(" 918381 065 ___Vun.unConn       %v", ___Vun.unConn)
	//_FpfN(" 918381 066 ___Vun.unAddr       %v", ___Vun.unAddr)
	_FpfN(" 918381 067 ___Vun.unLocalAddr  %v", ___Vun.unLocalAddr)
	//_FpfN(" 918381 068 ___Vun.unRemoteAddr %v", ___Vun.unRemoteAddr)

}

// _Fhandle_u01y__udpListen_Udp__read_main_loop
func (___Vun *_TudpNodeSt) _FsrvGroup__140201y__receive() {
	for {
		// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
		___Vun.unLen, ___Vun.unRemoteAddr, ___Vun.unRerr =
			___Vun.unConn.ReadFromUDP(___Vun.unBuf)
		if nil == ___Vun.unRerr {
			_FpfNhex(&___Vun.unBuf, 40, " 831818 01 rece:")
		} else {
			_FpfN(" 831818 09 rece error : [%v] ", ___Vun.unRerr)
		}
		//_Fsleep_1s()
	}
}

func (___Vun *_TudpNodeSt) _FsrvGroup__140201z__send() {
	for {
		_Fsleep_1s()
	}
}
