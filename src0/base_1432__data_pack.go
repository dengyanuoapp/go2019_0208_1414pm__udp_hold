package main

// _TloginReq
func (___VuConnPort *_TudpConnPort) _FdataPack__101__udpConnPort(___VuTokenMe *[]byte, ___VoutBuf *[]byte) {
	__Vreq := _TloginReq{
		MeTime:   _FtimeInt(),
		ReqStr:   " step01__reqNewLogin ",
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _VS.meSeq128,
		TokenL:   *___VuTokenMe,
		//ToIdx128 []byte,
		//ToSeq128 []byte,
	}
	__Vreq._FdataPack__100__loginReq(Cmd__loginS02genReplyToken2, ___VoutBuf)
}

func (___Vreq *_TloginReq) _FdataPack__100__loginReq(___Vcmd byte, ___VoutBuf *[]byte) {

	__Vb2, __Verr2 := _FencGob__only(___Vreq)
	if nil != __Verr2 {
		_FpfN(" 387191 01 : %v", __Verr2)
		*___VoutBuf = []byte{}
		return
	}

	__Vlen2 := len(__Vb2)
	//_FpfN(" 381923 02 : %x", __Vb2)
	//_FpfN(" 381923 03 : len %d: %v", __Vlen2, __Vb2)

	*___VoutBuf = make([]byte, __Vlen2+37)
	(*___VoutBuf)[0] = ___Vcmd
	copy((*___VoutBuf)[1:], _VersionProtocol01)
	copy((*___VoutBuf)[37:], __Vb2)

	//_FpfN(" 381923 05 : len %d: %v", len(___VoutBuf), ___VoutBuf)
}
