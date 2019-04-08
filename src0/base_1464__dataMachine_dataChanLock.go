package main

func _FdataMachin__1000501x__time_gap_dataChanLock(___Vdm *_TdataMachine) {
	var __Vcnt int

	for {
		_Fsleep(_T5s)
		___Vdm.dmMconn.mux.Lock()

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
					(*___Vdm.dmCHloginGenMachineIdLO) <- ___Vdm.dmMconn.M[__Vk2].dmmID
				}

			}
		}

		___Vdm.dmMconn.LockLast = ___Vdm.dmMconn.LockNow
		___Vdm.dmMconn.LockNow = make(map[[16]byte]int)

		___Vdm.dmMconn.mux.Unlock()
	}
}
