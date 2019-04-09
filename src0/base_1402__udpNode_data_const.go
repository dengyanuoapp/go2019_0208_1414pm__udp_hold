package main

import (
	"net"
)

type _TudpNodeDataReceX struct {
}
type _TudpNodeDataRece struct {
	urrLocalPort  int
	urrRemoteAddr net.UDPAddr
	urrBuf        []byte
	urrReceiveKey _Tkey256
}

func (___Vr *_TudpNodeDataRece) String() string {
	return _Spf("mp:%d ra:%s buf(%d){%x}[%s] k[%s]",
		___Vr.urrLocalPort,
		___Vr.urrRemoteAddr.String(),
		len(___Vr.urrBuf),
		_FgenMd5__5(&___Vr.urrBuf),
		String9(&___Vr.urrBuf),
		___Vr.urrReceiveKey.String(), // urrReceiveKey _Tkey256X
	)
}

type _TudpNodeDataSendX struct {
}
type _TudpNodeDataSend struct {
	usToAddr _TudpConnPort
	usOutBuf []byte
}

func (___Vs *_TudpNodeDataSend) String() string {
	return _Spf("send<%s>(%d){%x}",
		___Vs.usToAddr.String(),
		len(___Vs.usOutBuf),
		_FgenMd5__5(&___Vs.usOutBuf))

}
