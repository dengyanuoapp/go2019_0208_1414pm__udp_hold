package main

// _TudpGroupSt
// _TudpNodeSt
// _TudpConnPort
func _FudpGroup__650301__CHin_select_send__default(___Vug *_TudpGroupSt) {
	//_Fsleep(_T1s)
	//__Vlen := len(___Vug.ugCHuConnPortX)
	//_FpfNdb(" 838111 01 len : %d", __Vlen)
	for {
		select {
		case __VchPort := <-___Vug.ugCHuConnPortX: // _TudpConnPort
			//_FpfNdb(" 838111 03 : %s, %0x", __VchPort.DstAddr, __VchPort.K256)
			//_FpfNdb(" 838111 04 : %v", __VchPort.DstAddr)
			__VusData := _TudpNodeDataSend{
				usToAddr: __VchPort,
				usOutBuf: *(__VchPort._FdataPack__101__udpConnPort()),
			}
			___Vug.
				_FudpGroup__650301zzz__send_byteOnly(&__VusData)
		}
	}
}

//	ugRANDidx     []int
//	ugRANDremain  int
var ___Vug_ugLastSendIdx_Arr []int

func (___Vug *_TudpGroupSt) _FudpGroup__650301zzz__send_byteOnly(___Vs *_TudpNodeDataSend) {
	__Vlen := len(___Vug.ugCHtmpSendLX)
	if 0 == __Vlen {
		_FpfNex(" 838117 01 : why NULL ?")
		return
	}

	__Vidx := 0
	//_FpfN(" 838117 02 : idx 0 now. len %d", __Vlen)

	if 1 != __Vlen {
		for {
			__Vidx = _FgenRand_int()
			//_FpfN(" 838117 03 : idx %0x", __Vidx)

			__Vidx = __Vidx % __Vlen
			//_FpfN(" 838117 04 : idx %0x , %d ", __Vidx, __Vidx)

			if ___Vug.ugLastSendIdx != __Vidx {
				___Vug.ugLastSendIdx = __Vidx
				break
			}
		}
	}
	if len(___Vug_ugLastSendIdx_Arr) > 25 {
		___Vug_ugLastSendIdx_Arr = append(___Vug_ugLastSendIdx_Arr[2:], __Vidx)
	} else {
		___Vug_ugLastSendIdx_Arr = append(___Vug_ugLastSendIdx_Arr, __Vidx)
	}
	//_FpfN(" 838117 05 : idx %0x , %d , %v", __Vidx, __Vidx, ___Vug_ugLastSendIdx_Arr)

	__Vch := ___Vug.ugCHtmpSendLX[__Vidx]
	if nil == __Vch {
		_FpfNex(" 838117 03 : why NULL ?")
		return
	}

	(*__Vch) <- (*___Vs)
}
