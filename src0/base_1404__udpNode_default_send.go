package main

func (___Vun *_TudpNodeSt) _FudpNode__500101z__send() {
	for {
		select {
		case __VchSend := <-___Vun.unCHsendI: // _TudpNodeDataSendX
			//_FpfN(" 839118 01 send ") // usToAddr

			__VchSend._FudpNode__500101zz__try_Rand_buf_before_send()
			___Vun._FudpNode__500101zzz__send_buf_real(&__VchSend)
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

func (___Vun *_TudpNodeSt) _FudpNode__500101zzz__send_buf_real(___Vus *_TudpNodeDataSend) {
	if nil == ___Vus {
		_FpfN(" 839119 01 : why nil ?")
		return
	}
	if len(___Vus.usOutBuf) < _VdataPackageMinLen {
		_FpfN(" 839119 02 : why buf NIL ?")
		return
	}

	var __Verr2 error
	__VkLen := len(___Vus.usToAddr.K256) //
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
			if 3 == 3 {
				if 2 == 2 {
					_CpfN(
						" 839119 04 udp-send-rand %11d me<%s> to<%s> toK<%s> before:%d{%x}, rand:%d{%x} [%s]",
						_FtimeI64(),
						___Vun.unLocalAddr.String(),
						___Vus.usToAddr.DstAddr.String(),
						String5(&___Vus.usToAddr.K256),
						len(___Vus.usOutBuf),
						_FgenMd5__5(&___Vus.usOutBuf),
						len(__Vbuf),
						_FgenMd5__5(&__Vbuf),
						___Vus.String(),
					)
				}
				if 2 == 3 {
					_FpfNhex(&__Vbuf, 28,
						" 839119 06 udp send-rand %11d len:%d: dst<%s> to-key<%x> local<%s> md5{%x}",
						_FtimeI64(), len(___Vus.usOutBuf),
						___Vus.usToAddr.DstAddr.String(),
						___Vus.usToAddr.K256[:8],
						___Vun.unLocalAddr.String(),
						_FgenMd5__5(&___Vus.usOutBuf))
				}
				if 2 == 3 {
					_FpfNhex(&__Vbuf, 16, " 839119 07 send: %s", ___Vun.String())
				}
			}
			if 3 == 2 {
				_FpfN(" 839121 01 uSend %s", ___Vus.String())
			}
		} else {
			_FpfN(" 839121 06 : why key len error (%d) ?", __VkLen)
			return
		}
	} else {
		// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
		_, __Verr2 = ___Vun.unConn.WriteToUDP(___Vus.usOutBuf, &___Vus.usToAddr.DstAddr)
		_FpfNhex(&___Vus.usOutBuf, 40, " 839121 07 : %s : udp send direct :", ___Vun.unName)
	}

	if __Verr2 != nil {
		_FpfN(" 839121 08 : udp send error <%s>[%v]", ___Vus.usToAddr.DstAddr, __Verr2)
		return
	}

	if 3 == 2 {
		_Ppf("839121 11 dst<%v>, local<%v>, listen<%v>\n", ___Vus.usToAddr.DstAddr, ___Vun.unLocalAddr, ___Vun.unAddr)
	}
}
