package main

func (___Vgf *_TgapFilter) _FgapFilter__1200301x1__uDataSwap() {

	if nil == ___Vgf.gfCHudpNodeDataReceLO && nil == ___Vgf.gfCHbyteLO {
		_FpfNonce(" 381917 01 : gfCHdelay out-chan NULL")
		return
	}

	//_CpfN(" 381917 02 : ___Vgf.gfR.now_ {%#v} ", ___Vgf.gfR.now_)
	//_CpfN(" 381917 03 : ___Vgf.gfR.last {%#v} ", ___Vgf.gfR.last)
	//_CpfN(" 381917 04 : ___Vgf.gfR.las2 {%#v} ", ___Vgf.gfR.las2)

	//_FpfN(" 381917 05 ")
	//_CpfN(" 381917 06 ")
	for __Vk, __Vv := range ___Vgf.gfR.now_ { // _TgapFilter_ReceCntX
		//_FpfN(" 381917 07 ")
		//_CpfN(" 381917 08 ")
		if 1 <= __Vv.cnt && 3 >= __Vv.cnt {
			//_FpfN(" 381917 09 ")
			//_CpfN(" 381917 10 ")
			__Vcnt := 0
			__Vv2, __Vok2 := ___Vgf.gfR.last[__Vk]
			__Vv3, __Vok3 := ___Vgf.gfR.las2[__Vk]
			if __Vok2 {
				__Vcnt = __Vv2.cnt
				//_FpfN(" 381926 01 %d , %d, %d", __Vcnt, __Vv2.cnt, __Vv3.cnt)
				//_CpfN(" 381926 02 %d , %d, %d", __Vcnt, __Vv2.cnt, __Vv3.cnt)
			}
			if __Vok3 {
				__Vcnt += __Vv3.cnt
				//_FpfN(" 381926 03 %d , %d, %d", __Vcnt, __Vv2.cnt, __Vv3.cnt)
				//_CpfN(" 381926 04 %d , %d, %d", __Vcnt, __Vv2.cnt, __Vv3.cnt)
			}
			if 1 <= __Vcnt && 3 >= __Vcnt {
				//_FpfN(" 381926 05 %d , %d, %d", __Vcnt, __Vv2.cnt, __Vv3.cnt)
				//_CpfN(" 381926 06 %d , %d, %d", __Vcnt, __Vv2.cnt, __Vv3.cnt)
				if nil != ___Vgf.gfCHudpNodeDataReceLO {
					if 0 == len(__Vv.unb) {
						(*___Vgf.gfCHudpNodeDataReceLO) <- __Vv.unr
					} else {
						_FpfNonce(" 381926 07 error input / out : type error ")
					}
				} else {
					if 0 == len(__Vv.unb) {
						_FpfNonce(" 381926 08 error input / out : type error ")
					} else {
						(*___Vgf.gfCHbyteLO) <- __Vv.unb
					}
				}
			} else {
				//_FpfN(" 381926 09 %d , %d, %d", __Vcnt, __Vv2.cnt, __Vv3.cnt)
				//_CpfN(" 381926 10 %d , %d, %d", __Vcnt, __Vv2.cnt, __Vv3.cnt)
			}
		}
	}

	___Vgf.gfR.las2 = ___Vgf.gfR.last
	___Vgf.gfR.last = ___Vgf.gfR.now_
	___Vgf.gfR.now_ = make(map[string]_TgapFilter_ReceCnt)

	//_FpfN(" 381926 09 ")

	//_CpfN(" 381926 12 : ___Vgf.gfR.now_ {%#v} ", ___Vgf.gfR.now_)
	//_CpfN(" 381926 13 : ___Vgf.gfR.last {%#v} ", ___Vgf.gfR.last)
	//_CpfN(" 381926 14 : ___Vgf.gfR.las2 {%#v} ", ___Vgf.gfR.las2)
}
