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
		String5(&__Vtndr.tnrId128),
		String5(&__Vtndr.tnrK256),
		__Vtndr.tnrLen,
		len(__Vtndr.tnrBuf),
		String5(&__Vtndr.tnrBuf),
	)

}

type _TtcpNodeDataSend struct {
	tnsToAddr net.Addr
	tnsLen    int
	tnsId128  []byte // channel id
	tnsK256   []byte // AES key 256 using when send if not nil
	tnsBuf    []byte
}

func (__Vtndr *_TtcpNodeDataSend) String() string {
	//_FpfNex(" ttttt [%t] ", __Vtndr == nil)
	//	_FpfNex(" ttttt [%v] ", __Vtndr)
	__Vso := _Spf(
		//"addr:%s ", // id:%s k:%s (%d/%d)%s",
		//__Vtndr.tnsToAddr.String(),
		"addr:[%v] ", // id:%s k:%s (%d/%d)%s",
		__Vtndr.tnsToAddr)
	//	return _Spf(
	//		"addr:%s id:%s k:%s (%d/%d)%s",
	//		__Vtndr.tnsToAddr.String(),
	//		String5(&__Vtndr.tnsId128),
	//		String5(&__Vtndr.tnsK256),
	//		__Vtndr.tnsLen,
	//		len(__Vtndr.tnsBuf),
	//		String5(&__Vtndr.tnsBuf),
	//	)

	return __Vso
}
