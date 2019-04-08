package main

func (___Vug *_TudpGroupSt) _FprKey() {
	__Vlen := len(___Vug.ugNodeS)
	if __Vlen != ___Vug.ugAmount {
		_FpfNonce(" 838378 01 : why not  equal ?")
	}

	__Vi := 0
	for __Vi < __Vlen {
		___Vug.ugNodeS[__Vi]._FprKey()
		__Vi++
	}
}
