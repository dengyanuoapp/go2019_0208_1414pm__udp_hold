package main

import "net"
import "sync"

var ___VudpNode__500201y__mux sync.Mutex

// _Fhandle_u01y__udpListen_Udp__read_main_loop
func (___Vun *_TudpNodeSt) _FudpNode__500201y__receive__default() {
	// _FpfNdb(" 831818 00 rece start ")
	for {
		var __VuAddr *net.UDPAddr
		// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
		___Vun.unRlen, __VuAddr, ___Vun.unRerr = ___Vun.unConn.ReadFromUDP(___Vun.unRbuf)

		___VudpNode__500201y__mux.Lock()

		___Vun.unRemoteAddr = *__VuAddr

		___Vun._FudpNode__500201y01__receive__default()

		___VudpNode__500201y__mux.Unlock()
	}
}

func (___Vun *_TudpNodeSt) _FudpNode__500201y01__receive__default() {

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

	if nil == ___Vun.unCHreceLO { // *chan _TudpNodeDataReceX
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

	(*___Vun.unCHreceLO) <- __Vrece

	//		//_CpfN(" 831818 07 receOrigin: {%s} t:%11d ", __Vrece.String(), _FtimeI64())
	//
	//		__VreceB, __Verr3 := _FencGob__only(&__Vrece) // _TudpNodeDataReceX
	//		if nil != __Verr3 {
	//			_FpfN(" 831818 08 gobEnc error: {%v} t:%11d {%s}",
	//				__Verr3, _FtimeI64(),
	//				__Vrece.String(),
	//			)
	//			continue
	//		}
	//		___CpfN(" 831818 09 receOriginByte: (%d){%x}[%s] t:%11d {%s}",
	//			len(__VreceB),
	//			_FgenMd5__5(&__VreceB),
	//			String9(&__VreceB),
	//			_FtimeI64(),
	//			__Vrece.String(),
	//		)
	//
	//		___VudpNode__500201y__mux.Lock()
	//		// *chan _TudpNodeDataReceX
	//		(*___Vun.unCHreceLO) <- __VreceB
	//		___VudpNode__500201y__mux.Unlock()

	//_Fsleep_1s()
}
