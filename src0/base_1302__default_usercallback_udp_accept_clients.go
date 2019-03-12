package main

func (___VaccM *_TuAcceptClientMap) _FuserCallback__acceptClient01__default(___Vreq _TreqIneedToLogin, ___Vid [16]byte) {
	_FpfN(" 668191 01 : accepting %x :%v", ___Vid, ___Vreq)

	__VaccSt, __Vexist3 := ___VaccM.uMap[___Vid]
	if true == __Vexist3 {
		__VaccLen := len(__VaccSt)
		for __Vi0 := 0; __Vi0 < __VaccLen; __Vi0++ {
			__VaccSt[__Vi0].CexitAcc <- _FencJsonExit(" 668191 02 ", __Vi0)
		}
		var __Vexist4 bool = false
		for __Vi1 := 10; __Vi1 > 0; __Vi1-- {
			_Fsleep_1s()
			_, __Vexist4 = ___VaccM.uMap[___Vid]
			if false == __Vexist4 {
				break
			}
		}
		if true == __Vexist4 {
			_Fex1(" 668191 04 : unknow what happens: why not being delete ? ")
		}
		_Fex1(" 668191 05 debuging ")
	}

	_FpfN(" 668191 06 creat ... ")

	___VaccM.muxAcc.Lock()
	defer ___VaccM.muxAcc.Unlock()

	if ___VaccM.cntClient >= ___VaccM.maxClient {
		_FpfN(" 668191 07 maxClinet reach : %d / %d ", ___VaccM.cntClient, ___VaccM.maxClient)
		return
	}

	___VaccM.cntClient++

	for __Vi2 := ___VaccM.maxConnPerClient; __Vi2 > 0; __Vi2-- {
	}

} // _FuserCallback__acceptClient01__default
