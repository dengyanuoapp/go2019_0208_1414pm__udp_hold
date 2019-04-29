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
		_CFpfN(" 839291 02 : _TdataMachine : LoadT__data_01_special DDType__c2s underconstring :{%s}", ___Vdec.String())
	default:
		_CFpfN(" 839291 07 : _TdataMachine : LoadT__data_01_special unknown:{%s}", ___Vdec.String())
	}

	__FpfN(" 839291 09 : _TdataMachine : LoadT__data_01_special :{%s}", ___Vdec.String())
}
