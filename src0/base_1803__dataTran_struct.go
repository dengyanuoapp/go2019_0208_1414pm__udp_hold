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
		"mid:%s,%s tid:%s,%s tkD:%s cmd:%d off:%d B:%s",

		String5(&___Vlr.meIdx128),
		String5(&___Vlr.MySeq128),
		String5(&___Vlr.toIdx128),
		String5(&___Vlr.toSeq128),
		String5(&___Vlr.tokenD),
		___Vlr.Dcmd,
		___Vlr.Doffset,
		String9(&___Vlr.Dbuf))
	return __Vo
}
