package main

func (___Vun *_TudpNodeSt) _FudpNode__540211x__gap() {
	if 0 == ___Vun.unLoopGap {
		return
	}

	___Vun.unCBrece = _FudpNode__540211z__receiveCallBack_withTimeGap

	if nil == ___Vun.unCBgap {
		go _FudpNode__540211y__gap_default(___Vun)
	} else {
		go ___Vun.unCBgap(___Vun)
	}
}

func _FudpNode__540211y__gap_default(___Vun *_TudpNodeSt) {
	for {
		_Fsleep(___Vun.unLoopGap)
		___Vun._FudpNode__540211yy__gap_default()
	}
}

func (___Vun *_TudpNodeSt) _FudpNode__540211yy__gap_default() {
	_FpfNdb(" 848231 01 ")
	___Vun.unRmap.unrMux.Lock()

	___Vun.unRmap.unrMapLast = ___Vun.unRmap.unrMapNow
	___Vun.unRmap.unrMapNow = make(map[string]_TuNodeDataRcnt)

	___Vun.unRmap.unrMux.Unlock()
}

// replace the _FudpNode__540201yy3__receiveCallBack_default__randDecodeOut
func _FudpNode__540211z__receiveCallBack_withTimeGap(___Vun *_TudpNodeSt) {
	_FpfNhex(&___Vun.unRbuf, 30, " 848232 01 rece %d", ___Vun.unRlen)

	__Vrece := _TudpNodeDataRece{
		urrRemoteAddr: ___Vun.unRemoteAddr,
		urrLen:        ___Vun.unRlen,
		urrBuf:        ___Vun.unRbuf[:___Vun.unRlen],
	}
	__VrKey := __Vrece.urrRemoteAddr.String()

	_FpfNhex(&___Vun.unRbuf, 30, " 848232 02 <%s>", __VrKey)
	_FpfNhex(&__Vrece.urrBuf, 30, " 848232 03 ")

	if "" == __VrKey {
		_FpfN(" 848232 04 address error %v", __Vrece.urrRemoteAddr)
		return
	}

	__Vreply := false

	___Vun.unRmap.unrMux.Lock()

	__Vnow, __VokN := ___Vun.unRmap.unrMapNow[__VrKey]
	if __VokN {
		__Vnow.cnt++ // alreay exist ... so , skip
	} else {
		___Vun.unRmap.unrMapNow[__VrKey] = _TuNodeDataRcnt{
			cnt: 1,
			urr: __Vrece,
		}
		__Vlast, __VokL := ___Vun.unRmap.unrMapLast[__VrKey]
		if __VokL { // map[string]_TuNodeDataRcnt
			if 1 == __Vlast.cnt {
				__Vreply = true
			}
		}
	}

	___Vun.unRmap.unrMux.Unlock()

	if __Vreply {
		(*___Vun.unCHreceLX) <- __Vrece
	}
}
