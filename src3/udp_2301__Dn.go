package main

import (
	"flag"
)

var (
	_VserviceTcp_Dn     _TserviceTCP
	_VudpGroup_Dn       _TudpGroupSt
	_VudpDecode_Dn      _TuDecode
	_VloginGenerator_Dn _TloginGenerator
	_VdataMachine_Dn    _TdataMachine

	_VloginCheck_Dn _TloginCheck

	_Cexit   chan string
	_Clog    chan string
	_Vself   _Tself
	_Vconfig _Tconfig
)

func _Finit__2301() {

	_Vself.ProjName = "Dn"

	_Fbase_1203__init_self_All()

	_Fbase_104z__try_to_read_json_config_top()

	_Fbase_107__rand_init()

	_FPargs()

	_VserviceTcp_Dn = _TserviceTCP{
		name:        "TcpService__DebugLog__Md",
		hostPortStr: "127.0.0.1:50003",
		Cexit:       &_Cexit,
		Clog:        &_Clog,
		cAmount:     10,
	}

	_VudpDecode_Dn = _TuDecode{
		uTmDecodeCmdLO: &_VloginCheck_Dn.ulDecodeI, // _TloginCheck _Tdecode
	}

	_VloginCheck_Dn = _TloginCheck{
		ulCHSendLO:          &_VudpGroup_Dn.ugCHSendI,             // _TudpGroupSt _TudpNodeDataSend
		ulCHdataMachineIdLO: &_VdataMachine_Dn.dmCHdataMachineIdI, // *chan _TdataMachinEid
	}

	_VudpGroup_Dn = _TudpGroupSt{
		ugName:        "udpGroup_Dn",
		ugAmount:      10,
		ugHostPortStr: []string{":0"},
		ugCHreceLO:    &_VudpDecode_Dn.uTmCHunDataReceI,
	}

	flag.StringVar(&_VudpGroup_Dn.ugHostPortStr[0], "cn", ":0", _VudpGroup_Dn.ugName)

	flag.Parse()

	_VloginGenerator_Dn = _TloginGenerator{
		ulSrvDownInfoLX: &_TsrvDownInfo{
			name:         "srvDn2Fn",
			updateUri:    "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/FnWaitDn.gob.rand", // gob  json
			updatePasswd: _Vpasswd_udp_FnWaitDn_download_config,
		},
		//ulCHunSendLO: &_VudpGroup_Dn.ugCHSendI,
		ulCHuConnPortLO: &_VloginCheck_Dn.ulCHconnPortI, // chan _TudpConnPort
	}

	// _FdebugPrintTest()

	_Cexit = make(chan string, 3)
	_Clog = make(chan string, 100)

}

func main() {

	_Finit__2301()

	go _Frun(&_VserviceTcp_Dn, 200101)

	// _TudpNodeSt _TudpGroupSt
	go _Frun(&_VudpGroup_Dn, 600101) // IRun _FudpGroup__600101__main_init__default
	// _FudpGroup__600201__CHin_selecT_send__default
	// _FdataPack__101__udpConnPort

	// _FudpDecode__700201x__receive__default
	go _Frun(&_VudpDecode_Dn, 700101) // IRun _FudpDecode__700101x__init__default

	// _FloginCheck__900201x__standardCheck
	// _FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn
	go _Frun(&_VloginCheck_Dn, 900101) // _FloginCheck__900101x__init__default

	go _Frun(&_VloginGenerator_Dn, 800101) // IRun _FudpDecode__800101x__init__tryUdpLogin__default

	// _FdataMachin__1000501x__time_gap_dataChanLock
	go _Frun(&_VdataMachine_Dn, 1000101) // IRun _FdataMachin__1000101__main_init__default

	_Fex(" the reason exit : "+<-_Cexit, nil)
} // main
