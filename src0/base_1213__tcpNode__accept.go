// _TtcpNodE
package main

import (
//"time"
)

func _Fhandle_tcpAccept01(___VtcpNode2 *_TtcpNodE) {

	for __Vi := 0; __Vi < ___VtcpNode2.cAmount; __Vi++ {
		___VtcpNode2.acceptTCPs[__Vi].Vbuf = make([]byte, 2048) // silice : with var len
		___VtcpNode2.acceptTCPs[__Vi].idx = __Vi
		___VtcpNode2.acceptTCPs[__Vi].enabled = false
		___VtcpNode2.acceptTCPs[__Vi].serverTCP = ___VtcpNode2

		___VtcpNode2.acceptTCPs[__Vi].Clog = ___VtcpNode2.Clog
		___VtcpNode2.acceptTCPs[__Vi].Cexit = ___VtcpNode2.Cexit

		___VtcpNode2.acceptTCPs[__Vi].Cstart = make(chan string, 1)
		___VtcpNode2.acceptTCPs[__Vi].CreceiveMsg = make(chan []byte, 10)
		___VtcpNode2.acceptTCPs[__Vi].CchanMsg = make(chan []byte, 10)
		___VtcpNode2.acceptTCPs[__Vi].CreceiveErr = make(chan string, 1)

		go ___VtcpNode2.acceptTCPs[__Vi]._FhandleTcp_accept_dataReceiveMsg01()
		go ___VtcpNode2.acceptTCPs[__Vi]._FhandleTcp__accept_dataChan__main_top()
	}

	defer ___VtcpNode2.tcpListener.Close() //_FtcpNode__tryListen01
	for {
		_Fsleep_1ms()
		_FtcpAccept01_loop(___VtcpNode2)
	}
	(*___VtcpNode2.Cexit) <- "Error 381911: (" + ___VtcpNode2.hostPortStr + ")"

} // _Fhandle_tcpAccept01

func _FtcpAccept01_loop(___VtcpNode2 *_TtcpNodE) {

	// func (l *TCPListener) AcceptTCP() (*TCPConn, error)
	__Vconn, __Verr := ___VtcpNode2.tcpListener.AcceptTCP()
	_FerrExit(" 381810 : tcp accept error ", __Verr)

	//func (c *TCPConn) Write(b []byte) (int, error)
	//_FpfN( "381812 accepting 1 :%x\n" , _Vself.progSha )
	__Vconn.Write([]byte(_Pspf("%x\n", _Vself.progSha)))
	//_FpfN( "381814 accepting 2 :%s\n" , _Vself.progSha )

	//_FpfN( "381816 accepting : max %d , now %d" , ___VtcpNode2.cAmount , ___VtcpNode2.clientCnt )
	if ___VtcpNode2.cAmount > ___VtcpNode2.clientCnt {
		__Vcnt := ___VtcpNode2.clientCnt
		for __Vi := 0; __Vi < ___VtcpNode2.cAmount; __Vi++ {
			__VacceptTcp := &(___VtcpNode2.acceptTCPs[__Vi])
			if __VacceptTcp.enabled == false {
				// func (c *TCPConn) LocalAddr() Addr
				// func (c *TCPConn) RemoteAddr() Addr
				__VacceptTcp.VlocalAddr = __Vconn.LocalAddr()
				__VacceptTcp.VremoteAddr = __Vconn.RemoteAddr()
				//_FpfN( "381818 l:%s r:%s"     , __VacceptTcp.VlocalAddr , __VacceptTcp.VremoteAddr )

				// acceptTcpINC / acceptTcpDEC : begin
				___VtcpNode2.clientMux.Lock()

				___VtcpNode2.clientCnt++
				__VacceptTcp.enabled = true
				__VacceptTcp.connTCP = __Vconn

				___VtcpNode2.clientMux.Unlock()
				// acceptTcpINC / acceptTcpDEC : end

				//_FpfN( "381815 accept succeed : old %d , now %d" , __Vcnt , ___VtcpNode2.clientCnt )

				__VacceptTcp.Cstart <- " 183191 start: " + _FtimeNow()
				//__VacceptTcp.Cstart              = _FtimeNow()
				// _TtcpNodE
				break
			}
		}

		//_FpfN( "381817 after accept : old %d , now %d" , __Vcnt , ___VtcpNode2.clientCnt )
		_FeqExit("381819 err : why not inc ? ", __Vcnt, ___VtcpNode2.clientCnt)

	} else {
		__Vconn.Close()
		_FpfN("38181b refuse accept : max reach %d , now %d", ___VtcpNode2.cAmount, ___VtcpNode2.clientCnt)
	}
} // _FtcpAccept01_loop
