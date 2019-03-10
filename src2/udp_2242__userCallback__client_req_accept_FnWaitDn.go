package main

import (
//"fmt"
)

var __VreqIneedToLogin__FnWaitDn _TreqIneedToLogin
var __Vmap__FnWaitDn map[[16]byte]_TreqIneedToLogin

func (___VserviceUDP *_TserviceUDP) _FuserCallback__client_req_accept_FnWaitDn__01x() {
	//_FpfNhex(&___VserviceUDP.Vubuf, 40, " 1738182 01 :")
	//_FpfNhex(&___VserviceUDP.VuSrvInfo.K256, 80, " 1738182 02 : %s :", ___VserviceUDP.name)

	//__Vb401 := _FdecAesRandExit(" 1738182 03 ", &___VserviceUDP.VuSrvInfo.K256, &___VserviceUDP.Vubuf)
	__Vb401, __Verr401 := _FdecAesRand__only(&___VserviceUDP.VuSrvInfo.K256, &___VserviceUDP.Vubuf)
	if nil != __Verr401 {
		_FpfN(" 1738182 04 received error package , ignore it: %v :", __Verr401)
		return
	}

	_FpfN(" 1738182 05 : rcev : %d : %s", len(__Vb401), __Vb401)
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

	// if
	//___VserviceUDP.CuOut01 <-

} // _FuserCallback__client_req_accept_FnWaitDn__01y _TserviceUDP
