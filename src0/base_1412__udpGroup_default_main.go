package main

func (___Vug *_TudpGroupSt) IRun(___Vidx int) {
	switch ___Vidx {
	case 150201:
		if nil == ___Vug.ugCBinit {
			_FudpGroup__150201__main_init__default(___Vug)
		} else {
			___Vug.ugCBinit(___Vug)
		}
	default:
		_FpfNex(" 838911 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

// _TudpGroupSt
func _FudpGroup__150201__main_init__default(___Vug *_TudpGroupSt) {
	___Vug.ugRkeyS = make([]*_Tkey256, ___Vug.ugAmount)
	___Vug.ugNodeS = make([]_TudpNodeSt, ___Vug.ugAmount)
	if nil == ___Vug.ugRKeyLP {
		for __Vi := 0; __Vi < ___Vug.ugAmount; __vi++ {
			___Vug.ugNodeS[__Vi].unName = ___Vug.ugName
			___Vug.ugNodeS[__Vi].unHostPortStr = ___Vug.ugHostPortStr
			//___Vug.ugRkeyS[__Vi]    = *___Vug.ugRKeyLP
		}
	} else {
		for __Vi := 0; __Vi < ___Vug.ugAmount; __vi++ {
		}
	}

	_FpfNex(" 834811 99 : exit.[%#v]", ___Vug)
}
