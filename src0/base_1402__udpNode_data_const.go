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
	return _Spf("Ti:%d mp:%d ra:%s rk[%s] buf{%s} ",
		___Vr.Ti,
		___Vr.UrrLocalPort,
		___Vr.UrrRemoteAddr.String(),
		//___Vr.UrrReceiveKey.String(), // UrrReceiveKey _Tkey256X
		String5(&___Vr.UrrReceiveKey.Bkey), // UrrReceiveKey _Tkey256X
		String9(&___Vr.UrrBuf),
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
	return _Spf("Ti:%d sa:%s {%s}",
		___Vs.Ti,
		___Vs.usToAddr.String(),
		String9(&___Vs.usOutBuf),
	)

}
