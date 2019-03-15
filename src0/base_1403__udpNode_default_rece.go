package main

// _Fhandle_u01y__udpListen_Udp__read_main_loop
func (___Vun *_TudpNodeSt) _FudpNode__140201y__receive() {
	for {
		// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
		___Vun.unLen, ___Vun.unRemoteAddr, ___Vun.unRerr = ___Vun.unConn.ReadFromUDP(___Vun.unRbuf)
		if nil == ___Vun.unRerr {
			if nil == ___Vun.unCHrece {
				_FpfNhex(&___Vun.unRbuf, 40, " 831818 01 rece: %5d,%11d,noOutCH drop,", ___Vun.unLocalPort, _FtimeI64())
			} else {
				if nil == ___Vun.unCBrece {
					//___Vun._FudpNode__140201yy__receiveCallBack_default__directChanOut()
					___Vun._FudpNode__140201yyy__receiveCallBack_default__randDecodeOut()
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

func (___Vun *_TudpNodeSt) _FudpNode__140201yy__receiveCallBack_default__directChanOut() {

	__Vrece := _TudpNodeDataRece{
		unInRemoteAddr: *___Vun.unRemoteAddr,
		unInLen:        ___Vun.unLen,
		unInBuf:        make([]byte, ___Vun.unLen),
	}
	copy(__Vrece.unInBuf, ___Vun.unRbuf[:___Vun.unLen])

	_FpfNhex(&__Vrece.unInBuf, 38, " 839191 01 rece : %d ,%11d %v ", __Vrece.unInLen, _FtimeI64(), __Vrece.unInRemoteAddr)

	(*___Vun.unCHrece) <- __Vrece

	//__Vrece2 := <-(*___Vun.unCHrece)
	//_FpfNhex(__Vrece2.unInBuf, 38, " 839191 02 rece : %d ,%11d %v ", __Vrece2.unInLen, _FtimeI64(), __Vrece2.unInRemoteAddr)
}

func (___Vun *_TudpNodeSt) _FudpNode__140201yyy__receiveCallBack_default__randDecodeOut() {

	if nil == ___Vun.unRkey.Bkey {
		_FpfN(" 439191 01 key ERROR : len %d ,%11d ,addr %v , key %x.",
			___Vun.unLen, _FtimeI64(), ___Vun.unRemoteAddr, ___Vun.unRkey.Bkey)
		return
	}

	__Vtmp2, __Verr2 := _FdecAesRand__only(&___Vun.unRkey.Bkey, &___Vun.unRbuf)
	if nil != __Verr2 {
		_FpfNhex(&___Vun.unRbuf, 38, " 439191 03 rece Null : %d ,%11d %v %x:",
			___Vun.unLen, _FtimeI64(), ___Vun.unRemoteAddr, ___Vun.unRkey.Bkey)
		return
	}

	//	__Vrece := _TudpNodeDataRece{
	//		unInRemoteAddr: *___Vun.unRemoteAddr,
	//		unInLen:        ___Vun.unLen,
	//		unInBuf:        make([]byte, ___Vun.unLen),
	//	}
	//	copy(__Vrece.unInBuf, ___Vun.unRbuf[:___Vun.unLen])

	__Vrece := _TudpNodeDataRece{
		unInRemoteAddr: *___Vun.unRemoteAddr,
		unInLen:        len(__Vtmp2),
		unInBuf:        __Vtmp2,
	}

	if __Vrece.unInLen < 400 && __Vrece.unInLen > 32 {
		_FpfN(" 439191 05 rece : %d ,%11d %v : %s", __Vrece.unInLen, _FtimeI64(), __Vrece.unInRemoteAddr, __Vrece.unInBuf)
	} else {
		_FpfNhex(&__Vrece.unInBuf, 38, " 439191 06 rece : %d ,%11d %v ", __Vrece.unInLen, _FtimeI64(), __Vrece.unInRemoteAddr)
	}

	(*___Vun.unCHrece) <- __Vrece
}
