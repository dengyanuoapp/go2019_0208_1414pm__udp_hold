// _TacceptTCP
// _TtcpNodE
// _Fhandle_tcpAccept01
package main

import (
	"io"
)

func (___VtAcc2 *_TacceptTCP) _FhandleTcp_accept_dataReceiveMsg01() {
	//var __VcanReceiveMsg bool = true
	for {
		if 3 == 2 {
			<-___VtAcc2.Cstart
		} else {
			_FpfN(" 188118 rece run start at " + <-___VtAcc2.Cstart)
		}

		//if ( false == ___VtAcc2.enabled ) { continue }

		for {
			_Fsleep_1ms()

			//__VcanReceiveMsg = ___VtAcc2.enabled
			if false == ___VtAcc2.enabled {
				break
			}

			_FhandleTcp_accept_dataReceiveMsg01__loop(___VtAcc2)
		}
	}

} // _FhandleTcp_accept_dataReceiveMsg01

// func (c *TCPConn) Read(b []byte) (int, error)
// _TacceptTCP
func _FhandleTcp_accept_dataReceiveMsg01__loop(___VtAcc3 *_TacceptTCP) bool {
	if 4 == 2 {
		_FhandleTcp_accept_dataReceiveMsg01__debug(___VtAcc3)
	}

	___VtAcc3.cR.try++

	___VtAcc3.Vlen,
		___VtAcc3.Verr =
		___VtAcc3.connTCP.Read(___VtAcc3.Vbuf)

	// _Fhandle_tcpAccept01
	if ___VtAcc3.Verr == io.EOF { // lost the connect.
		___VtAcc3.cR.eofErr++
		// acceptTcpINC / acceptTcpDEC : begin
		___VtAcc3.serverTCP.clientMux.Lock()

		___VtAcc3.serverTCP.clientCnt--
		___VtAcc3.enabled = false
		___VtAcc3.connTCP.Close()

		___VtAcc3.serverTCP.clientMux.Unlock()
		// acceptTcpINC / acceptTcpDEC : end

		___VtAcc3.CreceiveErr <- _Pspf("EOF:%d", ___VtAcc3.idx)
		return false
	}

	_FerrExit(" reading from tcp 831911 ", ___VtAcc3.Verr)

	_FnullExit(" 183813 : why ___Vconn.ReadFromTCP addr error ?", ___VtAcc3.VremoteAddr)

	___VtAcc3.cR.ok++

	/*
	   _Fpf( " 183814 | l:%s | r:%s | " , ___VtAcc3.VlocalAddr , ___VtAcc3.VremoteAddr )
	   _PpdN( ___VtAcc3.Vlen , &(___VtAcc3.Vbuf) )
	*/

	_FcopyByte(&(___VtAcc3.Vbuf2), &(___VtAcc3.Vbuf), ___VtAcc3.Vlen)
	//___VtAcc3.Vbuf2 = make([]byte , ___VtAcc3.Vlen ); copy( ___VtAcc3.Vbuf2 , ___VtAcc3.Vbuf )

	// _FcallbackForDebugLog_accept
	_FnotNullRunTcp02_accept(___VtAcc3.serverTCP.tnCBaccDataRece, ___VtAcc3)

	return true
} // _FhandleTcp_accept_dataReceiveMsg01__loop

func _FhandleTcp_accept_dataReceiveMsg01__debug(___VtAcc4 *_TacceptTCP) {

	_Ppn(" 183891 : under constructing ", ___VtAcc4.cR.try, ___VtAcc4.cR.ok)

	if nil == ___VtAcc4.serverTCP {
		_Ppn(" 183892 : ___VtAcc4.serverTCP == nil ")
	} else {

		_Ppf(" 183893 : ")
		for __Vi := 0; __Vi < ___VtAcc4.serverTCP.tnAmount; __Vi++ {
			__VacceptTcp := &(___VtAcc4.serverTCP.acceptTCPs[__Vi])
			_Ppf(" %d,%d,%d,%d", __Vi, __VacceptTcp.idx, __VacceptTcp.cR.try, __VacceptTcp.cR.ok)
		}
		_Ppn()
	}
} // _FhandleTcp_accept_dataReceiveMsg01__debug
