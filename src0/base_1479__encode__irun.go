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
		if nil == ___Vuen.enCB1100201chanRece {
			_FuEncode__1100201x__chanRece__default(___Vuen)
		} else {
			___Vuen.enCB1100201chanRece(___Vuen)
		}
	default:
		_FpfNex(" 849191 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FuEncode__1100101__main_init__default(___Vuen *_TuEncode) {

	___Vuen.enCHencodeCkI = make(chan _Tencode, 50)
	___Vuen.enCHencodeDataCommI = make(chan _Tencode, 50)

	_Fsleep(_T1s)

	go _Frun(___Vuen, 1100201) // _FuEncode__1100201x__chanRece__default
}
