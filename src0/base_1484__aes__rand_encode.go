package main

func (___Vaes *_Taes) _FrandPatAddToOutput() bool {
	__Vlen := len(___Vaes.aesBufOut)
	__VpatLen = _FtimeInt() % 16
	if 0 == __VpatLen {
		return
	}

	__Vt2 := make([]byte, __Vlen+__VpatLen)
	__Vt3 := _FgenRand_nByte__(__VpatLen)
	if __VpatLen != len(__Vt3) {
		___Vaes.aesErr = _Spf(" 138144 02 : pad add error : %d , %d", __VpatLen, len(__Vt3))
		return false
	}

	copy(__Vt2, ___Vaes.aesBufOut)
	copy(__Vt2[__Vlen:], __Vt3)

	___Vaes.aesBufOut = __Vt2

	return true
}
