package main

import (
	//"bytes"
	"net"
)

const (
	LoadT__NULL                    = 101
	LoadT__loginS01genReplyTokenA  = 102
	LoadT__loginS02genReplyTokenB  = 103
	LoadT__loginS03acceptWithToken = 104
	LoadT__loginEnd                = 105
	LoadT__data_01_special         = 106
	LoadT__data_99_normal          = 107
	LoadT__dataEnd                 = 108
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

	case LoadT__data_01_special:
		return "LoadT__data_01_special"
	case LoadT__data_99_normal:
		return "LoadT__data_99_normal"

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
	SendAddr      net.UDPAddr
	SendPortKey   []byte
	Type          byte
	Dlogin        _TloginReq
	Ddata         _TdataTran
}

func (___Vd *_Tdecode) String() string {
	var __Vcontent string
	switch ___Vd.Type {
	case LoadT__loginS01genReplyTokenA, LoadT__loginS02genReplyTokenB, LoadT__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken:
		__Vcontent = "Dlogin:" + ___Vd.Dlogin.String()
	case LoadT__data_01_special, LoadT__data_99_normal: //
		__Vcontent = "Ddata:" + ___Vd.Ddata.String()
	default:
		__Vcontent = _Spf("Content:===under constructing %d===", ___Vd.Type)
	}
	__Vrs := _Spf(
		"Ti:%d ok:%T rm:%s rmk:%s sa:%s sk:%s type:%s {%s} t:%d ",
		___Vd.Ti,
		___Vd.Ok,
		___Vd.RemoteAddr.String(),
		String5(&___Vd.RemotePortKey),
		___Vd.SendAddr.String(),
		String5(&___Vd.SendPortKey),
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
	EnToId128    []byte        // one of to addr : use EnToConnPort if zero , or use this as to addess
	EnFromId128  []byte        // one of to addr : use EnToConnPort if zero , or use this as to addess
	EnToConnPort _TudpConnPort // another of to addr
	EnLoadType   byte
	EnLogin      _TloginReq
	EnData       _TdataTran
	EnDelay      int // a delay resend if not zero
	EnMultiSend  int // send multi timeS if not zero
}

func (___Ven *_Tencode) String() string {
	if nil == ___Ven {
		return ""
	}

	var __Vso string
	__Vso = _Spf("Ti:%d to:<%s> fromId:%s toId:%s T:%d", ___Ven.Ti, ___Ven.EnToConnPort.String(),
		String5(&___Ven.EnFromId128), String5(&___Ven.EnToId128), ___Ven.EnLoadType)

	switch ___Ven.EnLoadType {
	case LoadT__loginS01genReplyTokenA, LoadT__loginS02genReplyTokenB,
		LoadT__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken: // 15540362231554036223
		__Vso += _Spf(" cmd:{%s}", ___Ven.EnLogin.String()) // _TloginReq
	case LoadT__data_01_special, LoadT__data_99_normal:
		__Vso += _Spf(" data:{%s}", ___Ven.EnData.String()) // _TdataTran
	default:
		__Vso += _Spf("===839818918unknown(%d)===", ___Ven.EnLoadType)
	}
	__Vso += _Spf(" delay:%d multi:%d", ___Ven.EnDelay, ___Ven.EnMultiSend)

	return __Vso
}
