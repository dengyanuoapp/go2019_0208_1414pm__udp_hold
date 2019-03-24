package main

import "net"

// _Fhandle_u01y__udpListen_Udp__read_main_loop
func (___Vun *_TudpNodeSt) _FudpNode__540201y__receive() {
	var __VuAddr *net.UDPAddr
	for {
		// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
		___Vun.unRlen, __VuAddr, ___Vun.unRerr = ___Vun.unConn.ReadFromUDP(___Vun.unRbuf)
		___Vun.unRemoteAddr = *__VuAddr
		if nil == ___Vun.unRerr {
			if nil == ___Vun.unCHreceLX {
				_FpfNhex(&___Vun.unRbuf, 40, " 831818 01 rece: %5d,%11d,noOutCH drop,", ___Vun.unLocalPort, _FtimeI64())
			} else {
				if nil == ___Vun.unCBrece {
					___Vun._FudpNode__540201yy3__receiveCallBack_default__randDecodeOut()
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

func (___Vun *_TudpNodeSt) _FudpNode__540201yy3__receiveCallBack_default__randDecodeOut() {

	if ___Vun.unRkeyX.disable {
		__Vrece := _TudpNodeDataRece{
			urInRemoteAddr: ___Vun.unRemoteAddr,
			urInLen:        ___Vun.unRlen,
			urInBuf:        ___Vun.unRbuf[:___Vun.unRlen],
			//urInBuf:        make([]byte,___Vun.unRlen)
		}
		//copy(__Vrece.urInBuf,        ___Vun.unRbuf[:___Vun.unRlen])
		(*___Vun.unCHreceLX) <- __Vrece
		_FpfNhex(&___Vun.unRbuf, 30, " 439191 01 key disabled ,skip rece rand decode")
		return
	}

	if 32 != len(___Vun.unRkeyX.Bkey) {
		_FpfN(" 439191 02 key ERROR : len %d ,%11d ,addr %v , key %x.",
			___Vun.unRlen, _FtimeI64(), ___Vun.unRemoteAddr, ___Vun.unRkeyX.Bkey)
		return
	}

	__Vtmp2, __Verr2 := _FdecAesRand__only(&___Vun.unRkeyX.Bkey, &___Vun.unRbuf)
	if nil != __Verr2 {
		_FpfNhex(&___Vun.unRbuf, 68, " 439191 03 rece Null or error : %d ,%11d %v %x:",
			___Vun.unRlen, _FtimeI64(), ___Vun.unRemoteAddr, ___Vun.unRkeyX.Bkey)
		return
	}

	__Vrece := _TudpNodeDataRece{
		urInRemoteAddr: ___Vun.unRemoteAddr,
		urInLen:        len(__Vtmp2),
		urInBuf:        __Vtmp2,
	}

	if __Vrece.urInLen < 400 && __Vrece.urInLen > 32 {
		//_FpfN(" 439191 05 rece : %d ,%11d %v : %s", __Vrece.urInLen, _FtimeI64(), __Vrece.urInRemoteAddr, __Vrece.urInBuf)
		//_FpfNhex(&___Vun.unRbuf, 48, " 439191 06 origin len %d :", ___Vun.unRlen)
		_FpfNhex(&__Vrece.urInBuf, 33, " 439191 07 oldLen %d %11d from %v", ___Vun.unRlen, _FtimeI64(), __Vrece.urInRemoteAddr)
	} else {
		_FpfNhex(&__Vrece.urInBuf, 38, " 439191 08 rece : %d ,%11d %v ", __Vrece.urInLen, _FtimeI64(), __Vrece.urInRemoteAddr)
	}

	(*___Vun.unCHreceLX) <- __Vrece
}
