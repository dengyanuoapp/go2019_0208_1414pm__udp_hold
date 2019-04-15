package main

import "sync"

var ___V_FdataMachin__1000201x__receive__default__mux sync.Mutex

func _FdataMachin__1000201x__receive__default(___Vdm *_TdataMachine) {
	//var __VdmID _TdataMachinEid
	for {
		select {
		case __VdmID := <-___Vdm.dmCHdataMachineIdI: // chan _TdataMachinEid
			//_FpfNdb(" 839192 01 : reset-MachineID {%s}", __VdmID.String())
			___V_FdataMachin__1000201x__receive__default__mux.Lock()
			___Vdm.
				_FdataMachin__1000201x11__connMap_insertId(&__VdmID)
		case __Vdec := <-___Vdm.dmCHdecodeDataI: // from uDeCHdecodeDataLO  // _TdecodeX
			___V_FdataMachin__1000201x__receive__default__mux.Lock()
			_FpfN(" 839192 03 : _TdataMachine receive {%s}", __Vdec.String())
		}
		___V_FdataMachin__1000201x__receive__default__mux.Unlock()
	}
}
