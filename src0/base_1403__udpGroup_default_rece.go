package main

// _Fhandle_u01y__udpListen_Udp__read_main_loop
func (___Vun *_TudpNodeSt) _FsrvGroup__140201y__receive() {
	for {
		// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
		___Vun.unLen, ___Vun.unRemoteAddr, ___Vun.unRerr =
			___Vun.unConn.ReadFromUDP(___Vun.unBuf)
		if nil == ___Vun.unRerr {
			if nil == ___Vun.unCHrece {
				_FpfNhex(&___Vun.unBuf, 40, " 831818 01 rece: %5d,%11d,noOutCH drop,", ___Vun.unLocalPort, _FtimeI64())
			} else {
				if nil == ___Vun.unCBrece {
					___Vun._FsrvGroup__140201yy__receiveCallBack_default()
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

func (___Vun *_TudpNodeSt) _FsrvGroup__140201yy__receiveCallBack_default() {

	__Vrece := _TudpNodeDataRece{
		unInRemoteAddr: ___Vun.unRemoteAddr,
		unInLen:        ___Vun.unLen,
		unInBuf:        &___Vun.unBuf,
	}
	_FpfNhex(__Vrece.unInBuf, 40, " 839191 01 rece : %d , %v ", __Vrece.unInLen, __Vrece.unInRemoteAddr)

	(*___Vun.unCHrece) <- __Vrece
}
