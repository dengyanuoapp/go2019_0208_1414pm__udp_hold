package main

import (
	"net"
	"sync"
)

type _TacceptTCP struct {
	taId128      []byte
	taEnabled    bool
	taIdx        int // idx ot accArr:w
	taServerTCP  *_TtcpNodE
	taConnTCP    *net.TCPConn
	taRcnt       _Tcount
	taWcnt       _Tcount
	taRemoteAddr net.Addr
	taLocalAddr  net.Addr
	taOffset     int64
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

// LoadT__NULL
const (
	TcpNodeCmd__NULL = 31
	TcpNodeCmd__Eof  = 32
)

type _TtcpNodE struct {
	tnName               string
	tnHostPortStr        string
	tnAmount             int
	tnAddr               *net.TCPAddr
	tnListener           *net.TCPListener
	tnLisnAddr           net.Addr
	tnErr                error
	tnCBinit             func(*_TtcpNodE) // _FtcpNode__200101x__init_default
	tnCHtcpSendCmdI      [17]byte         // command of tunnel : byte 0:15 -> channelID, byte [16] -> cmd : // TcpNodeCmd__NULL
	tnCHsendToAllClientI chan _TtcpNodeDataSend
	tnCHdebugInfoLO      *chan byte     // if not-null , when receiving , send one byte into this byte to info the other monitorInfoByteChan
	tnCHtcpReceBLO       *chan []byte   // byte of _TtcpNodeDataRece
	tnCHtcpReceCmdLO     *chan [17]byte // command of tunnel : byte 0:15 -> channelID, byte [16] -> cmd : // TcpNodeCmd__NULL
	tnAcceptTCPs         []_TacceptTCP
	tnClientMux          sync.Mutex
	tnClientCnt          int
	tnCBaccDataChan      func(*_TacceptTCP) // _FuserCallback__accept_dataChan__Log_Fn
	tnCBsvrDataChan      func(*_TtcpNodE)   // _FuserCallback__service_dataChan__Log_Fn
} // _TtcpNodE
