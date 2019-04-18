package main

import (
	"bytes"
)

func (___Vdm *_TdataMachine) _FdataMachin__1000201x11__rece_machineId_check_and_insert(___VinsID *_TdataMachinEid) {

	defer ___Vdm.dmMconn.dcsMux.Unlock() // _TdataMachinEconnectSt
	___Vdm.dmMconn.dcsMux.Lock()

	if nil == ___VinsID || 16 != len(___VinsID.diIdx128) {
		_FpfN(" 839193 01 : input Error <%s>", ___VinsID.String())
	}

	__Vk := _FgenB16(&___VinsID.diIdx128)         // _TdataMachinEconnectClient
	__Vidx, __Vok := ___Vdm.dmMconn.dcsMidx[__Vk] // _TdataMachinEconnectSt

	//__VsetLastReadTime := _FtimeInt() // - _Vgap_skip_idle_send

	if false == __Vok { // if not exist
		__Vidx = ___Vdm.dmMconn.
			_FdataMachin__search_avaiable__TdataMachinEconnectClient() // _TdataMachinEconnectClient _TdataMachinEconnectSt
		if __Vidx < 0 {
			return
		}

		// _TdataMachinEconnectClient _TdataMachinEconnectSt
		___Vdm.dmMconn.dcsMm[__Vidx].dccID = *___VinsID
		___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortArr = []_TudpConnPort{___VinsID.diConnPort}
		___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortStrMap[___VinsID.diConnPort.DstAddr.String()] = 1
		___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortAmount = 1
		___Vdm.dmMconn.dcsMm[__Vidx].dccLastReceTime = _FtimeInt()
		___Vdm.dmMconn.dcsMm[__Vidx].dccLockCntNow = 1

		return
	}

	if // if exist , but  token / sequence not match , replace the old one
	false == bytes.Equal(___Vdm.dmMconn.dcsMm[__Vidx].dccID.diIdx128, ___VinsID.diIdx128) ||
		false == bytes.Equal(___Vdm.dmMconn.dcsMm[__Vidx].dccID.diSeq128, ___VinsID.diSeq128) ||
		false == bytes.Equal(___Vdm.dmMconn.dcsMm[__Vidx].dccID.diToken, ___VinsID.diToken) {
		//_FpfN(" 839193 03 : delete-old, gen New connPort Map hash.")
		___Vdm.dmMconn.dcsMm[__Vidx].
			Clear()

		___Vdm.dmMconn.dcsMm[__Vidx].dccID = *___VinsID
		___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortArr = []_TudpConnPort{___VinsID.diConnPort}
		___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortStrMap[___VinsID.diConnPort.DstAddr.String()] = 1
		___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortAmount = 1
		___Vdm.dmMconn.dcsMm[__Vidx].dccLastReceTime = _FtimeInt()
		___Vdm.dmMconn.dcsMm[__Vidx].dccLockCntNow = 1

		return
	}

	// if exist , and  token / sequence match , update the last time and inc the lock cnt
	//_FpfN(" 839193 05 : with the same token, so , try append to connPort Map hash.")

	//__VcpArr := ___VinsID.diConnPort._FdataMachin__1000201x12__appendConnPort(&___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortArr)
	//__VipStr := ___VinsID.diConnPort.DstAddr.String()

	_, __Vok3 := ___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortStrMap[___VinsID.diConnPort.DstAddr.String()]

	if false == __Vok3 {
		___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortArr = []_TudpConnPort{___VinsID.diConnPort}
		___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortStrMap[___VinsID.diConnPort.DstAddr.String()] = 1
		___Vdm.dmMconn.dcsMm[__Vidx].dccConnPortAmount++
	}

	___Vdm.dmMconn.dcsMm[__Vidx].dccLastReceTime = _FtimeInt()
	___Vdm.dmMconn.dcsMm[__Vidx].dccLockCntNow++
}

func (___VnewConnP *_TudpConnPort) _FdataMachin__1000201x12__appendConnPort(___VoldConnParr *[]_TudpConnPort) []_TudpConnPort {
	if nil == ___VoldConnParr || 0 == len(*___VoldConnParr) {
		_FpfN(" 839194 01 : nil , use old connPort Arr ")
		return []_TudpConnPort{*___VnewConnP}
	}

	__Vlen := len(*___VoldConnParr)
	__Vncp := make([]_TudpConnPort, __Vlen+1)
	__Vstr1 := ___VnewConnP.DstAddr.String()

	__Vi := 0
	for __Vi < __Vlen {
		__Vcp := (*___VoldConnParr)[__Vi]
		__Vstr2 := __Vcp.DstAddr.String()
		if __Vstr1 == __Vstr2 {
			//_FpfN(" 839194 03 : already exist , use old connPort Arr ")
			return (*___VoldConnParr)
		}
		__Vncp[__Vi] = __Vcp
		__Vi++
	}

	//_FpfN(" 839194 05 : insert 1 into connPort Arr ")
	__Vncp[__Vlen] = (*___VnewConnP)

	return __Vncp
}
