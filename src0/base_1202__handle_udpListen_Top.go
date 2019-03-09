// _TserviceUDP
package main

//import ()

func (___VsrvU *_TserviceUDP) IRun(___Vidx int) {
	switch ___Vidx {
	case 201:
		if nil == ___VsrvU.UsrvGoCall01 {
			___VsrvU._Fhandle_u01x__udpListen_Udp__read_main_top__default()
		} else {
			___VsrvU.UsrvGoCall01(___VsrvU)
		}
	case 202:
		if nil == ___VsrvU.UsrvGoCall02 {
			___VsrvU._Fhandle_u02x__udpListen__chanIn_main_top__default()
		} else {
			___VsrvU.UsrvGoCall02(___VsrvU)
		}
	case 203:
		if nil == ___VsrvU.UsrvGoCall03 {
			___VsrvU._Fhandle_u03x__udpListen__timer__main_top__default()
		} else {
			___VsrvU.UsrvGoCall03(___VsrvU)
		}

	case 211:
		if nil != ___VsrvU.UuserLoopCall11 {
			___VsrvU.UuserLoopCall11(___VsrvU)
		}
	case 221:
		if nil != ___VsrvU.UuserLoopCall21 {
			___VsrvU.UuserLoopCall21(___VsrvU)
		}

	default:
		_FpfNex(" 839181 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

// https://golang.org/pkg/net/#UDPConn.ReadFromUDP
// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
//func _Fhandle_udpListen_Udp_main_top(___VserviceUdp *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
func (___VserviceUdp *_TserviceUDP) _Fhandle_u01x__udpListen_Udp__read_main_top__default() {

	___VserviceUdp._FtryListenToUDP01()

	___VserviceUdp.Vubuf = make([]byte, 2048)    // silice : with var len
	___VserviceUdp.CuIn01 = make(chan []byte, 5) // silice : with var len

	//go ___VserviceUdp._Fhandle_u02x__udpListen__chanIn_main_top__default()
	go _Frun(___VserviceUdp, 202)

	//go ___VserviceUdp._Fhandle_u03x__udpListen__timer__main_top__default()
	go _Frun(___VserviceUdp, 203)

	for {
		___VserviceUdp._Fhandle_u01y__udpListen_Udp__read_main_loop()
	}
	(*___VserviceUdp.Cexit) <- "338191 Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _Fhandle_u01x__udpListen_Udp__read_main_top__default

func (___VserviceUdp *_TserviceUDP) _Fhandle_u01y__udpListen_Udp__read_main_loop() {
	___VserviceUdp.Vulen,
		___VserviceUdp.VuremoteAddr,
		___VserviceUdp.Vuerr =
		___VserviceUdp.udpConn.ReadFromUDP(___VserviceUdp.Vubuf)

	_FerrExit("err 338195 : udp reading : ", ___VserviceUdp.Vuerr)

	_FnullExit(" err 338197 : why ___Vconn.ReadFromUDP addr error ?", ___VserviceUdp.VuremoteAddr)

	//_FnotNullRunUdp01(___VserviceUdp.UuserLoopCall11, ___VserviceUdp)
	_Frun(___VserviceUdp, 211)
} // _Fhandle_u01y__udpListen_Udp__read_main_loop

func (___VserviceUdp *_TserviceUDP) _Fhandle_u02x__udpListen__chanIn_main_top__default() {

	for {
		_Fsleep_1s()
		//_FpfN( "438195 : udp rece Chan msg ." )
		//_FnotNullRunUdp01(___VserviceUdp.UuserLoopCall21, ___VserviceUdp)
		_Frun(___VserviceUdp, 221)
	}

	(*___VserviceUdp.Cexit) <- "438191 Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _Fhandle_u02x__udpListen__chanIn_main_top__default

func (___VserviceUdp *_TserviceUDP) _Fhandle_u03x__udpListen__timer__main_top__default() {

	__VUreqNewSessionTM := ___VserviceUdp.UreqNewSessionTM
	if nil == __VUreqNewSessionTM {
		return // no need to run
	}

	__VUreqNewSessionTM.serviceUdP = ___VserviceUdp

	for {
		//_Fsleep_1s()
		//___VserviceUdp._Fhandle_u03y__udpListen__timer__main_loop()
		//_FnotNullRunUdp02(__VUreqNewSessionTM.UcallbackTM, __VUreqNewSessionTM) //  _FuserCallback_u01__reqNewSession_in_Dn_to_fn
		_Frun(__VUreqNewSessionTM, 301)
	}

	(*___VserviceUdp.Cexit) <- "538191 Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _Fhandle_u03x__udpListen__timer__main_top__default
