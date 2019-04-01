package main

import (
	//"bytes"
	"net"
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

func _FcmdType(___Vc byte) string {
	switch ___Vc {
	case Cmd__NULL:
		return "Cmd__NULL"
	case Cmd__idle:
		return "Cmd__idle"
	case Cmd__data:
		return "Cmd__data"
	case Cmd__loginS1ReqTryNoToken:
		return "Cmd__loginS1ReqTryNoToken"
	case Cmd__loginS2ReplyTmpToken:
		return "Cmd__loginS2ReplyTmpToken"
	case Cmd__loginS3ReqWithToken:
		return "Cmd__loginS3ReqWithToken"
	default:
		return "Cmd__unknown"
	}
}

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

type _Tdecode struct {
	ok            bool
	receiveTime   int
	remoteAddr    net.UDPAddr
	remotePortKey []byte
	Type          byte
	Dlogin        _TloginReq
	Ddata         _TdataTran
}

func (___Vd *_Tdecode) String() string {
	var __Vcontent string
	switch ___Vd.Type {
	case Cmd__loginS1ReqTryNoToken:
		__Vcontent = ___Vd.Dlogin.String()
	default:
		__Vcontent = _Pspf("===under constructing %d===", ___Vd.Type)
	}
	__Vrs := _Pspf(
		"ok:%T rece %d addr %s key %x %s :%s",
		___Vd.ok,
		___Vd.receiveTime,
		___Vd.remoteAddr.String(),
		___Vd.remotePortKey,
		_FcmdType(___Vd.Type),
		__Vcontent)
	return __Vrs
}
