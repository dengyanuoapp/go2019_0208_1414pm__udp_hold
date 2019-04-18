package main

func (___Vdm *_TdataMachine) _FdataMachin__1000501y__clean_timeoutObj() {
	__Vnow2 := _FtimeInt()

	__VkDelArr := [][16]byte{}
	for __Vkey7, __Vidx7 := range ___Vdm.dmMconn.dcsMidx { // _TdataMachinEconnectSt
		if __Vnow2-___Vdm.dmMconn.dcsMm[__Vidx7].dccLastReceTime > _Vgap_nothingToLost {
			__VkDelArr = append(__VkDelArr, __Vkey7)
			___Vdm.dmMconn.dcsMm[__Vidx7].Clear()
		}
	}
	for _, __Vkey8 := range __VkDelArr {
		delete(___Vdm.dmMconn.dcsMidx, __Vkey8)
	}

	__VkDelArr = [][16]byte{}
	for __Vkey9, __Vidx9 := range ___Vdm.dmMdata.ddsMidx { // _TdataMachinEdataSt
		if __Vnow2-___Vdm.dmMdata.ddsMm[__Vidx9].ddcLastReceTime > _Vgap_nothingToLost {
			__VkDelArr = append(__VkDelArr, __Vkey9)

			___Vdm.dmMdata.ddsMm[__Vidx9].Clear()

			if nil == ___Vdm.dmCHloginGenMachineIdLO {
				_FpfNonce(" 381921 07 : %11d : try to stop lost connect %x in loginGen , but NULL out-chain", _FtimeInt(), __Vkey9)
			} else {
				_FpfNonce(" 381921 08 : %11d : try to stop lost connect %x in loginGen, ok ", _FtimeInt(), __Vkey9)
				_CpfN(" 381921 09 : %11d : try to stop lost connect %x in loginGen, ok ", _FtimeInt(), __Vkey9)

				(*___Vdm.dmCHloginGenMachineIdLO) <- _TdataMachinEid{
					diIdx128: __Vkey9[:],
				}
			}
		}
	}
	for _, __Vkey6 := range __VkDelArr {
		delete(___Vdm.dmMdata.ddsMidx, __Vkey6)
	}
}
