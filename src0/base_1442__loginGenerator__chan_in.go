package main

func _FudpDecode__800201x__chan_in__default(___Vlg *_TloginGenerator) {
	for {
		select {
		case __Vds := <-___Vlg.lgCHdataMachineIdI: // _TdataMachinEid
			_FpfN(" 388192 01 ")
			_FpfN(" 388192 02 : logGen received {%s}", __Vds.String())
		}
	}
}
