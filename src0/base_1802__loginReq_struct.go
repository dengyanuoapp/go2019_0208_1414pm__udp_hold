package main

type _TloginReq struct {
	MeRand5  []byte // rand , 5 byte , not for verify / recgonize , but for debug only
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
	return _Spf("%x", *___Vb)
}

func (___Vlr *_TloginReq) String() string {
	__Vo := _Spf(
		//" %d %s %s me:%x %x to:%x %x tokenAB: %s,%s ",
		" rand5:%x ti:%d req:%s me:%s mid:%x,%x tid:%x,%x tokenAB: %x,%x ",
		___Vlr.MeRand5,
		___Vlr.MeTime,
		___Vlr.ReqStr,
		___Vlr.MeName,

		String5(&___Vlr.MeIdx128),
		String5(&___Vlr.MeSeq128),
		String5(&___Vlr.ToIdx128),
		String5(&___Vlr.ToSeq128),
		String5(&___Vlr.TokenL),
		String5(&___Vlr.TokenR))
	return __Vo
}
