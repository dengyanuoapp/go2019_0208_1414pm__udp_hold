package main

import "sync"

var ___V_FdataMachin__1000201x__receive__default__mux sync.Mutex

func _FdataMachin__1000201x__receive__default(___Vdm *_TdataMachine) {
	//var __VdmID _TdataMachinEid
	for {
		select {
		case __VdmID := <-___Vdm.dmCHdataMachineIdI: // chan _TdataMachinEid

			___V_FdataMachin__1000201x__receive__default__mux.Lock()

			//_CFpfN(" 839192 02 : _TdataMachine receive dataMachinEid {%s}", __VdmID.String())

			___Vdm.
				_FdataMachin__1000201x11__rece_machineId_check_and_insert(&__VdmID)

		case __Vdec := <-___Vdm.dmCHdecodeDataI: // from uDeCHdecodeDataLO  // _TdecodeX

			___V_FdataMachin__1000201x__receive__default__mux.Lock()

			//_CFpfN(" 839192 04 : _TdataMachine receive _Tdecode {%s}", __Vdec.String())

			___Vdm.
				_FdataMachin__1000201x21__rece_encodeData(&__Vdec)

		case <-___Vdm.dmChSendIdleNoteInternalUSE:
			___V_FdataMachin__1000201x__receive__default__mux.Lock()

			//_CFpfN(" 839192 06 : _TdataMachine receive idle timeOut note . trying to send idle.")

			___Vdm.
				_FdataMachin__1000502y__dataSendIdle__packAndSendAll()

		case <-___Vdm.dmChSwapLoginCkInfoForLock: // chan byte // a 5s timer , send swap note to main receive loop. internal use only.
			___V_FdataMachin__1000201x__receive__default__mux.Lock()

			//_CFpfN(" 839192 08 : reset-MachineID : _TdataMachine receive swap loginCKinfo note . ")

			___Vdm.
				_FdataMachin__1000501y__swapLoginCkInfoForLock__swap()
		case <-___Vdm.dmChCheckTimeOutDieClient:
			___V_FdataMachin__1000201x__receive__default__mux.Lock()

			//_CFpfN(" 839192 09 : reset-MachineID : _TdataMachine receive swap loginCKinfo note . ")

			___Vdm.
				_FdataMachin__1000501y__clean_timeoutObj()
		case <-___Vdm.dmCHdebugInfoI:
			___V_FdataMachin__1000201x__receive__default__mux.Lock()

			___Vdm.
				_FdataMachin__1000508__printDebugInfo()
		}

		___V_FdataMachin__1000201x__receive__default__mux.Unlock()
	}
}

func (___Vdm *_TdataMachine) _FdataMachin__1000508__printDebugInfo() {
	_CFpfN(" 348181 01 : tcp rece ")
}