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

		__Vc.dccLockCntLast = __Vc.dccLockCntNow
		if __Vcnt >= 3 { // one connect must have the amount of connect-portS large then 2

			if nil == ___Vdm.dmCHloginGenMachineIdLO {
				switch _VS.ProjName {
				case "Fn":
					// do nothing
					_FpfNonce(" 839196 01 connected succeed, Fn do NOT have loginGen , so NULL.<%s>", _VS.ProjName)
					_CpfN(" 839196 02 connected succeed, Fn do NOT have loginGen , so NULL.<%s>", _VS.ProjName)
				default:
					_FpfN(" 839196 03 connected succeed, but why loginGen NULL ? <%s>", _VS.ProjName)
					_CpfN(" 839196 04 connected succeed, but why loginGen NULL ? <%s>", _VS.ProjName)
				}
			} else {
				//_FpfN(" 839196 05 connected succeed, then told loginGen to stop ( push chan)")
				_CpfN(" 839196 06 connected succeed, then told loginGen to stop ( push chan)")
				___Vout__dmCHloginGenMachineIdLO__mux.Lock()
				(*___Vdm.dmCHloginGenMachineIdLO) <- ___Vdm.dmMconn.dcsMm[__Vidx2].dccID
				___Vout__dmCHloginGenMachineIdLO__mux.Unlock()

			}

			___Vdm.
				_FdataMachin__1000501z__swapLoginCkInfoForLock__createNewDataClient(&__Vkey2, __Vidx2)
			__Vc.dccLockCntLast = 0 // clear the counter , forbit being next used

		} else {
			//_FpfN(" 839196 07 no-enough, swap only , no update . (%d)[%x]{%s}" , __Vidx2 , __Vkey2 , ___Vdm.dmMconn.String())
			_CpfN(" 839196 08 no-enough, swap only , update if possible. (%d)[%x]{%s}", __Vidx2, __Vkey2, ___Vdm.dmMconn.String())
			//__Vc.dccLockCntLast = 0 // clear the counter , forbit being next used
			___Vdm.
				_FdataMachin__1000501z__swapLoginCkInfoForLock__tryUpdateAddressArrIfSame(&__Vkey2, __Vidx2)
		}
		__Vc.dccLockCntNow = 0 // _TdataMachinEconnectSt
	}
	_CpfN(" 839196 09 after swap {%s} --- {%s} ", ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())
}

// _FdataMachin__1000201x11__rece_machineId_check_and_insert
func (___Vdm *_TdataMachine) _FdataMachin__1000501z__swapLoginCkInfoForLock__createNewDataClient(___Vkey *[16]byte, ___VcIdx int) {
	// ___Vdm.dmMconn.dcsMm[__Vidx2].
	__Vc := &(___Vdm.dmMconn.dcsMm[___VcIdx])
	__Vidx4, __VdOk4 := ___Vdm.dmMdata.ddsMidx[*___Vkey] // _TdataMachinEdataSt _TdataMachinEdataClient
	if __VdOk4 {
		_CpfN(" 183818 01 : re-create client of {%s} ", ___Vdm.dmMconn.String()) // _TdataMachinEconnectSt _TdataMachinEconnectClient
		_CpfN(" 183818 02 : re-create client of {%s} ", ___Vdm.dmMdata.String()) // _TdataMachinEdataSt    _TdataMachinEdataClient
	} else {
		_CpfN(" 183818 05 : create new client of {%s} ", ___Vdm.dmMconn.String()) // _TdataMachinEconnectSt _TdataMachinEconnectClient
		_CpfN(" 183818 06 : create new client of {%s} ", ___Vdm.dmMdata.String()) // _TdataMachinEdataSt    _TdataMachinEdataClient
		__Vidx4 = ___Vdm.dmMdata.
			_FdataMachin__search_avaiable__TdataMachinEdataClient() // _TdataMachinEdataClient _TdataMachinEdataSt
		if __Vidx4 < 0 {
			_CpfN(" 183818 07 : new conn-data insered <%s>", ___Vdm.dmMdata.String())
			return
		}
	}
	__Vd := &(___Vdm.dmMdata.ddsMm[__Vidx4])    // _TdataMachinEdataClient
	__Vd.Clear()                                // _TdataMachinEdataSt
	__Vd.ddcID = __Vc.dccID                     // _TdataMachinEid
	__Vd.ddcConnPortArr = __Vc.dccConnPortArr   // []_TudpConnPort
	for _, __Vv5 := range __Vd.ddcConnPortArr { // _TudpConnPort
		__Vd.ddcConnPortStrMap[__Vv5.DstAddr.String()] = 1
	}
	__Vd.ddcConnPortAmount = __Vc.dccConnPortAmount //  int
	//__Vd .  ddcLastReceTime   int
	//__Vd .  ddcLastSendTime   int
	//__Vd .  ddcLastSendIdx    int
	_FpfNex(" 183818 09 {%s} --- {%s} ", ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())
}
func (___Vdm *_TdataMachine) _FdataMachin__1000501z__swapLoginCkInfoForLock__tryUpdateAddressArrIfSame(___Vkey *[16]byte, ___VcIdx int) {
	_FpfNex(" 183819 01 {%s} --- {%s} ", ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())
}
