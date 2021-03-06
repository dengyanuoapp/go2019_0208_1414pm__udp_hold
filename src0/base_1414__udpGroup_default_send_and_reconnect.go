package main

import "sync"

var __VudpGroup__600201__mux sync.Mutex

// _TudpGroupSt
// _TudpNodeSt
func _FudpGroup__600201__CHin_selecT_send__default(___Vug *_TudpGroupSt) {
	//_Fsleep(_T1s)
	//_FpfNdb(" 838111 01 len : %d", __Vlen)
	var __Vunds _TudpNodeDataSend
	var __Vidx int
	for {
		select {
		case __Vunds = <-___Vug.ugCHSendI: // _TudpNodeDataSendX
			__VudpGroup__600201__mux.Lock()

		}

		__Vidx = ___Vug.
			_FudpGroup__600201www__send_genIdx()

		___Vug.
			_FudpGroup__600201zzz__send_byteOnly(&__Vunds, __Vidx)

		__VudpGroup__600201__mux.Unlock()
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
	if 0 >= __Vlen {
		_FpfNex(" 838117 02 : why NULL ?")
		return 0
	}
	if 1 == __Vlen {
		_FpfN(" 838117 03 : only one send-dst port. ")
		_CpfN(" 838117 04 : only one send-dst port. ")
		___Vug.ugLastSendIdx = 0
		return 0
	}

	__Vidx := 0
	//_CpfN(" 838117 05 : idx 0 now. len %d", __Vlen)

	for {
		__Vidx = _FgenRand_int()
		//_FpfN(" 838117 04 : idx %0x", __Vidx)

		__Vidx = __Vidx % __Vlen
		//_FpfN(" 838117 05 : idx %0x , %d ", __Vidx, __Vidx)

		if ___Vug.ugLastSendIdx != __Vidx {
			___Vug.ugLastSendIdx = __Vidx
			return __Vidx
		}
	}
	//return 0
}
