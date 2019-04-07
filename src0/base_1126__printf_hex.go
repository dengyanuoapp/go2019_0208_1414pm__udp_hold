package main

// dump hex
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

func _CpfNhex(___VbyteIn *[]byte, ___VmaxPrLen int, ___Vfmt string, ___Vobj ...interface{}) {
	var __Vs string
	__Vl := len(*___VbyteIn)

	__Vs += _Sph()
	__Vs += _Spf(___Vfmt, ___Vobj...)
	__Vs += _Spf(" (%d,%d)", __Vl, ___VmaxPrLen)

	if __Vl < ___VmaxPrLen {
		___VmaxPrLen = __Vl
	}

	//_Ppf(" (%d) %x\n", __Vl, (*___VbyteIn)[0:___VmaxPrLen])
	__Vst := 0
	for (16 + __Vst) < ___VmaxPrLen {
		__Vs += _Spf(" %x", (*___VbyteIn)[__Vst:__Vst+16])
		__Vst += 16
	}
	__Vs += _Spf(" %x\n", (*___VbyteIn)[__Vst:___VmaxPrLen])

} // _CpfNhex

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
