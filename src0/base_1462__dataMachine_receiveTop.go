package main

import "sync"

var ___V_FdataMachin__1000201x__receive__default__mux sync.Mutex

func _FdataMachin__1000201x__receive__default(___Vdm *_TdataMachine) {
	//var __VdmID _TdataMachinEid
	for {
		select {
		case __VdmID := <-___Vdm.dmCHdataMachineIdI: // chan _TdataMachinEid

			___V_FdataMachin__1000201x__receive__default__mux.Lock()

			_FpfNonce(" 839192 01 : reset-MachineID : _TdataMachine receive data {%s}", __VdmID.String())

			___Vdm.
				_FdataMachin__1000201x11__checkConnMap_insertId(&__VdmID)

		case __Vdec := <-___Vdm.dmCHdecodeDataI: // from uDeCHdecodeDataLO  // _TdecodeX

			___V_FdataMachin__1000201x__receive__default__mux.Lock()

			_FpfNonce(" 839192 03 : reset-MachineID : _TdataMachine receive data {%s}", __Vdec.String())

			___Vdm.
				_FdataMachin__1000201x21__rece_encodeData(&__Vdec)
		}
		___V_FdataMachin__1000201x__receive__default__mux.Unlock()
	}
}
