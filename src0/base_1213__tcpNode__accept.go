// _TtcpNodE
package main

import (
//"time"
)

func (___Vtn2 *_TtcpNodE) _FtcpNode__200401x_accept_default() {

	for __Vi := 0; __Vi < ___Vtn2.tnAmount; __Vi++ {
		___Vtn2.tnAcceptTCPs[__Vi].taBuf = make([]byte, 2048) // silice : with var len
		___Vtn2.tnAcceptTCPs[__Vi].taIdx = __Vi
		___Vtn2.tnAcceptTCPs[__Vi].taEnabled = false
		___Vtn2.tnAcceptTCPs[__Vi].taServerTCP = ___Vtn2

		___Vtn2.tnAcceptTCPs[__Vi].taClog = ___Vtn2.tnClog
		___Vtn2.tnAcceptTCPs[__Vi].taCexit = ___Vtn2.tnCexit

		___Vtn2.tnAcceptTCPs[__Vi].taCstart = make(chan string, 1)
		___Vtn2.tnAcceptTCPs[__Vi].taCreceiveMsg = make(chan []byte, 10)
		___Vtn2.tnAcceptTCPs[__Vi].taCchanMsg = make(chan []byte, 10)
		___Vtn2.tnAcceptTCPs[__Vi].taCreceiveErr = make(chan string, 1)

		go ___Vtn2.tnAcceptTCPs[__Vi]._FhandleTcp_accept_dataReceiveMsg01()
		go ___Vtn2.tnAcceptTCPs[__Vi]._FhandleTcp__accept_dataChan__main_top()
	}

	defer ___Vtn2.tnListener.Close() //_FtcpNode__tryListen01
	for {
		_Fsleep_1ms()
		_FtcpAccept01_loop(___Vtn2)
	}
	(*___Vtn2.tnCexit) <- "Error 381911: (" + ___Vtn2.tnHostPortStr + ")"

} // _FtcpNode__200401x_accept_default

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
			if __VacceptTcp.taEnabled == false {
				// func (c *TCPConn) LocalAddr() Addr
				// func (c *TCPConn) RemoteAddr() Addr
				__VacceptTcp.taLocalAddr = __Vconn.LocalAddr()
				__VacceptTcp.taRemoteAddr = __Vconn.RemoteAddr()
				//_FpfN( "381818 l:%s r:%s"     , __VacceptTcp.taLocalAddr , __VacceptTcp.taRemoteAddr )

				// acceptTcpINC / acceptTcpDEC : begin
				___Vtn3.tnClientMux.Lock()

				___Vtn3.tnClientCnt++
				__VacceptTcp.taEnabled = true
				__VacceptTcp.taConnTCP = __Vconn

				___Vtn3.tnClientMux.Unlock()
				// acceptTcpINC / acceptTcpDEC : end

				//_FpfN( "381815 accept succeed : old %d , now %d" , __Vcnt , ___Vtn3.tnClientCnt )

				__VacceptTcp.taCstart <- " 183191 start: " + _FtimeNow()
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
