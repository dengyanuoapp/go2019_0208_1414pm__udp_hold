package main

import "sync"

var ___FudpNode__500101z_mux sync.Mutex

func (___Vun *_TudpNodeSt) _FudpNode__500101z__send__default() {
	for {
		select {
		case __VchSend := <-___Vun.unCHsendI: // _TudpNodeDataSendX
			//_FpfN(" 839118 01 send ") // usToAddr

			___FudpNode__500101z_mux.Lock()

			__VchSend._FudpNode__500101zz__try_Rand_buf_before_send()
			___Vun._FudpNode__500101zzz__send_buf_real(__VchSend)

		}
		//_Fsleep_1s()
		___FudpNode__500101z_mux.Unlock()
	}
}

func (___Vus *_TudpNodeDataSend) _FudpNode__500101zz__try_Rand_buf_before_send() {
	__Vlen := len(___Vus.usToAddr.K256)
	if 0 == __Vlen {
		return
	}

	if 32 != __Vlen {
		_FpfNex(" 839117 01 key Len error %d ", __Vlen)
	}

}

func (___Vun *_TudpNodeSt) _FudpNode__500101zzz__send_buf_real(___Vus _TudpNodeDataSend) { // _TudpNodeDataSendX
	//	if nil == ___Vus {
	//		_FpfN(" 839116 01 : why nil ?")
	//		return
	//	}
	__VmyUs := ___Vus
	if len(__VmyUs.usOutBuf) < _VdataPackageMinLen {
		_FpfN(" 839116 02 : why buf NIL ?")
		return
	}

	//__VtraceInt := ___Vus.Ti
	__VtraceInt := _FgenRand_int()

	var __Verr2 error
	__VkLen := len(__VmyUs.usToAddr.K256) //
	//__VkLen = 0 // force disable the rand aes

	if 0 == __VkLen { // send direct
		// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
		_, __Verr2 = ___Vun.unConn.WriteToUDP(__VmyUs.usOutBuf, &__VmyUs.usToAddr.DstAddr)
		_FpfNhex(&__VmyUs.usOutBuf, 40, " 839116 07 : %s : udp send direct :", ___Vun.unName)
	}

	if 32 != __VkLen {
		_FpfN(" 839116 06 : why key len error (%d) ?", __VkLen)
		return
	}

	//_FpfN(" 839119 01 key %x %x", ___Vun.unRkeyX.B32, ___Vun.unRkeyX.Bkey)
	//copy(__VmyUs.usOutBuf[_VdataPackageKeyStart:], __VmyUs.usToAddr.K256)
	___CpfN(" 839119 03 : tN:%d before copy-key Mp:%d Mrk:%s to<%v> buf[%x]", __VtraceInt, ___Vun.unLocalPort, // keykey
		String5(&___Vun.unRkeyX.Bkey), __VmyUs.usToAddr.String(), __VmyUs.usOutBuf)

	copy(__VmyUs.usOutBuf[_VdataPackageKeyStart:], ___Vun.unRkeyX.Bkey)

	//_CpfN(" 839119 05 : tN:%d after copy-key me:%d <%v> %x", __VtraceInt, ___Vun.unLocalPort, __VmyUs.usToAddr.String(), __VmyUs.usOutBuf) //keykey

	//_CpfN(" 839119 06 : tN:%d udpNode: {%s}", __VtraceInt, ___Vun.String()) // _TudpNodeSt //keykey

	__Vbuf, __Verr := _FencAesRand__only(__VmyUs.usToAddr.K256, __VmyUs.usOutBuf, __VtraceInt)

	if nil != __Verr {
		_FpfN(" 839119 07 : why error ? %v", __Verr)
		return
	}

	// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
	_, __Verr2 = ___Vun.unConn.WriteToUDP(__Vbuf, &__VmyUs.usToAddr.DstAddr)
	if nil != __Verr2 {
		_CpfN(" 839119 08 error [%v] ", __Verr2)
		return
	}

	_NpfN( //keykey
		" 839119 09 udp-send-rand %11d me<:%d> mreceK:<%s> to<%s> toK<%s> before:%d{%x}, rand:%d{%x} [%s]",
		//                          1      2            3     4       5          6   7        8   9   10
		_FtimeI64(),                       // 1
		___Vun.unLocalPort,                // 2
		String5(&___Vun.unRkeyX.Bkey),     // 3
		__VmyUs.usToAddr.DstAddr.String(), // 4
		String5(&__VmyUs.usToAddr.K256),   // 5
		len(__VmyUs.usOutBuf),             // 6
		_Fmd5__5x(&__VmyUs.usOutBuf),      // 7
		len(__Vbuf),                       // 8
		_Fmd5__5x(&__Vbuf),                // 9
		__VmyUs.String(),                  // 10 _TudpNodeDataSendX
	)
}
