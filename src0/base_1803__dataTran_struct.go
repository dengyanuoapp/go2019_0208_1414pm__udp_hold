package main

type _TdataTran struct {
	meIdx128 []byte
	meSeq128 []byte
	toIdx128 []byte
	toSeq128 []byte
	tokenD   []byte
	Dcmd     byte
	Doffset  uint64
	Dbuf     []byte
}

func (___Vlr *_TdataTran) String() string {
	__Vo := _Pspf(
		" me:%x %x to:%x %x tokenD: %x cmd %d offset 0x%x %x",

		___Vlr.meIdx128,
		___Vlr.meSeq128,
		___Vlr.toIdx128,
		___Vlr.toSeq128,

		___Vlr.tokenD,

		___Vlr.Dcmd,
		___Vlr.Doffset,
		___Vlr.Dbuf[:10])
	return __Vo
}
