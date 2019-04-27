package main

import (
//"bytes"
//"net"
)

//func (__Vundr *_TudpNodeDataRece) _FdataPack__dePack_endecode__from_udpNodeDataRece(___Vencode *_Tencode) {
//	_FdataPack__deGob__encode(___Vencode, __Vundr.UrrBuf)
//}

func _FdataPack__deGob__encode(___Vencode *_Tencode, ___Vbuf []byte) {
	if nil == ___Vencode {
		return
	}

	_CFpfN(" 387194 001 _FdataPack__deGob__encode <%v> <%s>", ___Vencode, String9(&___Vbuf))

	(*___Vencode) = _Tencode{}

	_CFpfN(" 387194 002 _FdataPack__deGob__encode <%v>", ___Vencode)

	__Verr2 :=
		_FdecGob___(" 387194 01 ", ___Vbuf, ___Vencode)
	if nil != __Verr2 {
		_CFpfN(" 387194 02 decodeGob error <%v>", __Verr2)
	}
	_CFpfN(" 387194 03 buf{%s} , enc{%s}", String9(&___Vbuf), ___Vencode.String())
}
