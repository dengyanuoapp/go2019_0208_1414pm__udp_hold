// _TtcpNodE
package main

import (
//"time"
)

func _Fhandle_tcpAccept01(___Vtn2 *_TtcpNodE) {

	for __Vi := 0; __Vi < ___Vtn2.cAmount; __Vi++ {
		___Vtn2.acceptTCPs[__Vi].Vbuf = make([]byte, 2048) // silice : with var len
		___Vtn2.acceptTCPs[__Vi].idx = __Vi
		___Vtn2.acceptTCPs[__Vi].enabled = false
		___Vtn2.acceptTCPs[__Vi].serverTCP = ___Vtn2

		___Vtn2.acceptTCPs[__Vi].Clog = ___Vtn2.Clog
		___Vtn2.acceptTCPs[__Vi].Cexit = ___Vtn2.Cexit

		___Vtn2.acceptTCPs[__Vi].Cstart = make(chan string, 1)
		___Vtn2.acceptTCPs[__Vi].CreceiveMsg = make(chan []byte, 10)
		___Vtn2.acceptTCPs[__Vi].CchanMsg = make(chan []byte, 10)
		___Vtn2.acceptTCPs[__Vi].CreceiveErr = make(chan string, 1)

		go ___Vtn2.acceptTCPs[__Vi]._FhandleTcp_accept_dataReceiveMsg01()
		go ___Vtn2.acceptTCPs[__Vi]._FhandleTcp__accept_dataChan__main_top()
	}

	defer ___Vtn2.tcpListener.Close() //_FtcpNode__tryListen01
	for {
		_Fsleep_1ms()
		_FtcpAccept01_loop(___Vtn2)
	}
	(*___Vtn2.Cexit) <- "Error 381911: (" + ___Vtn2.hostPortStr + ")"

} // _Fhandle_tcpAccept01

func _FtcpAccept01_loop(___Vtn3 *_TtcpNodE) {

	// func (l *TCPListener) AcceptTCP() (*TCPConn, error)
	__Vconn, __Verr := ___Vtn3.tcpListener.AcceptTCP()
	_FerrExit(" 381810 : tcp accept error ", __Verr)

	//func (c *TCPConn) Write(b []byte) (int, error)
	//_FpfN( "381812 accepting 1 :%x\n" , _Vself.progSha )
	__Vconn.Write([]byte(_Pspf("%x\n", _Vself.progSha)))
	//_FpfN( "381814 accepting 2 :%s\n" , _Vself.progSha )

	//_FpfN( "381816 accepting : max %d , now %d" , ___Vtn3.cAmount , ___Vtn3.clientCnt )
	if ___Vtn3.cAmount > ___Vtn3.clientCnt {
		__Vcnt := ___Vtn3.clientCnt
		for __Vi := 0; __Vi < ___Vtn3.cAmount; __Vi++ {
			__VacceptTcp := &(___Vtn3.acceptTCPs[__Vi])
			if __VacceptTcp.enabled == false {
				// func (c *TCPConn) LocalAddr() Addr
				// func (c *TCPConn) RemoteAddr() Addr
				__VacceptTcp.VlocalAddr = __Vconn.LocalAddr()
				__VacceptTcp.VremoteAddr = __Vconn.RemoteAddr()
				//_FpfN( "381818 l:%s r:%s"     , __VacceptTcp.VlocalAddr , __VacceptTcp.VremoteAddr )

				// acceptTcpINC / acceptTcpDEC : begin
				___Vtn3.clientMux.Lock()

				___Vtn3.clientCnt++
				__VacceptTcp.enabled = true
				__VacceptTcp.connTCP = __Vconn

				___Vtn3.clientMux.Unlock()
				// acceptTcpINC / acceptTcpDEC : end

				//_FpfN( "381815 accept succeed : old %d , now %d" , __Vcnt , ___Vtn3.clientCnt )

				__VacceptTcp.Cstart <- " 183191 start: " + _FtimeNow()
				//__VacceptTcp.Cstart              = _FtimeNow()
				// _TtcpNodE
				break
			}
		}

		//_FpfN( "381817 after accept : old %d , now %d" , __Vcnt , ___Vtn3.clientCnt )
		_FeqExit("381819 err : why not inc ? ", __Vcnt, ___Vtn3.clientCnt)

	} else {
		__Vconn.Close()
		_FpfN("38181b refuse accept : max reach %d , now %d", ___Vtn3.cAmount, ___Vtn3.clientCnt)
	}
} // _FtcpAccept01_loop
