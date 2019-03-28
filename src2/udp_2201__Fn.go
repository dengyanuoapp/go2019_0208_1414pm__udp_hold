package main

import (
	"flag"
)

var (
	_VudpNode_FunWaitDun _TudpNodeSt

	_VserviceTcpMf _TserviceTCP

	_VudpTimer01 _TgapTimer

	_Cexit   chan string
	_Clog    chan string
	_Vself   _Tself
	_Vconfig _Tconfig
)

func _Finit_2201() {

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
		unName:    "_VudpNode_FunWaitDun",
		unRKeyLP:  &_Vpasswd_udp_Fn_waitForCliens01,
		unLoopGap: _T10s,
	}

	_VudpTimer01 = _TgapTimer{
		uTmCHunDataReceLI: &_VudpNode_FunWaitDun,
	}

	flag.StringVar(&_VudpNode_FunWaitDun.unHostPortStr, "FunWdun", ":32001", _VudpNode_FunWaitDun.unName)

	flag.Parse()

	// _FdebugPrintTest()

	_Cexit = make(chan string, 3)
	_Clog = make(chan string, 100)
}

func main() {
	_Finit_2201()

	// ------------------- tcp for debug monitor log --- begin
	// IRun _Fhandle_udpListen_Tcp__main_top()
	go _Frun(&_VserviceTcpMf, 200101)
	// ------------------- tcp for debug monitor log --- end

	// _FudpNode__540211z__receiveCallBack_withTimeGap
	// _FdataPack__301__dataDecode_loginS1ReqTryNoToken
	go _Frun(&_VudpNode_FunWaitDun, 540201) // IRun _FudpNode__540201__main_init__default

	// _FudpTimer__750201x__gap_receive__default
	go _Frun(&_VudpTimer01, 750101) // IRun _FudpTimer__750101x__init__default

	_Fex(" 893189 99 : the reason exit : "+<-_Cexit, nil)
} // main
