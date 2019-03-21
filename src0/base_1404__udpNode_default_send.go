package main

func (___Vun *_TudpNodeSt) _FudpNode__540201z__send() {
	for {
		select {
		case __VchSend := <-___Vun.unCHsendX: // _TudpNodeDataSend
			_FpfN(" 839118 01 send ")

			// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
			_, __Verr2 := ___Vun.unConn.WriteToUDP(__VchSend.usOutBuf, &__VchSend.usDstAddr)
			if __Verr2 != nil {
				_FpfN(" 839118 06 : udp send error <%s>[%v]", __VchSend.usDstAddr, __Verr2)
				return
			}
			_FpfN(" 839118 07 : %s : udp send succeed (%d) 01 dst<%s>, local<%v>, listen<%v>", ___Vun.unName,
				len(__VchSend.usOutBuf), __VchSend.usDstAddr, ___Vun.unLocalAddr, ___Vun.unAddr)
		}
		//_Fsleep_1s()
	}
}
