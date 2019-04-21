package main

func (___Vdm *_TdataMachine) _FdataMachin__1000502y__dataSendIdle__packAndSendAll() {

	if 0 == len(___Vdm.dmMdata.ddsMidx) {
		_FpfN(" 381921 01 : %11d : try send idle , but why no client ? {%s} --- {%s} ", _FtimeInt(), ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())
		_CpfN(" 381921 02 : %11d : try send idle , but why no client ? {%s} --- {%s} ", _FtimeInt(), ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())
		return
	}

	//dmmLastReceTime
	//__VkDelArr := [][16]byte{}
	__Vnow2 := _FtimeInt()

	defer ___Vout__dmCHloginGenMachineIdLO__mux.Unlock()
	defer ___Vdm.dmMconn.dcsMux.Unlock()

	___Vout__dmCHloginGenMachineIdLO__mux.Lock()
	___Vdm.dmMconn.dcsMux.Lock()

	for __Vkey6, __Vidx6 := range ___Vdm.dmMdata.ddsMidx { // _TdataMachinEdataClient _TdataMachinEdataSt
		if __Vnow2-___Vdm.dmMdata.ddsMm[__Vidx6].ddcLastSendTime > _Vgap_skip_idle_send {
			// pack as _TdataTran -->  _TdecodeX .  Ddata // _TencodeX
			//_CFpfN(" 381921 04 : %11d : try send idle %x {%s}", _FtimeInt(), __Vkey6, ___Vdm.dmMdata.ddsMm[__Vidx6].String())

			___Vdm.
				_FdataMachin__1000502z__dataSendIdle__packAndSendToOneID(&(___Vdm.dmMdata.ddsMm[__Vidx6]))

		} else {
			_FpfN(" 381921 07 : %11d : try send idle %x , but in 10s sent already. Skip. {%s}",
				_FtimeInt(), __Vkey6, ___Vdm.dmMdata.ddsMm[__Vidx6].String())
		}
	}
}

func (___Vdm *_TdataMachine) _FdataMachin__1000502z__dataSendIdle__packAndSendToOneID(___Vdmdc *_TdataMachinEdataClient) {
	__VtoAddrAmount := len(___Vdmdc.ddcConnPortArr) // _TdataMachinEdataClient

	if __VtoAddrAmount < 1 {
		_FpfN(" 381922 01 len error :%d ", __VtoAddrAmount)
		return
	}

	__Vidx := _FgenRand_int() % __VtoAddrAmount
	if ___Vdmdc.ddcLastSendIdx == __Vidx {
		__Vidx = _FgenRand_int() % __VtoAddrAmount
	}
	___Vdmdc.ddcLastSendIdx = __Vidx

	// _TdataMachinEdataClient
	__Vconn := ___Vdmdc.ddcConnPortArr[__Vidx]

	__Venc := _Tencode{ // _TencodeX // Cmd__loginS01genReplyTokenA
		Ti:           _FtimeInt(),
		enType:       Cmd__data_01_idle,
		enToId128:    ___Vdmdc.ddcID.diIdx128, // _TdataMachinEid
		enToConnPort: __Vconn,                 // _TudpConnPort
		enData: _TdataTran{
			MEidx128: _VC.MyId128,
			MYseq128: _VS.MySeq128,
			TOidx128: ___Vdmdc.ddcID.diIdx128,
			TOseq128: ___Vdmdc.ddcID.diSeq128,
			DDcmd:    Cmd__data_01_idle,
		},
	}

	//_FpfN(" 381922 02 {%#v} ", ___Vdmdc) // _TdataMachinEconnectClient
	//_FpfN(" 381922 03 {%s} ", ___Vdmdc.String())
	//_FpfN(" 381922 04 : myID:%x mySeq:%x id:%x", _VC.MyId128, _VS.MySeq128, ___Vid)

	//_FpfN(" 381922 05 _TdataMachine -> _Tencode {%s}", __Venc.String()) // _TencodeX
	//_CpfN(" 381922 06 _TdataMachine -> _Tencode {%s}", __Venc.String()) // _TencodeX

	//_Fex("381922 07")

	___Vdmdc.ddcLastSendTime = _FtimeInt() // _TdataMachinEconnectClient

	___Vdm.
		_FdataMachin__1000601x__encodeData_sendMux(&__Venc)
}
