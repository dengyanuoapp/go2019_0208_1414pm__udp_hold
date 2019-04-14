package main

// _TudpNodeDataReceX
func (___Vgf *_TgapFilter) _FgapFilter__1200301x2__uData_rece(___Vur *_TudpNodeDataRece) {
	//_FpfN(" 381927 12 : gfCHudpNodeDataReceI : {%s}", ___Vur.String())

	//__Vk := ___Vur.UrrRemoteAddr.String()
	__Vk := ___Vur.UrrRemoteAddr.IP.String()

	if "" == __Vk {
		_FpfNonce(" 381927 13 : why blank string ?")
		return
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

func (___Vgf *_TgapFilter) _FgapFilter__1200301x3__Byte_rece(__VbyteIn *[]byte) {
	//_FpfN(" 381917 23 : gfCHbyteI :[%x]", __VbyteIn)

	__VunRece := _TudpNodeDataRece{}
	__Verr4 := _FdecGob___(" 381928 02 ", *__VbyteIn, &__VunRece)
	if nil != __Verr4 {
		_FpfN(" 381928 11 : why error <%v> ?", __Verr4)
		return
	}

	__Vk := __VunRece.UrrRemoteAddr.IP.String()

	if "" == __Vk {
		_FpfN(" 381928 13 : why blank string ?")
		return
	}

	__Vv, __Vok := ___Vgf.gfR.now_[__Vk]
	if false == __Vok {
		___Vgf.gfR.now_[__Vk] = _TgapFilter_ReceCnt{ // _TgapFilter_ReceStX
			cnt: 1,
			//unr: __Vur, // _TudpNodeDataReceX
			unb: (*__VbyteIn),
		}
		//_FpfN(" 381928 15 : first")
		return
	}

	___Vgf.gfR.now_[__Vk] = _TgapFilter_ReceCnt{ // _TgapFilter_ReceStX
		cnt: __Vv.cnt + 1,
		//unr: __Vur, // _TudpNodeDataReceX
		unb: (*__VbyteIn),
	}
}
