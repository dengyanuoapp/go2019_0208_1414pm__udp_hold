package main

import (
	"net"
)

type _TudpNodeDataReceX struct {
}
type _TudpNodeDataRece struct {
	urrRemoteAddr net.UDPAddr
	urrLen        int
	urrBuf        []byte
}

func (___Vr *_TudpNodeDataRece) String() string {
	return _Spf("[%s](%d len:%d %x)",
		___Vr.urrRemoteAddr.String(),
		___Vr.urrLen,
		len(___Vr.urrBuf),
		___Vr.urrBuf[:40])
}

type _TudpNodeDataSendX struct {
}
type _TudpNodeDataSend struct {
	usToAddr _TudpConnPort
	usOutBuf []byte
}

func (___Vs *_TudpNodeDataSend) String() string {
	return _Spf("{%s}(%d %5)",
		___Vs.usToAddr.String(),
		len(___Vs.usOutBuf),
		String5(&___Vs.usOutBuf))

}
