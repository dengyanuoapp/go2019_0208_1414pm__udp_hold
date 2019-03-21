package main

func (___Vun *_TudpNodeSt) _FudpNode__540201z__send() {
	for {
		select {
		case __VchSend := <-___Vun.unCHsendX: // _TudpNodeDataSend
			_FpfN(" 839118 01 send ")

			//usToAddr _TudpConnPort
			// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
			_, __Verr2 := ___Vun.unConn.WriteToUDP(__VchSend.usOutBuf, &__VchSend.usToAddr.DstAddr)
			if __Verr2 != nil {
				_FpfN(" 839118 06 : udp send error <%s>[%v]", __VchSend.usToAddr.DstAddr, __Verr2)
				return
			}
			_FpfN(" 839118 07 : %s : udp send succeed (%d) ", ___Vun.unName, len(__VchSend.usOutBuf))
			_Ppf("01 dst<%v>, local<%v>, listen<%v>", __VchSend.usToAddr.DstAddr, ___Vun.unLocalAddr, ___Vun.unAddr)
		}
		//_Fsleep_1s()
	}
}
