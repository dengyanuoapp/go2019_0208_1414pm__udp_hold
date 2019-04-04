package main

// _TudpGroupSt
// _TudpNodeSt
func _FudpGroup__600201__CHin_select_send__default(___Vug *_TudpGroupSt) {
	//_Fsleep(_T1s)
	//__Vlen := len(___Vug. ugCHuConnPortI)
	//_FpfNdb(" 838111 01 len : %d", __Vlen)
	var (
		__VusData _TudpNodeDataSend
		__Vidx    int
	)
	for {
		select {
		//		case __VchPort := <-___Vug.ugCHuConnPortI: //
		//			//_FpfNdb(" 838111 03 : %s, %0x", __VchPort.DstAddr, __VchPort.K256)
		//			//_FpfNdb(" 838111 04 : %v", __VchPort.DstAddr)
		//			__VusData.usToAddr = __VchPort
		//			//__VusData .  usOutBuf= *(__VchPort._FdataPack__101__udpConnPort()),
		//			__VchPort._FdataPack__101__udpConnPort(&__VusData.usOutBuf)
		//			__Vidx = ___Vug._FudpGroup__600201www__send_genIdx()
		//			//__Vidx = 0
		//
		//		}
		case __VusData = <-___Vug.ugCHSendI: // _TudpNodeDataSend
			__Vidx = ___Vug._FudpGroup__600201www__send_genIdx()
		}
		___Vug._FudpGroup__600201zzz__send_byteOnly(&__VusData, __Vidx)
	}
}

//	ugRANDidx     []int
//	ugRANDremain  int
var ___Vug_ugLastSendIdx_Arr []int

func (___Vug *_TudpGroupSt) _FudpGroup__600201zzz__send_byteOnly(___Vs *_TudpNodeDataSend, ___Vidx int) {
	if (nil == ___Vs) || (0 == len(___Vs.usOutBuf)) {
		_FpfN(" 838117 01 : why NULL ?")
		return
	}

	if len(___Vug_ugLastSendIdx_Arr) > 25 {
		___Vug_ugLastSendIdx_Arr = append(___Vug_ugLastSendIdx_Arr[2:], ___Vidx)
	} else {
		___Vug_ugLastSendIdx_Arr = append(___Vug_ugLastSendIdx_Arr, ___Vidx)
	}
	//_FpfN(" 838117 06 : idx %0x , %d , %v", ___Vidx, ___Vidx, ___Vug_ugLastSendIdx_Arr)

	__Vch := ___Vug.ugChTmpSendLI[___Vidx]
	if nil == __Vch {
		_FpfNex(" 838117 07 : why NULL ?")
		return
	}

	(*__Vch) <- (*___Vs) // send plainText to the single node : the rand-gen being process in specified-node
}

func (___Vug *_TudpGroupSt) _FudpGroup__600201www__send_genIdx() int {
	__Vlen := len(___Vug.ugChTmpSendLI)
	if 0 == __Vlen {
		_FpfNex(" 838117 02 : why NULL ?")
		return 0
	}

	__Vidx := 0
	//_FpfN(" 838117 03 : idx 0 now. len %d", __Vlen)

	if 1 != __Vlen {
		for {
			__Vidx = _FgenRand_int()
			//_FpfN(" 838117 04 : idx %0x", __Vidx)

			__Vidx = __Vidx % __Vlen
			//_FpfN(" 838117 05 : idx %0x , %d ", __Vidx, __Vidx)

			if ___Vug.ugLastSendIdx != __Vidx {
				___Vug.ugLastSendIdx = __Vidx
				break
			}
		}
	}
	return __Vidx
}
