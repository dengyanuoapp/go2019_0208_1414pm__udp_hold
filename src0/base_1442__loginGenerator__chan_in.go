package main

import "sync"

var ___VudpDecode__800201x__mux sync.Mutex

func _Flogin__800201x__chan_in__default(___Vlg *_TloginGenerator) {
	for {
		select {
		case __Vdmid := <-___Vlg.lgCHdataMachineIdI: // _TdataMachinEid
			//_FpfN(" 388192 01 ")
			//_FpfN(" 388192 02 : logGen received {%s}", __Vdmid.String())
			___VudpDecode__800201x__mux.Lock()
			__Vlen2 := len(__Vdmid.diSeq128)
			switch __Vlen2 {
			case 0:
				_FpfN("388192 03 logGen lost {%s} t:%11d ", __Vdmid.String(), _FtimeInt())
				_CpfN("388192 04 logGen lost {%s} t:%11d ", __Vdmid.String(), _FtimeInt())
				___Vlg.lgConnected = false
			case 16:
				_FpfNt("388192 05 logGen ok {%s} t:%11d ", __Vdmid.String(), _FtimeInt())
				_CpfN("388192 06 logGen ok {%s} t:%11d ", __Vdmid.String(), _FtimeInt())
				___Vlg.lgConnected = true
			default:
				_FpfN(" 388192 07 : why len %d ? t:%11d ", __Vlen2)
				_CpfN(" 388192 08 : why len %d ? t:%11d ", __Vlen2)
			}
		}
		___VudpDecode__800201x__mux.Unlock()
	}
}
