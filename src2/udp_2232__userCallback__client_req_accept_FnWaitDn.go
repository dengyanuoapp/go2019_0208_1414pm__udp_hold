package main

func (___VserviceUDP *_TserviceUDP) _FuserCallback__client_req_accept_FnWaitDn() {
	//_FpfNhex(&___VserviceUDP.Vubuf, 40, " 1738182 03 :")
	//_FpfNhex(&___VserviceUDP.VuSrvInfo.K256, 80, " 1738182 04 : %s :", ___VserviceUDP.name)

	//__Vb3 := _FdecAesRandExit(" 1738182 05 ", &___VserviceUDP.VuSrvInfo.K256, &___VserviceUDP.Vubuf)
	__Vb3, __Verr3 := _FdecAesRand__only(&___VserviceUDP.VuSrvInfo.K256, &___VserviceUDP.Vubuf)
	if nil != __Verr3 {
		_FpfN(" 1738182 06 received error package , ignore it: %v :", __Verr3)
		return
	}

	//_FpfNhex(&__Vb3, 80, " 1738182 07 : %s :", __Vb3)
	_FpfN(" 1738182 08 : %d / %d : %s \n", ___VserviceUDP.Vulen, len(__Vb3), __Vb3)

	if nil != ___VserviceUDP.CuOut01 {
	}

} // _FuserCallback__client_req_accept_FnWaitDn _TserviceUDP
