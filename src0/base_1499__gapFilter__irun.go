package main

import (
//"net"
)

func (___Vgf *_TgapFilter) IRun(___Vidx int) {
	switch ___Vidx {
	case 1200101:
		if nil == ___Vgf.gfCBinit {
			go _FgapFilter__1200101x__init_default(___Vgf)
		} else {
			go ___Vgf.gfCBinit(___Vgf)
		}
	case 1200201:
		if nil == ___Vgf.gfCBdelay {
			go _FgapFilter__1200201x__gapLoop_delay_timeOut__generator(___Vgf)
		} else {
			go ___Vgf.gfCBdelay(___Vgf)
		}
	case 1200301:
		if nil == ___Vgf.gfCBudpNodeDataRece {
			go _FgapFilter__1200301x__udpNodeDataRece(___Vgf)
		} else {
			go ___Vgf.gfCBudpNodeDataRece(___Vgf)
		}
	case 1200302:
		if nil == ___Vgf.gfCBbyte {
			go _FgapFilter__1200302x__byte(___Vgf)
		} else {
			go ___Vgf.gfCBbyte(___Vgf)
		}
	default:
		_FpfNex(" 839171 99 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FgapFilter__1200101x__init_default(___Vgf *_TgapFilter) {

	if 0 == ___Vgf.gfGap {
		_Fex(" 839171 01 : why Gap ZERO ? ")
	}
	__VfCnt := 0

	if nil != ___Vgf.gfCHudpNodeDataReceLO {
		___Vgf.gfCHudpNodeDataReceI = make(chan _TudpNodeDataRece, 50)
		__VfCnt++
	}
	if nil != ___Vgf.gfCHbyteLO {
		___Vgf.gfCHbyteI = make(chan []byte, 50)
		__VfCnt++
	}

	if 0 == __VfCnt {
		_Fex(" 839171 03 : why Chan Amount ZERO ? ")
	}

	___Vgf.gfCHdelay = make(chan byte, 50)

	_Fsleep(_T100ms)

	go _Frun(___Vgf, 1200201) // _FgapFilter__1200201x__gapLoop_delay_timeOut__generator

	if nil != ___Vgf.gfCHudpNodeDataReceLO {
		go _Frun(___Vgf, 1200301) // _FgapFilter__1200301x__udpNodeDataRece
	}
	if nil != ___Vgf.gfCHbyteLO {
		go _Frun(___Vgf, 1200302) // _FgapFilter__1200302x__byte
	}
}

func _FgapFilter__1200201x__gapLoop_delay_timeOut__generator(___Vgf *_TgapFilter) {
	for {
		_Fsleep(___Vgf.gfGap)
		___Vgf.gfCHdelay <- 1
	}
}
func _FgapFilter__1200301x__udpNodeDataRece(___Vgf *_TgapFilter) {
}
func _FgapFilter__1200302x__byte(___Vgf *_TgapFilter) {
}
