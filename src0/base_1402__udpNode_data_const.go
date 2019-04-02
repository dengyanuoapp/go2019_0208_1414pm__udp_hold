package main

import (
	"net"
)

type _TudpNodeDataRece struct {
	urrRemoteAddr net.UDPAddr
	urrLen        int
	urrBuf        []byte
}

type _TudpNodeDataSend struct {
	usToAddr _TudpConnPort
	usOutBuf []byte
}

func (___Vs *_TudpNodeDataSend) String() string {
	return _Pspf("{%s}(%d %x)",
		___Vs.usToAddr.String(),
		len(___Vs.usOutBuf),
		___Vs.usOutBuf[:40])

}
func (___Vr *_TudpNodeDataRece) String() string {
	return _Pspf("[%s](%d %x)",
		___Vr.urrRemoteAddr.String(),
		___Vr.urrLen,
		___Vr.urrBuf[:40])
}
