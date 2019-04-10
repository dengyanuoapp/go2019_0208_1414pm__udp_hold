package main

func (___Vun *_TudpNodeSt) _FudpNode__500101z__send__default() {
	for {
		select {
		case __VchSend := <-___Vun.unCHsendI: // _TudpNodeDataSendX
			//_FpfN(" 839118 01 send ") // usToAddr

			__VchSend._FudpNode__500101zz__try_Rand_buf_before_send()
			___Vun._FudpNode__500101zzz__send_buf_real(__VchSend)
		}
		//_Fsleep_1s()
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
	//		_FpfN(" 839119 01 : why nil ?")
	//		return
	//	}
	__VmyUs := ___Vus
	if len(__VmyUs.usOutBuf) < _VdataPackageMinLen {
		_FpfN(" 839119 02 : why buf NIL ?")
		return
	}

	__VtraceInt := _FgenRand_int()

	var __Verr2 error
	__VkLen := len(__VmyUs.usToAddr.K256) //
	//__VkLen = 0 // force disable the rand aes
	if 0 != __VkLen {
		if 32 == __VkLen {
			//_FpfN(" 839119 03 key %x %x", ___Vun.unRkeyX.B32, ___Vun.unRkeyX.Bkey)
			//copy(__VmyUs.usOutBuf[_VdataPackageKeyStart:], __VmyUs.usToAddr.K256)
			_CpfN(" 839119 04 : tN:%d before me:%d <%v> %x", __VtraceInt, ___Vun.unLocalPort, __VmyUs.usToAddr.String(), __VmyUs.usOutBuf)
			copy(__VmyUs.usOutBuf[_VdataPackageKeyStart:], ___Vun.unRkeyX.Bkey)
			_CpfN(" 839119 05 : tN:%d after me:%d <%v> %x", __VtraceInt, ___Vun.unLocalPort, __VmyUs.usToAddr.String(), __VmyUs.usOutBuf)
			_CpfN(" 839119 06 : tN:%d udpNode: {%#v}", __VtraceInt, *___Vun)
			__Vbuf, __Verr := _FencAesRand__only(&__VmyUs.usToAddr.K256, &__VmyUs.usOutBuf, __VtraceInt)
			if nil != __Verr {
				_FpfN(" 839119 07 : why error ? %v", __Verr)
				return
			}
			_, __Verr2 = ___Vun.unConn.WriteToUDP(__Vbuf, &__VmyUs.usToAddr.DstAddr)
			if 3 == 3 {
				if 2 == 2 {
					_CpfN(
						" 839119 08 udp-send-rand %11d me<:%d> mreceK:<%s> to<%s> toK<%s> before:%d{%x}, rand:%d{%x} [%s]",
						_FtimeI64(),
						___Vun.unLocalPort,
						String5(&___Vun.unRkeyX.Bkey),
						__VmyUs.usToAddr.DstAddr.String(),
						String5(&__VmyUs.usToAddr.K256),
						len(__VmyUs.usOutBuf),
						_FgenMd5__5(&__VmyUs.usOutBuf),
						len(__Vbuf),
						_FgenMd5__5(&__Vbuf),
						__VmyUs.String(),
					)
				}
				if 2 == 3 {
					_FpfNhex(&__Vbuf, 28,
						" 839119 09 udp send-rand %11d len:%d: dst<%s> to-key<%x> local<%s> md5{%x}",
						_FtimeI64(), len(__VmyUs.usOutBuf),
						__VmyUs.usToAddr.DstAddr.String(),
						__VmyUs.usToAddr.K256[:8],
						___Vun.unLocalAddr.String(),
						_FgenMd5__5(&__VmyUs.usOutBuf))
				}
				if 2 == 3 {
					_FpfNhex(&__Vbuf, 16, " 839121 07 send: %s", ___Vun.String())
				}
			}
			if 3 == 2 {
				_FpfN(" 839121 01 uSend %s", __VmyUs.String())
			}
		} else {
			_FpfN(" 839121 06 : why key len error (%d) ?", __VkLen)
			return
		}
	} else {
		// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
		_, __Verr2 = ___Vun.unConn.WriteToUDP(__VmyUs.usOutBuf, &__VmyUs.usToAddr.DstAddr)
		_FpfNhex(&__VmyUs.usOutBuf, 40, " 839121 07 : %s : udp send direct :", ___Vun.unName)
	}

	if __Verr2 != nil {
		_FpfN(" 839121 08 : udp send error <%s>[%v]", __VmyUs.usToAddr.DstAddr, __Verr2)
		return
	}

	if 3 == 2 {
		_Ppf("839121 11 dst<%v>, local<%v>, listen<%v>\n", __VmyUs.usToAddr.DstAddr, ___Vun.unLocalAddr, ___Vun.unAddr)
	}
}
