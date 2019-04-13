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
		___Vun.unRlen, __VuAddr, ___Vun.unRerr = ___Vun.unConn.ReadFromUDP(___Vun.unRbuf)

		___VudpNode__500201y__mux.Lock()

		___Vun.unRemoteAddr = *__VuAddr

		_CpfN(" 331818 01 Origin rece: me<%d> ra:<%s> (%d/%d)<%x>",
			___Vun.unLocalPort, ___Vun.unRemoteAddr.String(),
			___Vun.unRlen, len(___Vun.unRbuf), ___Vun.unRbuf[:___Vun.unRlen])

		___Vun._FudpNode__500201y01__receive__default()

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

	if 1500 != len(___Vun.unRbuf) {
		_FpfN(
			" 831818 03 buf len error :{%s}",
			___Vun.String())
		return
	}

	if nil == ___Vun.unCHreceByteLO { // *chan _TudpNodeDataReceX
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
		UrrBuf:        ___Vun.unRbuf[:___Vun.unRlen], // []byte
		UrrReceiveKey: ___Vun.unRkeyX,                // _Tkey256
	}

	_CpfN(" 831818 07 Origin rece: me<%d> ra:<%s> (%d/%d){%x}<%x>",
		___Vun.unLocalPort, ___Vun.unRemoteAddr.String(),
		___Vun.unRlen, len(__Vrece.UrrBuf), _FgenMd5__5(&__Vrece.UrrBuf), __Vrece.UrrBuf)

	//(*___Vun.unCHreceByteLO) <- __Vrece
	__VreceB, __Verr3 := _FencGob__only(&__Vrece)

	if nil != __Verr3 {
		_FpfN(" 831818 08 gob En error <%v>", __Verr3)
		return
	}

	(*___Vun.unCHreceByteLO) <- __VreceB

}
