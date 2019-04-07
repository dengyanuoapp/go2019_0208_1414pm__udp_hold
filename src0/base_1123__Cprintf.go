package main

func _CpfS(___Vstr *string) {
	if nil == ___Vstr {
		return
	}

	if nil == _CHpr {
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

func _CpfN(___Vfmt string, ___Vpara ...interface{}) {
	__Vs := _Sph()
	__Vs += _Spf(___Vfmt+"\n", ___Vpara...)
	_CpfS(&__Vs)
} // _CpfN
