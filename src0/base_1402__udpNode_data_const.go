package main

import (
	"net"
)

type _TudpNodeDataReceX struct {
}
type _TudpNodeDataRece struct {
	UrrLocalPort  int
	UrrRemoteAddr net.UDPAddr
	UrrBuf        []byte
	UrrReceiveKey _Tkey256
}

func (___Vr *_TudpNodeDataRece) String() string {
	return _Spf("mp:%d ra:%s buf(%d){%x}[%s] k[%s]",
		___Vr.UrrLocalPort,
		___Vr.UrrRemoteAddr.String(),
		len(___Vr.UrrBuf),
		_FgenMd5__5(&___Vr.UrrBuf),
		String9(&___Vr.UrrBuf),
		___Vr.UrrReceiveKey.String(), // UrrReceiveKey _Tkey256X
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
