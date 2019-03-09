package main

import (
	"flag"
)

var (
	_VserviceUdpDn _TserviceUDP

	_VfilterCn2dn _TfilterDelay

	_VserviceTcpMd _TserviceTCP

	_Cexit chan string
	_Clog  chan string

	_Vself   _Tself
	_Vconfig _Tconfig
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
	_VserviceUdpDn = _TserviceUDP{
		name:   "UdpService__Dn",
		uExtMR: &_VuExtMR_Dn,

		UreqNewSessionTM: &_VUreqNewSession_Dn,

		Cexit: &_Cexit,
		Clog:  &_Clog,
	}

	flag.StringVar(&_VserviceUdpDn.hostPortStr, "cn", ":0", _VserviceUdpDn.name)

	flag.Parse()

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
	go _Frun(&_VserviceUdpDn, 201)
	// ------------------- udp for worker clinet : Cn , Dn , Sn --------- end

	// ------------------- filter between workers --------- begin
	//	_VfilterCn2dn = _TfilterDelay{
	//		sleepGap: 1,
	//		udpIn:    &_VserviceUdpDn,
	//		udpOut:   &_VserviceUdpDC,
	//		//FcallbackMainDelayGen   : _FuserCallback__FilterDelay__main_swap_signal_gen__Fn,
	//		//FcallbackFilterChan     : _FuserCallback__FilterDelay__chan_filter__Fn,
	//	}

	//go _VfilterCn2dn . _FfilterDelayGen01_main_top()
	// ------------------- filter between workers --------- end

	_Fex(" the reason exit : "+<-_Cexit, nil)
} // main
