package main

import (
	"net"
	"sync"
)

/*
// https://www.callicoder.com/golang-basic-types-operators-type-conversion/
// Type    Size        Range
// int8    8 bits      -128 to 127
// int16   16 bits     -2^^15 to 2^^15 -1
// int32   32 bits     -2^^31 to 2^^31 -1
// int64   64 bits     -2^^63 to 2^^63 -1
// int     Platform dependent     Platform dependent
//
// Unsigned Integers
// Type    Size        Range
// uint8   8 bits      0 to 255
// uint16  16 bits     0 to 2^^16 -1
// uint32  32 bits     0 to 2^^32 -1
// uint64  64 bits     0 to 2^^64 -1
// uint    Platform dependent     Platform dependent

//                --====----====----====----====----
//	_FnPasswd = "718291a81b893a8a8a8921a8dad892a1"
//	_DnPasswd = "b87288938a81838e828291823719aba2"
//	_SnPasswd = "ca79319381829379abe9abe889ac8aa3"
//	_CnPasswd = "3019303f301ab921b8a8c8a9018a78a4"
*/

type _TudpConnPort struct {
	//Uri  string
	DstAddr net.UDPAddr
	K256    []byte
	TK      []byte // token
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

type _TloginReq struct {
	MeTime   int
	ReqStr   string
	MeName   string
	MeIdx128 []byte
	MeSeq128 []byte
	ToIdx128 []byte
	ToSeq128 []byte
	TokenA   []byte
	TokenB   []byte
}

func (___Vlr *_TloginReq) String() string {
	return "===0aaeeb6ff2c8cd980641fdecf4f640b2==="
}

//type _TreqLoginCNT struct {
//	cntL int
//	reqL _TloginReq
//}
//
//type _TreqLogintMap struct {
//	muxLogin   sync.Mutex
//	reqMapNow  map[[16]byte]_TreqLoginCNT
//	reqMapLast map[[16]byte]_TreqLoginCNT
//}

type _TuAcceptClientSt struct {
	cId128   [16]byte
	reqA     _TloginReq
	CexitAcc chan []byte
}

type _TuAcceptClientMap struct {
	muxAcc           sync.Mutex
	maxClient        int
	cntClient        int
	maxConnPerClient int
	uMap             map[[16]byte][]_TuAcceptClientSt
}
