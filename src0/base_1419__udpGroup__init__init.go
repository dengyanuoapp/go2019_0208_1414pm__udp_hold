package main

func (___Vug *_TudpGroupSt) IRun(___Vidx int) {
	switch ___Vidx {
	case 650201:
		if nil == ___Vug.ugCBinit {
			_FudpGroup__650201__main_init__default(___Vug)
		} else {
			___Vug.ugCBinit(___Vug)
		}
	case 650301:
		if nil == ___Vug.ugCBportConn {
			_FudpGroup__650301__connPort__default(___Vug)
		} else {
			___Vug.ugCBportConn(___Vug)
		}
	default:
		_FpfNex(" 838911 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

// _FudpNode__540201__main_init__default
// _FudpNode__540201yyy__receiveCallBack_default__randDecodeOut
// _TudpGroupSt
// _TudpNodeSt
func _FudpGroup__650201__main_init__default(___Vug *_TudpGroupSt) {

	go _Frun(___Vug, 650301) //

	___Vug.ugNodeS = make([]_TudpNodeSt, ___Vug.ugAmount)
	__VunPortLen := len(___Vug.ugHostPortStr)
	__VunKeyLen := len(___Vug.ugRkeyLP)
	for __Vi := 0; __Vi < ___Vug.ugAmount; __Vi++ {
		__Vun := &(___Vug.ugNodeS[__Vi])
		__Vun.unIdx = __Vi
		__Vun.unName = _Pspf("%s__%02d", ___Vug.ugName, __Vi)
		if __Vi >= __VunPortLen {
			__Vun.unHostPortStr = ":0"
		} else {
			__Vun.unHostPortStr = ___Vug.ugHostPortStr[__Vi]
		}
		if __Vi >= __VunKeyLen {
			__Vb := _FgenRand_nByte__(32)
			__Vun.unRKeyLP = &__Vb
		} else {
			__Vun.unRKeyLP = ___Vug.ugRkeyLP[__Vi]
		}
	}

	//_FpfN(" 834811 97 : exit.[%#v]", ___Vug)
	//_FpfNex(" 834811 98 : exit. ")
	for {
		_FpfNdb(" 834811 99 : udpGroup %s runing.... ", ___Vug.ugName)
		//_Fsleep_1s(_T100s)
		_Fsleep(_T1200s)

	}
}
