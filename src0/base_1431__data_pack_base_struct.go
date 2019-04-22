package main

import (
	//"bytes"
	"net"
)

const (
	LoadT__NULL                    = iota // 0
	LoadT__loginS01genReplyTokenA         // 1
	LoadT__loginS02genReplyTokenB         // 2
	LoadT__loginS03acceptWithToken        // 3
	LoadT__loginEnd                       // 4
	LoadT__data_01_idle                   // 5
	LoadT__data_11_chan_new_req           // 6
	LoadT__dataEnd                        // 7
)

func _FcmdType(___Vc byte) string {
	switch ___Vc {
	case LoadT__NULL:
		return "LoadT__NULL"

	case LoadT__loginS01genReplyTokenA:
		return "LoadT__loginS01genReplyTokenA"
	case LoadT__loginS02genReplyTokenB:
		return "LoadT__loginS02genReplyTokenB"
	case LoadT__loginS03acceptWithToken:
		return "LoadT__loginS03acceptWithToken"

	case LoadT__loginEnd:
		return "LoadT__loginEnd"

	case LoadT__data_01_idle:
		return "LoadT__data_01_idle"
	case LoadT__data_11_chan_new_req:
		return "LoadT__data_11_chan_new_req"

	case LoadT__dataEnd:
		return "LoadT__dataEnd"
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
	Ok            bool
	ReceiveTime   int
	RemoteAddr    net.UDPAddr
	RemotePortKey []byte
	Type          byte
	Dlogin        _TloginReq
	Ddata         _TdataTran
}

func (___Vd *_Tdecode) String() string {
	var __Vcontent string
	switch ___Vd.Type {
	case LoadT__loginS01genReplyTokenA, LoadT__loginS02genReplyTokenB, LoadT__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken:
		__Vcontent = "Dlogin:" + ___Vd.Dlogin.String()
	case LoadT__data_01_idle, LoadT__data_11_chan_new_req: //
		__Vcontent = "Ddata:" + ___Vd.Ddata.String()
	default:
		__Vcontent = _Spf("Content:===under constructing %d===", ___Vd.Type)
	}
	__Vrs := _Spf(
		"Ti:%d ok:%T rm:%s rmk:%s type:%s {%s} t:%d ",
		___Vd.Ti,
		___Vd.Ok,
		___Vd.RemoteAddr.String(),
		String5(&___Vd.RemotePortKey),
		_FcmdType(___Vd.Type),
		__Vcontent,
		___Vd.ReceiveTime)
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
	__Vso = _Spf("Ti:%d to:<%s> toId:%s T:%d", ___Ven.Ti, ___Ven.enToConnPort.String(), String5(&___Ven.enToId128), ___Ven.enType)

	switch ___Ven.enType {
	case LoadT__loginS01genReplyTokenA, LoadT__loginS02genReplyTokenB,
		LoadT__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken: // 15540362231554036223
		__Vso += _Spf(" cmd:{%s}", ___Ven.enLogin.String()) // _TloginReq
	case LoadT__data_01_idle, LoadT__data_11_chan_new_req:
		__Vso += _Spf(" data:{%s}", ___Ven.enData.String()) // _TdataTran
	default:
		__Vso += "===839818918unknown==="
	}
	__Vso += _Spf(" delay:%d multi:%d", ___Ven.enDelay, ___Ven.enMultiSend)

	return __Vso
}
