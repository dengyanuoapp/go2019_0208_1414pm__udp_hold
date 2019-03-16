package main

import (
	"flag"
)

var (
	_VserviceTcpMd   _TserviceTCP
	_VudpGroup_Dn2Fn _TudpGroupSt
	_Cexit           chan string
	_Clog            chan string
	_Vself           _Tself
	_Vconfig         _Tconfig
)

func init() {

	_Vself.ProjName = "Dn"

	_Fbase_1203__init_self_All()

	_Fbase_104z__try_to_read_json_config_top()

	_Fbase_107__rand_init()

	_FPargs()

	_VserviceTcpMd = _TserviceTCP{
		name:        "TcpService__DebugLog__Md",
		hostPortStr: "127.0.0.1:56782",
		//TcallbackSvrDataChan    : _FuserCallback__Log_service_dataChan__Dn,
		//TcallbackAccDataRece    : _FuserCallback__LogAccept_dataReceive__Dn,
		//TcallbackAccDataChan    : _FuserCallback__accept_dataChan__Log_Dn,

		Cexit:   &_Cexit,
		Clog:    &_Clog,
		cAmount: 10,
	}

	// _FuserCallback_UdataRece_Dn
	// _FuserCallback_UdataMain_Dn
	//	_VserviceUdpDn = _TserviceUDP{
	//		name:   "UdpService__Dn",
	//		uExtMR: &_VuExtMR_Dn,
	//
	//		UreqNewSessionTM: &_VUreqNewSession_Dn,
	//
	//		Cexit: &_Cexit,
	//		Clog:  &_Clog,
	//	}
	//
	//	flag.StringVar(&_VserviceUdpDn.hostPortStr, "cn", ":0", _VserviceUdpDn.name)

	_VudpGroup_Dn2Fn := _TudpGroupSt{
		ugName:   "udpGroup_Dn2Fn",
		ugAmount: 10,
	}
	flag.StringVar(&_VudpGroup_Dn2Fn.ugHostPortStr, "cn", ":0", _VudpGroup_Dn2Fn.ugName)

	flag.Parse()
	_FpfN(" 834811 97 : exit.[%#v]", _VudpGroup_Dn2Fn)

	// _FdebugPrintTest()

	_Cexit = make(chan string, 3)
	_Clog = make(chan string, 100)
}

func main() {

	// ------------------- tcp for debug monitor log --- begin
	// _Fhandle_tcpAccept01
	// _FhandleTcp_accept_dataReceiveMsg01
	//go _VserviceTcpMd._Fhandle_udpListen_Tcp__main_top()
	go _Frun(&_VserviceTcpMd, 101)
	// ------------------- tcp for debug monitor log --- end

	// ------------------- udp for worker clinet : Cn , Dn , Sn --------- begin
	// _TserviceUDP
	//go _VserviceUdpDn._Fhandle_u01x__udpListen_Udp__read_main_top__default()
	//go _Frun(&_VserviceUdpDn, 201)
	// ------------------- udp for worker clinet : Cn , Dn , Sn --------- end

	_FpfNex(" 834811 98 : exit.[%#v]", _VudpGroup_Dn2Fn)
	// _TudpNodeSt _TudpGroupSt
	go _Frun(&_VudpGroup_Dn2Fn, 150201) // IRun _FudpGroup__150201__main_init__default

	// ------------------- filter between workers --------- begin
	//	_VfilterCn2dn := _TfilterDelay{
	//		sleepGap: 1,
	//		udpIn:    &_VserviceUdpDn,
	//		udpOut:   &_VserviceUdpDC,
	//		//Fusercallback__521_delayGapAction   : _FuserCallback__521_filterGapAction_gen_a_signal_to_swapChan_when_timeout__default,
	//		//Fusercallback__511_filterTheChanIn     : _FuserCallback__511_filterDelay_chan_from_FnWaitCn_to_FnWaitDn,
	//	}

	//go _VfilterCn2dn . _FfilterDelay501__main_top__default()
	// ------------------- filter between workers --------- end

	_Fex(" the reason exit : "+<-_Cexit, nil)
} // main
