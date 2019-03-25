package main

const (
	Cmd__NULL = iota
	Cmd__idle
	Cmd__data
	Cmd__loginS1ReqTryNoToken
	Cmd__loginS2ReplyTmpToken
	Cmd__loginS3ReqWithToken
)

var (
	_VersionProtocol01 = []byte{0x83, 0x20, 0x71, 0xc8}
)

/*
type _TdataPack_991 struct {
	C byte    // cmd
	V [4]byte // version
	D []byte  // gob.package
}
*/

// _TreqIneedToLogin
func (___VuConnPort *_TudpConnPort) _FdataPack__101__udpConnPort(___VoutBuf *[]byte) {
	__Vreq := _TreqIneedToLogin{
		MeTime:   _FtimeI64(),
		ReqStr:   "try_to_login01",
		MeName:   _VC.Name,
		MeIdx128: _VC.MyId128,
		MeSeq128: _VS.meSeq128,
		//ToIdx128 []byte,
		//ToSeq128 []byte,
	}

	__Vb2, __Verr2 := _FencGob__only(&__Vreq)
	if nil != __Verr2 {
		_FpfN(" 387191 01 : %v", __Verr2)
		*___VoutBuf = []byte{}
		return
	}

	__Vlen2 := len(__Vb2)
	//_FpfN(" 381923 02 : %x", __Vb2)
	//_FpfN(" 381923 03 : len %d: %v", __Vlen2, __Vb2)

	*___VoutBuf = make([]byte, __Vlen2+5)
	(*___VoutBuf)[0] = Cmd__loginS1ReqTryNoToken
	copy((*___VoutBuf)[1:], _VersionProtocol01)
	copy((*___VoutBuf)[5:], __Vb2)

	//_FpfN(" 381923 05 : len %d: %v", len(___VoutBuf), ___VoutBuf)
}
func (__Vundr *_TudpNodeDataRece) _FdataPack__301__dataDecode() {
	_FpfNdb(" 387192 01 : data decode start ")
}
