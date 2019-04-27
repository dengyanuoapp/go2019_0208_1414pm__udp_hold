package main

func (___Vdm *_TdataMachine) _FdataMachin__1000503x__cnNeedDnRetranOrDirectConnect(___Vc2sEncodeB *[]byte) {
	// _FrecePackThenEncodeAsLoad__1400201y__decode_and_check_and_repack pelCHc2sEncodeBLO

	_CFpfN(" 839199 01 : _TdataMachine : cnNeedDnRetranOrDirectConnect :<%s> ", String9(___Vc2sEncodeB))

	__Venc := _Tencode{} // _TencodeX
	_FdataPack__dataDecode_common(&__Venc, *___Vc2sEncodeB)
	if __Venc.Type != LoadT__data_01_special {
		_CFpfN(" 839199 02 : _TdataMachine : cnNeedDnRetranOrDirectConnect :<%s> {%s}",
			String9(___Vc2sEncodeB), __Venc.String())
		return
	}

	__Venc := _Tencode{ // _TencodeX
		Ti:         __VunRece.Ti,              // _TdecodeX
		enToId128:  __Vdecode.Dlogin.MeIdx128, // _TloginReq
		enLoadType: LoadT__data_01_special,    //byte
	}

}
