package main

const (
	DDType__NULL = 121
	DDType__idle = 122
	DDType__fndn = 123
	DDType__s2c  = 124
	DDType__c2s  = 125
	DDType__End  = 126
)

type _TdataTran struct {
	DtRole     string
	DtMEidx128 []byte
	DtMYseq128 []byte
	DtTOidx128 []byte
	DtTOseq128 []byte
	DtTTokenD  []byte
	DtDDcmd    byte // DDType__idle
	DtDDoffset uint64
	DtDDbuf    []byte
}

func (___Vlr *_TdataTran) String() string {
	__Vo := _Spf(
		"%s mid:%s,%s tid:%s,%s tkD:%s Dcmd:%s off:%d B:%s",

		___Vlr.DtRole,
		String5s(&___Vlr.DtMEidx128),
		String5s(&___Vlr.DtMYseq128),
		String5s(&___Vlr.DtTOidx128),
		String5s(&___Vlr.DtTOseq128),
		String5s(&___Vlr.DtTTokenD),
		_FddCmdtype(___Vlr.DtDDcmd),
		___Vlr.DtDDoffset,
		String9s(&___Vlr.DtDDbuf))
	return __Vo
}

func _FddCmdtype(___Vtype byte) string {
	__Vso := _Spf("%d:", ___Vtype)
	switch ___Vtype {
	case DDType__NULL:
		__Vso += "DTnull"
	case DDType__idle:
		__Vso += "DTidle"
	case DDType__fndn:
		__Vso += "DTfsdn"
	case DDType__s2c:
		__Vso += "DTs2c"
	case DDType__c2s:
		__Vso += "DTc2s"
	case DDType__End:
		__Vso += "DTend"
	default:
		__Vso += "DTunknown"
	}

	return __Vso
}
