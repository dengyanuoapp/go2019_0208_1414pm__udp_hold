package main

func (___Vdm *_TdataMachine) _FdataMachin__1000503x__cnNeedDnRetranOrDirectConnect(___Vc2sEncodeB *[]byte) {
	// _FrecePackThenEncodeAsLoad__1400201y__decode_and_check_and_repack pelCHc2sEncodeBLO

	_CFpfN(" 839199 01 : _TdataMachine : cnNeedDnRetranOrDirectConnect :<%s> <%x>", String9(___Vc2sEncodeB), _Fmd5__5x(___Vc2sEncodeB))

	__Venc := _Tencode{} // _TencodeX
	//_FdataPack__dePackUdpNodeRece__decode(&__Venc, *___Vc2sEncodeB)
	_FdataPack__deGob__encode(&__Venc, *___Vc2sEncodeB)
	if __Venc.enLoadType != LoadT__data_01_special {
		_CFpfN(" 839199 02 : _TdataMachine : enLoadType error :<%s> {%s}",
			String9(___Vc2sEncodeB), __Venc.String())
		return
	}

	if __Venc.enData.DDcmd != DDType__c2s { // byte
		_CFpfN(" 839199 03 : _TdataMachine : DDcmd error :<%s> {%s}",
			String9(___Vc2sEncodeB), __Venc.String())
		return
	}

	_CFpfN(" 839199 09 : _TdataMachine : cnNeedDnRetranOrDirectConnect :<%s> ", String9(___Vc2sEncodeB))
}
