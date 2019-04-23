package main

import "time"
import "net"

// _TencodeX
// _TdataTran
type _TrecePackThenEncodeAsLoad struct {
	pelGap                 time.Duration // _T1s
	pelCHudpNodeDataReceBI chan []byte
	pelCHencodeBLO         *chan []byte // byte of _TencodeX
	pelCBinit              func(*_TrecePackThenEncodeAsLoad)
	pelChOutputGenGap      chan byte // every 1s gen output
}

// DDbuf    : , // __Vtmp3out , byte of _Tdecode ; __VunRece _TudpNodeDataReceX
type _Trepack__c2s struct {
	C2sDe   _Tdecode
	C2sAddr net.UDPAddr
}

func (___VrpC2s *_Trepack__c2s) String() string {
	return _Spf("client:%s de{%s}", ___VrpC2s.C2sAddr.String(), ___VrpC2s.String())
}
