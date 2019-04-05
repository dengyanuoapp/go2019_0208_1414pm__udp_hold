package main

// _TloginReq
func (___VuConnPort *_TudpConnPort) _FdataPack__101__udpConnPort(___VuTokenMe *[]byte, ___VoutBuf *[]byte) {
	__Vreq := _TloginReq{
		MeTime:   _FtimeInt(),
		ReqStr:   " loginS01genReplyTokenA ",
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _VS.MySeq128,
		TokenL:   *___VuTokenMe,
		//ToIdx128 []byte,
		//ToSeq128 []byte,
	}
	if 0 != len(__Vreq.TokenR) {
		_FpfNex(" 138185 01 : why len is not ZERO ?")
	}
	_FpfN(" 138185 03 : gen loginReq , trying to package to bytes {%s}", __Vreq.String())
	__Vreq._FdataPack__100__loginReq(Cmd__loginS01genReplyTokenA, ___VoutBuf)
	//return &__Vreq
}

func (___Vreq *_TloginReq) _FdataPack__100__loginReq(___Vcmd byte, ___VoutBuf *[]byte) {

	__Vb2, __Verr2 := _FencGob__only(___Vreq)
	if nil != __Verr2 {
		_FpfN(" 138186 01 : %v", __Verr2)
		*___VoutBuf = []byte{}
		return
	}

	__Vlen2 := len(__Vb2)
	//_FpfN(" 138186 02 : %x", __Vb2)
	//_FpfN(" 138186 03 : len %d: %v", __Vlen2, __Vb2)

	*___VoutBuf = make([]byte, __Vlen2+37)
	(*___VoutBuf)[0] = ___Vcmd
	copy((*___VoutBuf)[1:], _VersionProtocol01)
	copy((*___VoutBuf)[37:], __Vb2)

	//_FpfN(" 138186 05 : len %d: %v", len(___VoutBuf), ___VoutBuf)
}
