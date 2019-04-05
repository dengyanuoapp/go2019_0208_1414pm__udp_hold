// _TacceptTCP
// _TtcpNodE
// _Fhandle_tcpAccept01
package main

import (
	"io"
)

func (___VacceptTcp *_TacceptTCP) _FhandleTcp_accept_dataReceiveMsg01() {
	//var __VcanReceiveMsg bool = true
	for {
		if 3 == 2 {
			<-___VacceptTcp.Cstart
		} else {
			_FpfN(" 188118 rece run start at " + <-___VacceptTcp.Cstart)
		}

		//if ( false == ___VacceptTcp.enabled ) { continue }

		for {
			_Fsleep_1ms()

			//__VcanReceiveMsg = ___VacceptTcp.enabled
			if false == ___VacceptTcp.enabled {
				break
			}

			_FhandleTcp_accept_dataReceiveMsg01__loop(___VacceptTcp)
		}
	}

} // _FhandleTcp_accept_dataReceiveMsg01

// func (c *TCPConn) Read(b []byte) (int, error)
// _TacceptTCP
func _FhandleTcp_accept_dataReceiveMsg01__loop(___VacceptTcp *_TacceptTCP) bool {
	if 4 == 2 {
		_FhandleTcp_accept_dataReceiveMsg01__debug(___VacceptTcp)
	}

	___VacceptTcp.r64try++

	___VacceptTcp.Vlen,
		___VacceptTcp.Verr =
		___VacceptTcp.connTCP.Read(___VacceptTcp.Vbuf)

	// _Fhandle_tcpAccept01
	if ___VacceptTcp.Verr == io.EOF { // lost the connect.
		___VacceptTcp.r64eof++
		// acceptTcpINC / acceptTcpDEC : begin
		___VacceptTcp.serverTCP.clientMux.Lock()

		___VacceptTcp.serverTCP.clientCnt--
		___VacceptTcp.enabled = false
		___VacceptTcp.connTCP.Close()

		___VacceptTcp.serverTCP.clientMux.Unlock()
		// acceptTcpINC / acceptTcpDEC : end

		___VacceptTcp.CreceiveErr <- _Pspf("EOF:%d", ___VacceptTcp.idx)
		return false
	}

	_FerrExit(" reading from tcp 831911 ", ___VacceptTcp.Verr)

	_FnullExit(" 183813 : why ___Vconn.ReadFromTCP addr error ?", ___VacceptTcp.VremoteAddr)

	___VacceptTcp.r64ok++

	/*
	   _Fpf( " 183814 | l:%s | r:%s | " , ___VacceptTcp.VlocalAddr , ___VacceptTcp.VremoteAddr )
	   _PpdN( ___VacceptTcp.Vlen , &(___VacceptTcp.Vbuf) )
	*/

	_FcopyByte(&(___VacceptTcp.Vbuf2), &(___VacceptTcp.Vbuf), ___VacceptTcp.Vlen)
	//___VacceptTcp.Vbuf2 = make([]byte , ___VacceptTcp.Vlen ); copy( ___VacceptTcp.Vbuf2 , ___VacceptTcp.Vbuf )

	// _FcallbackForDebugLog_accept
	_FnotNullRunTcp02_accept(___VacceptTcp.serverTCP.TcallbackAccDataRece, ___VacceptTcp)

	return true
} // _FhandleTcp_accept_dataReceiveMsg01__loop

func _FhandleTcp_accept_dataReceiveMsg01__debug(___VacceptTcp *_TacceptTCP) {

	_Ppn(" 183891 : under constructing ", ___VacceptTcp.r64try, ___VacceptTcp.r64ok)

	if nil == ___VacceptTcp.serverTCP {
		_Ppn(" 183892 : ___VacceptTcp.serverTCP == nil ")
	} else {

		_Ppf(" 183893 : ")
		for __Vi := 0; __Vi < ___VacceptTcp.serverTCP.cAmount; __Vi++ {
			__VacceptTcp := &(___VacceptTcp.serverTCP.acceptTCPs[__Vi])
			_Ppf(" %d,%d,%d,%d", __Vi, __VacceptTcp.idx, __VacceptTcp.r64try, __VacceptTcp.r64ok)
		}
		_Ppn()
	}
} // _FhandleTcp_accept_dataReceiveMsg01__debug
