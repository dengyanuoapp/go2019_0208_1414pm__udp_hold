package main

import (
	"net"
)

type _TudpNodeDataReceX struct {
}
type _TudpNodeDataRece struct {
	Ti            int
	UrrLocalPort  int
	UrrRemoteAddr net.UDPAddr
	UrrBuf        []byte
	UrrReceiveKey _Tkey256 // _Tkey256X
}

func (___Vr *_TudpNodeDataRece) String() string {
	return _Spf("Ti:%d mp:%d ra:%s buf(%d){%x}[%s] k[%s]",
		___Vr.Ti,
		___Vr.UrrLocalPort,
		___Vr.UrrRemoteAddr.String(),
		len(___Vr.UrrBuf),
		_Fmd5__5x(&___Vr.UrrBuf),
		String9(&___Vr.UrrBuf),
		___Vr.UrrReceiveKey.String(), // UrrReceiveKey _Tkey256X
	)
}

type _TudpNodeDataSendX struct {
}
type _TudpNodeDataSend struct {
	Ti       int
	usToAddr _TudpConnPort
	usOutBuf []byte
}

func (___Vs *_TudpNodeDataSend) String() string {
	return _Spf("Ti:%d send<%s>(%d){%x}",
		___Vs.Ti,
		___Vs.usToAddr.String(),
		len(___Vs.usOutBuf),
		_Fmd5__5x(&___Vs.usOutBuf))

}
