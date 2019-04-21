package main

func (___Vbnb *_TbyteNoteBuf) IRun(___Vidx int) {
	switch ___Vidx {
	case 1300101:
		if nil == ___Vbnb.bnbCBinit {
			go _FbyteNoteBuf__1300101x__init(___Vbnb)
		} else {
			go ___Vbnb.
				bnbCBinit(___Vbnb)
		}
	case 1300201:
		go ___Vbnb.
			_FbyteNoteBuf__1300201x__chan_rece__default()
	default:
		_FpfNex(" 834821 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FbyteNoteBuf__1300101x__init(___Vbnb *_TbyteNoteBuf) {
	___Vbnb.bnbCHinI = make(chan byte, 50)

	_Fsleep(_T1s)

	go _Frun(___Vbnb, 1300201) // _FbyteNoteBuf__1300201x__chan_rece__default
}
