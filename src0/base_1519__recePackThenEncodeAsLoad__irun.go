package main

func (___Vpel *_TrecePackThenEncodeAsLoad) IRun(___Vidx int) {
	switch ___Vidx {
	case 1400101:
		if nil == ___Vpel.pelCBinit {
			go _FrecePackThenEncodeAsLoad__1400101x__init(___Vpel)
		} else {
			go ___Vpel.
				pelCBinit(___Vpel)
		}
	case 1400201:
		go ___Vpel.
			_FrecePackThenEncodeAsLoad__1400201x__chan_rece__default()
	case 1400301:
		go ___Vpel.
			_FrecePackThenEncodeAsLoad__1400301x__oututTimeGapGen()
	default:
		_FpfNex(" 638193 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FrecePackThenEncodeAsLoad__1400101x__init(___Vpel *_TrecePackThenEncodeAsLoad) {
	___Vpel.pelCHudpNodeDataReceBI = make(chan []byte, 50) // _TdataTran, 50)
	___Vpel.pelChOutputGenGap = make(chan byte, 50)        // every 1s gen output
	_Fsleep(_T1s)

	go _Frun(___Vpel, 1400201) // _FrecePackThenEncodeAsLoad__1400201x__chan_rece__default

	go _Frun(___Vpel, 1400301) // _FrecePackThenEncodeAsLoad__1400301x__oututTimeGapGen
}

func (___Vpel *_TrecePackThenEncodeAsLoad) _FrecePackThenEncodeAsLoad__1400301x__oututTimeGapGen() {

	if 0 == ___Vpel.pelGap {
		return
	}

	if 0 > ___Vpel.pelGap {
		___Vpel.pelGap = _T1s
	}

	for {
		_Fsleep(___Vpel.pelGap)
		___Vpel.pelChOutputGenGap <- 1
	}
}
