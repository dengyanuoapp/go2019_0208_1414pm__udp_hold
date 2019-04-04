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
		__VstrO += _Pspf("%x ", ___Vk[:5])
	}
	return __VstrO
}

func String5(___Vb *[]byte) string {
	if (nil == ___Vb) || (0 == len(*___Vb)) {
		return ""
	}
	if len(*___Vb) < 5 {
		return _Pspf("%x", *___Vb)
	}
	return _Pspf("%x", (*___Vb)[:5])
}
