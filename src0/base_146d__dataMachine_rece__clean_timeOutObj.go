package main

func (___Vdm *_TdataMachine) _FdataMachin__1000501y__clean_timeoutObj() {
	___Vdm._FdataMachin__1000501y1__clean_timeoutConn()
	___Vdm._FdataMachin__1000501y2__clean_timeoutData()
}

func (___Vdm *_TdataMachine) _FdataMachin__1000501y1__clean_timeoutConn() {
	__Vnow2 := _FtimeInt()

	__VkDelArr := [][16]byte{}
	for __Vkey7, __Vidx7 := range ___Vdm.dmMconn.dcsMidx { // _TdataMachinEconnectSt
		if __Vnow2-___Vdm.dmMconn.dcsMm[__Vidx7].dccLastReceTime > _Vgap_connectLostTimeOut {
			__VkDelArr = append(__VkDelArr, __Vkey7)
			___Vdm.dmMconn.dcsMm[__Vidx7].Clear()
		}
	}
	for _, __Vkey8 := range __VkDelArr {
		delete(___Vdm.dmMconn.dcsMidx, __Vkey8)
		___Vdm.dmMconn.dcsMusedAmount--
		___Vdm.dmMconn.dcsMfreeAmount++
	}
	if 0 > ___Vdm.dmMconn.dcsMusedAmount {
		_CFpfN(" 381933 02 : %11d : used Amount calc error , check what hapens {%s}", ___Vdm.dmMdata.String())
	}
	//_CFpfN(" 381933 01 : %11d : Checking timeOutObj, conn{%s}", _FtimeInt(), ___Vdm.dmMconn.String())
}

func (___Vdm *_TdataMachine) _FdataMachin__1000501y2__clean_timeoutData() {
	__Vnow2 := _FtimeInt()
	__VkDelArr := [][16]byte{}                             // _TdataMachinEdataClient
	for __Vkey3, __Vidx3 := range ___Vdm.dmMdata.ddsMidx { // _TdataMachinEdataSt
		if __Vnow2-___Vdm.dmMdata.ddsMm[__Vidx3].ddcLastReceTime > _Vgap_connectLostTimeOut {
			__VkDelArr = append(__VkDelArr, __Vkey3)

			_CFpfN(" 381932 01 : %11d : try to stop lost connect %x in loginAutoGenerator, {%s}",
				_FtimeInt(), __Vkey3, ___Vdm.dmMdata.ddsMm[__Vidx3].String()) // _TdataMachinEdataClient

			//___Vdm.dmMdata.ddsMm[__Vidx3].Clear()

		}
	}

	for _, __Vkey9 := range __VkDelArr {
		if nil == ___Vdm.dmCHloginGenMachineIdLO {
			_CFpfN(" 381932 02 : %11d : try to stop lost connect %x in loginAutoGenerator , but NULL out-chain(%s)",
				_FtimeInt(), __Vkey9, ___Vdm.dmMdata.String())
		} else {
			__Vidx9 := ___Vdm.dmMdata.ddsMidx[__Vkey9] // _TdataMachinEdataSt

			switch ___Vdm.dmMdata.ddsMm[__Vidx9].ddcRole { // _TdataMachinEdataClient
			case "Fn":
				_CFpfN(" 381932 03 : Fn lost , so need to re_connect , send msg %x to loginAutoGenerator, {%s}",
					__Vkey9, ___Vdm.dmMdata.ddsMm[__Vidx9].String()) // _TdataMachinEdataClient

				(*___Vdm.dmCHloginGenMachineIdLO) <- _TdataMachinEid{
					diIdx128: __Vkey9[:],
				}

			case "Cn", "Dn":
				_CFpfN(" 381932 05 : Cn / Dn lost , so need to re_connect , send msg %x to loginAutoGenerator, {%s}",
					__Vkey9, ___Vdm.dmMdata.ddsMm[__Vidx9].String()) // _TdataMachinEdataClient

			default:
				_CFpfN(" 381932 07 : connect lost , key %x , don't know what happens : {%s}",
					__Vkey9, ___Vdm.dmMdata.ddsMm[__Vidx9].String()) // _TdataMachinEdataClient
			}

			___Vdm.dmMdata.ddsMm[__Vidx9].Clear() // _TdataMachinEdataClient
		}
	}

	for _, __Vkey6 := range __VkDelArr {
		delete(___Vdm.dmMdata.ddsMidx, __Vkey6) // _TdataMachine
		___Vdm.dmMdata.ddsMusedAmount--
		___Vdm.dmMdata.ddsMfreeAmount++
	}
	if 0 > ___Vdm.dmMdata.ddsMusedAmount {
		_CFpfN(" 381932 08 : %11d : used Amount calc error , check what hapens {%s}", ___Vdm.dmMdata.String())
	}

	if 0 != len(__VkDelArr) {
		_CFpfN(" 381932 09 : %11d : after delete objS, data{%s}", _FtimeInt(), ___Vdm.dmMdata.String())
	}
}
