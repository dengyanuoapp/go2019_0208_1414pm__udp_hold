package main

func _CpfS(___Vstr *string) {
	if nil == ___Vstr {
		return
	}

	if nil == _CHpr {
		return
	}

	if 0 == cap(*_CHpr) {
		_FpfNonce(" 839181 01 why len ZERO ? len %d, cap %d", len(*_CHpr), cap(*_CHpr))
		return
	}

	__Vts := _TtcpNodeDataSend{
		//tnsId128: _FgenRand_nByte__(16),
		//tnsK256:  _FgenRand_nByte__(32),
		tnsBuf: []byte(*___Vstr),
	}

	__Vts.tnsLen = len(__Vts.tnsBuf)

	//_FpfN(" 839181 03 : %d : testing : push into send-chan:{%s}", _FtimeInt(), __Vts.String())

	(*_CHpr) <- __Vts
}

func _CpfNt(___Vfmt string, ___Vpara ...interface{}) {
	__Vs := "\n" + _Sph()
	__Vs += _Spf("%d:", _FtimeInt())
	__Vs += _Spf(___Vfmt+"\n", ___Vpara...)
	_CpfS(&__Vs)
}
func _CpfN(___Vfmt string, ___Vpara ...interface{}) {
	__Vs := _Sph()
	__Vs += _Spf(___Vfmt+"\n", ___Vpara...)
	_CpfS(&__Vs)
} // _CpfN

func ___CpfN(___Vfmt string, ___Vpara ...interface{}) {
}
func _NpfN(___Vfmt string, ___Vpara ...interface{}) {
}
