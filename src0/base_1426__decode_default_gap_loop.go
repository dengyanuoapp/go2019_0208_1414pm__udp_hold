package main

func _FudpDecode__700301x__loop__default(___Vutm *_TuDecode) {
	_FpfNdb(" 388194 01 : gap main loop start ")
	if 0 == ___Vutm.uTmGapX {
		_FpfN(" 388194 02 : Warn : why ZERO gap ? ")
		return
	}

	for {
		_FpfNdb(" 388194 02 : gap main loop ing ")
		_Fsleep(___Vutm.uTmGapX)
	}
}