package main

import "net"

type _TtcpNodeDataRece struct {
	tnrRemoteAddr net.Addr
	tnrLen        int
	tnrChanId128  []byte
	tnrBuf        []byte
}

func (__Vtndr *_TtcpNodeDataRece) String() string {
	return _Spf(
		"addr:%s id:%s (%d/%d)%s",
		__Vtndr.tnrRemoteAddr.String(),
		String5(&__Vtndr.tnrChanId128),
		__Vtndr.tnrLen,
		len(__Vtndr.tnrBuf),
		String5(&__Vtndr.tnrBuf),
	)

}
