package main

type _TtestLeng01 struct {
	I01 int64
	I02 int64
	S01 string
	S02 string
	B01 []byte
	B02 []byte
	BBB []byte
}

const _VbbbSize = 2048

type _TtestLeng02 struct {
	I01 int64
	I02 int64
	S01 string
	S02 string
	B01 [16]byte
	B02 [32]byte
	BBB [_VbbbSize]byte
}

func _FtestLen01() {
	var (
		__Vtl01 _TtestLeng01
		__Vtl02 _TtestLeng02
	)

	__Vtl01.I01 = _FgenRand_int64()
	__Vtl01.I02 = _FgenRand_int64()

	__Vtl01.S01 = _Spf("%x", _FgenRand_nByte__(10))
	__Vtl01.S02 = _Spf("%x", _FgenRand_nByte__(5))

	__Vtl01.B01 = _FgenRand_nByte__(16)
	__Vtl01.B02 = _FgenRand_nByte__(32)
	__Vtl01.BBB = _FgenRand_nByte__(_VbbbSize)

	__Vtl02.I01 = __Vtl01.I01
	__Vtl02.I02 = __Vtl01.I02

	__Vtl02.S01 = __Vtl01.S01
	__Vtl02.S02 = __Vtl01.S02

	copy(__Vtl02.B01[:], __Vtl01.B01)
	copy(__Vtl02.B02[:], __Vtl01.B02)
	copy(__Vtl02.BBB[:], __Vtl01.BBB)

	_FpfN("838118 01 : __Vtl01 %x", __Vtl01)
	_FpfN("838118 02 : __Vtl02 %x", __Vtl02)
	__Vb01, _ := _FencJson___(__Vtl01)
	__Vb02, _ := _FencJson___(__Vtl02)
	_FpfN("838118 03 : _FencJson %d %d", len(__Vb01), len(__Vb02))

	__Vb03, _ := _FencGob__only(__Vtl01)
	__Vb04, _ := _FencGob__only(__Vtl02)
	_FpfN("838118 06 : _FencGob %d %d", len(__Vb03), len(__Vb04))
	_FpfNhex(&__Vb03, 30, "838118 07 : __Vb03 ")
	_FpfNhex(&__Vb04, 30, "838118 08 : __Vb04 ")

	//__Vb05 := _FencBin1(__Vtl01)
	//__Vb06 := _FencBin2(__Vtl02)
	//_FpfN("_FencBin %d %d", len(__Vb05), len(__Vb06))

}
