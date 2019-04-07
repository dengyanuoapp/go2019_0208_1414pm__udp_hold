package main

import "net"

// _TsrvDownInfo _TsrvInfo _TudpConnPort
func (___VnewSession *_TsrvDownInfo) _FudpDecode__800401x__tryfillSendChan() *_TudpConnPort {
	__Vlen2 := len(___VnewSession.srvInfo.UriArrs)
	__Vlen3 := len(___VnewSession.srvInfo.K256)
	__Vidx2 := ___VnewSession.tryCnt % __Vlen2

	//Uri: ___VnewSession.srvInfo.UriArrs[__Vidx2],
	// func ResolveUDPAddr(network, address string) (*UDPAddr, error)
	//__Verr2 error
	__VuAddr, __Verr2 := net.ResolveUDPAddr("udp4", ___VnewSession.srvInfo.UriArrs[__Vidx2])
	if nil != __Verr2 {
		_FpfN(" 388181 01 : %v", __Verr2)
		return nil
	}

	__VuConn := _TudpConnPort{}
	__VuConn.DstAddr = *__VuAddr

	if __Vidx2 >= __Vlen3 {
		__VuConn.K256 = ___VnewSession.srvInfo.K256[0]
	} else {
		__VuConn.K256 = ___VnewSession.srvInfo.K256[__Vidx2]
	}
	return &__VuConn
}
