package main

import "sync"

var ___VtcpBufMachine__LastTime int

var ___VtcpBufMachine__1500201__mutex sync.Mutex

func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__1500201x__chan_rece__default() {
	for {

		select {

		case __Vb := <-___Vtbm.tbmCHtcpReceBI:
			___VtcpBufMachine__1500201__mutex.Lock()

			__FpfN(" 398381 01 : _TtcpBufMachine received {%s}", String9s(&__Vb))
			___Vtbm.
				_FtcpBufMachine__1500201y1__chan_rece__decode(&__Vb)

		}

		___VtcpBufMachine__1500201__mutex.Unlock()
	}
}
