// _TacceptTCP
package main

import ()

type _TfilterDelay struct {
	sleepGap int
	udpIn    *_TserviceUDP
	udpOut   *_TserviceUDP

	CfSwap01 chan string
	CfIn01   chan []byte  // Filter's inChain,  to the sender's OutChain , make in Filter , then add to the sender's outChain-pointer
	CfOut01  *chan []byte // Filter's outChain, to the receiver's InChain

	Fusercallback__501_fileterMainTop  func(*_TfilterDelay) // _FfilterDelay501__main_top__default
	Fusercallback__511_filterTheChanIn func(*_TfilterDelay) // _FuserCallback__filterDelay_chan_from_FnWaitCn_to_FnWaitDn
	Fusercallback__521_delayGapAction  func(*_TfilterDelay) // _FuserCallback__filterGapAction_gen_a_signal_to_swapChan_when_timeout

	Cexit *chan string
	Clog  *chan string
} // _TfilterDelay

func (___Vf *_TfilterDelay) IRun(___Vidx int) {
	switch ___Vidx {
	case 501:
		if nil == ___Vf.Fusercallback__501_fileterMainTop {
			_FfilterDelay501__main_top__default(___Vf)
		} else {
			___Vf.Fusercallback__501_fileterMainTop(___Vf)
		}
	default:
		_FpfNex(" 848182 09 : unknown IRun : %d ", ___Vidx)
	} // switch ___Vidx
}

//    _VserviceUdpWdn._Fhandle_u01x__udpListen_Udp__read_main_top__default
func _FfilterDelay501__main_top__default(___Vf *_TfilterDelay) {

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
} // _FfilterDelay501__main_top__default

func (___Vf *_TfilterDelay) _FfilterDelayGen01_main_loop() {
	//_FpfN( " 418113 : filter main " )
	if nil != ___Vf.Fusercallback__521_delayGapAction {
		//_FpfN( " 418115 : filter main " )
		___Vf.Fusercallback__521_delayGapAction(___Vf) // _FuserCallback__filterGapAction_gen_a_signal_to_swapChan_when_timeout
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
	if nil != ___Vf.Fusercallback__511_filterTheChanIn {
		___Vf.Fusercallback__511_filterTheChanIn(___Vf) // _FuserCallback__filterDelay_chan_from_FnWaitCn_to_FnWaitDn
	}
} // _FfilterDelayGen01_filter_loop
