package main

import "net"

type _TtcpNodeDataRece struct {
	tnrAddr  net.Addr
	tnrLen   int
	tnrId128 []byte // channel id
	tnrK256  []byte // AES key 256 using when receive if not nil
	tnrBuf   []byte
}

func (__Vtndr *_TtcpNodeDataRece) String() string {
	return _Spf(
		"addr:%s id:%s k:%s (%d/%d)%s",
		__Vtndr.tnrAddr.String(),
		String5s(&__Vtndr.tnrId128),
		String5s(&__Vtndr.tnrK256),
		__Vtndr.tnrLen,
		len(__Vtndr.tnrBuf),
		String5s(&__Vtndr.tnrBuf),
	)

}

type _TtcpNodeDataSend struct {
	tnsToAddr net.Addr
	tnsId128  []byte // channel id
	tnsK256   []byte // AES key 256 using when send if not nil
	tnsLen    int
	tnsBuf    []byte
}

func (__Vtndr *_TtcpNodeDataSend) String() string {
	if nil == __Vtndr.tnsToAddr {
		return _Spf(
			"addr:<null> id:%s k:%s (%d/%d)%s",
			String5s(&__Vtndr.tnsId128),
			String5s(&__Vtndr.tnsK256),
			__Vtndr.tnsLen,
			len(__Vtndr.tnsBuf),
			String5s(&__Vtndr.tnsBuf),
		)
	} else {
		return _Spf(
			"addr:%s id:%s k:%s (%d/%d)%s",
			__Vtndr.tnsToAddr.String(),
			String5s(&__Vtndr.tnsId128),
			String5s(&__Vtndr.tnsK256),
			__Vtndr.tnsLen,
			len(__Vtndr.tnsBuf),
			String5s(&__Vtndr.tnsBuf),
		)
	}
}
