package main

import (
	"time"
	//"fmt"
)

var (
	__VreqIneedToLogin__FnWaitDn _TreqIneedToLogin
	__Vmap__FnWaitDn             map[[16]byte]_TreqIneedToLogin
	__Vhex16                     [16]byte
)

func (___VserviceUDP *_TserviceUDP) _FuserCallback__client_req_accept_FnWaitDn__01x() {
	//_FpfNhex(&___VserviceUDP.Vubuf, 40, " 1738182 01 :")
	//_FpfNhex(&___VserviceUDP.VuSrvInfo.K256, 80, " 1738182 02 : %s :", ___VserviceUDP.name)

	//__Vb401 := _FdecAesRandExit(" 1738182 03 ", &___VserviceUDP.VuSrvInfo.K256, &___VserviceUDP.Vubuf)
	__Vb401, __Verr401 := _FdecAesRand__only(&___VserviceUDP.VuSrvInfo.K256, &___VserviceUDP.Vubuf)
	if nil != __Verr401 {
		_FpfN(" 1738182 04 received error package , ignore it: %v :", __Verr401)
		return
	}

	_FpfN(" 1738182 05 : rcev : %d : %s : %d", len(__Vb401), __Vb401, time.Now().Unix())
	__Verr402 := _FdecJson___(" 1738182 06 ", &__Vb401, &__VreqIneedToLogin__FnWaitDn)
	if nil != __Verr402 {
		_FpfN(" 1738182 07 decJson error: %v :", __Verr402)
		return
	}
	if __VreqIneedToLogin__FnWaitDn.ReqStr != "reqLogin_Dn2Fn" || __VreqIneedToLogin__FnWaitDn.MeName != "Dn" {
		_FpfN(" 1738182 08 not req to me . ")
		return
	}

	_FpfN(" 1738182 09 decJson ok, for Fn(me).")
	___VserviceUDP._FuserCallback__client_req_accept_FnWaitDn__01y()
} // _FuserCallback__client_req_accept_FnWaitDn__01x _TserviceUDP

func (___VserviceUDP *_TserviceUDP) _FuserCallback__client_req_accept_FnWaitDn__01y() {

	copy(__Vhex16[:], __VreqIneedToLogin__FnWaitDn.MeIdx128)
	_FpfN(" 1738183 01 __Vhex16 : %x", __Vhex16)

	// if
	//___VserviceUDP.CuOut01 <-

	_Pn()
} // _FuserCallback__client_req_accept_FnWaitDn__01y _TserviceUDP