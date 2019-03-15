package main

type _TtestLeng01 struct {
	I01   int64
	I02   int64
	S01   string
	S02   string
	B01   []byte
	B02   []byte
	B1024 []byte
}
type _TtestLeng02 struct {
	I01   int64
	I02   int64
	S01   string
	S02   string
	B01   [16]byte
	B02   [32]byte
	B1024 [1024]byte
}

func _FtestLen01() {
	var (
		__Vtl01 _TtestLeng01
		__Vtl02 _TtestLeng02
	)

	__Vtl01.I01 = _FgenRand_int64()
	__Vtl01.I02 = _FgenRand_int64()

	__Vtl01.S01 = _Pspf("%x", _FgenRand_nByte__(10))
	__Vtl01.S02 = _Pspf("%x", _FgenRand_nByte__(5))

	__Vtl01.B01 = _FgenRand_nByte__(16)
	__Vtl01.B02 = _FgenRand_nByte__(32)
	__Vtl01.B1024 = _FgenRand_nByte__(1024)

	__Vtl02.I01 = __Vtl01.I01
	__Vtl02.I02 = __Vtl01.I02

	__Vtl02.S01 = __Vtl01.S01
	__Vtl02.S02 = __Vtl01.S02

	copy(__Vtl02.B01[:], __Vtl01.B01)
	copy(__Vtl02.B02[:], __Vtl01.B02)
	copy(__Vtl02.B1024[:], __Vtl01.B1024)

	_FpfN("__Vtl01 %v", __Vtl01)
	_FpfN("__Vtl02 %v", __Vtl02)
}
