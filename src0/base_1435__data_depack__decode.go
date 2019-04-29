package main

import (
	"bytes"
	//"net"
)

func (__Vundr *_TudpNodeDataRece) _FdataPack__dePack_decode__from_udpNodeDataRece(___Vdecode *_Tdecode) {
	_FdataPack__dePackUdpNodeRece__decode(___Vdecode, __Vundr.UrrBuf)
}

func _FdataPack__dePackUdpNodeRece__decode(___Vdecode *_Tdecode, ___Vbuf []byte) {
	if nil == ___Vdecode {
		return
	}
	(*___Vdecode) = _Tdecode{}

	__Vlen := len(___Vbuf)
	if __Vlen < (1 + 4 + 32 + 32) {
		_FpfNdb(" 387192 02 : data decode start ")
		return
	}

	if false == bytes.Equal(___Vbuf[1:5], _VersionProtocol01) {
		_FpfNdb(" 387192 03 : protocol version error ")
		return
	}

	var __Verr2 error
	__Vbuf2 := ___Vbuf[37:]

	switch ___Vbuf[0] {
	case LoadT__loginS01genReplyTokenA, LoadT__loginS02genReplyTokenB,
		LoadT__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken:
		//___Vdecode.DElogin = _TloginReq{}
		__Verr2 =
			_FdecGob___(" 387193 01 ", __Vbuf2, &___Vdecode.DElogin)
	case LoadT__data_01_special, // 5
		LoadT__data_99_normal: // 6
		__Verr2 =
			_FdecGob___(" 387193 02 ", __Vbuf2, &___Vdecode.DEdata) // _Tdecode _TdataTran

	default:
		_FpfNhex(&___Vbuf, 500, " 387193 05 : under constructing : type %d ,", ___Vbuf[0])
		return
	}

	if nil != __Verr2 {
		_FpfNdb(" 387193 06 :error :%v", __Verr2)
		return
	}

	___Vdecode.DEremotePortKey = ___Vbuf[5:37] // _TdecodeX
	___Vdecode.DEsendPortKey = ___Vbuf[5:37]   // _TdecodeX
	___Vdecode.DEok = true
	___Vdecode.DEtype = ___Vbuf[0]
	___Vdecode.DEreceiveTime = _FtimeInt()

	_NpfN(" 387193 07 : uDecode decode reult (%s)", ___Vdecode.String()) // _Tdecode _TdataTran
	//_FpfNdb(" 387193 08 : %#v, key %x", ___Vdecode, ___Vbuf[5:37])
}
