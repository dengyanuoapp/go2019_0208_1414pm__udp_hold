package main

func (___Vdm *_TdataMachine) _FdataMachin__1000501y__clean_timeoutObj() {
	___Vdm._FdataMachin__1000501y1__clean_timeoutConn()
	___Vdm._FdataMachin__1000501y2__clean_timeoutData()
}

func (___Vdm *_TdataMachine) _FdataMachin__1000501y1__clean_timeoutConn() {
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
}

func (___Vdm *_TdataMachine) _FdataMachin__1000501y2__clean_timeoutData() {
	__Vnow2 := _FtimeInt()
	__VkDelArr := [][16]byte{}
	for __Vkey3, __Vidx3 := range ___Vdm.dmMdata.ddsMidx { // _TdataMachinEdataSt
		if __Vnow2-___Vdm.dmMdata.ddsMm[__Vidx3].ddcLastReceTime > _Vgap_nothingToLost {
			__VkDelArr = append(__VkDelArr, __Vkey3)

			___Vdm.dmMdata.ddsMm[__Vidx3].Clear()

		}
	}
	for _, __Vkey6 := range __VkDelArr {
		delete(___Vdm.dmMdata.ddsMidx, __Vkey6)
	}
	for _, __Vkey9 := range __VkDelArr {
		if nil == ___Vdm.dmCHloginGenMachineIdLO {
			_FpfN(" 381932 04 : %11d : try to stop lost connect %x in loginGen , but NULL out-chain(%s)", _FtimeInt(), __Vkey9, ___Vdm.dmMdata.String())
			_CpfN(" 381932 05 : %11d : try to stop lost connect %x in loginGen , but NULL out-chain(%s)", _FtimeInt(), __Vkey9, ___Vdm.dmMdata.String())
		} else {
			_FpfN(" 381932 06 : %11d : try to stop lost connect %x in loginGen, ok {%s}", _FtimeInt(), __Vkey9, ___Vdm.dmMdata.String())
			_CpfN(" 381932 07 : %11d : try to stop lost connect %x in loginGen, ok {%s}", _FtimeInt(), __Vkey9, ___Vdm.dmMdata.String())

			(*___Vdm.dmCHloginGenMachineIdLO) <- _TdataMachinEid{
				diIdx128: __Vkey9[:],
			}
		}
	}
}
