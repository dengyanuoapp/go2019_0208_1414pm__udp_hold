package main

// _TudpNodeSt
func (___Vun *_TudpNodeSt) IRun(___Vidx int) {
	switch ___Vidx {
	case 140201:
		_FsrvGroup__140201__main_top__default(___Vun)
	default:
		_FpfNex(" 937191 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FsrvGroup__140201__main_top__default(___Vun *_TudpNodeSt) {
	if nil != ___Vun.unPwIn256LP {
		copy(___Vun.unPwIn256[:], *___Vun.unPwIn256LP)
	}
	_FpfN(" 918381 01 : %x ", ___Vun.unPwIn256)
}
