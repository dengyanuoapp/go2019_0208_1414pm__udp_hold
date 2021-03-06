package main

import "net"
import "sync"

var ___VudpNode__500201y__mux sync.Mutex

// _Fhandle_u01y__udpListen_Udp__read_main_loop
func (___Vun *_TudpNodeSt) _FudpNode__500201y__receive__default() {
	// _FpfNdb(" 331818 00 rece start ")
	for {
		var __VuAddr *net.UDPAddr
		// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
		___Vun.unRlen, __VuAddr, ___Vun.unRerr = ___Vun.unConn.ReadFromUDP(___Vun.unRbuf1500)

		if nil != ___Vun.unRerr {
			_FpfN(
				" 331818 01 rece error : [%v] ",
				___Vun.unRerr)
			continue
		}

		___VudpNode__500201y__mux.Lock()

		___Vun.unRemoteAddr = *__VuAddr
		___Vun.unRbufTmp = ___Vun.unRbuf1500[:___Vun.unRlen]

		_FpfNonce(" 331818 02 Origin rece: me<%d> ra:<%s>", ___Vun.unLocalPort, ___Vun.unRemoteAddr.String())
		___CpfN(" 331818 03 Origin rece: me<%d> ra:<%s> (%d/%d)<%x>",
			___Vun.unLocalPort, ___Vun.unRemoteAddr.String(),
			___Vun.unRlen, len(___Vun.unRbufTmp), ___Vun.unRbufTmp)

		___Vun.
			_FudpNode__500201y01__receive__default()

		___VudpNode__500201y__mux.Unlock()
	}
}

var ___VudpNode__500201y01__receive__default__mux sync.Mutex

func (___Vun *_TudpNodeSt) _FudpNode__500201y01__receive__default() {
	defer ___VudpNode__500201y01__receive__default__mux.Unlock()
	___VudpNode__500201y01__receive__default__mux.Lock()

	if nil != ___Vun.unRerr {
		_FpfN(
			" 831818 02 rece error : [%v] ",
			___Vun.unRerr)
		return
	}

	if 1500 != len(___Vun.unRbuf1500) {
		_FpfN(
			" 831818 03 buf len error :{%s}",
			___Vun.String())
		return
	}

	if nil == ___Vun.unCHreceByteLO && nil == ___Vun.unCHreceStructLO { // *chan _TudpNodeDataReceX
		_FpfNonce(
			" 831818 05 rece: port:%5d: outChan Null , so drop the data package only. %11d ",
			___Vun.unLocalPort,
			_FtimeI64())
		return
	}

	__Vrece := _TudpNodeDataRece{ // _TudpNodeDataReceX
		Ti:            _FgenRand_int(),
		UrrLocalPort:  ___Vun.unLocalPort, // _TudpNodeSt
		UrrRemoteAddr: ___Vun.unRemoteAddr,
		UrrBuf:        ___Vun.unRbuf1500[:___Vun.unRlen], // []byte
		UrrReceiveKey: ___Vun.unRkeyX,                    // _Tkey256
	}

	_FpfNonce(" 831819 01 Origin rece: me<%d> ra:<%s> ", ___Vun.unLocalPort, ___Vun.unRemoteAddr.String())
	___CpfN(" 831819 02 Origin rece: me<%d> ra:<%s> (%d/%d){%x}<%x>",
		___Vun.unLocalPort, ___Vun.unRemoteAddr.String(),
		___Vun.unRlen, len(__Vrece.UrrBuf), _Fmd5__5x(&__Vrece.UrrBuf), __Vrece.UrrBuf)

	if nil != ___Vun.unCHreceStructLO {
		_FpfN(" 831819 03 Origin rece direct sent.")
		_CpfN(" 831819 04 Origin rece direct sent.")
		(*___Vun.unCHreceStructLO) <- __Vrece

	} else {
		//(*___Vun.unCHreceByteLO) <- __Vrece
		__VreceB, __Verr3 := _FencGob__only(&__Vrece)

		if nil != __Verr3 {
			_CFpfN(" 831819 06 gob En error <%v>", __Verr3)
			return
		}

		_FpfNonce(" 831819 07 udpNode : push to Byte-chan : <%s>", String9s(&__VreceB))
		___CpfN(" 831819 08 udpNode : push to Byte-chan : <%s>", String9s(&__VreceB))

		(*___Vun.unCHreceByteLO) <- __VreceB
	}

}
