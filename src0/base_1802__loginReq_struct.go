package main

type _TloginReq struct {
	MeTime   int
	ReqStr   string
	MeName   string
	MeIdx128 []byte
	MeSeq128 []byte
	ToIdx128 []byte
	ToSeq128 []byte
	TokenA   []byte
	TokenB   []byte
}

func _Fbyte2str(___Vb *[]byte) string {
	if (nil == ___Vb) || (0 == len(*___Vb)) {
		return ""
	}
	return _Pspf("%x", *___Vb)
}

func (___Vlr *_TloginReq) String() string {
	__Vo := _Pspf(
		//" %d %s %s me:%x %x to:%x %x tokenAB: %s,%s ",
		" %d %s %s me:%x %x to:%x %x tokenAB: %x,%x ",
		___Vlr.MeTime,
		___Vlr.ReqStr,
		___Vlr.MeName,

		___Vlr.MeIdx128,
		___Vlr.MeSeq128,
		___Vlr.ToIdx128,
		___Vlr.ToSeq128,

		//_Fbyte2str(&___Vlr.TokenA),
		//_Fbyte2str(&___Vlr.TokenB))
		___Vlr.TokenA,
		___Vlr.TokenB)
	return __Vo
}
