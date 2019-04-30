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

type _TdecodeX struct{}
type _Tdecode struct {
	DEti            int // traceInfo
	DErole          string
	DEok            bool
	DEreceiveTime   int
	DEremoteAddr    net.UDPAddr
	DEremotePortKey []byte
	DEtype          byte
	DElogin         _TloginReq
	DEdata          _TdataTran
}

func (___Vd *_Tdecode) String() string {
	var __Vcontent string
	switch ___Vd.DEtype {
	case LoadT__loginS01genReplyTokenA, LoadT__loginS02genReplyTokenB, LoadT__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken:
		__Vcontent = "Dlogin:" + ___Vd.DElogin.String()
	case LoadT__data_01_special, LoadT__data_99_normal: //
		__Vcontent = "Ddata:" + ___Vd.DEdata.String()
	default:
		__Vcontent = _Spf("Content:===under constructing %d===", ___Vd.DEtype)
	}
	__Vrs := _Spf(
		"W:%s Ti:%d ok:%T rm:%s rmk:%s type:%s {%s} t:%d ",
		___Vd.DErole,
		___Vd.DEti,
		___Vd.DEok,
		___Vd.DEremoteAddr.String(),
		String5s(&___Vd.DEremotePortKey),
		_FcmdType(___Vd.DEtype),
		__Vcontent,
		___Vd.DEreceiveTime)
	return __Vrs
}

func (___Vd *_Tdecode) Count128() []int {
	return []int{
		len(___Vd.DElogin.LgToIdx128),
		len(___Vd.DElogin.LgToSeq128),
		len(___Vd.DElogin.LgMeIdx128),
		len(___Vd.DElogin.LgMeSeq128),
		len(___Vd.DElogin.LgTokenL),
		len(___Vd.DElogin.LgTokenR)}
}

type _TencodeX struct{}
type _Tencode struct {
	EnTi         int
	EnRole       string
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
	__Vso = _Spf("W:%s Ti:%d to:<%s> fromId:%s toId:%s T:%d", ___Ven.EnRole, ___Ven.EnTi, ___Ven.EnToConnPort.String(),
		String5s(&___Ven.EnFromId128), String5s(&___Ven.EnToId128), ___Ven.EnLoadType)

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
