package main

func (___Vug *_TudpGroupSt) IRun(___Vidx int) {
	switch ___Vidx {
	case 600101:
		if nil == ___Vug.ugCBinit {
			_FudpGroup__600101__main_init__default(___Vug)
		} else {
			___Vug.ugCBinit(___Vug)
		}
	case 600201:
		if nil == ___Vug.ugCBchInSend {
			_FudpGroup__600201__CHin_selecT_send__default(___Vug)
		} else {
			___Vug.ugCBchInSend(___Vug)
		}
	default:
		_FpfNex(" 838911 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

// _FudpNode__500101__main_init__default
// _FudpNode__500101yy3__receiveCallBack_default__randDecodeOut_noKeyWillDirect
// _TudpGroupSt
// _TudpNodeSt
//	ugCHreceLO  *chan _TudpNodeDataRece      // unCHreceLO    *chan _TudpNodeDataRece //
//	ugChTmpSendLI [](*(chan _TudpNodeDataSend)) // unCHsendI     chan _TudpNodeDataSend  //
func _FudpGroup__600101__main_init__default(___Vug *_TudpGroupSt) {

	//___Vug.ugCHuConnPortI = make(chan _TudpConnPort, 8)
	___Vug.ugCHSendI = make(chan _TudpNodeDataSend, 8)
	___Vug.ugNodeS = make([]_TudpNodeSt, ___Vug.ugAmount)
	___Vug.ugChTmpSendLI = make([](*(chan _TudpNodeDataSend)), ___Vug.ugAmount)

	if 0 == ___Vug.ugAmount {
		_FpfNex(" 834811 01 : why zero amounn ?")
	}

	__VunPortLen := len(___Vug.ugHostPortStr)
	__VunKeyLen := len(___Vug.ugRkeyLP)
	for __Vi := 0; __Vi < ___Vug.ugAmount; __Vi++ {
		__Vun := &(___Vug.ugNodeS[__Vi])
		__Vun.unIdx = __Vi
		__Vun.unName = _Spf("%s__%02d", ___Vug.ugName, __Vi)
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
		__Vun.unCHreceLO = ___Vug.ugCHreceLO
		___Vug.ugChTmpSendLI[__Vi] = &(__Vun.unCHsendI)
		_Frun(__Vun, 500101) // IRun // _FudpNode__500101__main_init__default
	}

	_Fsleep(_T1s)

	go _Frun(___Vug, 600201) // _FudpGroup__600201__CHin_selecT_send__default

	//_FpfN(" 834811 97 : exit.[%#v]", ___Vug)
	//_FpfNex(" 834811 98 : exit. ")
	//	for {
	//		_FpfNdb(" 834811 99 : udpGroup %s runing.... ", ___Vug.ugName)
	//		//_Fsleep_1s(_T100s)
	//		_Fsleep(_T1200s)
	//
	//	}
}
