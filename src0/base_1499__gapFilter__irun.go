package main

import (
//"net"
)

func (___Vgf *_TgapFilter) IRun(___Vidx int) {
	switch ___Vidx {
	case 1200101:
		if nil == ___Vgf.gfCBinit {
			go ___Vgf.
				_FgapFilter__1200101x__init_default()
		} else {
			go ___Vgf.
				gfCBinit(___Vgf)
		}
	case 1200201:
		if nil == ___Vgf.gfCBdelay {
			go ___Vgf.
				_FgapFilter__1200201x__gapLoop_delay_timeOut__generator()
		} else {
			go ___Vgf.
				gfCBdelay(___Vgf)
		}
	case 1200301:
		if nil == ___Vgf.gfCBudpNodeDataRece {
			go ___Vgf.
				_FgapFilter__1200301x__Chan_rece()
		} else {
			go ___Vgf.
				gfCBudpNodeDataRece(___Vgf)
		}
	default:
		_FpfNex(" 839171 99 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func (___Vgf *_TgapFilter) _FgapFilter__1200101x__init_default() {

	if 0 == ___Vgf.gfGap {
		//_Fex(" 839171 01 : why Gap ZERO ? ")
		___Vgf.gfGap = _T10s
	}

	___Vgf.gfCHudpNodeDataReceI = make(chan _TudpNodeDataRece, 50)
	___Vgf.gfCHbyteI = make(chan []byte, 50)
	___Vgf.gfCHdelay = make(chan byte, 50)

	__VfCnt := 0

	if nil != ___Vgf.gfCHudpNodeDataReceLO {
		__VfCnt++
	}
	if nil != ___Vgf.gfCHbyteLO {
		__VfCnt++
	}

	if 0 == __VfCnt {
		_Fex(" 839171 03 : why Chan Amount ZERO ? ")
	}

	_Fsleep(_T100ms)

	go _Frun(___Vgf, 1200201) // _FgapFilter__1200201x__gapLoop_delay_timeOut__generator

	go _Frun(___Vgf, 1200301) // _FgapFilter__1200301x__Chan_rece
}

func (___Vgf *_TgapFilter) _FgapFilter__1200201x__gapLoop_delay_timeOut__generator() {
	for {
		_Fsleep(___Vgf.gfGap)
		___Vgf.gfCHdelay <- 1
		_FpfN(" 839171 05 len(%d) cap(%d) ", len(___Vgf.gfCHdelay), cap(___Vgf.gfCHdelay))
	}
}
