package main

func (___Vuen *_TuEncode) IRun(___Vidx int) {
	switch ___Vidx {
	case 1100101:
		if nil == ___Vuen.enCB1100101init {
			_FuEncode__1100101__main_init__default(___Vuen)
		} else {
			___Vuen.enCB1100101init(___Vuen)
		}
	case 1100201:
		if nil == ___Vuen.enCB1100201packSend {
			_FuEncode__1100201x__packSend__default(___Vuen)
		} else {
			___Vuen.enCB1100201packSend(___Vuen)
		}
	default:
		_FpfNex(" 849191 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FuEncode__1100101__main_init__default(___Vuen *_TuEncode) {

	___Vuen.enCHencodeCmdI = make(chan _Tencode, 20)
	___Vuen.enCHencodeDataI = make(chan _Tencode, 20)

	_Fsleep(_T1s)

	go _Frun(___Vuen, 1100201) // _FuEncode__1100201x__packSend__default
}
