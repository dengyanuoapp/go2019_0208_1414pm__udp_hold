package main

import (
	"flag"
)

var (
	_VserviceUdp_FnWaitCn _TserviceUDP
	//_VserviceUdp_FnWaitDn _TserviceUDP

	//_VfilterFn2dn _TfilterDelay

	_VudpNode_FunWaitDun _TudpNodeSt

	_VserviceTcpMf _TserviceTCP

	_VudpTimer01 _TudpTimer

	_Cexit   chan string
	_Clog    chan string
	_Vself   _Tself
	_Vconfig _Tconfig
)

func init() {

	_Vself.ProjName = "Fn"

	_Fbase_1203__init_self_All()

	_Fbase_104z__try_to_read_json_config_top()
	_Fbase_107__rand_init()

	_FPargs()

	_VserviceTcpMf = _TserviceTCP{
		name:                 "servicePortDebugLog",
		hostPortStr:          "127.0.0.1:56781",
		TcallbackSvrDataChan: _FuserCallback__service_dataChan__Log_Fn,
		TcallbackAccDataRece: _FuserCallback__Accept_dataReceive__Log_Fn,
		TcallbackAccDataChan: _FuserCallback__accept_dataChan__Log_Fn,
		// _FhandleTcp__accept_dataChan__main_top

		Cexit:   &_Cexit,
		Clog:    &_Clog,
		cAmount: 10,
	}

	_VudpNode_FunWaitDun = _TudpNodeSt{
		name:     "_VudpNode_FunWaitDun",
		unRKeyLP: &_Vpasswd_udp_Fn_waitForCliens01,
	}

	_VserviceUdp_FnWaitCn = _TserviceUDP{
		name:             "FnServicePortForCn",
		UuserLoopCall211: _FuserCallback_u01M__dataRece_Cn,
		Cexit:            &_Cexit,
		Clog:             &_Clog,
	}
	//	_VserviceUdp_FnWaitDn = _TserviceUDP{
	//		name:             "FnServicePortForDn",
	//		UuserLoopCall211: _FuserCallback__211_dataRece__main_top__FnWaitDn,
	//		UuserLoopCall221: _FuserCallback__221_chanIn__main_top__FnWaitDn,
	//
	//		VuSrvInfo: &_TsrvInfo{K256: _Vpasswd_udp_Fn_waitForCliens01},
	//
	//		Cexit: &_Cexit,
	//		Clog:  &_Clog,
	//	}

	_VudpTimer01 = _TudpTimer{
		gap:          _T10s,
		uTmCHudpRece: &_VudpNode_FunWaitDun,
		//uTmCHudpRece      chan _TudpNodeDataRece,
		//uTmCHudpSend      : &_VudpNode_FunWaitDun,
		//uTmCB150001init   func(*_TudpTimer)       ,
		//uTmCB150201filter func(*_TudpTimer)       ,
	}

	flag.StringVar(&_VserviceUdp_FnWaitCn.hostPortStr, "FnWcn", ":53535", _VserviceUdp_FnWaitCn.name)
	//flag.StringVar(&_VserviceUdp_FnWaitDn.hostPortStr, "FnWdn", ":32001", _VserviceUdp_FnWaitDn.name)
	//flag.StringVar(&_VudpNode_FunWaitDun.hostPortStr, "FunWdun", ":42001", _VudpNode_FunWaitDun.name)
	flag.StringVar(&_VudpNode_FunWaitDun.hostPortStr, "FunWdun", ":32001", _VudpNode_FunWaitDun.name)

	flag.Parse()

	// _FdebugPrintTest()

	_Cexit = make(chan string, 3)
	_Clog = make(chan string, 100)
}

func main() {

	// ------------------- tcp for debug monitor log --- begin
	// IRun _Fhandle_udpListen_Tcp__main_top()
	go _Frun(&_VserviceTcpMf, 101)
	// ------------------- tcp for debug monitor log --- end

	// ------------------- udp for worker clinet : Cn , Dn , Sn --------- begin
	// _TserviceUDP
	go _Frun(&_VserviceUdp_FnWaitCn, 201) // IRun // _Fhandle_u01x__udpListen_Udp__read_main_top__default
	//go _Frun(&_VserviceUdp_FnWaitDn, 201)
	// ------------------- udp for worker clinet : Cn , Dn , Sn --------- end

	// ------------------- filter between workers --------- begin
	//	_VfilterFn2dn = _TfilterDelay{
	//		name:                               "filter_FnReceFromCnDelayNoteDn",
	//		sleepGap:                           1,
	//		udpIn:                              &_VserviceUdp_FnWaitCn,
	//		udpOut:                             &_VserviceUdp_FnWaitDn,
	//		Fusercallback__511_filterTheChanIn: _FuserCallback__511_filterDelay_chan_from_FnWaitCn_to_FnWaitDn,
	//	}

	// IRun _FudpNode__140201__main_init__default
	go _Frun(&_VudpNode_FunWaitDun, 140201)

	// IRun _FudpTimer__150001x__init__default
	go _Frun(&_VudpTimer01, 150001)

	//go _VfilterFn2dn. _FfilterDelay501__main_top__default
	//go _Frun(&_VfilterFn2dn, 501) // IRun
	// ------------------- filter between workers --------- end

	_Fex(" 893189 99 : the reason exit : "+<-_Cexit, nil)
} // main
