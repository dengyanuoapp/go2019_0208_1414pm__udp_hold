package main

func _FdataMachin__1000201x__receive__default(___Vdm *_TdataMachine) {
	var __VdmID _TdataMachinEid
	for {
		select {
		case __VdmID = <-___Vdm.dmCHdataMachineIdI: // chan _TdataMachinEid
			//_FpfNdb(" 839192 01 : reset-MachineID {%s}", __VdmID.String())
			___Vdm.
				_FdataMachin__1000201x11__connMap_insertId(&__VdmID)

		}
	}
}
