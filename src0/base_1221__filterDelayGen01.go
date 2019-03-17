// _TacceptTCP
package main

import ()

type _TfilterDelay struct {
	name     string
	sleepGap int
	udpIn    *_TserviceUDP
	udpOut   *_TserviceUDP

	CfSwap01 chan string
	CfIn01   chan []byte  // Filter's inChain,  to the sender's OutChain , make in Filter , then add to the sender's outChain-pointer
	CfOut01  *chan []byte // Filter's outChain, to the receiver's InChain

	Fusercallback__12501_fileterMainTop func(*_TfilterDelay) // _FfilterDelay501__main_top__default
	Fusercallback__12521_delayGapAction func(*_TfilterDelay) // _FuserCallback__521_filterGapAction_gen_a_signal_to_swapChan_when_timeout__default

	Fusercallback__12511_filterTheChanIn func(*_TfilterDelay) // _FuserCallback__511_filterDelay_chan_from_FnWaitCn_to_FnWaitDn

	Cexit *chan string
	Clog  *chan string
} // _TfilterDelay

func (___Vf *_TfilterDelay) IRun(___Vidx int) {
	switch ___Vidx {
	case 12501:
		if nil == ___Vf.Fusercallback__12501_fileterMainTop {
			_FfilterDelay501__main_top__default(___Vf)
		} else {
			___Vf.Fusercallback__12501_fileterMainTop(___Vf)
		}
	case 12511:
		if nil == ___Vf.Fusercallback__12511_filterTheChanIn {
			_Fex1(" 848182 511 : you must define Fusercallback__12511_filterTheChanIn ")
		} else {
			___Vf.Fusercallback__12511_filterTheChanIn(___Vf)
		}
	case 12521:
		if nil == ___Vf.Fusercallback__12521_delayGapAction {
			_FuserCallback__521_filterGapAction_gen_a_signal_to_swapChan_when_timeout__default(___Vf)
		} else {
			___Vf.Fusercallback__12521_delayGapAction(___Vf)
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
		_Frun(___Vf, 521) // _FuserCallback__521_filterGapAction_gen_a_signal_to_swapChan_when_timeout__default
	}
} // _FfilterDelay501__main_top__default

var __V521_filteSwapCNT int

func _FuserCallback__521_filterGapAction_gen_a_signal_to_swapChan_when_timeout__default(___Vf *_TfilterDelay) {
	//_FpfN(" 818391 01: filter swap user start : %d =========== ", __V521_filteSwapCNT)
	__V521_filteSwapCNT++
	___Vf.CfSwap01 <- " 818391 03 : ok, time Gap reach, action please." + _FtimeNow() // It's time to swap
} // _FuserCallback__521_filterGapAction_gen_a_signal_to_swapChan_when_timeout__default

func (___Vf *_TfilterDelay) _FfilterDelayGen01_filter_top() {
	//_FpfN( " 421191 : filter main " )
	for {
		//_Fsleep_1s() ;
		_Fsleep_1ms()
		_Frun(___Vf, 511) // _FuserCallback__511_filterDelay_chan_from_FnWaitCn_to_FnWaitDn
	}
} // _FfilterDelayGen01_filter_top
