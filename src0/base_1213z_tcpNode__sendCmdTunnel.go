package main

// _TacceptTCP
// _TtcpNodE
func (___Vacc *_TacceptTCP) _FtcpNodeAccept_sendOneCmdTunnel(___Vcmd [17]byte) {
	defer ___Vacc.taServerTCP.tnClientMux.Unlock()
	___Vacc.taServerTCP.tnClientMux.Lock()

	___Vacc.taEnabled = false
	___Vacc.
		taConnTCP.
		Close()
}
