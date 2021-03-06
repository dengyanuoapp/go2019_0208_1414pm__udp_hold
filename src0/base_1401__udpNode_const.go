package main

import (
	"net"
	"sync"
	"time"
)

type _TuNodeDataRcnt struct {
	cnt int
	urr _TudpNodeDataRece
}
type _TuNodeDataRmap struct {
	unrMux     sync.Mutex
	unrMapNow  map[string]_TuNodeDataRcnt
	unrMapLast map[string]_TuNodeDataRcnt
	unrMapLas2 map[string]_TuNodeDataRcnt
}

// _TserviceUDP
// a udp node: input accept AES256 rand only , can output anything
type _TudpNodeSt struct {
	unMux            sync.Mutex
	unRKeyLP         *[]byte
	unRkeyX          _Tkey256
	unAddr           *net.UDPAddr
	unConn           *net.UDPConn
	unLoopGap        time.Duration
	unLocalAddr      net.Addr
	unLocalPort      int
	unErr            error
	unRerr           error
	unSerr           error
	unRbuf1500       []byte // receive-buf
	unRbufTmp        []byte // receive-buf
	unRlen           int
	unRemoteAddr     net.UDPAddr
	unName           string
	unHostPortStr    string
	unIdx            int
	unGroup          *_TudpGroupSt
	unRmap           _TuNodeDataRmap
	unCHsendI        chan _TudpNodeDataSend  // try get data from chan, then send it out.
	unCBinit         func(*_TudpNodeSt)      //
	unCBrece         func(*_TudpNodeSt)      // if nil , use the default procedure to deal with receive
	unCBsend         func(*_TudpNodeSt)      // if nil , use the default procedure to deal with send
	unCHreceByteLO   *chan []byte            // if nil , drop it ; not-nil , put the received data into this chan, note: only one of chan will be active
	unCHreceStructLO *chan _TudpNodeDataRece // if nil , drop it ; not-nil , put the received data into this chan, note: only one of chan will be active
}

func (___Vun _TudpNodeSt) String() string {
	return _Spf(
		"key %x,addr %s,gap %d,Laddr %s,Rbuf(len:%d)%x, Raddr %s name %s,%s,%d",
		___Vun.unRkeyX.B32,
		___Vun.unAddr.String(),
		___Vun.unLoopGap,
		___Vun.unLocalAddr,
		___Vun.unRlen,
		___Vun.unRbuf1500[:20],
		___Vun.unRemoteAddr.String(),
		___Vun.unName,
		___Vun.unHostPortStr,
		___Vun.unIdx)
}
