package main

const (
	Cmd__NULL = iota
	Cmd__loginReqTryNoToken
	Cmd__loginTmpToken
	Cmd__loginReqWithToken
	Cmd__data
)

var (
	_VersionProtocol01 = []byte{0x83, 0x20, 0x71, 0xc8}
)

type _TdataPack_991 struct {
	V [4]byte // version
	C byte    // cmd
	D []byte  // gob.package
}

// _TreqIneedToLogin
func (___VuConnPort *_TudpConnPort) _FdataPack__101__udpConnPort() *[]byte {
	__Vreq := _TreqIneedToLogin{
		MeTime:   _FtimeI64(),
		ReqStr:   "try_to_login01",
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _VS.meSeq128,
		//ToIdx128 []byte,
		//ToSeq128 []byte,
	}

	__Vb2, __Verr2 := _FencGob(&__Vreq)
	if nil != __Verr2 {
		_FpfN(" 381923 01 : %v", __Verr2)
		return nil
	}

	__Vlen2 := len(__Vb2)
	//_FpfN(" 381923 02 : %x", __Vb2)
	//_FpfN(" 381923 03 : len %d: %v", __Vlen2, __Vb2)

	__VbOut := make([]byte, __Vlen2+5)
	__VbOut[0] = Cmd__loginReqTryNoToken
	copy(__VbOut[1:], _VersionProtocol01)
	copy(__VbOut[5:], __Vb2)

	//_FpfN(" 381923 05 : len %d: %v", len(__VbOut), __VbOut)

	return &__VbOut
}
