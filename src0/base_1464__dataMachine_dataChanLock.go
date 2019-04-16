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

func (___Vdm *_TdataMachine) _FdataMachin__1000501y__swapLoginCkInfoForLock__swap() {
	defer ___Vdm.dmMconn.mux.Unlock()
	___Vdm.dmMconn.mux.Lock()

	var __Vcnt int
	for __Vk2, __Vv2 := range ___Vdm.dmMconn.LockNow {
		__Vv3, __Vok3 := ___Vdm.dmMconn.LockLast[__Vk2]
		if __Vok3 {
			__Vcnt = __Vv2 + __Vv3
		} else {
			__Vcnt = __Vv2
		}
		if __Vcnt >= 3 { // one connect must have the amount of connect-portS large then 2
			//_FpfN(" 839196 03 connected succeed : %d : %d", __Vcnt, _FtimeInt())

			if nil != ___Vdm.dmCBprReceKey {
				___Vdm.dmCBprReceKey(___Vdm)
				//___Vlg.FprKey()
			}

			if nil == ___Vdm.dmCHloginGenMachineIdLO {
				_FpfNonce(" 839196 08 connected succeed, but why loginGen NULL ?")
			} else {
				//_FpfN(" 839196 09 connected succeed, then told loginGen to stop ( push chan)")
				___Vout__dmCHloginGenMachineIdLO__mux.Lock()
				(*___Vdm.dmCHloginGenMachineIdLO) <- ___Vdm.dmMconn.M[__Vk2].dmmID
				___Vout__dmCHloginGenMachineIdLO__mux.Unlock()
			}

		}
	}

	___Vdm.dmMconn.LockLast = ___Vdm.dmMconn.LockNow
	___Vdm.dmMconn.LockNow = make(map[[16]byte]int)

}
