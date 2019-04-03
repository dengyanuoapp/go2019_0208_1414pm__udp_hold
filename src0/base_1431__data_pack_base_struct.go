package main

import (
	//"bytes"
	"net"
)

const (
	Cmd__NULL                   = iota // 0
	Cmd__idle                          // 1
	Cmd__loginS01genReplyTokenA        // 2

	Cmd__loginS02genReplyTokenB  // 3
	Cmd__loginS03acceptWithToken // 4
	//Cmd__loginS04acceptWithToken // 5

	Cmd__data // 6
	Cmd__end  // 7
)

func _FcmdType(___Vc byte) string {
	switch ___Vc {
	case Cmd__NULL:
		return "Cmd__NULL"
	case Cmd__idle:
		return "Cmd__idle"
	case Cmd__data:
		return "Cmd__data"
	case Cmd__loginS01genReplyTokenA:
		return "Cmd__loginS01genReplyTokenA"
	case Cmd__loginS02genReplyTokenB:
		return "Cmd__loginS02genReplyTokenB"
	case Cmd__loginS03acceptWithToken:
		return "Cmd__loginS03acceptWithToken"
	//case Cmd__loginS04acceptWithToken:
	//	return "Cmd__loginS04acceptWithToken"
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
	case Cmd__loginS02genReplyTokenB, Cmd__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken:
		__Vcontent = ___Vd.Dlogin.String()
	default:
		__Vcontent = _Pspf("===under constructing %d===", ___Vd.Type)
	}
	__Vrs := _Pspf(
		"ok:%T rece %d addr %s key %x %s :%s",
		___Vd.ok,
		___Vd.receiveTime,
		___Vd.remoteAddr.String(),
		___Vd.remotePortKey[:8],
		_FcmdType(___Vd.Type),
		__Vcontent)
	return __Vrs
}

func (___Vd *_Tdecode) Count128() []int {
	return []int{
		len(___Vd.Dlogin.ToIdx128),
		len(___Vd.Dlogin.ToSeq128),
		len(___Vd.Dlogin.MeIdx128),
		len(___Vd.Dlogin.MeSeq128),
		len(___Vd.Dlogin.TokenL),
		len(___Vd.Dlogin.TokenR)}
}
