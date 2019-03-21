package main

// _TudpGroupSt
// _TudpNodeSt
// _TudpConnPort
func _FudpGroup__650301__connPort__default(___Vug *_TudpGroupSt) {
	//_Fsleep(_T1s)
	//__Vlen := len(___Vug.ugCHuConnPortX)
	//_FpfNdb(" 838111 01 len : %d", __Vlen)
	for {
		select {
		case __VchPort := <-___Vug.ugCHuConnPortX:
			//_FpfNdb(" 838111 03 : %s, %0x", __VchPort.DstAddr, __VchPort.K256)
			_FpfNdb(" 838111 04 : %s", __VchPort.DstAddr)
			___Vug.
				_FudpGroup__650301z__send_using_rand_port(__VchPort.
					_FdataPack__101__udpConnPort())
		}
	}
}

//	ugRANDidx     []int
//	ugRANDremain  int
func (___Vug *_TudpGroupSt) _FudpGroup__650301zzz__send_byteOnly(___VbIn *[]byte) {
	__Vlen := len(___Vug.ugCHtmpSendLX)
	if 0 == __Vlen {
		_FpfNex(" 838117 01 : why NULL ?")
		return
	}
	__Vidx := 0
	if 1 == __Vlen {
		for {
			__Vidx = int(_FgenRand_int32())
			__Vidx = __Vidx % __Vlen
			if ___Vug.ugLastSendIdx != __Vidx {
				___Vug.ugLastSendIdx = __Vidx
				break
			}
		}
	}
	__Vch := ___Vug.ugCHtmpSendLX[__Vidx]
	if nil == __Vch {
		_FpfNex(" 838117 03 : why NULL ?")
		return
	}

	__VuData := _TudpNodeDataSend{
		//unDstAddr : ,// net.UDPAddr
		unOutBuf: (*___VbIn),
	}

	(*__Vch) <- __VuData
}

func (___Vug *_TudpGroupSt) _FudpGroup__650301z__send_using_rand_port(___VbIn *[]byte) {
	if nil == ___VbIn {
		_FpfNex(" 838112 01 : why NULL ?")
		return
	}
	___Vug._FudpGroup__650301zzz__send_byteOnly(___VbIn)
}
