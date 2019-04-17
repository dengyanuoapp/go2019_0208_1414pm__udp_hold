package main

import "sync"

var ___Vout__dmCHloginGenMachineIdLO__mux sync.Mutex

func (___Vdm *_TdataMachine) _FdataMachin__1000501x__swapLoginCkInfoForLock__timeoutGen() {

	for {
		_Fsleep(_T5s)
		___Vdm.
			dmChSwapLoginCkInfoForLock <- 1 // a 5s timer , send swap note to main receive loop. internal use only.
	}
}

// _FdataMachin__1000201x11__rece_machineId_check_and_insert
func (___Vdm *_TdataMachine) _FdataMachin__1000501y__swapLoginCkInfoForLock__swap() {
	defer ___Vdm.dmMconn.dcsMux.Unlock() // _TdataMachinEconnectSt
	___Vdm.dmMconn.dcsMux.Lock()         // _TdataMachinEconnectClient

	var __Vcnt int // _TdataMachinEconnectSt
	for __Vkey2, __Vidx2 := range ___Vdm.dmMconn.dcsMidx {
		__Vc := ___Vdm.dmMconn.dcsM[__Vidx2] // _TdataMachinEconnectClient
		__Vcnt = __Vc.dccLockCntLast + __Vc.dccLockCntNow

		if __Vcnt >= 3 { // one connect must have the amount of connect-portS large then 2
			if nil == ___Vdm.dmCHloginGenMachineIdLO {
				_FpfNonce(" 839196 08 connected succeed, but why loginGen NULL ?")
			} else {
				//_FpfN(" 839196 09 connected succeed, then told loginGen to stop ( push chan)")
				___Vout__dmCHloginGenMachineIdLO__mux.Lock()
				(*___Vdm.dmCHloginGenMachineIdLO) <- ___Vdm.dmMconn.M[__Vk2].dmmID
				___Vout__dmCHloginGenMachineIdLO__mux.Unlock()
			}
			__Vc.dccLockCntLast = 0 // clear the counter , forbit being next used
			__Vc.dccLockCntNow = 0  // clear the counter , forbit being next used

			__VdIdx4, __VdOk4 := ___Vdm.dmMdata.dcsMidx[__Vkey2] // _TdataMachinEdataSt _TdataMachinEdataClient
		}
	}

	___Vdm.dmMconn.LockLast = ___Vdm.dmMconn.LockNow
	___Vdm.dmMconn.LockNow = make(map[[16]byte]int)

}
