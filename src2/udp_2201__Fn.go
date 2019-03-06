package main

import (
	"flag"
)

type _Tconfig struct {
	Name    string
	MyId128 []byte
}

var (
	_VserviceUdpWcn _TserviceUDP
	_VserviceUdpWdn _TserviceUDP

	_VfilterFn2dn _TfilterDelay

	_VserviceTcpMf _TserviceTCP

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

	_VserviceUdpWcn = _TserviceUDP{
		name:           "servicePortForCn",
		UsrvLoopCall11: _FuserCallback_u01M__dataRece_Cn,
		Cexit:          &_Cexit,
		Clog:           &_Clog,
	}
	_VserviceUdpWdn = _TserviceUDP{
		name:           "servicePortForCD",
		UsrvLoopCall11: _FuserCallback_u01M__dataRece_Fn__main_top,
		UsrvLoopCall12: _FuserCallback_chanIn_Fn__main_top,
		Cexit:          &_Cexit,
		Clog:           &_Clog,
	}

	//	_VserviceUdpFS = _TserviceUDP{
	//		name:        "servicePortForCS",
	//		UsrvLoopCall11: _FuserCallback_u01M__dataRece_Sn,
	//		Cexit:       &_Cexit,
	//		Clog:        &_Clog,
	//	}

	flag.StringVar(&_VserviceUdpWcn.hostPortStr, "FnWcn", ":53535", _VserviceUdpWcn.name)
	flag.StringVar(&_VserviceUdpWdn.hostPortStr, "cd", ":32001", _VserviceUdpWdn.name)

	flag.Parse()

	// _FdebugPrintTest()

	_Cexit = make(chan string, 3)
	_Clog = make(chan string, 100)
}

func main() {

	// ------------------- tcp for debug monitor log --- begin
	// _Fhandle_tcpAccept01
	// _FhandleTcp_accept_dataReceiveMsg01
	//go _VserviceTcpMf._Fhandle_udpListen_Tcp__main_top()
	go _Frun(&_VserviceTcpMf, 1)
	// ------------------- tcp for debug monitor log --- end

	// ------------------- udp for worker clinet : Cn , Dn , Sn --------- begin
	// _TserviceUDP
	go _Frun(&_VserviceUdpWcn, 1) // IRun // _Fhandle_u01x__udpListen_Udp__read_main_top__default
	go _Frun(&_VserviceUdpWdn, 1)
	// ------------------- udp for worker clinet : Cn , Dn , Sn --------- end

	// ------------------- filter between workers --------- begin
	_VfilterFn2dn = _TfilterDelay{
		sleepGap:              1,
		udpIn:                 &_VserviceUdpWcn,
		udpOut:                &_VserviceUdpWdn,
		FcallbackMainDelayGen: _FuserCallback__FilterDelay__main_swap_signal_gen__Fn,
		FcallbackFilterChan:   _FuserCallback__FilterDelay__chan_filter__Fn,
	}

	go _VfilterFn2dn._FfilterDelayGen01_main_top()
	// ------------------- filter between workers --------- end

	_Fex(" the reason exit : "+<-_Cexit, nil)
} // main
