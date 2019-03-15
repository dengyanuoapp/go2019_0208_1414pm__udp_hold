package main

import (
//"net"
)

// _TudpNodeSt
// _Fhandle_u01x__udpListen_Udp__read_main_top__default
func (___Vun *_TudpNodeSt) IRun(___Vidx int) {
	switch ___Vidx {
	case 140201:
		_FudpNode__140201__main_init__default(___Vun)
	default:
		_FpfNex(" 937191 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

func _FudpNode__140201__main_init__default(___Vun *_TudpNodeSt) {

	___Vun.unRbuf = make([]byte, 1500)
	_FpfN(" 918381 01 : %x ", ___Vun.unPwIn256LP)

	if nil == ___Vun.unPwIn256LP {
		copy(___Vun.unPwIn256B, _FgenRand_nByte__(32))
		copy(___Vun.unPwIn256F[:], ___Vun.unPwIn256B)
	} else {
		copy(___Vun.unPwIn256B, *___Vun.unPwIn256LP)
		copy(___Vun.unPwIn256F[:], ___Vun.unPwIn256B)
	}
	_FpfN(" 918381 02 : LP %x, B %x, F %x ", ___Vun.unPwIn256LP, ___Vun.unPwIn256B, ___Vun.unPwIn256F)

	___Vun._FudpNode__140201x__listen()

	go ___Vun._FudpNode__140201y__receive()

	go ___Vun._FudpNode__140201z__send()

	//_Fex1(" 918381 09 ")
}
