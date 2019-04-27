package main

import (
	"sync"
)

type _TcmdMap struct {
	M   map[[16]byte]_Tdecode // key by 16byte : MeIdx128
	mux sync.Mutex
}

func (___Vcm *_TcmdMap) String() string {
	//	if 0 == len((*___Vcm).M) {
	//		return ""
	//	}

	var __VstrO string

	for ___Vk := range (*___Vcm).M {
		__VstrO += _Spf("%x ", ___Vk[:5])
	}
	return __VstrO
}

func String5(___Vb *[]byte) string {
	if (nil == ___Vb) || (0 == len(*___Vb)) {
		return ""
	}
	if len(*___Vb) < 5 {
		return _Spf("%x", *___Vb)
	}
	__Vt := *___Vb
	return _Spf("%x", __Vt[:5])
}
func String9(___Vb *[]byte) string {
	if (nil == ___Vb) || (0 == len(*___Vb)) {
		return ""
	}
	if len(*___Vb) < 9 {
		return _Spf("(%d){%x}[%x]", len(*___Vb), _Fmd5__5x(___Vb), *___Vb)
	}
	__Vt := *___Vb
	return _Spf("(%d){%x}[%x]", len(*___Vb), _Fmd5__5x(___Vb), __Vt[:9])
}
