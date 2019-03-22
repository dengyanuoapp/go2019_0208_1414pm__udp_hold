package main

func (___Vun *_TudpNodeSt) _FudpNode__540201z__send() {
	for {
		select {
		case __VchSend := <-___Vun.unCHsendX: // _TudpNodeDataSend // usToAddr _TudpConnPort
			//_FpfN(" 839118 01 send ")

			___Vun._FudpNode__540201zz__try_Rand_buf_before_send()
			___Vun._FudpNode__540201zzz__send_buf_real()
		}
		//_Fsleep_1s()
	}
}

func (___Vun *_TudpNodeSt) _FudpNode__540201zz__try_Rand_buf_before_send() {
	__Vlen := len(__VchSend.usToAddr.K256)
	if 0 == __Vlen {
		return
	}

	if 32 != __Vlen {
		_FpfN(" 839117 01 key Len error %d ", __Vlen)
	}

}

func (___Vun *_TudpNodeSt) _FudpNode__540201zzz__send_buf_real() {
	//usToAddr _TudpConnPort
	// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
	_, __Verr2 := ___Vun.unConn.WriteToUDP(__VchSend.usOutBuf, &__VchSend.usToAddr.DstAddr)
	if __Verr2 != nil {
		_FpfN(" 839119 06 : udp send error <%s>[%v]", __VchSend.usToAddr.DstAddr, __Verr2)
		return
	}
	if 2 == 2 {
		_Fpf(" 839119 07 : %s : udp send succeed (%d) ", ___Vun.unName, len(__VchSend.usOutBuf))
		_Ppf("01 dst<%v>, local<%v>, listen<%v>\n", __VchSend.usToAddr.DstAddr, ___Vun.unLocalAddr, ___Vun.unAddr)
	}
}
