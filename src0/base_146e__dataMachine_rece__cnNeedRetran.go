package main

func (___Vdm *_TdataMachine) _FdataMachin__1000503x__cnNeedDnRetranOrDirectConnect(___Vc2sEncodeB *[]byte) {
	// _FrecePackThenEncodeAsLoad__1400201y__decode_and_check_and_repack pelCHc2sEncodeBLO

	__FpfN(" 839199 01 : _TdataMachine : cnNeedDnRetranOrDirectConnect :<%s> ", String9s(___Vc2sEncodeB))

	__Venc := _Tencode{} // _TencodeX
	//_FdataPack__dePackUdpNodeRece__decode(&__Venc, *___Vc2sEncodeB)
	_FdataPack__deGob__encode(&__Venc, *___Vc2sEncodeB)
	if __Venc.EnLoadType != LoadT__data_01_special {
		_CFpfN(" 839199 02 : _TdataMachine : EnLoadType error :<%s> {%s}",
			String9s(___Vc2sEncodeB), __Venc.String())
		return
	}

	if __Venc.EnData.DtDDcmd != DDType__c2s { // byte
		_CFpfN(" 839199 03 : _TdataMachine : DtDDcmd error :<%s> {%s}",
			String9s(___Vc2sEncodeB), __Venc.String())
		return
	}

	__Vlen2 := len(___Vdm.dmMdata.ddsMidx) // _TdataMachinEdataSt
	if 1 > __Vlen2 {
		_CFpfN(" 839199 04 : _TdataMachine : c2s received , but , not Dn avaiable. ")
		return
	}

	__VidxRand := _FgenRand_int() % __Vlen2

	var __VdmClient *_TdataMachinEdataClient = nil

	__Vi := 0
	for __Vidx, __Vv3 := range ___Vdm.dmMdata.ddsMidx { // _TdataMachinEdataSt // ddsMidx map[[16]byte]int
		if __Vi == __VidxRand {
			_CFpfN(" 839199 05 : _TdataMachine : %s rePackAndTran data and send to %s : __Vidx %x , __Vv3 %d , __Vi %d",
				_VS.RoleName, ___Vdm.dmMdata.ddsMm[__Vv3].ddcRole, __Vidx[:5], __Vv3, __Vi)
			__VdmClient = &(___Vdm.dmMdata.ddsMm[__Vv3]) // _TdataMachinEdataClient
			break
		}
		__Vi++
	}

	if nil == __VdmClient {
		_CFpfN(" 839199 06 : _TdataMachine : NULL ! what happens ? __Vlen2:%d , __VidxRand:%d", __Vlen2, __VidxRand)
		_CFpfN(" 839199 07 : ___Vdm.dmMdata.ddsMidx {%v} ", ___Vdm.dmMdata.ddsMidx)
		_CFpfN(" 839199 08 : ___Vdm.dmMdata {%s} ", ___Vdm.dmMdata.String())
		return
	} // _TencodeX

	__FpfN(" 839199 09 : _TdataMachine : dmClient {%s} === Vc2sEncodeB<%s> === Venc<%s>",
		__VdmClient.String(), String9s(___Vc2sEncodeB), __Venc.String())

	___Vdm.
		_FdataMachin__1000503y__cnNeedDnRetranOrDirectConnect(__VdmClient, &__Venc)

}

func (___Vdm *_TdataMachine) _FdataMachin__1000503y__cnNeedDnRetranOrDirectConnect(___Vdmdc *_TdataMachinEdataClient, ___Venc *_Tencode) {
	/*
		__Venc := _Tencode{ // _TencodeX // LoadT__loginS01genReplyTokenA
			EnLoadType: LoadT__data_01_special,
			EnToId128:  ___Vdmdc.ddcID.diIdx128, // _TdataMachinEid
			//EnToConnPort: __Vconn,                 // _TudpConnPort
			EnData: _TdataTran{
				DtMEidx128: _VC.MyId128,
				DtMYseq128: _VS.MySeq128,
				DtTOidx128: ___Vdmdc.ddcID.diIdx128,
				DtTOseq128: ___Vdmdc.ddcID.diSeq128,
				DtDDcmd:    DDType__idle,
			},
		}
	*/

	___Vdm.
		_FdataMachin__1000601y__encodeData_sendEnc_only(___Vdmdc, ___Venc)
}
