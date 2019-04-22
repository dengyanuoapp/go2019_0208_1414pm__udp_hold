package main

const (
	DDType__NULL = iota // 0
	DDType__idle        // 1
	DDType__fndn        // 2
	DDType__s2c         // 3
	DDType__c2s         // 4
	DDType__End         // 5
)

type _TdataTran struct {
	MEidx128 []byte
	MYseq128 []byte
	TOidx128 []byte
	TOseq128 []byte
	TTokenD  []byte
	DDcmd    byte
	DDoffset uint64
	DDbuf    []byte
}

func (___Vlr *_TdataTran) String() string {
	__Vo := _Spf(
		"mid:%s,%s tid:%s,%s tkD:%s Dcmd:%d off:%d B:%s",

		String5(&___Vlr.MEidx128),
		String5(&___Vlr.MYseq128),
		String5(&___Vlr.TOidx128),
		String5(&___Vlr.TOseq128),
		String5(&___Vlr.TTokenD),
		___Vlr.DDcmd,
		___Vlr.DDoffset,
		String9(&___Vlr.DDbuf))
	return __Vo
}
