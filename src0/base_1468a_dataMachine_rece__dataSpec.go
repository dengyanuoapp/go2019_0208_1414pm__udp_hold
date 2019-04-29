package main

func (___Vdm *_TdataMachine) _FdataMachin__rece__dataSpec(___Vdec *_Tdecode) { // _TdecodeX
	if ___Vdec.DEtype == DDType__idle {
		__FpfN(" 839291 01 : _TdataMachine : LoadT__data_01_special , idle , ignore :{%s}", ___Vdec.String())
		return
	}
	_CFpfN(" 839291 02 : _TdataMachine : LoadT__data_01_special :{%s}", ___Vdec.String())
}
