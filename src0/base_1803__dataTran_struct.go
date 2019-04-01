package main

type _TdataTran struct {
	MeIdx128 []byte
	MeSeq128 []byte
	ToIdx128 []byte
	ToSeq128 []byte
	TokenD   []byte
	Dcmd     byte
	Doffset  uint64
	Dbuf     []byte
}

func (___Vlr *_TdataTran) String() string {
	__Vo := _Pspf(
		" me:%x %x to:%x %x tokenD: %x cmd %d offset 0x%x %x",

		___Vlr.MeIdx128,
		___Vlr.MeSeq128,
		___Vlr.ToIdx128,
		___Vlr.ToSeq128,

		___Vlr.TokenD,

		___Vlr.Dcmd,
		___Vlr.Doffset,
		___Vlr.Dbuf[:10])
	return __Vo
}
