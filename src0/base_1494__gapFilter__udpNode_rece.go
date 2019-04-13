package main

// _TudpNodeDataReceX
func (___Vgf *_TgapFilter) _FgapFilter__1200301x2__uData_rece(___Vur *_TudpNodeDataRece) {
	//_FpfN(" 381927 12 : gfCHudpNodeDataReceI : {%s}", ___Vur.String())

	//__Vk := ___Vur.UrrRemoteAddr.String()
	__Vk := ___Vur.UrrRemoteAddr.IP.String()

	if "" == __Vk {
		//_FpfN(" 381927 13 : why blank string ?")
	}

	__Vv, __Vok := ___Vgf.gfR.now_[__Vk]
	if false == __Vok {
		___Vgf.gfR.now_[__Vk] = _TgapFilter_ReceCnt{ // _TgapFilter_ReceStX
			cnt: 1,
			unr: (*___Vur), // _TudpNodeDataReceX
		}
		//_FpfN(" 381927 15 : first")
		return
	}

	___Vgf.gfR.now_[__Vk] = _TgapFilter_ReceCnt{ // _TgapFilter_ReceStX
		cnt: __Vv.cnt + 1,
		unr: (*___Vur), // _TudpNodeDataReceX
	}
	//_FpfN(" 381927 17 : cnt %d", ___Vgf.gfR.now_[__Vk].cnt)
}
