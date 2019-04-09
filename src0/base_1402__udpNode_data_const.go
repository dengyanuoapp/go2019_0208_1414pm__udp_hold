package main

import (
	"net"
)

type _TudpNodeDataReceX struct {
}
type _TudpNodeDataRece struct {
	urrRemoteAddr net.UDPAddr
	urrBuf        []byte
	urrReceiveKey _Tkey256
}

func (___Vr *_TudpNodeDataRece) String() string {
	return _Spf("rm:%s (%d){%x} buf[%s] k[%s]",
		___Vr.urrRemoteAddr.String(),
		len(___Vr.urrBuf),
		_FgenMd5__5(&___Vr.urrBuf),
		String9(&___Vr.urrBuf),
		___Vr.urrReceiveKey.String(),
	)
}

type _TudpNodeDataSendX struct {
}
type _TudpNodeDataSend struct {
	usToAddr _TudpConnPort
	usOutBuf []byte
}

func (___Vs *_TudpNodeDataSend) String() string {
	return _Spf("{%s}(%d %s)",
		___Vs.usToAddr.String(),
		len(___Vs.usOutBuf),
		String5(&___Vs.usOutBuf))

}
