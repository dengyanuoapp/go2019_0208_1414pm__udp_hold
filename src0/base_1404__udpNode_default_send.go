package main

func (___Vun *_TudpNodeSt) _FudpNode__140201z__send() {
	for {
		select {
		case __VchSend := <-___Vun.unCHsend: // _TudpNodeDataSend
			_FpfN(" 839118 01 send ")

			// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
			_, __Verr2 := ___Vun.unConn.WriteToUDP(__VchSend.unOutBuf, &__VchSend.unDstAddr)
			if __Verr2 != nil {
				_FpfN(" 839118 06 : udp send error <%s>[%v]", __VchSend.unDstAddr, __Verr2)
				return
			}
			_FpfN(" 839118 07 : %s : udp send succeed (%d) 01 dst<%s>, local<%v>, listen<%v>", ___Vun.name,
				len(__VchSend.unOutBuf), __VchSend.unDstAddr, ___Vun.unLocalAddr, ___Vun.unAddr)
		}
		//_Fsleep_1s()
	}
}