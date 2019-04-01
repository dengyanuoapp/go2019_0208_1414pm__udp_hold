package main

import (
	"flag"
)

var (
	_VudpNode_FunWaitDun   _TudpNodeSt
	_VserviceTcpMf         _TserviceTCP
	_VudpDecode01          _TuDecode
	_VloginCheck_FnWaitDun _TloginCheck
	_VudpGroup_Fn          _TudpGroupSt
	_Cexit                 chan string
	_Clog                  chan string
	_Vself                 _Tself
	_Vconfig               _Tconfig
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
		TcallbackAccDataChan: _FuserCallback__accept_dataChan__Log_Fn, // _FhandleTcp__accept_dataChan__main_top
		Cexit:                &_Cexit,
		Clog:                 &_Clog,
		cAmount:              10,
	}

	_VudpNode_FunWaitDun = _TudpNodeSt{
		unName:    "_VudpNode_FunWaitDun",
		unRKeyLP:  &_Vpasswd_udp_Fn_waitForCliens01,
		unLoopGap: _T10s,
	}

	_VudpDecode01 = _TuDecode{
		uTmCHunDataReceLI: &_VudpNode_FunWaitDun,
		uTmDecodeCmdLO:    &_VloginCheck_FnWaitDun.ucDecodeI, // _TloginCheck _Tdecode
	}

	_VloginCheck_FnWaitDun = _TloginCheck{
		ucCHSendLO: &_VudpGroup_Fn.ugCHSendI, // _TudpGroupSt _TudpNodeDataSend
	}

	_VudpGroup_Fn = _TudpGroupSt{
		ugName:        "udpGroup_Fn",
		ugAmount:      20,
		ugHostPortStr: []string{":0"},
	}
	flag.StringVar(&_VudpGroup_Fn.ugHostPortStr[0], "cn", ":0", _VudpGroup_Fn.ugName)

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
	// _FdataPack__decode_from_udpNodeDataRece
	go _Frun(&_VudpNode_FunWaitDun, 500101) // IRun _FudpNode__500101__main_init__default

	// _FudpDecode__700201x__receive__default
	go _Frun(&_VudpDecode01, 700101) // IRun _FudpDecode__700101x__init__default

	// _FloginCheck__900201x__standardCheck
	// _FloginCheck_step900201y__sReply_tokenB
	go _Frun(&_VloginCheck_FnWaitDun, 900101) // _FloginCheck__900101x__init__default

	// _TudpNodeSt _TudpGroupSt
	go _Frun(&_VudpGroup_Fn, 600101) // IRun _FudpGroup__600101__main_init__default
	// _FudpGroup__600201__CHin_select_send__default
	// _FdataPack__101__udpConnPort

	_Fex(" 893189 99 : the reason exit : "+<-_Cexit, nil)
} // main
