package main

import "net"

// _Fhandle_u01y__udpListen_Udp__read_main_loop
func (___Vun *_TudpNodeSt) _FudpNode__500101y__receive() {
	var __VuAddr *net.UDPAddr
	for {
		// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
		___Vun.unRlen, __VuAddr, ___Vun.unRerr = ___Vun.unConn.ReadFromUDP(___Vun.unRbuf)
		___Vun.unRemoteAddr = *__VuAddr
		if nil == ___Vun.unRerr {
			if nil == ___Vun.unCHreceLO {
				_FpfNhex(&___Vun.unRbuf, 40, " 831818 01 rece: %5d,%11d,noOutCH drop,", ___Vun.unLocalPort, _FtimeI64())
			} else {
				if nil == ___Vun.unCBrece {
					_FpfN(" 831818 03 default receive:{%s}", ___Vun.String())
					___Vun._FudpNode__500101yy3__receiveCallBack_default__randDecodeOut_noKeyWillDirect() // if gap is not set , default
					//___Vun._FudpNode__500101yy4__receiveCallBack_default__randDecodeOut_mustDecode()    // if gap set , use this
				} else {
					_FpfN(" 831818 08 custom receive")
					___Vun.unCBrece(___Vun)
				}
			}
		} else {
			_FpfN(" 831818 09 rece error : [%v] ", ___Vun.unRerr)
		}
		//_Fsleep_1s()
	}
}

func (___Vun *_TudpNodeSt) _FudpNode__500101yy3__receiveCallBack_default__randDecodeOut_noKeyWillDirect() {

	var __Vrece _TudpNodeDataRece
	if ___Vun.unRkeyX.disable {
		__Vrece = _TudpNodeDataRece{
			urrRemoteAddr: ___Vun.unRemoteAddr,
			urrLen:        ___Vun.unRlen,
			urrBuf:        ___Vun.unRbuf[:___Vun.unRlen],
			//urrBuf:        make([]byte,___Vun.unRlen)
		}
		//copy(__Vrece.urrBuf,        ___Vun.unRbuf[:___Vun.unRlen])
		(*___Vun.unCHreceLO) <- __Vrece
		_FpfNhex(&___Vun.unRbuf, 30, " 439191 01 key disabled ,skip rece rand decode")
		return
	}
	_FpfN(" 439191 02 custom receive")
	___Vun._FudpNode__500101yy4__receiveCallBack_default__randDecodeOut_mustDecode(&__Vrece.urrBuf)

}

func (___Vun *_TudpNodeSt) _FudpNode__500101yy4__receiveCallBack_default__randDecodeOut_mustDecode(___VbufIn *[]byte) {

	if ___Vun.unRkeyX.disable {
		_FpfNhex(___VbufIn, 30, " 439192 01 key disabled ,skip rece rand decode")
		return
	}

	if 32 != len(___Vun.unRkeyX.Bkey) {
		_FpfN(" 439192 02 key ERROR : len %d ,%11d ,addr %v , key %x.",
			___Vun.unRlen, _FtimeI64(), ___Vun.unRemoteAddr, ___Vun.unRkeyX.Bkey)
		return
	}

	__Vtmp2, __Verr2 := _FdecAesRand__only(&___Vun.unRkeyX.Bkey, ___VbufIn)
	if nil != __Verr2 {
		_FpfN(" 439192 03 rece buf: %v ", ___VbufIn)
		_FpfNhex(___VbufIn, 68, " 439192 04 rece Null or error : %d ,%11d %v %x. error:%v ",
			___Vun.unRlen, _FtimeI64(), ___Vun.unRemoteAddr, ___Vun.unRkeyX.Bkey, __Verr2)
		return
	}

	__Vunr := _TudpNodeDataRece{
		urrRemoteAddr: ___Vun.unRemoteAddr,
		urrLen:        len(__Vtmp2),
		urrBuf:        __Vtmp2,
	}

	if __Vunr.urrLen < 400 && __Vunr.urrLen > 32 {
		//_FpfN(" 439192 05 rece : %d ,%11d %v : %s", __Vunr.urrLen, _FtimeI64(), __Vunr.urrRemoteAddr, __Vunr.urrBuf)
		//_FpfNhex(___VbufIn, 48, " 439192 06 origin len %d :", ___Vun.unRlen)
		//_FpfNhex(&__Vunr.urrBuf, 33, " 439192 07 oldLen %d %11d from %v", ___Vun.unRlen, _FtimeI64(), __Vunr.urrRemoteAddr)
	} else {
		_FpfNhex(&__Vunr.urrBuf, 38, " 439192 08 rece : %d ,%11d %v ", __Vunr.urrLen, _FtimeI64(), __Vunr.urrRemoteAddr)
	}

	(*___Vun.unCHreceLO) <- __Vunr
}
