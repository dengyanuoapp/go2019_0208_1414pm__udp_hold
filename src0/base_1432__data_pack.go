package main

// _TloginReq
func (___VuConnPort *_TudpConnPort) _FdataPack__101__udpConnPort(___VuTokenMe *[]byte, ___VoutBuf *[]byte) {
	__Vreq := _TloginReq{
		LgMeRand5:  _FgenRand_nByte__(5),
		LgMeTime:   _FtimeInt(),
		LgReqStr:   " loginS01genReplyTokenA ",
		LgMeName:   _VC.Name,
		LgMeIdx128: _VC.MyId128,
		LgMeSeq128: _VS.MySeq128,
		LgTokenL:   *___VuTokenMe,
		//LgToIdx128 []byte,
		//LgToSeq128 []byte,
	}
	if 0 != len(__Vreq.LgTokenR) {
		_FpfNex(" 138185 01 : why len is not ZERO ?")
	}

	//_FpfN(" 138185 03 : gen loginReq , trying to package to bytes {%s}", __Vreq.String())

	__Vencode := _Tencode{
		EnLogin: __Vreq,
	}
	__Vencode.
		_FdataPack__100__loginReq(LoadT__loginS01genReplyTokenA, ___VoutBuf)

	//_FpfN(" 138185 04 : gen (origin:%d) byte to buf:{%s}", len(*___VoutBuf), __Vreq.String())

	//return &__Vreq
}

// _TencodeX
func (___Vencode *_Tencode) _FdataPack__100__loginReq(___Vtype byte, ___VoutBuf *[]byte) {

	if nil == ___VoutBuf {
		_Fex(" 138186 01 : why NULL ?")
	}

	var (
		__Vb2   []byte
		__Verr2 error
	)

	switch ___Vtype { // _TloginReq
	case LoadT__loginS01genReplyTokenA,
		LoadT__loginS02genReplyTokenB,
		LoadT__loginS03acceptWithToken:

		___Vencode.EnLogin.LgRole = _VS.RoleName // _TloginReq

		__Vb2, __Verr2 =
			_FencGob__only(&___Vencode.EnLogin) // _TloginReq
	case LoadT__data_01_special, // 5
		LoadT__data_99_normal: // 6
		__Vb2, __Verr2 =
			_FencGob__only(&___Vencode.EnData) // _TdataTran

		___Vencode.EnData.DtRole = _VS.RoleName // _TloginReq _TencodeX

		//_CpfN(" 138186 02 : _Tencode encode (%s) .", ___Vencode.EnData.String())
		//_FpfN(" 138186 03 : _Tencode encode (%s) .", ___Vencode.EnData.String())
	default:
		_FpfN(" 138186 04 : encode error ?")
		*___VoutBuf = []byte{}
		return
	}

	if nil != __Verr2 {
		_FpfN(" 138186 05 : %v", __Verr2)
		*___VoutBuf = []byte{}
		return
	}

	__Vlen2 := len(__Vb2)
	//_FpfN(" 138186 06 : %x", __Vb2)
	//_FpfN(" 138186 07 : len %d: %v", __Vlen2, __Vb2)

	*___VoutBuf = make([]byte, __Vlen2+37)
	(*___VoutBuf)[0] = ___Vtype
	copy((*___VoutBuf)[1:], _VersionProtocol01)
	copy((*___VoutBuf)[37:], __Vb2)

	//_FpfN(" 138186 09 : len %d: %v", len(___VoutBuf), ___VoutBuf)
}
