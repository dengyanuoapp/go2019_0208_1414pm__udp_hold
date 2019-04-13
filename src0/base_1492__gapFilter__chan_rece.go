package main

import "sync"

var ___VgapFilter__1200301x__mux sync.Mutex

func (___Vgf *_TgapFilter) _FgapFilter__1200301x__Chan_rece() {
	for {
		select {
		case <-___Vgf.gfCHdelay: // _TudpNodeDataReceX
			___VgapFilter__1200301x__mux.Lock()
			___Vgf.
				_FgapFilter__1200301x1__uDataSwap()

		case __Vur := <-___Vgf.gfCHudpNodeDataReceI: // _TudpNodeDataReceX
			___VgapFilter__1200301x__mux.Lock()

			___Vgf.
				_FgapFilter__1200301x2__uData_rece(&__Vur)

		case __Vbyte := <-___Vgf.gfCHbyteI: // []byte
			___VgapFilter__1200301x__mux.Lock()
			___Vgf.
				_FgapFilter__1200301x3__Byte_rece(&__Vbyte)
		} // end Select
		___VgapFilter__1200301x__mux.Unlock()
	} // end For
} // _TgapFilter

func (___Vgf *_TgapFilter) _FgapFilter__1200301x3__Byte_rece(__Vbyte *[]byte) {
	_FpfN(" 381917 23 : gfCHbyteI :[%x]", __Vbyte)

}

func (___Vgf *_TgapFilter) _FgapFilter__1200301x1__uDataSwap() {
	// _FpfN(" 381917 01 : gfCHdelay ")
	for __Vk, __Vv := range ___Vgf.gfR.now_ {
	}

	___Vgf.gfR.las2 = ___Vgf.gfR.last
	___Vgf.gfR.last = ___Vgf.gfR.now_
}
