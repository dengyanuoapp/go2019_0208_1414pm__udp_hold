// _TtcpNodE
package main

import (
//"time"
)

func (___Vtn2 *_TtcpNodE) _FtcpNode__200401x_accept_default() {

	for __Vi := 0; __Vi < ___Vtn2.tnAmount; __Vi++ {
		___Vtn2.tnAcceptTCPs[__Vi].taIdx = __Vi
		___Vtn2.tnAcceptTCPs[__Vi].taEnabled = false
		___Vtn2.tnAcceptTCPs[__Vi].taServerTCP = ___Vtn2

		//___Vtn2.tnAcceptTCPs[__Vi].taCreceiveMsg = make(chan []byte, 10)
		//___Vtn2.tnAcceptTCPs[__Vi].taCchanMsg = make(chan []byte, 10)
		//___Vtn2.tnAcceptTCPs[__Vi].taCreceiveErr = make(chan string, 1)
	}

	for __Vi := 0; __Vi < ___Vtn2.tnAmount; __Vi++ {
		go ___Vtn2.tnAcceptTCPs[__Vi].
			_FtcpNodeAccept__200401x4__dataReceiveMsg01()
	}

	defer ___Vtn2.tnListener.Close() //_FtcpNode__tryListen01
	for {
		_Fsleep_1ms()
		_FtcpNode__200401x1_accept_loop__default(___Vtn2)
	}
	_Fex("Error 381911: (" + ___Vtn2.tnHostPortStr + ")")

} // _FtcpNode__200401x_accept_default

func _FtcpNode__200401x1_accept_loop__default(___Vtn3 *_TtcpNodE) {

	// func (l *TCPListener) AcceptTCP() (*TCPConn, error)
	__Vconn, __Verr := ___Vtn3.tnListener.AcceptTCP()
	_FerrExit(" 381818 01 : tcp accept error ", __Verr)

	//func (c *TCPConn) Write(b []byte) (int, error)
	//__Vconn.Write([]byte(_Pspf("%x\n", _Vself.progSha)))

	//_FpfN( "381818 02 accepting : max %d , now %d" , ___Vtn3.tnAmount , ___Vtn3.tnClientCnt )
	if ___Vtn3.tnAmount > ___Vtn3.tnClientCnt {
		__VcntOld := ___Vtn3.tnClientCnt
		for __Vi := 0; __Vi < ___Vtn3.tnAmount; __Vi++ {
			__VaccTcp := &(___Vtn3.tnAcceptTCPs[__Vi]) // _TacceptTCP
			if __VaccTcp.taEnabled == false {
				__VaccTcp.taId128 = _FgenRand_nByte__(16)

				// func (c *TCPConn) LocalAddr() Addr
				// func (c *TCPConn) RemoteAddr() Addr
				__VaccTcp.taLocalAddr = __Vconn.LocalAddr()
				__VaccTcp.taRemoteAddr = __Vconn.RemoteAddr()

				// acceptTcpINC / acceptTcpDEC : begin
				___Vtn3.tnClientMux.Lock()

				___Vtn3.tnClientCnt++
				__VaccTcp.taEnabled = true
				__VaccTcp.taConnTCP = __Vconn
				__VaccTcp.taOffset = 0

				___Vtn3.tnClientMux.Unlock()
				// acceptTcpINC / acceptTcpDEC : end

				_CFpfN("381818 04 new tcp accept(%d) l:%s r:%s id:{%s}",
					__Vi, __VaccTcp.taLocalAddr.String(), __VaccTcp.taRemoteAddr.String(),
					String9s(&__VaccTcp.taId128))

				break
			}
		}

		//_FpfN( "381818 06 after accept : old %d , now %d" , __VcntOld , ___Vtn3.tnClientCnt )
		_FeqExit("381818 07 err : why not inc ? ", __VcntOld, ___Vtn3.tnClientCnt)

	} else {
		__Vconn.Close()
		_CFpfN("381818 09 refuse accept : max amount of client reach %d , now %d", ___Vtn3.tnAmount, ___Vtn3.tnClientCnt)
	}
} // _FtcpNode__200401x1_accept_loop__default
