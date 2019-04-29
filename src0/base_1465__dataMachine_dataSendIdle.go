package main

func (___Vdm *_TdataMachine) _FdataMachin__1000502y__dataSendIdle__packAndSendAll() {

	if 0 == len(___Vdm.dmMdata.ddsMidx) {
		//_CFpfN(" 381921 01 : %11d : try send idle , but why no client ? {%s} ======######====== {%s} ", _FtimeInt(), ___Vdm.dmMconn.String(), ___Vdm.dmMdata.String())
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
			_CpfN(" 381921 07 : %11d : try send idle %x , but in 10s sent already. Skip. {%s}",
				_FtimeInt(), __Vkey6, ___Vdm.dmMdata.ddsMm[__Vidx6].String())
		}
	}
}

func (___Vdm *_TdataMachine) _FdataMachin__1000502z__dataSendIdle__packAndSendToOneID(___Vdmdc *_TdataMachinEdataClient) {
	__Venc := _Tencode{ // _TencodeX // LoadT__loginS01genReplyTokenA
		EnLoadType: LoadT__data_01_special,
		EnToId128:  ___Vdmdc.ddcID.diIdx128, // _TdataMachinEid
		//EnToConnPort: __Vconn,                 // _TudpConnPort
		EnData: _TdataTran{
			MEidx128: _VC.MyId128,
			MYseq128: _VS.MySeq128,
			TOidx128: ___Vdmdc.ddcID.diIdx128,
			TOseq128: ___Vdmdc.ddcID.diSeq128,
			DDcmd:    DDType__idle,
		},
	}

	___Vdm.
		_FdataMachin__1000601y__encodeData_sendEnc_only(___Vdmdc, &__Venc)
}
