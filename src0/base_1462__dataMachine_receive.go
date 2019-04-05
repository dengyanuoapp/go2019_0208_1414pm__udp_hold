package main

func _FdataMachin__1000201x__receive__default(___Vdm *_TdataMachine) {
	var __VchDataMachineID _TdataMachinEid
	for {
		select {
		case __VchDataMachineID = <-___Vdm.dmCHdataMachineIdI: // chan _TdataMachinEid
			_FpfNdb(" 839192 01 : reset-MachineID {%s}", __VchDataMachineID.String())
		}
	}
}
