package main

func (___Vgf *_TgapFilter) _FgapFilter__1200301x__Chan_rece() {
	for {
		select {
		case <-___Vgf.gfCHdelay: // _TudpNodeDataReceX
		case __Vur := <-___Vgf.gfCHudpNodeDataReceI: // _TudpNodeDataReceX
		case __Vbyte := <-___Vgf.gfCHbyteI: // []byte
		}
	}
}
