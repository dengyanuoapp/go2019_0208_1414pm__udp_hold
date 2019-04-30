package main

import "sync"

var ___VudpDecode__800201x__mux sync.Mutex

func _Flogin__800201x__chan_in__default(___Vlg *_TloginGenerator) {
	for {
		select {
		case __Vdmid := <-___Vlg.lgCHdataMachineIdI: // _TdataMachinEid
			//_FpfN(" 388192 01 ")
			//_FpfN(" 388192 02 : loginAutoGenerator received {%s}", __Vdmid.String())
			___VudpDecode__800201x__mux.Lock()
			__Vlen2 := len(__Vdmid.diSeq128)
			switch __Vlen2 {
			case 0:
				_CFpfN("388192 04 loginAutoGenerator lost {%s} t:%11d ", __Vdmid.String(), _FtimeInt())
				___Vlg.lgConnected = false
			case 16:
				_CFpfN("388192 06 loginAutoGenerator ok {%s} t:%11d ", __Vdmid.String(), _FtimeInt())
				___Vlg.lgConnected = true
			default:
				_CFpfN(" 388192 08 : why len %d ? t:%11d ", __Vlen2)
			}
		}
		___VudpDecode__800201x__mux.Unlock()
	}
}
