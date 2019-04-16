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

		case <-___Vdm.dmChSendIdleNoteInternalUSE:
			___V_FdataMachin__1000201x__receive__default__mux.Lock()
			_FpfNonce(" 839192 05 : reset-MachineID : _TdataMachine receive idle timeOut note . trying to send idle.")
			___Vdm.
				_FdataMachin__1000502y__dataSendIdle__packAndSendAll()
		case <-___Vdm.dmChSwapLoginCkInfoForLock: // chan byte // a 5s timer , send swap note to main receive loop. internal use only.
			___V_FdataMachin__1000201x__receive__default__mux.Lock()
			_FpfNonce(" 839192 07 : reset-MachineID : _TdataMachine receive swap loginCKinfo note . trying to send idle.")
			___Vdm.
				_FdataMachin__1000501y__swapLoginCkInfoForLock__swap()
		}
		___V_FdataMachin__1000201x__receive__default__mux.Unlock()
	}
}
