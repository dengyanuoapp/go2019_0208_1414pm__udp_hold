package main

import "sync"

var ___VtcpBufMachine__LastTime int

var ___VtcpBufMachine__1500201__mutex sync.Mutex

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__1500201x__chan_rece__default() {
	for {

		select {

		case __Vb := <-___Vtbm.tbmCHtcpReceBI:
			___VtcpBufMachine__1500201__mutex.Lock()

			__FpfN(" 398381 01 : _TtcpBufMachine received Data{%s}", String9s(&__Vb))

			___Vtbm.
				_FtcpBufMachine__1500201y1__chan_rece__Local2Remote(&__Vb)

		case __VcmdB17 := <-___Vtbm.tbmCHtcpReceCmdI:
			___VtcpBufMachine__1500201__mutex.Lock()

			_CFpfN(" ###### 398381 03 : _TtcpBufMachine received Cmmd %d , {%x},", __VcmdB17[16], __VcmdB17[:3])

			___Vtbm.
				_FtcpBufMachine__local2remote_remove_an_tunnel(&__VcmdB17)

		}

		___VtcpBufMachine__1500201__mutex.Unlock()
	}
}
