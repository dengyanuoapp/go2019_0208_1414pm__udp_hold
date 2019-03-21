package main

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
	var __VbOut []byte
	return &__VbOut
}
