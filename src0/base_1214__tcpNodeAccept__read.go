// _TacceptTCP
// _TtcpNodE
// _FtcpNode__200401x_accept_default
package main

import (
	"io"
	//"sync"
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

//var ___VtcpNodeAccept__mux sync.Mutex

// func (c *TCPConn) Read(b []byte) (int, error)
// _TacceptTCP
func _FtcpNodeAccept__200401x5__dataReceiveMsg01_default(___VtAcc3 *_TacceptTCP) bool {

	__VtrBuf := make([]byte, 1024)
	__VtrLen, __VtrErr := ___VtAcc3.taConnTCP.Read(__VtrBuf)

	___TtcpNodE__mux.Lock()

	//defer ___VtcpNodeAccept__mux.Unlock()
	//___VtcpNodeAccept__mux.Lock()

	___VtAcc3.taRcnt.try++ // _TacceptTCP

	if false == ___VtAcc3.taEnabled || __VtrErr == io.EOF { // lost the connect. or another thread close it.
		___VtAcc3._FtcpNodeAccept__200401x6__dataReceiveMsg01_removeLostconnect()
		return false
	}

	_FerrExit(" 183813 08 reading from tcp ", __VtrErr)

	_FnullExit(" 183813 09 : why ___Vconn.ReadFromTCP addr error ?", ___VtAcc3.taRemoteAddr)

	___VtAcc3.taRcnt.ok++

	__VtrData := _TtcpNodeDataRece{
		TnrRaddr:  ___VtAcc3.taRemoteAddr.String(),
		TnrLen:    __VtrLen,
		TnrId128:  ___VtAcc3.taId128,
		TnrBuf:    __VtrBuf[:__VtrLen],
		TnrOffset: ___VtAcc3.taOffset,
	}
	___VtAcc3.taOffset += int64(__VtrLen)

	//func _FencGobExit(___VeMsg string, ___V interface{}) []byte {
	__VtrDataB := _FencGobExit(" 183813 05 ", &__VtrData)

	_FnullExit(" 183813 06 : why tcpNode pointer NULL ? ", ___VtAcc3.taServerTCP)

	if _FchanNullCheckOk(" 183813 07 "+___VtAcc3.taServerTCP.tnName, ___VtAcc3.taServerTCP.tnCHtcpReceBLO) {

		__FpfN(" 183813 09 : tcp rece :{%s} ============= acc{%s} ",
			__VtrData.String(), ___VtAcc3.String()) // _TtcpNodeDataRece

		(*___VtAcc3.taServerTCP.tnCHtcpReceBLO) <- __VtrDataB // _TtcpNodE
	}

	if nil != ___VtAcc3.taServerTCP.tnCHdebugInfoLO {
		(*(___VtAcc3.taServerTCP.tnCHdebugInfoLO)) <- 1
	}
	___TtcpNodE__mux.Unlock()
	return true
} //

func (___VtAcc3 *_TacceptTCP) _FtcpNodeAccept__200401x6__dataReceiveMsg01_removeLostconnect() {
	___VtAcc3.taRcnt.eofErr++
	// acceptTcpINC / acceptTcpDEC : begin
	___VtAcc3.taServerTCP.tnClientMux.Lock()

	___VtAcc3.taServerTCP.tnClientCnt--
	if true == ___VtAcc3.taEnabled {
		___VtAcc3.taConnTCP.Close()
	} // else already closed.
	___VtAcc3.taEnabled = false

	_CFpfN(" 183815 02 : eof, delete old-id:{%s}", String9s(&___VtAcc3.taId128))

	if nil != ___VtAcc3.taServerTCP.tnCHtcpReceCmdLO {
		_CFpfN(" ######################3 183815 03 : send note to the Users:{%s}", String9s(&___VtAcc3.taId128))
		__VcmdB := [17]byte{}
		copy(__VcmdB[:], ___VtAcc3.taId128)
		__VcmdB[16] = TcpNodeCmd__Eof
		(*(___VtAcc3.taServerTCP.tnCHtcpReceCmdLO)) <- __VcmdB // _TtcpNodE
	} else {
		_CFpfN(" ######################3 183815 04 : why no CmdLO ? {%s}", String9s(&___VtAcc3.taId128))
	}

	//___VtAcc3.taId128 = _FgenRand_nByte__(16) // regen new id // no need to do here, the new accetp will re-gen
	___VtAcc3.taId128 = []byte{}

	___VtAcc3.taServerTCP.tnClientMux.Unlock()
	// acceptTcpINC / acceptTcpDEC : end

	_CFpfN(" 183815 07 : eof, TCP end l:%v , r:%v , id:{%s}",
		___VtAcc3.taLocalAddr, ___VtAcc3.taRemoteAddr, String9s(&___VtAcc3.taId128))
	//___VtAcc3.taCreceiveErr <- _Pspf("EOF:%d", ___VtAcc3.taIdx)
	___TtcpNodE__mux.Unlock()
}

func _FtcpNodeAccept__200401x5__dataReceiveMsg01_debug(___VtAcc4 *_TacceptTCP) {

	_Ppln(" 183891 : under constructing ", ___VtAcc4.taRcnt.try, ___VtAcc4.taRcnt.ok)

	if nil == ___VtAcc4.taServerTCP {
		_Ppln(" 183892 : ___VtAcc4.taServerTCP == nil ")
	} else {

		_Ppf(" 183893 : ")
		for __Vi := 0; __Vi < ___VtAcc4.taServerTCP.tnAmount; __Vi++ {
			__VacceptTcp := &(___VtAcc4.taServerTCP.tnAcceptTCPs[__Vi])
			_Ppf(" %d,%d,%d,%d", __Vi, __VacceptTcp.taIdx, __VacceptTcp.taRcnt.try, __VacceptTcp.taRcnt.ok)
		}
		_Ppln()
	}
} // _FtcpNodeAccept__200401x5__dataReceiveMsg01_debug
