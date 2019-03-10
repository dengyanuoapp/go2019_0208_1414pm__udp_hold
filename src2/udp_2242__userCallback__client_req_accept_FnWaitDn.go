package main

func (___VserviceUDP *_TserviceUDP) _FuserCallback__client_req_accept_FnWaitDn() {
	//_FpfNhex(&___VserviceUDP.Vubuf, 40, " 1738182 01 :")
	//_FpfNhex(&___VserviceUDP.VuSrvInfo.K256, 80, " 1738182 02 : %s :", ___VserviceUDP.name)

	//__Vb3 := _FdecAesRandExit(" 1738182 03 ", &___VserviceUDP.VuSrvInfo.K256, &___VserviceUDP.Vubuf)
	__Vb3, __Verr3 := _FdecAesRand__only(&___VserviceUDP.VuSrvInfo.K256, &___VserviceUDP.Vubuf)
	if nil != __Verr3 {
		_FpfN(" 1738182 04 received error package , ignore it: %v :", __Verr3)
		return
	}

	_FpfN(" 1738182 05 : rcev : %d : %s", len(__Vb3), __Vb3)

	// if
	//___VserviceUDP.CuOut01 <-

} // _FuserCallback__client_req_accept_FnWaitDn _TserviceUDP
