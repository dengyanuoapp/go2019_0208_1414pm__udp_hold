package main

import "sync"

var ___Vout__dmCHloginGenMachineIdLO__mux sync.Mutex

// _FdataMachin__1000201x11__rece_machineId_check_and_insert
func (___Vdm *_TdataMachine) _FdataMachin__1000501y__swapLoginCkInfoForLock__swap() {

	defer ___Vdm.dmMconn.dcsMux.Unlock() // _TdataMachinEconnectSt
	___Vdm.dmMconn.dcsMux.Lock()         // _TdataMachinEconnectClient

	if 0 == len(___Vdm.dmMconn.dcsMidx) {
		_CpfN(" 839196 00 zero-len conn-Arr , nothing need to swap.")
		return
	}

	var __Vcnt int // _TdataMachinEconnectSt
	for __Vkey2, __Vidx2 := range ___Vdm.dmMconn.dcsMidx {
		__Vc := &(___Vdm.dmMconn.dcsMm[__Vidx2]) // _TdataMachinEconnectClient
		__Vcnt = __Vc.dccLockCntLast + __Vc.dccLockCntNow

		if __Vcnt >= 3 { // one connect must have the amount of connect-portS large then 2
			if nil == ___Vdm.dmCHloginGenMachineIdLO {
				_FpfNonce(" 839196 01 connected succeed, but why loginGen NULL ?")
				_CpfN(" 839196 02 connected succeed, but why loginGen NULL ?")
			} else {
				//_FpfN(" 839196 03 connected succeed, then told loginGen to stop ( push chan)")
				_CpfN(" 839196 04 connected succeed, then told loginGen to stop ( push chan)")
				___Vout__dmCHloginGenMachineIdLO__mux.Lock()
				(*___Vdm.dmCHloginGenMachineIdLO) <- ___Vdm.dmMconn.dcsMm[__Vidx2].dccID
				___Vout__dmCHloginGenMachineIdLO__mux.Unlock()
			}
			__Vc.dccLockCntLast = 0 // clear the counter , forbit being next used
			__Vc.dccLockCntNow = 0  // clear the counter , forbit being next used

			___Vdm.
				_FdataMachin__1000501z__swapLoginCkInfoForLock__createNewDataClient(&__Vkey2, __Vidx2)
		} else {
			__Vc.dccLockCntLast = __Vc.dccLockCntNow
			__Vc.dccLockCntNow = 0
			//_FpfN(" 839196 07 no-enough, no swap. " )
			_CpfN(" 839196 08 no-enough, no swap. ")
		}
	}
	_CpfN(" 839196 09 after swap {%s}", ___Vdm.dmMconn.String())
}

// _FdataMachin__1000201x11__rece_machineId_check_and_insert
func (___Vdm *_TdataMachine) _FdataMachin__1000501z__swapLoginCkInfoForLock__createNewDataClient(___Vkey *[16]byte, ___VcIdx int) {
	// ___Vdm.dmMconn.dcsMm[__Vidx2].
	_, __VdOk4 := ___Vdm.dmMdata.ddsMidx[*___Vkey] // _TdataMachinEdataSt _TdataMachinEdataClient
	if __VdOk4 {
		_FpfN(" 183818 01 : re-create client of {%s} ", ___Vdm.dmMconn.dcsMm[___VcIdx].dccID.String()) // _TdataMachinEdataSt _TdataMachinEdataClient
		_FpfN(" 183818 02 : under-constructing")
	} else {
		_FpfN(" 183818 03 : create new client of {%s} ", ___Vdm.dmMconn.dcsMm[___VcIdx].dccID.String()) // _TdataMachinEdataSt _TdataMachinEdataClient
		_FpfN(" 183818 04 : under-constructing")
	}
}
