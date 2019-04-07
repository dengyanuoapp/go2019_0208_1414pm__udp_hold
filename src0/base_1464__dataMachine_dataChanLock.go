package main

func _FdataMachin__1000501x__time_gap_dataChanLock(___Vdm *_TdataMachine) {
	var __Vcnt int

	for {
		_Fsleep(_T5s)
		___Vdm.dmMlock.mux.Lock()

		for __Vk2, __Vv2 := range ___Vdm.dmMlock.LockNow {
			__Vv3, __Vok3 := ___Vdm.dmMlock.LockLast[__Vk2]
			if __Vok3 {
				__Vcnt = __Vv2 + __Vv3
			} else {
				__Vcnt = __Vv2
			}
			if __Vcnt >= 3 { // one connect must have the amount of connect-portS large then 2
				_FpfN(" 839196 03 connected succeed : %d : %d", __Vcnt, _FtimeInt())
			}
		}

		___Vdm.dmMlock.LockLast = ___Vdm.dmMlock.LockNow
		___Vdm.dmMlock.LockNow = make(map[[16]byte]int)

		___Vdm.dmMlock.mux.Unlock()
	}
}
