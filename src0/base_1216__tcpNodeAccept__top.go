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

		___Vtn2.tnAcceptTCPs[__Vi].taCreceiveMsg = make(chan []byte, 10)
		___Vtn2.tnAcceptTCPs[__Vi].taCchanMsg = make(chan []byte, 10)
		___Vtn2.tnAcceptTCPs[__Vi].taCreceiveErr = make(chan string, 1)
	}

	for __Vi := 0; __Vi < ___Vtn2.tnAmount; __Vi++ {
		go ___Vtn2.tnAcceptTCPs[__Vi]._FtcpNodeAccept__200401x4__dataReceiveMsg01()
		go ___Vtn2.tnAcceptTCPs[__Vi]._FhandleTcp__accept_dataChan__main_top()
	}

	defer ___Vtn2.tnListener.Close() //_FtcpNode__tryListen01
	for {
		_Fsleep_1ms()
		_FtcpNode__200401x1_accept_loop__default(___Vtn2)
	}
	_Fex1("Error 381911: (" + ___Vtn2.tnHostPortStr + ")")

} // _FtcpNode__200401x_accept_default

func _FtcpNode__200401x1_accept_loop__default(___Vtn3 *_TtcpNodE) {

	// func (l *TCPListener) AcceptTCP() (*TCPConn, error)
	__Vconn, __Verr := ___Vtn3.tnListener.AcceptTCP()
	_FerrExit(" 381810 : tcp accept error ", __Verr)

	//func (c *TCPConn) Write(b []byte) (int, error)
	//__Vconn.Write([]byte(_Pspf("%x\n", _Vself.progSha)))

	//_FpfN( "381816 accepting : max %d , now %d" , ___Vtn3.tnAmount , ___Vtn3.tnClientCnt )
	if ___Vtn3.tnAmount > ___Vtn3.tnClientCnt {
		__VcntOld := ___Vtn3.tnClientCnt
		for __Vi := 0; __Vi < ___Vtn3.tnAmount; __Vi++ {
			__VaccTcp := &(___Vtn3.tnAcceptTCPs[__Vi])
			if __VaccTcp.taEnabled == false {
				__VaccTcp.taIdx128 = _FgenRand_nByte__(16)

				// func (c *TCPConn) LocalAddr() Addr
				// func (c *TCPConn) RemoteAddr() Addr
				__VaccTcp.taLocalAddr = __Vconn.LocalAddr()
				__VaccTcp.taLocalAddr = __Vconn.LocalAddr()
				__VaccTcp.taRemoteAddr = __Vconn.RemoteAddr()
				//_FpfN( "381818 l:%s r:%s"     , __VaccTcp.taLocalAddr , __VaccTcp.taRemoteAddr )

				// acceptTcpINC / acceptTcpDEC : begin
				___Vtn3.tnClientMux.Lock()

				___Vtn3.tnClientCnt++
				__VaccTcp.taEnabled = true
				__VaccTcp.taConnTCP = __Vconn

				___Vtn3.tnClientMux.Unlock()
				// acceptTcpINC / acceptTcpDEC : end

				break
			}
		}

		//_FpfN( "381817 after accept : old %d , now %d" , __VcntOld , ___Vtn3.tnClientCnt )
		_FeqExit("381819 err : why not inc ? ", __VcntOld, ___Vtn3.tnClientCnt)

	} else {
		__Vconn.Close()
		_FpfN("38181b refuse accept : max amount of client reach %d , now %d", ___Vtn3.tnAmount, ___Vtn3.tnClientCnt)
	}
} // _FtcpNode__200401x1_accept_loop__default
