package main

func (___Vacc *_TuAcceptClientMap) _FuserCallback__acceptClient01__default(___Vreq _TreqIneedToLogin, ___Vid [16]byte) {
	_FpfN(" 668191 01 : accepting %x :%v", ___Vid, ___Vreq)

	_Fex1(" 83891918383 debuging ")

	for {
		__VaccMap, __Vexist3 := ___Vacc.uMap[___Vid]
		if true == __Vexist3 {
			__VaccMap.CexitAcc <- _FencJsonExit(" 668191 03 ", ___Vreq)
			_Fsleep_1s()
		} else {
			break // don't exist , create direct
		}
	}
} // _FuserCallback__acceptClient01__default
