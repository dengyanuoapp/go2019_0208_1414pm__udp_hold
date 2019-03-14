package main

var __Vbuf140201y []byte

// _Fhandle_u01y__udpListen_Udp__read_main_loop
func (___Vun *_TudpNodeSt) _FudpNode__140201y__receive() {
	for {
		// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
		__Vbuf140201y = make([]byte, 2048)
		___Vun.unBuf = &__Vbuf140201y
		___Vun.unLen, ___Vun.unRemoteAddr, ___Vun.unRerr =
			___Vun.unConn.ReadFromUDP(__Vbuf140201y)
		if nil == ___Vun.unRerr {
			if nil == ___Vun.unCHrece {
				_FpfNhex(___Vun.unBuf, 40, " 831818 01 rece: %5d,%11d,noOutCH drop,", ___Vun.unLocalPort, _FtimeI64())
			} else {
				if nil == ___Vun.unCBrece {
					___Vun._FudpNode__140201yy__receiveCallBack_default()
				} else {
					___Vun.unCBrece(___Vun)
				}
			}
		} else {
			_FpfN(" 831818 09 rece error : [%v] ", ___Vun.unRerr)
		}
		//_Fsleep_1s()
	}
}

func (___Vun *_TudpNodeSt) _FudpNode__140201yy__receiveCallBack_default() {

	__Vrece := _TudpNodeDataRece{
		unInRemoteAddr: ___Vun.unRemoteAddr,
		unInLen:        ___Vun.unLen,
		unInBuf:        ___Vun.unBuf,
	}
	_FpfNhex(__Vrece.unInBuf, 38, " 839191 01 rece : %d ,%11d %v ", __Vrece.unInLen, _FtimeI64(), __Vrece.unInRemoteAddr)

	(*___Vun.unCHrece) <- __Vrece
}
