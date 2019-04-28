package main

func (___Vdm *_TdataMachine) _FdataMachin__1000503x__cnNeedDnRetranOrDirectConnect(___Vc2sEncodeB *[]byte) {
	// _FrecePackThenEncodeAsLoad__1400201y__decode_and_check_and_repack pelCHc2sEncodeBLO

	__FpfN(" 839199 01 : _TdataMachine : cnNeedDnRetranOrDirectConnect :<%s> ", String9(___Vc2sEncodeB))

	__Venc := _Tencode{} // _TencodeX
	//_FdataPack__dePackUdpNodeRece__decode(&__Venc, *___Vc2sEncodeB)
	_FdataPack__deGob__encode(&__Venc, *___Vc2sEncodeB)
	if __Venc.EnLoadType != LoadT__data_01_special {
		_CFpfN(" 839199 02 : _TdataMachine : EnLoadType error :<%s> {%s}",
			String9(___Vc2sEncodeB), __Venc.String())
		return
	}

	if __Venc.EnData.DDcmd != DDType__c2s { // byte
		_CFpfN(" 839199 03 : _TdataMachine : DDcmd error :<%s> {%s}",
			String9(___Vc2sEncodeB), __Venc.String())
		return
	}

	__Vlen2 := len(___Vdm.dmMdata.ddsMidx) // _TdataMachinEdataSt
	if 1 > __Vlen2 {
		_CFpfN(" 839199 04 : _TdataMachine : c2s received , but , not Dn avaiable. ")
		return
	}

	__VidxRand := _FgenRand_int() % __Vlen2

	var __VdmClient *_TdataMachinEdataClient = nil

	for _, __Vv3 := range ___Vdm.dmMdata.ddsMidx { // _TdataMachinEdataSt // ddsMidx map[[16]byte]int
		if __Vv3 == __VidxRand {
			__VdmClient = &(___Vdm.dmMdata.ddsMm[__Vv3]) // _TdataMachinEdataClient
			break
		}
	}

	if nil == __VdmClient {
		_CFpfN(" 839199 05 : _TdataMachine : NULL ! what happens ? ")
		return
	}

	_CFpfN(" 839199 06 : _TdataMachine : dmClient {%s} ", __VdmClient.String())

	_CFpfN(" 839199 09 : _TdataMachine : cnNeedDnRetranOrDirectConnect :<%s> <%s>", String9(___Vc2sEncodeB), __Venc.String())
}
