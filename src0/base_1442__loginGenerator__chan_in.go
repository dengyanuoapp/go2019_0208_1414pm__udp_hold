package main

func _FudpDecode__800201x__chan_in__default(___Vlg *_TloginGenerator) {
	for {
		select {
		case __Vdmid := <-___Vlg.lgCHdataMachineIdI: // _TdataMachinEid
			//_FpfN(" 388192 01 ")
			_FpfN(" 388192 02 : logGen received {%s}", __Vdmid.String())
			__Vlen2 := len(__Vdmid.diIdx128)
			switch __Vlen2 {
			case 0:
				___Vlg.lgConnected = false
			case 16:
				___Vlg.lgConnected = true
			default:
				_FpfN(" 388192 05 : why len %d ?", __Vlen2)
			}
		}
	}
}
