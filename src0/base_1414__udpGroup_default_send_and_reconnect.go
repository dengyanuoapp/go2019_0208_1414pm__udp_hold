package main

// _TudpGroupSt
// _TudpNodeSt
// _TudpConnPort
func _FudpGroup__650301__connPort__default(___Vug *_TudpGroupSt) {
	//_Fsleep(_T1s)
	//__Vlen := len(___Vug.ugCHuConnPortX)
	//_FpfNdb(" 838111 01 len : %d", __Vlen)
	for {
		select {
		case __VchPort := <-___Vug.ugCHuConnPortX:
			//_FpfNdb(" 838111 03 : %s, %0x", __VchPort.Uri, __VchPort.K256)
			_FpfNdb(" 838111 04 : %s", __VchPort.Uri)
			___Vug._FudpGroup__650301z__send_using_rand_port(
				__VchPort._FudpGroup__650301x01__packData())
		}
	}
}

//	ugRANDidx     []int
//	ugRANDremain  int
func (___Vug *_TudpGroupSt) _FudpGroup__650301z__send_using_rand_port(___VbIn *[]byte) {
}
