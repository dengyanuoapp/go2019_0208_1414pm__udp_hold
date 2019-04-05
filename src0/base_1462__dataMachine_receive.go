package main

func _FdataMachin__1000201x__receive__default(___Vdm *_TdataMachine) {
	var __VdmID _TdataMachinEid
	for {
		select {
		case __VdmID = <-___Vdm.dmCHdataMachineIdI: // chan _TdataMachinEid
			_FpfNdb(" 839192 01 : reset-MachineID {%s}", __VdmID.String())

			//__Vk := _FgenB16(&__VdmID.diIdx128)
			//___Vdm.dmMconn
		}
	}
}
