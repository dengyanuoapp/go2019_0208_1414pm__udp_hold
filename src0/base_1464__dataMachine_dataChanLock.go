package main

import "sync"
import "bytes"

var ___Vout__dmCHloginGenMachineIdLO__mux sync.Mutex

// _FdataMachin__1000201x11__rece_machineId_check_and_insert
func (___Vdm *_TdataMachine) _FdataMachin__1000501y__swapLoginCkInfoForLock__swap() {

	defer ___Vdm.dmMconn.dcsMux.Unlock() // _TdataMachinEconnectSt
	___Vdm.dmMconn.dcsMux.Lock()         // _TdataMachinEconnectClient

	if 0 == len(___Vdm.dmMconn.dcsMidx) {
		//_CFpfN(" 839196 00 zero-len conn-Arr , nothing need to swap.")
		return
	}

	var __Vcnt int // _TdataMachinEconnectSt
	for __Vkey2, __Vidx2 := range ___Vdm.dmMconn.dcsMidx {
		__Vc := &(___Vdm.dmMconn.dcsMm[__Vidx2]) // _TdataMachinEconnectClient
		__Vcnt = __Vc.dccLockCntLast + __Vc.dccLockCntNow

		__Vc.dccLockCntLast = __Vc.dccLockCntNow
		if 0 == __Vc.dccLockCntNow {
			_CpfN(" 839196 01 no-enough, swap only , update if possible. (%d)[%x]{%s}", __Vidx2, __Vkey2, ___Vdm.dmMconn.String())
			__Vc.dccLockCntNow = 0 // _TdataMachinEconnectSt
			continue
		}

		if __Vcnt <= 2 { // one connect must have the amount of connect-portS large then 2
			_CpfN(" 839196 02 no-enough, swap only , update if possible. (%d)[%x]{%s}", __Vidx2, __Vkey2, ___Vdm.dmMconn.String())
			//__Vc.dccLockCntLast = 0 // clear the counter , forbit being next used
			___Vdm.
				_FdataMachin__1000501z__swapLoginCkInfoForLock__tryUpdateAddressArrIfSame(&__Vkey2, __Vc)
			__Vc.dccLockCntNow = 0 // _TdataMachinEconnectSt
			continue
		}

		___Vdm.
			_FdataMachin__1000501z__swapLoginCkInfoForLock__createNewDataClient(&__Vkey2, __Vc)

		if nil == ___Vdm.dmCHloginGenMachineIdLO {
			switch _VS.RoleName {
			case "Fn":
				// do nothing
				_FpfN(" 839196 03 connected succeed, Fn do NOT have loginGen , so NULL.<%s>", _VS.RoleName)
				_CpfN(" 839196 04 connected succeed, Fn do NOT have loginGen , so NULL.<%s>", _VS.RoleName)
			default:
				_FpfN(" 839196 05 connected succeed, but why loginGen NULL ? <%s>", _VS.RoleName)
				_CpfN(" 839196 06 connected succeed, but why loginGen NULL ? <%s>", _VS.RoleName)
			}
		} else {
			_FpfN(" 839196 07 connected succeed, then told loginGen to stop ( push chan)")
			_CpfN(" 839196 08 connected succeed, then told loginGen to stop ( push chan)")
			___Vout__dmCHloginGenMachineIdLO__mux.Lock()
			(*___Vdm.dmCHloginGenMachineIdLO) <- ___Vdm.dmMconn.dcsMm[__Vidx2].dccID
			___Vout__dmCHloginGenMachineIdLO__mux.Unlock()

		}

		__Vc.dccLockCntLast = 0 // clear the counter , forbit being next used
		__Vc.dccLockCntNow = 0  // _TdataMachinEconnectSt
	}
	_CpfN(" 839196 09 after swap {%s} ======######====== {%s} ", ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())
}

