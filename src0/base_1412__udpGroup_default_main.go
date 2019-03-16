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
	_FpfNex(" 834811 99 : exit.")
}
