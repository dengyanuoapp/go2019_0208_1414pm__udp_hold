package main

type _TdataTran struct {
	meIdx128 []byte
	MySeq128 []byte
	toIdx128 []byte
	toSeq128 []byte
	tokenD   []byte
	Dcmd     byte
	Doffset  uint64
	Dbuf     []byte
}

func (___Vlr *_TdataTran) String() string {
	__Vo := _Spf(
		" mid:%x,%x tid:%x,%x tkD:%x cmd:%d off:%x B:%s",

		___Vlr.meIdx128[:5],
		___Vlr.MySeq128[:5],
		___Vlr.toIdx128[:5],
		___Vlr.toSeq128[:5],

		___Vlr.tokenD[:5],

		___Vlr.Dcmd,
		___Vlr.Doffset,
		String5(&___Vlr.Dbuf))
	return __Vo
}
