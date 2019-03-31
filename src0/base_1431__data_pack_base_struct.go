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
