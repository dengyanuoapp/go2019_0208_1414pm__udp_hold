package main

// _TudpGroupSt
// _TudpNodeSt
func _FudpGroup__650301__connPort__default(___Vug *_TudpGroupSt) {
	//_Fsleep(_T1s)
	//__Vlen := len(___Vug.ugCHuConnPortX)
	//_FpfNdb(" 838111 01 len : %d", __Vlen)
	for {
		select {
		case __VchPort := <-___Vug.ugCHuConnPortX:
			_FpfNdb(" 838111 03 : %v", __VchPort)
		}
	}
}
