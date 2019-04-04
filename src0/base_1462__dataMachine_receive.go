package main

func _FdataMachin__1000201x__receive__default(___Vdm *_TdataMachine) {
	var __VchDataMachineID _TdataMachinEid
	for {
		select {
		case __VchDataMachineID = <-___Vdm.dmCHdataMachineIdI: // chan _TdataMachinEid
			_FpfN(" 839192 01 : rece {%s}", __VchDataMachineID.String())
		}
	}
}
