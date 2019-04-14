package main

import (
	//"bytes"
	"net"
)

const (
	Cmd__NULL                    = iota // 0
	Cmd__loginS01genReplyTokenA         // 1
	Cmd__loginS02genReplyTokenB         // 2
	Cmd__loginS03acceptWithToken        // 3
	Cmd__loginEnd                       // 4
	Cmd__data_01_idle                   // 5
	Cmd__data_11_chan_new_req           // 6
	Cmd__dataEnd                        // 7
)

func _FcmdType(___Vc byte) string {
	switch ___Vc {
	case Cmd__NULL:
		return "Cmd__NULL"

	case Cmd__loginS01genReplyTokenA:
		return "Cmd__loginS01genReplyTokenA"
	case Cmd__loginS02genReplyTokenB:
		return "Cmd__loginS02genReplyTokenB"
	case Cmd__loginS03acceptWithToken:
		return "Cmd__loginS03acceptWithToken"

	case Cmd__loginEnd:
		return "Cmd__loginEnd"

	case Cmd__data_01_idle:
		return "Cmd__data_01_idle"
	case Cmd__data_11_chan_new_req:
		return "Cmd__data_11_chan_new_req"

	case Cmd__dataEnd:
		return "Cmd__dataEnd"
	default:
		return "Cmd__unknown"
	}
}

var (
	_VersionProtocol01    = []byte{0x11, 0x22, 0x33, 0x44}
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

type _TdecodeX struct {
}
type _Tdecode struct {
	Ti            int // traceInfo
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
	case Cmd__loginS01genReplyTokenA, Cmd__loginS02genReplyTokenB, Cmd__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken:
		__Vcontent = "Dlogin:" + ___Vd.Dlogin.String()
	default:
		__Vcontent = _Spf("Content:===under constructing %d===", ___Vd.Type)
	}
	__Vrs := _Spf(
		"Ti:%d ok:%T rm:%s rmk:%x type:%s {%s} t:%d ",
		___Vd.Ti,
		___Vd.ok,
		___Vd.remoteAddr.String(),
		___Vd.remotePortKey[:8],
		_FcmdType(___Vd.Type),
		__Vcontent,
		___Vd.receiveTime)
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

type _TencodeX struct {
}
type _Tencode struct {
	Ti           int
	enToId128    []byte        // one of to addr : use enToConnPort if zero , or use this as to addess
	enToConnPort _TudpConnPort // another of to addr
	enType       byte
	enLogin      _TloginReq
	enData       _TdataTran
	enDelay      int // a delay resend if not zero
	enMultiSend  int // send multi timeS if not zero
}

func (___Ven *_Tencode) String() string {
	if nil == ___Ven {
		return ""
	}

	var __Vso string
	__Vso = _Spf("Ti:%d to:%s toId:%s T:%d", ___Ven.Ti, ___Ven.enToConnPort.String(), String5(&___Ven.enToId128), ___Ven.enType)

	switch ___Ven.enType {
	case Cmd__loginS01genReplyTokenA, Cmd__loginS02genReplyTokenB,
		Cmd__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken: // 15540362231554036223
		__Vso += _Spf(" cmd:{%s}", ___Ven.enLogin.String()) // _TloginReq
	case Cmd__data_01_idle, Cmd__data_11_chan_new_req:
		__Vso += _Spf(" data:{%s}", ___Ven.enData.String()) // _TdataTran
	default:
		__Vso += "===839818918unknown==="
	}

	return __Vso
}
