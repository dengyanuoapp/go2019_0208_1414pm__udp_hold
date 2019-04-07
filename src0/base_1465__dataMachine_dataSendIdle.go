package main

func _FdataMachin__1000502x__time_gap_dataSendIdle(___Vdm *_TdataMachine) {

	for {
		_FsleepRand_10_to_15s()
		//dmmLastAccTime
		__VkDelArr := [][16]byte{}
		__Vnow2 := _FtimeInt()

		___Vdm.dmMconn.mux.Lock()
		for __Vk2, __Vv2 := range ___Vdm.dmMconn.M {
			if __Vnow2-__Vv2.dmmLastAccTime > 57 {
				__VkDelArr = append(__VkDelArr, __Vk2)
			} else {
				_FpfN(" 381921 01 : try send idle %x ", __Vk2)
			}
		}

		for _, __Vk3 := range __VkDelArr {
			_FpfN(" 381921 05 : try delete lost connect %x in dataMachine", __Vk3)

			if nil == ___Vdm.dmCHloginGenMachineIdLO {
				_FpfNonce(" 381921 07 : try to stop lost connect %x in loginGen , but NULL out-chain", __Vk3)
			} else {
				_FpfNonce(" 381921 08 : try to stop lost connect %x in loginGen, ok ", __Vk3)
				(*___Vdm.dmCHloginGenMachineIdLO) <- _TdataMachinEid{
					diIdx128: ___Vdm.dmMconn.M[__Vk3].dmmID.diIdx128,
				}
			}

			delete(___Vdm.dmMconn.M, __Vk3)
			delete(___Vdm.dmMconn.LockLast, __Vk3)
			delete(___Vdm.dmMconn.LockNow, __Vk3)
		}
		___Vdm.dmMconn.mux.Unlock()
	}
}
