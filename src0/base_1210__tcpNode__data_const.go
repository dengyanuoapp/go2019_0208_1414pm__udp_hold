package main

import "net"

type _TtcpNodeDataRece struct {
	TnrRaddr  string // string of net.Addr
	TnrLen    int
	TnrId128  []byte // channel id
	TnrK256   []byte // AES key 256 using when receive if not nil
	TnrBuf    []byte
	TnrOffset int64
}

func (__Vtndr *_TtcpNodeDataRece) String() string {
	return _Spf(
		"addr:%s id:%s k:%s os:%d (%d/%d)%s",
		__Vtndr.TnrRaddr,
		String5s(&__Vtndr.TnrId128),
		String5s(&__Vtndr.TnrK256),
		__Vtndr.TnrOffset,
		__Vtndr.TnrLen,
		len(__Vtndr.TnrBuf),
		String5s(&__Vtndr.TnrBuf),
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
