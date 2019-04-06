// _TacceptTCP
// _TtcpNodE
// _FtcpNode__200401x_accept_default
package main

import (
	"io"
)

func (___VtAcc2 *_TacceptTCP) _FtcpNodeAccept__200401x4__dataReceiveMsg01() {
	//var __VcanReceiveMsg bool = true
	for {
		//_Fsleep_1ms()

		//__VcanReceiveMsg = ___VtAcc2.taEnabled
		if false == ___VtAcc2.taEnabled {
			_Fsleep(_T1s)
		} else {
			if 2 == 2 {
				_FtcpNodeAccept__200401x5__dataReceiveMsg01_default(___VtAcc2) // default when inpu
			} else {
				_FtcpNodeAccept__200401x5__dataReceiveMsg01_debug(___VtAcc2) // if try debug when input
			}
		}
	}
} //

// func (c *TCPConn) Read(b []byte) (int, error)
// _TacceptTCP
func _FtcpNodeAccept__200401x5__dataReceiveMsg01_default(___VtAcc3 *_TacceptTCP) bool {
	if 4 == 2 {
	}

	___VtAcc3.taR.try++

	___VtAcc3.taLen,
		___VtAcc3.taErr =
		___VtAcc3.taConnTCP.Read(___VtAcc3.taBuf)

	// _FtcpNode__200401x_accept_default
	if ___VtAcc3.taErr == io.EOF { // lost the connect.
		___VtAcc3.taR.eofErr++
		// acceptTcpINC / acceptTcpDEC : begin
		___VtAcc3.taServerTCP.tnClientMux.Lock()

		___VtAcc3.taServerTCP.tnClientCnt--
		___VtAcc3.taEnabled = false
		___VtAcc3.taConnTCP.Close()

		___VtAcc3.taServerTCP.tnClientMux.Unlock()
		// acceptTcpINC / acceptTcpDEC : end

		___VtAcc3.taCreceiveErr <- _Pspf("EOF:%d", ___VtAcc3.taIdx)
		return false
	}

	_FerrExit(" reading from tcp 831911 ", ___VtAcc3.taErr)

	_FnullExit(" 183813 : why ___Vconn.ReadFromTCP addr error ?", ___VtAcc3.taRemoteAddr)

	___VtAcc3.taR.ok++

	/*
	   _Fpf( " 183814 | l:%s | r:%s | " , ___VtAcc3.VlocalAddr , ___VtAcc3.taRemoteAddr )
	   _PpdN( ___VtAcc3.taLen , &(___VtAcc3.taBuf) )
	*/

	_FcopyByte(&(___VtAcc3.taBuf2), &(___VtAcc3.taBuf), ___VtAcc3.taLen)
	//___VtAcc3.taBuf2 = make([]byte , ___VtAcc3.taLen ); copy( ___VtAcc3.taBuf2 , ___VtAcc3.taBuf )

	// _FcallbackForDebugLog_accept
	_FnotNullRunTcp02_accept(___VtAcc3.taServerTCP.tnCBaccDataRece, ___VtAcc3)

	return true
} // _FtcpNodeAccept__200401x5__dataReceiveMsg01_default

func _FtcpNodeAccept__200401x5__dataReceiveMsg01_debug(___VtAcc4 *_TacceptTCP) {

	_Ppn(" 183891 : under constructing ", ___VtAcc4.taR.try, ___VtAcc4.taR.ok)

	if nil == ___VtAcc4.taServerTCP {
		_Ppn(" 183892 : ___VtAcc4.taServerTCP == nil ")
	} else {

		_Ppf(" 183893 : ")
		for __Vi := 0; __Vi < ___VtAcc4.taServerTCP.tnAmount; __Vi++ {
			__VacceptTcp := &(___VtAcc4.taServerTCP.tnAcceptTCPs[__Vi])
			_Ppf(" %d,%d,%d,%d", __Vi, __VacceptTcp.taIdx, __VacceptTcp.taR.try, __VacceptTcp.taR.ok)
		}
		_Ppn()
	}
} // _FtcpNodeAccept__200401x5__dataReceiveMsg01_debug