// _FdataMachin__1000201x11__rece_machineId_check_and_insert
func (___Vdm *_TdataMachine) _FdataMachin__1000501z__swapLoginCkInfoForLock__createNewDataClient(___Vkey *[16]byte, ___Vc *_TdataMachinEconnectClient) {
	// ___Vdm.dmMconn.dcsMm[__Vidx2].
	//__Vc := &(___Vdm.dmMconn.dcsMm[___VcIdx])
	__Vidx4, __VdOk4 := ___Vdm.dmMdata.ddsMidx[*___Vkey] // _TdataMachinEdataSt _TdataMachinEdataClient
	if __VdOk4 {
		_CpfN(" 839197 01 : re-create client of  conn{%s} ", ___Vdm.dmMconn.String()) // _TdataMachinEconnectSt _TdataMachinEconnectClient
		_CpfN(" 839197 02 : re-create client of  data{%s} ", ___Vdm.dmMdata.String()) // _TdataMachinEdataSt    _TdataMachinEdataClient
	} else {
		_CpfN(" 839197 03 : create new client of conn{%s} ", ___Vdm.dmMconn.String()) // _TdataMachinEconnectSt _TdataMachinEconnectClient
		_CpfN(" 839197 04 : create new client of data{%s} ", ___Vdm.dmMdata.String()) // _TdataMachinEdataSt    _TdataMachinEdataClient
		__Vidx4 = ___Vdm.dmMdata.
			_FdataMachin__search_avaiable__TdataMachinEdataClient() // _TdataMachinEdataClient _TdataMachinEdataSt
		if __Vidx4 < 0 {
			_CpfN(" 839197 07 : new conn-data insered <%s>", ___Vdm.dmMdata.String())
			return
		}
		___Vdm.dmMdata.ddsMidx[*___Vkey] = __Vidx4
	}
	__Vd := &(___Vdm.dmMdata.ddsMm[__Vidx4])    // _TdataMachinEdataClient
	__Vd.Clear()                                // _TdataMachinEdataSt
	__Vd.ddcID = ___Vc.dccID                    // _TdataMachinEid
	__Vd.ddcRole = ___Vc.dccRole                // _TdataMachinEid
	__Vd.ddcConnPortArr = ___Vc.dccConnPortArr  // []_TudpConnPort
	for _, __Vv5 := range __Vd.ddcConnPortArr { // _TudpConnPort
		__Vd.ddcConnPortStrMap[__Vv5.DstAddr.String()] = 1
	}
	__Vd.ddcConnPortAmount = ___Vc.dccConnPortAmount //  int
	//__Vd .  ddcLastReceTime   int
	//__Vd .  ddcLastSendTime   int
	//__Vd .  ddcLastSendIdx    int
	___Vdm.dmMdata.ddsMusedAmount++
	___Vdm.dmMdata.ddsMfreeAmount--
	_CpfN(" 839197 08 : after create         conn{%s} ", ___Vdm.dmMconn.String()) // _TdataMachinEconnectSt _TdataMachinEconnectClient
	_CpfN(" 839197 09 : after create         data{%s} ", ___Vdm.dmMdata.String()) // _TdataMachinEdataSt    _TdataMachinEdataClient
}

func (___Vdm *_TdataMachine) _FdataMachin__1000501z__swapLoginCkInfoForLock__tryUpdateAddressArrIfSame(___Vkey *[16]byte, ___Vc *_TdataMachinEconnectClient) {
	_CpfN(" 839198 01 {%s} ======######====== {%s} ", ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())

	__Vidx4, __VdOk4 := ___Vdm.dmMdata.ddsMidx[*___Vkey] // _TdataMachinEdataSt _TdataMachinEdataClient
	if false == __VdOk4 {
		_CpfN(" 839198 02 {%s} ======######====== {%s} not exist, do nothing.", ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())
		return
	}

	__Vd := &(___Vdm.dmMdata.ddsMm[__Vidx4])                     // _TdataMachinEdataClient
	if bytes.Equal(__Vd.ddcID.diIdx128, ___Vc.dccID.diIdx128) && // _TdataMachinEid
		bytes.Equal(__Vd.ddcID.diSeq128, ___Vc.dccID.diSeq128) && // _TdataMachinEid
		bytes.Equal(__Vd.ddcID.diToken, ___Vc.dccID.diToken) { // _TdataMachinEid
		for _, __Vv5 := range ___Vc.dccConnPortArr { // _TudpConnPort
			_, __Vok6 := __Vd.ddcConnPortStrMap[__Vv5.DstAddr.String()]
			if false == __Vok6 { // _TdataMachinEdataSt
				__Vd.ddcConnPortArr = append(__Vd.ddcConnPortArr, __Vv5) // []_TudpConnPort
				__Vd.ddcConnPortStrMap[__Vv5.DstAddr.String()] = 1
				__Vd.ddcConnPortAmount++
			}
		}
		_CpfN(" 839198 07 {%s} ======######====== {%s} ", ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())
		return
	}

	__Vd.Clear()                                // _TdataMachinEdataSt
	__Vd.ddcID = ___Vc.dccID                    // _TdataMachinEid
	__Vd.ddcRole = ___Vc.dccRole                // _TdataMachinEid
	__Vd.ddcConnPortArr = ___Vc.dccConnPortArr  // []_TudpConnPort
	for _, __Vv5 := range __Vd.ddcConnPortArr { // _TudpConnPort
		__Vd.ddcConnPortStrMap[__Vv5.DstAddr.String()] = 1
	}
	__Vd.ddcConnPortAmount = ___Vc.dccConnPortAmount //  int
	_CpfN(" 839198 09 {%s} ======######====== {%s} ", ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())

}
