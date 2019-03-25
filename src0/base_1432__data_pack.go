package main

import (
	"bytes"
)

const (
	Cmd__NULL = iota
	Cmd__idle
	Cmd__data
	Cmd__loginS1ReqTryNoToken
	Cmd__loginS2ReplyTmpToken
	Cmd__loginS3ReqWithToken
	Cmd__end
)

var (
	_VersionProtocol01    = []byte{0x83, 0x20, 0x71, 0xc8}
	_VdataPackageMinLen   = 1 + 4 + 32 + 32
	_VdataPackageKeyStart = 1 + 4
	_VdataPackageObjStart = 1 + 4 + 32
)

/*
type _TdataPack_991 struct {
	C byte    // cmd
	V [4]byte // version
	K [32]byte // TheSendPortReceiveKey
	D []byte  // gob.package
}
*/

// _TreqIneedToLogin
func (___VuConnPort *_TudpConnPort) _FdataPack__101__udpConnPort(___VoutBuf *[]byte) {
	__Vreq := _TreqIneedToLogin{
		MeTime:   _FtimeI64(),
		ReqStr:   "try_to_login01",
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _VS.meSeq128,
		Token:    ___VuConnPort.TK,
		//ToIdx128 []byte,
		//ToSeq128 []byte,
	}

	__Vb2, __Verr2 := _FencGob__only(&__Vreq)
	if nil != __Verr2 {
		_FpfN(" 387191 01 : %v", __Verr2)
		*___VoutBuf = []byte{}
		return
	}

	__Vlen2 := len(__Vb2)
	//_FpfN(" 381923 02 : %x", __Vb2)
	//_FpfN(" 381923 03 : len %d: %v", __Vlen2, __Vb2)

	*___VoutBuf = make([]byte, __Vlen2+37)
	(*___VoutBuf)[0] = Cmd__loginS1ReqTryNoToken
	copy((*___VoutBuf)[1:], _VersionProtocol01)
	copy((*___VoutBuf)[37:], __Vb2)

	//_FpfN(" 381923 05 : len %d: %v", len(___VoutBuf), ___VoutBuf)
}

func (__Vundr *_TudpNodeDataRece) _FdataPack__301__dataDecode() (*_TreqIneedToLogin, bool) {
	//_FpfNdb(" 387192 01 : data decode start ")
	if __Vundr.urrLen < (1 + 4 + 32 + 32) {
		_FpfNdb(" 387192 02 : data decode start ")
		return nil, false
	}

	if false == bytes.Equal(__Vundr.urrBuf[1:5], _VersionProtocol01) {
		_FpfNdb(" 387192 03 : protocol version error ")
		return nil, false
	}

	if __Vundr.urrBuf[0] >= Cmd__end {
		_FpfNdb(" 387192 04 : unknown type ")
		return nil, false
	}

	if __Vundr.urrBuf[0] != Cmd__loginS1ReqTryNoToken {
		_FpfNdb(" 387192 05 : under constructing ")
		return nil, false
	}

	_FpfNdb(" 387192 06 : Cmd__loginS1ReqTryNoToken decode start ")

	var __VloginS1ReqTryNoToken _TreqIneedToLogin

	__Vbuf2 := __Vundr.urrBuf[37:]
	__Verr2 := _FdecGob___(" 387193 01 ", &__Vbuf2, &__VloginS1ReqTryNoToken)
	if nil != __Verr2 {
		_FpfNdb(" 387193 03 :error :%v", __Verr2)
		return nil, false
	}

	_FpfNdb(" 387193 05 : %#v, key %x", __VloginS1ReqTryNoToken, __Vundr.urrBuf[5:37])
	return &__VloginS1ReqTryNoToken, true
}
