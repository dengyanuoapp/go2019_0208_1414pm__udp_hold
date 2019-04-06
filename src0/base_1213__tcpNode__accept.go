// _TtcpNodE
package main

import (
//"time"
)

func _Fhandle_tcpAccept01(___Vtn2 *_TtcpNodE) {

	for __Vi := 0; __Vi < ___Vtn2.tnAmount; __Vi++ {
		___Vtn2.tnAcceptTCPs[__Vi].Vbuf = make([]byte, 2048) // silice : with var len
		___Vtn2.tnAcceptTCPs[__Vi].idx = __Vi
		___Vtn2.tnAcceptTCPs[__Vi].enabled = false
		___Vtn2.tnAcceptTCPs[__Vi].serverTCP = ___Vtn2

		___Vtn2.tnAcceptTCPs[__Vi].Clog = ___Vtn2.tnClog
		___Vtn2.tnAcceptTCPs[__Vi].Cexit = ___Vtn2.tnCexit

		___Vtn2.tnAcceptTCPs[__Vi].Cstart = make(chan string, 1)
		___Vtn2.tnAcceptTCPs[__Vi].CreceiveMsg = make(chan []byte, 10)
		___Vtn2.tnAcceptTCPs[__Vi].CchanMsg = make(chan []byte, 10)
		___Vtn2.tnAcceptTCPs[__Vi].CreceiveErr = make(chan string, 1)

		go ___Vtn2.tnAcceptTCPs[__Vi]._FhandleTcp_accept_dataReceiveMsg01()
		go ___Vtn2.tnAcceptTCPs[__Vi]._FhandleTcp__accept_dataChan__main_top()
	}

	defer ___Vtn2.tnListener.Close() //_FtcpNode__tryListen01
	for {
		_Fsleep_1ms()
		_FtcpAccept01_loop(___Vtn2)
	}
	(*___Vtn2.tnCexit) <- "Error 381911: (" + ___Vtn2.tnHostPortStr + ")"

} // _Fhandle_tcpAccept01

func _FtcpAccept01_loop(___Vtn3 *_TtcpNodE) {

	// func (l *TCPListener) AcceptTCP() (*TCPConn, error)
	__Vconn, __Verr := ___Vtn3.tnListener.AcceptTCP()
	_FerrExit(" 381810 : tcp accept error ", __Verr)

	//func (c *TCPConn) Write(b []byte) (int, error)
	//_FpfN( "381812 accepting 1 :%x\n" , _Vself.progSha )
	__Vconn.Write([]byte(_Pspf("%x\n", _Vself.progSha)))
	//_FpfN( "381814 accepting 2 :%s\n" , _Vself.progSha )

	//_FpfN( "381816 accepting : max %d , now %d" , ___Vtn3.tnAmount , ___Vtn3.tnClientCnt )
	if ___Vtn3.tnAmount > ___Vtn3.tnClientCnt {
		__Vcnt := ___Vtn3.tnClientCnt
		for __Vi := 0; __Vi < ___Vtn3.tnAmount; __Vi++ {
			__VacceptTcp := &(___Vtn3.tnAcceptTCPs[__Vi])
			if __VacceptTcp.enabled == false {
				// func (c *TCPConn) LocalAddr() Addr
				// func (c *TCPConn) RemoteAddr() Addr
				__VacceptTcp.VlocalAddr = __Vconn.LocalAddr()
				__VacceptTcp.VremoteAddr = __Vconn.RemoteAddr()
				//_FpfN( "381818 l:%s r:%s"     , __VacceptTcp.VlocalAddr , __VacceptTcp.VremoteAddr )

				// acceptTcpINC / acceptTcpDEC : begin
				___Vtn3.tnClientMux.Lock()

				___Vtn3.tnClientCnt++
				__VacceptTcp.enabled = true
				__VacceptTcp.connTCP = __Vconn

				___Vtn3.tnClientMux.Unlock()
				// acceptTcpINC / acceptTcpDEC : end

				//_FpfN( "381815 accept succeed : old %d , now %d" , __Vcnt , ___Vtn3.tnClientCnt )

				__VacceptTcp.Cstart <- " 183191 start: " + _FtimeNow()
				//__VacceptTcp.Cstart              = _FtimeNow()
				// _TtcpNodE
				break
			}
		}

		//_FpfN( "381817 after accept : old %d , now %d" , __Vcnt , ___Vtn3.tnClientCnt )
		_FeqExit("381819 err : why not inc ? ", __Vcnt, ___Vtn3.tnClientCnt)

	} else {
		__Vconn.Close()
		_FpfN("38181b refuse accept : max reach %d , now %d", ___Vtn3.tnAmount, ___Vtn3.tnClientCnt)
	}
} // _FtcpAccept01_loop
