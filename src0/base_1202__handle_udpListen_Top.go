// _TserviceUDP
package main

//import ()

func (___VsrvU *_TserviceUDP) IRun(___Vidx int) {
	switch ___Vidx {
	case 100201:
		if nil == ___VsrvU.UsrvGoCall100201 {
			___VsrvU._Fhandle_u100201x__udpListen_Udp__read_main_top__default()
		} else {
			___VsrvU.UsrvGoCall100201(___VsrvU)
		}
	case 100202:
		if nil == ___VsrvU.UsrvGoCall100202 {
			___VsrvU._Fhandle_u100202x__udpListen__chanIn_main_top__default()
		} else {
			___VsrvU.UsrvGoCall100202(___VsrvU)
		}
	case 100203:
		if nil == ___VsrvU.UsrvGoCall100203 {
			___VsrvU._Fhandle_u100203x__udpListen__timer__main_top__default()
		} else {
			___VsrvU.UsrvGoCall100203(___VsrvU)
		}

	case 100211:
		if nil != ___VsrvU.UuserLoopCall100211 {
			___VsrvU.UuserLoopCall100211(___VsrvU)
		}
	case 100221:
		if nil != ___VsrvU.UuserLoopCall100221 {
			___VsrvU.UuserLoopCall100221(___VsrvU)
		}

	default:
		_FpfNex(" 839181 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

// https://golang.org/pkg/net/#UDPConn.ReadFromUDP
// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
//func _Fhandle_udpListen_Udp_main_top(___VserviceUdp *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
func (___VserviceUdp *_TserviceUDP) _Fhandle_u100201x__udpListen_Udp__read_main_top__default() {

	___VserviceUdp._FtryListenToUDP01()

	___VserviceUdp.Vubuf = make([]byte, 2048)    // silice : with var len
	___VserviceUdp.CuIn01 = make(chan []byte, 5) // silice : with var len

	//go ___VserviceUdp. _Fhandle_u100202x__udpListen__chanIn_main_top__default()
	go _Frun(___VserviceUdp, 100202)

	//go ___VserviceUdp._Fhandle_u100203x__udpListen__timer__main_top__default()
	go _Frun(___VserviceUdp, 100203)

	for {
		___VserviceUdp._Fhandle_u01y__udpListen_Udp__read_main_loop()
	}
	(*___VserviceUdp.Cexit) <- "338191 Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _Fhandle_u100201x__udpListen_Udp__read_main_top__default

func (___VserviceUdp *_TserviceUDP) _Fhandle_u01y__udpListen_Udp__read_main_loop() {
	___VserviceUdp.Vulen,
		___VserviceUdp.VuremoteAddr,
		___VserviceUdp.Vuerr =
		___VserviceUdp.udpConn.ReadFromUDP(___VserviceUdp.Vubuf)

	_FerrExit("err 338195 : udp reading : ", ___VserviceUdp.Vuerr)

	_FnullExit(" err 338197 : why ___Vconn.ReadFromUDP addr error ?", ___VserviceUdp.VuremoteAddr)

	//_FnotNullRunUdp01(___VserviceUdp.UuserLoopCall100211, ___VserviceUdp)
	_Frun(___VserviceUdp, 100211)
} // _Fhandle_u01y__udpListen_Udp__read_main_loop

func (___VserviceUdp *_TserviceUDP) _Fhandle_u100202x__udpListen__chanIn_main_top__default() {

	for {
		_Fsleep_1s()
		//_FpfN( "438195 : udp rece Chan msg ." )
		//_FnotNullRunUdp01(___VserviceUdp.UuserLoopCall100221, ___VserviceUdp)
		_Frun(___VserviceUdp, 100221)
	}

	(*___VserviceUdp.Cexit) <- "438191 Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _Fhandle_u100202x__udpListen__chanIn_main_top__default

func (___VserviceUdp *_TserviceUDP) _Fhandle_u100203x__udpListen__timer__main_top__default() {

	__VUreqNewSessionTM := ___VserviceUdp.UreqNewSessionTM
	if nil == __VUreqNewSessionTM {
		return // no need to run
	}

	__VUreqNewSessionTM.serviceUdP = ___VserviceUdp

	for {
		_Frun(__VUreqNewSessionTM, 430301) // IRun _Fconnect_to_server_430301x__req_new_sessionID__default
	}

	(*___VserviceUdp.Cexit) <- "538191 Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _Fhandle_u100203x__udpListen__timer__main_top__default
