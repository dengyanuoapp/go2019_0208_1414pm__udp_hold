package main

/*
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
	//_FpfNdb(" 848231 01 ")
	defer ___Vun.unRmap.unrMux.Unlock()
	___Vun.unRmap.unrMux.Lock()

	//_FpfN(" 848231 03 las2 %v", ___Vun.unRmap.unrMapLas2)
	//_FpfN(" 848231 04 last %v", ___Vun.unRmap.unrMapLast)
	//_FpfN(" 848231 05 now  %v", ___Vun.unRmap.unrMapNow)

	___Vun.unRmap.unrMapLas2 = ___Vun.unRmap.unrMapLast
	___Vun.unRmap.unrMapLast = ___Vun.unRmap.unrMapNow
	___Vun.unRmap.unrMapNow = make(map[string]_TuNodeDataRcnt)

	//_FpfN(" 848231 06 las2 %v", ___Vun.unRmap.unrMapLas2)
	//_FpfN(" 848231 07 last %v", ___Vun.unRmap.unrMapLast)
	//_FpfN(" 848231 08 now  %v", ___Vun.unRmap.unrMapNow)

}

// replace the _FudpNode__500101yy3__receiveCallBack_default__randDecodeOut_noKeyWillDirect
func _FudpNode__540211z__receiveCallBack_withTimeGap(___Vun *_TudpNodeSt) {
	//_FpfNhex(&___Vun.unRbuf1500, 30, " 848236 01 rece %d", ___Vun.unRlen)

	__Vrece := _TudpNodeDataRece{
		UrrRemoteAddr: ___Vun.unRemoteAddr,
		UrrBuf:        ___Vun.unRbuf1500[:___Vun.unRlen],
	}
	__VrKey := __Vrece.UrrRemoteAddr.IP.String()

	//_FpfNhex(&___Vun.unRbuf1500, 30, " 848236 02 <%s>", __VrKey)
	//_FpfNhex(&__Vrece.UrrBuf, 30, " 848236 03 ")

	if "" == __VrKey || "<nil>" == __VrKey {
		_FpfN(" 848236 04 address error %v", __Vrece.UrrRemoteAddr)
		return
	}
	//_FpfNdb(" 848236 05 %s", __VrKey)

	__Vreply := false

	___Vun.unRmap.unrMux.Lock()

	__Vnow, __VokN := ___Vun.unRmap.unrMapNow[__VrKey]
	if __VokN {
		__Vnow.cnt++ // alreay exist ... so , skip
		_FpfN(" 848237 01 ")
	} else {
		___Vun.unRmap.unrMapNow[__VrKey] = _TuNodeDataRcnt{
			cnt: 1,
			urr: __Vrece,
		}
		__Vlast, __VokL := ___Vun.unRmap.unrMapLast[__VrKey]
		if __VokL { // map[string]_TuNodeDataRcnt
			if 1 == __Vlast.cnt {
				__Vreply = true
				//_FpfN(" 848237 03 ")
			} else {
				_FpfN(" 848237 04 ")
			}
		} else {
			__Vlas2, __Vok2 := ___Vun.unRmap.unrMapLas2[__VrKey]
			if __Vok2 { // map[string]_TuNodeDataRcnt
				if 1 == __Vlas2.cnt {
					//_FpfN(" 848237 06 ")
					__Vreply = true
				} else {
					_FpfN(" 848237 07 ")
				}
			} else {
				//_FpfN(" 848237 09 ")
			}
		}
		//_FpfN(" 848238 01 las2 %v", ___Vun.unRmap.unrMapLas2)
		//_FpfN(" 848238 02 last %v", ___Vun.unRmap.unrMapLast)
		//_FpfN(" 848238 03 now  %v", ___Vun.unRmap.unrMapNow)
	}

	___Vun.unRmap.unrMux.Unlock()

	//_FpfN(" 848238 06 :  __Vreply %t", __Vreply)

	if __Vreply {
		//(*___Vun.unCHreceByteLO) <- __Vrece
		//_FpfN(" 848238 09 custom_receive 02 checkok,start to loginIn")
		___Vun.
			_FudpNode__500101yy4__receiveCallBack_default__randDecodeOut_mustDecode(&__Vrece.UrrBuf)
	}
}
*/
