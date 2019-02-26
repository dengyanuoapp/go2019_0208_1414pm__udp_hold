// _TacceptTCP
package main

import ()

type _TfilterDelay struct {
	sleepGap int
	udpIn    *_TserviceUDP
	udpOut   *_TserviceUDP

	CfSwap01 chan string
	CfIn01   chan []byte
	CfOut01  *chan []byte

	FcallbackMainDelayGen func(*_TfilterDelay) // _FuserCallback__FilterDelay__main_swap_signal_gen__Fn
	FcallbackFilterChan   func(*_TfilterDelay) // _FuserCallback__FilterDelay__chan_filter__Fn

	Cexit *chan string
	Clog  *chan string
} // _TfilterDelay

//    _Fhandle_u01x__udpListen_Udp__read_main_top
func (___Vf *_TfilterDelay) _FfilterDelayGen01_main_top() {

	if 1 > ___Vf.sleepGap {
		_Fex(" 418111 : error sleep gap ", nil)
	}

	___Vf.CfSwap01 = make(chan string, 4)
	___Vf.CfIn01 = make(chan []byte, 5)
	___Vf.udpIn.CuOut01 = &(___Vf.CfIn01)

	go ___Vf._FfilterDelayGen01_filter_top()

	//_Fsleep_1s()
	___Vf.CfOut01 = &(___Vf.udpOut.CuIn01)

	for {
		_Fsleep_10sX(___Vf.sleepGap)
		___Vf._FfilterDelayGen01_main_loop()
	}
} // _FfilterDelayGen01_main_top

func (___Vf *_TfilterDelay) _FfilterDelayGen01_main_loop() {
	//_FpfN( " 418113 : filter main " )
	if nil != ___Vf.FcallbackMainDelayGen {
		//_FpfN( " 418115 : filter main " )
		___Vf.FcallbackMainDelayGen(___Vf) // _FuserCallback__FilterDelay__main_swap_signal_gen__Fn
	}
} // _FfilterDelayGen01_main_loop

func (___Vf *_TfilterDelay) _FfilterDelayGen01_filter_top() {
	//_FpfN( " 421191 : filter main " )
	for {
		//_Fsleep_1s() ;
		_Fsleep_1ms()
		___Vf._FfilterDelayGen01_filter_loop()
	}
} // _FfilterDelayGen01_filter_top

func (___Vf *_TfilterDelay) _FfilterDelayGen01_filter_loop() {
	//_FpfN( " 421193 : filter main " )
	if nil != ___Vf.FcallbackFilterChan {
		___Vf.FcallbackFilterChan(___Vf) // _FuserCallback__FilterDelay__chan_filter__Fn
	}
} // _FfilterDelayGen01_filter_loop
