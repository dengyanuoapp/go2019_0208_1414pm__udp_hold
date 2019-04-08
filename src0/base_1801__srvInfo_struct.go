package main

import (
	"net"
	//"sync"
)

type _TudpConnPort struct {
	DstAddr net.UDPAddr
	K256    []byte
}

func (___Vucp *_TudpConnPort) String() string {
	return _Spf("%s k:%s",
		___Vucp.DstAddr.String(),
		String5(&___Vucp.K256))
}

type _TsrvInfo struct {
	ok      bool
	UriArrs []string // try-Uris
	//UdstAddr   []net.UDPAddr
	K256       [][]byte // passwd to connect the this server
	name       string   // name
	refreshUri string   // refresh-uri-Github
	refreshPwd []byte
} // _TsrvInfo
func (___Vsi *_TsrvInfo) String() string {
	__Vos := _Spf("ok:%t uri[", ___Vsi.ok)
	for __Vidx2, __Vuri2 := range ___Vsi.UriArrs {
		if 0 == __Vidx2 {
			__Vos += __Vuri2
		} else {
			__Vos += " " + __Vuri2
		}
	}

	__Vos += _Spf("] k[", ___Vsi.ok)
	for __Vidx3, __Vuri3 := range ___Vsi.K256 {
		if 0 == __Vidx3 {
			__Vos += String5(&__Vuri3)
		} else {
			__Vos += " " + String5(&__Vuri3)
		}
	}

	__Vos += "] n:" + ___Vsi.name + " rUri[" + ___Vsi.refreshUri + "] rPwd[" + String5(&___Vsi.refreshPwd) + "]"

	return __Vos
}
