package main

type _TloginReq struct {
	MeTime   int
	ReqStr   string
	MeName   string
	MeIdx128 []byte
	MeSeq128 []byte
	ToIdx128 []byte
	ToSeq128 []byte
	TokenL   []byte // token Local
	TokenR   []byte // token Remote
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

		___Vlr.MeIdx128[:8],
		___Vlr.MeSeq128[:8],
		___Vlr.ToIdx128[:8],
		___Vlr.ToSeq128[:8],

		//_Fbyte2str(&___Vlr.TokenL),
		//_Fbyte2str(&___Vlr.TokenR))
		___Vlr.TokenL[:8],
		___Vlr.TokenR[:8])
	return __Vo
}
