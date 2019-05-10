package main

import "bytes"

func (___VtcpNode4 *_TtcpNodE) _FtcpNode__200301x2_sendAll__default(___Vds *_TtcpNodeDataSend) {
	if 0 == len(___Vds.tnsId128) {
		for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
			if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled {
				//_FpfN("283822 03 get from send-chain :{%s}", ___Vds.String())
				___VtcpNode4.
					tnAcceptTCPs[__Vi].
					_FtcpNodeAccept_sendOneTunnel(___Vds)
			}
		}
	} else {
		for __Vi := 0; __Vi < ___VtcpNode4.tnAmount; __Vi++ {
			if ___VtcpNode4.tnAcceptTCPs[__Vi].taEnabled && (bytes.Equal(___Vds.tnsId128, ___VtcpNode4.tnAcceptTCPs[__Vi].taId128)) {
				_FpfN("283822 05 get from send-chain :{%s}", ___Vds.String())
				___VtcpNode4.
					tnAcceptTCPs[__Vi].
					_FtcpNodeAccept_sendOneTunnel(___Vds)
			}
		}
	}
}
