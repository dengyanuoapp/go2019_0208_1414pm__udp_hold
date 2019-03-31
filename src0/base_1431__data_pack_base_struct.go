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

func CmdType(___Vc int) string {
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
	ok                      bool
	receiveTime             int
	remoteAddr              net.UDPAddr
	remotePortKey           []byte
	Type                    byte
	D__loginS1ReqTryNoToken _TloginReq
}
