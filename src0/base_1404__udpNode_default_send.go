package main

func (___Vun *_TudpNodeSt) _FudpNode__540201z__send() {
	for {
		select {
		case __VchSend := <-___Vun.unCHsendX: // _TudpNodeDataSend
			//_FpfN(" 839118 01 send ") // usToAddr _TudpConnPort

			__VchSend._FudpNode__540201zz__try_Rand_buf_before_send()
			___Vun._FudpNode__540201zzz__send_buf_real(&__VchSend)
		}
		//_Fsleep_1s()
	}
}

func (___Vus *_TudpNodeDataSend) _FudpNode__540201zz__try_Rand_buf_before_send() {
	__Vlen := len(___Vus.usToAddr.K256)
	if 0 == __Vlen {
		return
	}

	if 32 != __Vlen {
		_FpfNex(" 839117 01 key Len error %d ", __Vlen)
	}

}

func (___Vun *_TudpNodeSt) _FudpNode__540201zzz__send_buf_real(___Vus *_TudpNodeDataSend) {
	if nil == ___Vus {
		_FpfN(" 839119 01 : why nil ?")
		return
	}
	if len(___Vus.usOutBuf) < _VdataPackageMinLen {
		_FpfN(" 839119 02 : why buf NIL ?")
		return
	}

	var __Verr2 error
	__VkLen := len(___Vus.usToAddr.K256) //_TudpConnPort
	//__VkLen = 0 // force disable the rand aes
	if 0 != __VkLen {
		if 32 == __VkLen {
			//_FpfN(" 839119 03 key %x %x", ___Vun.unRkeyX.B32, ___Vun.unRkeyX.Bkey)
			//copy(___Vus.usOutBuf[_VdataPackageKeyStart:], ___Vus.usToAddr.K256)
			copy(___Vus.usOutBuf[_VdataPackageKeyStart:], ___Vun.unRkeyX.Bkey)
			__Vbuf, __Verr := _FencAesRand__only(&___Vus.usToAddr.K256, &___Vus.usOutBuf)
			if nil != __Verr {
				_FpfN(" 839119 03 : why error ? %v", __Verr)
				return
			}
			_, __Verr2 = ___Vun.unConn.WriteToUDP(__Vbuf, &___Vus.usToAddr.DstAddr)
			_FpfNhex(&__Vbuf, 6, " 839119 04 : %s : udp send rand %11d:%d: dst<%v>, local<%v>, listen<%v>",
				___Vun.unName, _FtimeI64(), len(___Vus.usOutBuf), ___Vus.usToAddr.DstAddr, ___Vun.unLocalAddr, ___Vun.unAddr)
		} else {
			_FpfN(" 839119 05 : why key len error (%d) ?", __VkLen)
			return
		}
	} else {
		//usToAddr _TudpConnPort
		// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
		_, __Verr2 = ___Vun.unConn.WriteToUDP(___Vus.usOutBuf, &___Vus.usToAddr.DstAddr)
		_FpfNhex(&___Vus.usOutBuf, 40, " 839119 06 : %s : udp send direct :", ___Vun.unName)
	}

	if __Verr2 != nil {
		_FpfN(" 839119 08 : udp send error <%s>[%v]", ___Vus.usToAddr.DstAddr, __Verr2)
		return
	}

	if 3 == 2 {
		_Ppf("839119 11 dst<%v>, local<%v>, listen<%v>\n", ___Vus.usToAddr.DstAddr, ___Vun.unLocalAddr, ___Vun.unAddr)
	}
}
