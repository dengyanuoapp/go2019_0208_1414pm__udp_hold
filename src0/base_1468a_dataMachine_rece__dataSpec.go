package main

func (___Vdm *_TdataMachine) _FdataMachin__rece__dataSpec(___Vdec *_Tdecode) { // _TdecodeX
	if ___Vdec.DEtype != LoadT__data_01_special {
		__FpfN(" 839291 01 : _TdataMachine : DEtype error : {%s}", ___Vdec.String())
		return
	}

	if ___Vdec.DEdata.DDcmd == DDType__idle { // ignore , no more process is needed.
		__FpfN(" 839291 01 : _TdataMachine : LoadT__data_01_special , idle , ignore :{%s}", ___Vdec.String())
		return
	}

	switch ___Vdec.DEdata.DDcmd { // _TdataTran
	case DDType__c2s:
		_CFpfN(" 839291 02 : _TdataMachine : LoadT__data_01_special decode from DDType__c2s' buf to origin-Tdecode :{%s}", ___Vdec.String())
		___Vdm._FdataMachin__rece__dataRepack_c2s(___Vdec) // _TdecodeX
	default:
		_CFpfN(" 839291 07 : _TdataMachine : LoadT__data_01_special unknown:{%s}", ___Vdec.String())
	}

	__FpfN(" 839291 09 : _TdataMachine : LoadT__data_01_special :{%s}", ___Vdec.String())
}

func (___Vdm *_TdataMachine) _FdataMachin__rece__dataRepack_c2s(___Vdec *_Tdecode) { // _TdecodeX
	__Vdecode := _Tdecode{}                                                 // _TdecodeX
	__Verr5 := _FdecGob___(" 839292 11 ", ___Vdec.DEdata.DDbuf, &__Vdecode) // _TdataTran
	if nil != __Verr5 {
		_CFpfN(" 839292 12 rece__dataRepack_c2 error: {%v} ", __Verr5)
	}

	_CFpfN(" 839292 13 rece__dataRepack_c2 {%s} ", __Vdecode.String())
}
