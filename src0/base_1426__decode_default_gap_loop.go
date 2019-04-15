package main

func _FudpDecode__700301x__loop__default(___Vutm *_TuDecode) {
	_FpfNdb(" 388194 01 : gap main loop start ")
	if 0 == ___Vutm.uDeGapX {
		_FpfNdb(" 388194 02 : Warn : why ZERO gap ? loop disabled")
		return
	}

	for {
		_FpfNdb(" 388194 03 : gap main loop ing ")
		_Fsleep(___Vutm.uDeGapX)
	}
}
