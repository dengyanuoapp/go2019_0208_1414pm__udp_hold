package main

import (
	"net"
	//"sync"
)

type _TudpConnPort struct {
	//Uri  string
	DstAddr net.UDPAddr
	K256    []byte
	TKme    []byte // token
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
