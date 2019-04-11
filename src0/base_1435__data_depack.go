package main

import (
	"bytes"
	//"net"
)

func (__Vundr *_TudpNodeDataRece) _FdataPack__decode_from_udpNodeDataRece(___Vdecode *_Tdecode) {
	_FdataPack__dataDecode_common(___Vdecode, __Vundr.UrrBuf)
}

func _FdataPack__dataDecode_common(___Vdecode *_Tdecode, ___Vbuf []byte) {
	if nil == ___Vdecode {
		return
	}
	//___Vdecode.ok = false
	//___Vdecode.Type = 0
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

	if ___Vbuf[0] >= Cmd__loginEnd {
		_FpfNdb(" 387192 04 : unknown type ")
		return
	}

	switch ___Vbuf[0] {
	case Cmd__loginS01genReplyTokenA, Cmd__loginS02genReplyTokenB,
		Cmd__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken:
		__Vbuf2 := ___Vbuf[37:]
		//___Vdecode.Dlogin = _TloginReq{}
		__Verr2 := _FdecGob___(" 387193 01 ", __Vbuf2, &___Vdecode.Dlogin)
		if nil != __Verr2 {
			_FpfNdb(" 387193 03 :error :%v", __Verr2)
			return
		}

	default:
		_FpfNhex(&___Vbuf, 500, " 387193 05 : under constructing : type %d ,", ___Vbuf[0])
	}

	___Vdecode.remotePortKey = ___Vbuf[5:37]
	___Vdecode.ok = true
	___Vdecode.Type = ___Vbuf[0]
	___Vdecode.receiveTime = _FtimeInt()

	//_FpfNdb(" 387193 07 : %#v, key %x", ___Vdecode, ___Vbuf[5:37])
}
