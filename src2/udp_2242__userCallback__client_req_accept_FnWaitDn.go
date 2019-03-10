package main

import (
//"time"
//"fmt"
//"sync"
)

var (
	__V2242_req  _TreqIneedToLogin
	__V2242_rMap _TreqAcceptMap = _TreqAcceptMap{
		reqMapNow:  make(map[[16]byte]_TreqLoginCNT),
		reqMapLast: make(map[[16]byte]_TreqLoginCNT),
	}
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

	//_FpfN(" 1738182 05 : rcev : %d : %s : %d", len(__Vb401), __Vb401, time.Now().Unix())
	__Verr402 := _FdecJson___(" 1738182 06 ", &__Vb401, &__V2242_req)
	if nil != __Verr402 {
		_FpfN(" 1738182 07 decJson error: %v :", __Verr402)
		return
	}
	if __V2242_req.ReqStr != "reqLogin_Dn2Fn" || __V2242_req.MeName != "Dn" {
		_FpfN(" 1738182 08 not req to me . ")
		return
	}

	//_FpfN(" 1738182 09 decJson ok, for Fn(me).")
	___VserviceUDP._FuserCallback__client_req_accept_FnWaitDn__01y()
} // _FuserCallback__client_req_accept_FnWaitDn__01x _TserviceUDP

func (___VserviceUDP *_TserviceUDP) _FuserCallback__client_req_accept_FnWaitDn__01y() {

	var __VreqId128 [16]byte
	copy(__VreqId128[:], __V2242_req.MeIdx128)
	_FpfN(" 1738183 01 __VreqId128 : %x", __VreqId128)

	__V2242_rMap.mux.Lock()
	__Vreq2, __Vexist2 := __V2242_rMap.reqMapNow[__VreqId128]
	if true == __Vexist2 {
		_FpfN(" 1738183 02 %t , %v", __Vexist2, __Vreq2)
	} else {
		_FpfN(" 1738183 03 %t ", __Vexist2)
	}

	__V2242_rMap.mux.Unlock()
	// if
	//___VserviceUDP.CuOut01 <-

	_Pn()
} // _FuserCallback__client_req_accept_FnWaitDn__01y _TserviceUDP
