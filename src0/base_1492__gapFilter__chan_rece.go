package main

import "sync"

var ___VgapFilter__1200301x__mux sync.Mutex

func (___Vgf *_TgapFilter) _FgapFilter__1200301x__Chan_rece() {
	for {
		select {
		case <-___Vgf.gfCHdelay: // _TudpNodeDataReceX
			___VgapFilter__1200301x__mux.Lock()
		// _FpfN(" 381917 01 : gfCHdelay ")
		case __Vur := <-___Vgf.gfCHudpNodeDataReceI: // _TudpNodeDataReceX
			___VgapFilter__1200301x__mux.Lock()
			_FpfN(" 381917 12 : gfCHudpNodeDataReceI : {%s}", __Vur.String())
		case __Vbyte := <-___Vgf.gfCHbyteI: // []byte
			___VgapFilter__1200301x__mux.Lock()
			_FpfN(" 381917 23 : gfCHbyteI :[%x]", __Vbyte)
		} // end Select
		___VgapFilter__1200301x__mux.Unlock()
	} // end For
}
