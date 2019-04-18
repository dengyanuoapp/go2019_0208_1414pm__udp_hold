package main

import "sync"

var ___Vout__dmCHloginGenMachineIdLO__mux sync.Mutex

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

// _FdataMachin__1000201x11__rece_machineId_check_and_insert
func (___Vdm *_TdataMachine) _FdataMachin__1000501y__swapLoginCkInfoForLock__swap() {
	defer ___Vdm.dmMconn.dcsMux.Unlock() // _TdataMachinEconnectSt
	___Vdm.dmMconn.dcsMux.Lock()         // _TdataMachinEconnectClient

	var __Vcnt int // _TdataMachinEconnectSt
	for __Vkey2, __Vidx2 := range ___Vdm.dmMconn.dcsMidx {
		__Vc := ___Vdm.dmMconn.dcsMm[__Vidx2] // _TdataMachinEconnectClient
		__Vcnt = __Vc.dccLockCntLast + __Vc.dccLockCntNow

		if __Vcnt >= 3 { // one connect must have the amount of connect-portS large then 2
			if nil == ___Vdm.dmCHloginGenMachineIdLO {
				_FpfNonce(" 839196 08 connected succeed, but why loginGen NULL ?")
			} else {
				//_FpfN(" 839196 09 connected succeed, then told loginGen to stop ( push chan)")
				___Vout__dmCHloginGenMachineIdLO__mux.Lock()
				(*___Vdm.dmCHloginGenMachineIdLO) <- ___Vdm.dmMconn.dcsMm[__Vidx2].dccID
				___Vout__dmCHloginGenMachineIdLO__mux.Unlock()
			}
			__Vc.dccLockCntLast = 0 // clear the counter , forbit being next used
			__Vc.dccLockCntNow = 0  // clear the counter , forbit being next used

			___Vdm.
				_FdataMachin__1000501z__swapLoginCkInfoForLock__createNewDataClient(&__Vkey2, __Vidx2)
		}
	}
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
