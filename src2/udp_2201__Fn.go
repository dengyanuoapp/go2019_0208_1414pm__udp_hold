package main

import (
	"flag"
)

var (
	_VserviceUdpFn _TserviceUDP
	_VserviceUdpFp _TserviceUDP
	_VserviceUdpFD _TserviceUDP
	_VserviceUdpFS _TserviceUDP

	_VfilterFn2dn _TfilterDelay

	_VserviceTcpMf _TserviceTCP

	_Cexit chan string
	_Clog  chan string
)

func init() {

	_VprojectName = "Fn"

	_Fbase_1203__gen_self_md5_sha()
	_Fbase_103__gen_rand_seed()
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

	_VserviceUdpFn = _TserviceUDP{
		name:        "servicePortForCn",
		UcallbackMR: _FuserCallback_u01M__dataRece_Cn,
		Cexit:       &_Cexit,
		Clog:        &_Clog,
	}
	_VserviceUdpFp = _TserviceUDP{
		name: "servicePortForCp",
		//UcallbackMR            : _FuserCallback_u01M__dataRece_Cp,
		Cexit: &_Cexit,
		Clog:  &_Clog,
	}

	_VserviceUdpFD = _TserviceUDP{
		name:        "servicePortForCD",
		UcallbackMR: _FuserCallback_u01M__dataRece_Fn__main_top,
		UcallbackCI: _FuserCallback_chanIn_Fn__main_top,
		Cexit:       &_Cexit,
		Clog:        &_Clog,
	}

	_VserviceUdpFS = _TserviceUDP{
		name:        "servicePortForCS",
		UcallbackMR: _FuserCallback_u01M__dataRece_Sn,
		Cexit:       &_Cexit,
		Clog:        &_Clog,
	}

	flag.StringVar(&_VserviceUdpFn.hostPortStr, "cn", ":5353", _VserviceUdpFn.name)
	flag.StringVar(&_VserviceUdpFp.hostPortStr, "cp", ":32001", _VserviceUdpFp.name)
	flag.StringVar(&_VserviceUdpFD.hostPortStr, "cd", ":32002", _VserviceUdpFD.name)
	flag.StringVar(&_VserviceUdpFS.hostPortStr, "cs", ":32003", _VserviceUdpFS.name)

	flag.Parse()

	// _FdebugPrintTest()

	_Cexit = make(chan string, 3)
	_Clog = make(chan string, 100)
}

func main() {

	// ------------------- tcp for debug monitor log --- begin
	// _Fhandle_tcpAccept01
	// _FhandleTcp_accept_dataReceiveMsg01
	go _VserviceTcpMf._Fhandle_udpListen_Tcp__main_top()
	// ------------------- tcp for debug monitor log --- end

	// ------------------- udp for worker clinet : Cn , Dn , Sn --------- begin
	// _TserviceUDP
	go _VserviceUdpFn._Fhandle_u01x__udpListen_Udp__read_main_top()
	go _VserviceUdpFp._Fhandle_u01x__udpListen_Udp__read_main_top()
	go _VserviceUdpFD._Fhandle_u01x__udpListen_Udp__read_main_top()
	go _VserviceUdpFS._Fhandle_u01x__udpListen_Udp__read_main_top()
	// ------------------- udp for worker clinet : Cn , Dn , Sn --------- end

	// ------------------- filter between workers --------- begin
	_VfilterFn2dn = _TfilterDelay{
		sleepGap:              1,
		udpIn:                 &_VserviceUdpFn,
		udpOut:                &_VserviceUdpFD,
		FcallbackMainDelayGen: _FuserCallback__FilterDelay__main_swap_signal_gen__Fn,
		FcallbackFilterChan:   _FuserCallback__FilterDelay__chan_filter__Fn,
	}

	go _VfilterFn2dn._FfilterDelayGen01_main_top()
	// ------------------- filter between workers --------- end

	_Fex(" the reason exit : "+<-_Cexit, nil)
} // main
