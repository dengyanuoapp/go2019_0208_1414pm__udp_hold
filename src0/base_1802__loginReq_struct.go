package main

type _TloginReq struct {
	LgRole     string
	LgMeRand5  []byte // rand , 5 byte , not for verify / recgonize , but for debug only
	LgMeTime   int
	LgReqStr   string
	LgMeName   string
	LgMeIdx128 []byte
	LgMeSeq128 []byte
	LgToIdx128 []byte
	LgToSeq128 []byte
	LgTokenL   []byte // token Local
	LgTokenR   []byte // token Remote
}

func _Fbyte2str(___Vb *[]byte) string {
	if (nil == ___Vb) || (0 == len(*___Vb)) {
		return ""
	}
	return _Spf("%x", *___Vb)
}

func (___Vlr *_TloginReq) String() string {
	__Vo := _Spf(
		//" %d %s %s me:%x %x to:%x %x tkAB: %s,%s ",
		"%s rand5:%x ti:%d req:%s mn:%s mid:%s,%s tid:%s,%s tkAB:%s,%s",
		___Vlr.LgRole,
		___Vlr.LgMeRand5,
		___Vlr.LgMeTime,
		___Vlr.LgReqStr,
		___Vlr.LgMeName,

		String5s(&___Vlr.LgMeIdx128),
		String5s(&___Vlr.LgMeSeq128),
		String5s(&___Vlr.LgToIdx128),
		String5s(&___Vlr.LgToSeq128),
		String5s(&___Vlr.LgTokenL),
		String5s(&___Vlr.LgTokenR))
	return __Vo
}
