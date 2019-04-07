package main

func (___Vlg *_TloginGenerator) IRun(___Vidx int) {
	switch ___Vidx {
	case 800101:
		if nil == ___Vlg.lgCB850101init {
			_FudpDecode__800101x__init__tryUdpLogin__default(___Vlg)
		} else {
			___Vlg.lgCB850101init(___Vlg)
		}
	case 800201:
		if nil == ___Vlg.lgCB850201chRece {
			_FudpDecode__800201x__chan_in__default(___Vlg)
		} else {
			___Vlg.lgCB850201chRece(___Vlg)
		}
	case 800301:
		if nil == ___Vlg.lgCB850301ConnGen {
			_FudpDecode__800301x__connGen__default(___Vlg)
		} else {
			___Vlg.lgCB850301ConnGen(___Vlg)
		}
	default:
		_FpfNex(" 739182 99 unknow :idx %d", ___Vidx)
	}
}

// _TloginReq
// _TsrvInfo
// _TuDecode
// _TloginGenerator
// lgSrvDownInfoLX *_TsrvDownInfo
func _FudpDecode__800101x__init__tryUdpLogin__default(___Vlg *_TloginGenerator) {

	___Vlg.lgCHdataMachineIdI = make(chan _TdataMachinEid, 10)

	// ___Vlg.lgSrvDownInfoLX
	go _Frun(___Vlg, 800301) // _FudpDecode__800301x__connGen__default

	go _Frun(___Vlg, 800201) // _FudpDecode__800201x__chan_in__default
}
