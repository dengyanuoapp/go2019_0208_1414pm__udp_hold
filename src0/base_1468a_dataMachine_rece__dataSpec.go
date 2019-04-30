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
		__FpfN(" 839291 02 : _TdataMachine : LoadT__data_01_special decode from DDType__c2s' buf to origin-Tdecode :{%s}", ___Vdec.String())
		___Vdm._FdataMachin__rece__dataRepacked_c2s(___Vdec) // _TdecodeX
	default:
		_CFpfN(" 839291 07 : _TdataMachine : LoadT__data_01_special unknown:{%s}", ___Vdec.String())
	}

	__FpfN(" 839291 09 : _TdataMachine : LoadT__data_01_special :{%s}", ___Vdec.String())
}

func (___Vdm *_TdataMachine) _FdataMachin__rece__dataRepacked_c2s(___Vdec *_Tdecode) { // _TdecodeX
	__Vdecode := _Tdecode{}                                                 // _TdecodeX
	__Verr5 := _FdecGob___(" 839292 01 ", ___Vdec.DEdata.DDbuf, &__Vdecode) // _TdataTran
	if nil != __Verr5 {
		_CFpfN(" 839292 02 rece__dataRepacked_c2s error: {%v} ", __Verr5)
		return
	}

	if nil == ___Vdm.dmCHrepackDecodeC2sLO { // _TencodeX , get repacked-c2s decode as spec-data , then send to loginChecker
		_CFpfN(" 839292 05 rece__dataRepacked_c2s error: Why ___Vdm . dmCHrepackDecodeC2sLO NULL ?")
		return
	}

	(*(___Vdm.dmCHrepackDecodeC2sLO)) <- __Vdecode

	_CFpfN(" 839292 09 rece__dataRepacked_c2s repacked_C2s finished. {%s} ", __Vdecode.String())
}
