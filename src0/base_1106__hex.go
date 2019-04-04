package main

// dump hex
func _PpdL(___VmaxPrLen int, ___Vbuf *[]byte) {
	__Xlen := len(*___Vbuf)
	if __Xlen > ___VmaxPrLen {
		__Xlen = ___VmaxPrLen
	}
	_Ppf("(%d)", __Xlen)
	__Xlen--
	for __Vi := 0; __Vi <= __Xlen; __Vi++ {
		if __Vi == __Xlen {
			_Ppf("%02x", (*___Vbuf)[__Vi])
		} else {
			_Ppf("%02x ", (*___Vbuf)[__Vi])
		}
	}
} // _PpdL

func _FpfNhex(___VbyteIn *[]byte, ___VmaxPrLen int, ___Vfmt string, ___Vobj ...interface{}) {
	__Vl := len(*___VbyteIn)

	_Fpf(___Vfmt, ___Vobj...)
	_Ppf(" (%d,%d)", __Vl, ___VmaxPrLen)

	if __Vl < ___VmaxPrLen {
		___VmaxPrLen = __Vl
	}

	//_Ppf(" (%d) %x\n", __Vl, (*___VbyteIn)[0:___VmaxPrLen])
	__Vst := 0
	for (16 + __Vst) < ___VmaxPrLen {
		_Ppf(" %x", (*___VbyteIn)[__Vst:__Vst+16])
		__Vst += 16
	}
	_Ppf(" %x\n", (*___VbyteIn)[__Vst:___VmaxPrLen])

} // _FpfNhex

func _FpfhexlastN(___VbyteIn *[]byte, ___VmaxPrLen int, ___Vfmt string, ___Vobj ...interface{}) {
	__Vl := len(*___VbyteIn)

	_Fpf(___Vfmt, ___Vobj...)
	_Ppf(" (%d,%d)", __Vl, ___VmaxPrLen)

	if __Vl < ___VmaxPrLen {
		___VmaxPrLen = __Vl
	}
	__Vst := __Vl - ___VmaxPrLen

	//_Ppf(" (%d) %x\n", __Vl, (*___VbyteIn)[0:___VmaxPrLen])
	for (16 + __Vst) < __Vl {
		_Ppf(" %x", (*___VbyteIn)[__Vst:__Vst+16])
		__Vst += 16
	}
	_Ppf(" %x\n", (*___VbyteIn)[__Vst:__Vl])

} // _FpfhexlastN

//func _FcopyByte( &((*___VacceptTCP).Vbuf2) , &((*___VacceptTCP).Vbuf), (*___VacceptTCP).Vlen )
//(*___VacceptTCP).Vbuf2 = make([]byte , (*___VacceptTCP).Vlen ); copy( (*___VacceptTCP).Vbuf2 , (*___VacceptTCP).Vbuf )
func _FcopyByte(___dst *[]byte, ___src *[]byte, ___len int) {
	__len2 := cap((*___src))
	if ___len > __len2 {
		___len = __len2
	}
	(*___dst) = make([]byte, ___len)
	copy((*___dst), (*___src))
} // _FcopyByte

func _PpdLN(___VmaxPrLen int, ___Vbuf *[]byte) {
	_PpdL(___VmaxPrLen, ___Vbuf)
	_Pn()
} // _PpdLN
func _Fpd(___VmaxPrLen int, ___Vbuf *[]byte) {
	_Fph()
	_PpdL(___VmaxPrLen, ___Vbuf)
} // _Fpd
func _FpdN(___VmaxPrLen int, ___Vbuf *[]byte) {
	_Fpd(___VmaxPrLen, ___Vbuf)
	_Pn()
} // _FpdN

func _FmakeByte(___VbyteSlice []byte) []byte {
	__Vlen := len(___VbyteSlice)
	__Vbyte := make([]byte, __Vlen)
	copy(__Vbyte, ___VbyteSlice)
	return __Vbyte
} // _FmakeByte

func _FmakeByte16(___VbyteArr [16]byte) []byte {
	__Vbyte := make([]byte, 16)
	copy(__Vbyte, ___VbyteArr[:])
	return __Vbyte
} // _FmakeByte16

func _FmakeByte32(___VbyteArr [32]byte) []byte {
	__Vbyte := make([]byte, 32)
	copy(__Vbyte, ___VbyteArr[:])
	return __Vbyte
} // _FmakeByte32
