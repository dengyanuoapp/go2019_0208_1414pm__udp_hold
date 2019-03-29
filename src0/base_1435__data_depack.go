package main

import (
	"bytes"
	//"net"
)

func (__Vundr *_TudpNodeDataRece) _FdataPack__301__dataDecode_loginS1ReqTryNoToken(___Vdecode *_Tdecode) {
	_FdataPack__dataDecode_common(___Vdecode, __Vundr.urrLen, __Vundr.urrBuf)
}

func _FdataPack__dataDecode_common(___Vdecode *_Tdecode, ___Vlen int, ___Vbuf []byte) {
	if nil == ___Vdecode {
		return
	}
	___Vdecode.ok = false
	___Vdecode.Type = 0

	if ___Vlen < (1 + 4 + 32 + 32) {
		_FpfNdb(" 387192 02 : data decode start ")
		return
	}

	if false == bytes.Equal(___Vbuf[1:5], _VersionProtocol01) {
		_FpfNdb(" 387192 03 : protocol version error ")
		return
	}

	if ___Vbuf[0] >= Cmd__end {
		_FpfNdb(" 387192 04 : unknown type ")
		return
	}

	if ___Vbuf[0] != Cmd__loginS1ReqTryNoToken {
		_FpfNdb(" 387192 05 : under constructing ")
		return
	}

	//_FpfNdb(" 387192 06 : Cmd__loginS1ReqTryNoToken decode start ")

	__Vbuf2 := ___Vbuf[37:]
	__Verr2 := _FdecGob___(" 387193 01 ", &__Vbuf2, &___Vdecode.D__loginS1ReqTryNoToken)
	if nil != __Verr2 {
		_FpfNdb(" 387193 03 :error :%v", __Verr2)
		return
	}

	___Vdecode.remotePortKey = ___Vbuf[5:37]
	___Vdecode.ok = true
	___Vdecode.Type = ___Vbuf[0]
	___Vdecode.receiveTime = _FtimeInt()

	//_FpfNdb(" 387193 05 : %#v, key %x", ___Vdecode, ___Vbuf[5:37])
}