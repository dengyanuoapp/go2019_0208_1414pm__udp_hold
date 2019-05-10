package main

// _TacceptTCP
// _TtcpNodE
func (___Vacc *_TacceptTCP) _FtcpNodeAccept_sendOneTunnel(___Vs *_TtcpNodeDataSend) {
	defer ___Vacc.taServerTCP.tnClientMux.Unlock()
	___Vacc.taServerTCP.tnClientMux.Lock()

	if false == ___Vacc.taEnabled {
		_FpfN(" 838187 01 : already lost connect:%s", ___Vacc.String())
		return
	}

	if nil == ___Vs {
		_FpfN(" 838187 02 : why Null ? ")
		return
	}

	if 0 == ___Vs.tnsLen {
		_FpfN(" 838187 03 : why ZERO len ? ")
		return
	}

	if ___Vs.tnsLen != len(___Vs.tnsBuf) {
		_FpfN(" 838187 04 : why diff len ? ")
		return
	}

	if nil == ___Vacc.taConnTCP {
		_FpfN(" 838187 05 : why conn NULL ? ")
		return
	}

	//_FpfN(" 838187 06 : <%v>sending %s", ___Vacc.taConnTCP, ___Vacc.String())

	//func (c *TCPConn) Write(b []byte) (int, error)
	___Vacc.
		taConnTCP.
		Write(___Vs.tnsBuf)

}
