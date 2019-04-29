package main

import (
	"net"
	"sync"
)

type _TacceptTCP struct {
	taId128      []byte
	taEnabled    bool
	taIdx        int
	taServerTCP  *_TtcpNodE
	taConnTCP    *net.TCPConn
	taErr        error
	taCHreceLO   *chan _TtcpNodeDataRece
	taRcnt       _Tcount
	taWcnt       _Tcount
	taBuf        []byte
	taRdata      _TtcpNodeDataRece
	taLen        int
	taErr2       error
	taRemoteAddr net.Addr
	taLocalAddr  net.Addr
	//taCreceiveMsg chan []byte
	//taCchanMsg    chan []byte
	//taCreceiveErr chan string
} // _TacceptTCP
func (___Vacc *_TacceptTCP) String() string {
	return _Spf(
		"id:%s ena:%T idx:%d R{%s} W{%s} r:%s l:%s",
		String5s(&___Vacc.taId128),
		___Vacc.taEnabled,
		___Vacc.taIdx,
		___Vacc.taRcnt.String(),
		___Vacc.taWcnt.String(),
		___Vacc.taRemoteAddr.String(),
		___Vacc.taLocalAddr.String(),
	)
}

type _TtcpNodE struct {
	tnName               string
	tnHostPortStr        string
	tnAmount             int
	tnAddr               *net.TCPAddr
	tnListener           *net.TCPListener
	tnLisnAddr           net.Addr
	tnErr                error
	tnCBinit             func(*_TtcpNodE) // _FtcpNode__200101x__init_default
	tnCHsendToAllClientI chan _TtcpNodeDataSend
	tnCHdebugInfoLO      *chan byte // if not-null , when receiving , send one byte into this byte to info the other monitorInfoByteChan
	tnAcceptTCPs         []_TacceptTCP
	tnClientMux          sync.Mutex
	tnClientCnt          int
	tnCBaccDataChan      func(*_TacceptTCP) // _FuserCallback__accept_dataChan__Log_Fn
	tnCBsvrDataChan      func(*_TtcpNodE)   // _FuserCallback__service_dataChan__Log_Fn
} // _TtcpNodE
