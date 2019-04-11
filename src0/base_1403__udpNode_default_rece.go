package main

import "net"
import "sync"

var ___lock001 sync.Mutex

// _Fhandle_u01y__udpListen_Udp__read_main_loop
func (___Vun *_TudpNodeSt) _FudpNode__500201y__receive__default() {
	// _FpfNdb(" 831818 00 rece start ")
	for {
		// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
		//___Vun.unRlen, __VuAddr, ___Vun.unRerr = ___Vun.unConn.ReadFromUDP(___Vun.unRbuf)
		var __VuAddr *net.UDPAddr = nil
		___Vun.unRlen, __VuAddr, ___Vun.unRerr = ___Vun.unConn.ReadFromUDP(___Vun.unRbuf)
		___Vun.unRemoteAddr = *__VuAddr

		if nil != ___Vun.unRerr {
			_FpfN(
				" 831818 02 rece error : [%v] ",
				___Vun.unRerr)
			continue
		}

		if 1500 != len(___Vun.unRbuf) {
			_FpfN(
				" 831818 03 buf len error :{%s}",
				___Vun.String())
			continue
		}

		if nil == ___Vun.unCHreceLO { // *chan _TudpNodeDataReceX
			_FpfNonce(
				" 831818 05 rece: port:%5d: outChan Null , so drop the data package only. %11d ",
				___Vun.unLocalPort,
				_FtimeI64())
			continue
		}

		__Vrece := _TudpNodeDataRece{ // _TudpNodeDataReceX
			Ti:            _FgenRand_int(),
			UrrLocalPort:  ___Vun.unLocalPort,            // _TudpNodeSt
			UrrRemoteAddr: (*__VuAddr),                   // net.UDPAddr
			UrrBuf:        ___Vun.unRbuf[:___Vun.unRlen], // []byte
			UrrReceiveKey: ___Vun.unRkeyX,                // _Tkey256
		}

		//_CpfN(" 831818 07 receOrigin: {%s} t:%11d ", __Vrece.String(), _FtimeI64())

		__VreceB, __Verr3 := _FencGob__only(&__Vrece) // _TudpNodeDataReceX
		if nil != __Verr3 {
			_FpfN(" 831818 08 gobEnc error: {%v} t:%11d {%s}",
				__Verr3, _FtimeI64(),
				__Vrece.String(),
			)
			continue
		}
		_CpfN(" 831818 09 receOriginByte: (%d){%x}[%s] t:%11d {%s}",
			len(__VreceB),
			_FgenMd5__5(&__VreceB),
			String9(&__VreceB),
			_FtimeI64(),
			__Vrece.String(),
		)

		___lock001.Lock()
		// *chan _TudpNodeDataReceX
		//(*___Vun.unCHreceLO) <- __Vrece
		(*___Vun.unCHreceLO) <- __VreceB
		___lock001.Unlock()

		//_Fsleep_1s()
	}
}

/*

func (___Vun *_TudpNodeSt) _FudpNode__500101yy3__receiveCallBack_default__randDecodeOut_noKeyWillDirect() {

	if ___Vun.unRkeyX.disable {
		var __Vrece _TudpNodeDataRece
		__Vrece = _TudpNodeDataRece{ // _TudpNodeDataReceX
			UrrRemoteAddr: ___Vun.unRemoteAddr,
			UrrBuf:        ___Vun.unRbuf[:___Vun.unRlen],
			UrrReceiveKey: ___Vun.unRkeyX,
		}
		//copy(__Vrece.UrrBuf,        ___Vun.unRbuf[:___Vun.unRlen])
		(*___Vun.unCHreceLO) <- __Vrece
		_FpfNhex(&___Vun.unRbuf, 30, " 439191 01 key disabled ,skip rece rand decode")
		return
	}
	//_FpfNhex(&___Vun.unRbuf, 40, " 439191 02 custom receive ")
	//___Vun._FudpNode__500101yy4__receiveCallBack_default__randDecodeOut_mustDecode(&__Vrece.UrrBuf)
	___Vun._FudpNode__500101yy4__receiveCallBack_default__randDecodeOut_mustDecode(&___Vun.unRbuf)

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

	if len(*___VbufIn) < ___Vun.unRlen {
		_Fex("439192 03 why len error ?")
	}

	__Vtmp2 := (*___VbufIn)[:___Vun.unRlen]

	//__Vtmp3, __Verr2 := _FdecAesRand__only(&___Vun.unRkeyX.Bkey, ___VbufIn)
	__Vtmp3, __Verr2 := _FdecAesRand__only(&___Vun.unRkeyX.Bkey, &__Vtmp2)
	if nil != __Verr2 {
		//_FpfN(" 439192 04 rece buf: %v ", ___VbufIn)
		_FpfN(" 439192 05 AES ERROR : %d ", _FtimeInt())
		_CpfN(" 439192 06 rece Null or AES-decode error : rL:%d-bufL:%d/%d receM:%x ti:%11d remote:%s local:%s ReceKey:%s. error:%v ",
			___Vun.unRlen,
			len(*___VbufIn),
			len(__Vtmp2),
			_FgenMd5__5(&__Vtmp2),
			_FtimeI64(),
			___Vun.unRemoteAddr.String(),
			___Vun.unLocalAddr.String(),
			String5(&___Vun.unRkeyX.Bkey),
			__Verr2)
		_CpfN(" 439192 07 udpNodeInfo:{%s}", ___Vun.String())
		_CpfN(" 439192 08 len(%d/%d):md5sum{%x}", ___Vun.unRlen, len(___Vun.unRbuf), _FgenMd5__5(&__Vtmp3))
		//_Fex1(" 439192 09 : data_error , maybe re-run is needed")
		return
	}

	__VuRece := _TudpNodeDataRece{ // _TudpNodeDataReceX
		UrrRemoteAddr: ___Vun.unRemoteAddr,
		UrrBuf:        __Vtmp3,
	}
	if 3 == 3 {

		_CpfN("439195 01:after AESdec: %11d : me:%s ,remote:%s, before(%d/%d){%x} , after:%d/%d{%x} ",
			_FtimeI64(),
			___Vun.unLocalAddr.String(),
			___Vun.unRemoteAddr.String(),
			___Vun.unRlen,
			len(__Vtmp2),
			_FgenMd5__5(&__Vtmp2),
			___Vun.unRlen,
			len(__Vtmp3),
			_FgenMd5__5(&__Vtmp3),
		)
		//_CpfN("439195 03 me:%s : before AES:<%x> after:<%x>", ___Vun.unLocalAddr.String(), __Vtmp2, __Vtmp3)
	}

	if nil == ___Vun.unCHreceLO {
		_FpfN(" 439196 07 udpNodeDataRece can NOT output , for outChan is null : %s", __VuRece.String())
	} else {
		if 2 == 3 {
			_CpfN(" 439196 08 udpNodeDataRece %11d la:%v allBuf:(%d){%x} lenBuf(%d){%x}: [%s]",
				_FtimeInt(),
				___Vun.unLocalAddr,
				len(*___VbufIn),
				_FgenMd5__5(___VbufIn),
				len(__Vtmp3),
				_FgenMd5__5(&__Vtmp3),
				__VuRece.String()) // _TudpNodeDataReceX
		}
		//_FpfN(" 439196 09 udpNodeDataRece : %s", __VuRece.String())
		(*___Vun.unCHreceLO) <- __VuRece
	}
}
*/
