package main

import (
	"flag"
)

var (
	_VtcpDebugLog__Dn   _TtcpNodE
	_VudpGroup_Dn       _TudpGroupSt
	_VudpDecode_Dn      _TuDecode
	_VloginGenerator_Dn _TloginGenerator
	_VdataMachine_Dn    _TdataMachine

	_VloginCheck_Dn _TloginCheck

	_Vself   _Tself
	_Vconfig _Tconfig
)

func _Finit__2301() {

	_Vself.ProjName = "Dn"

	_Fbase_1203__init_self_All()

	_Fbase_104z__try_to_read_json_config_top()

	_Fbase_107__rand_init()

	_FPargs()

	_VtcpDebugLog__Dn = _TtcpNodE{
		tnName:        "TcpService__DebugLog__Md",
		tnHostPortStr: "127.0.0.1:50003",
		tnAmount:      10,
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

}

func main() {

	_Finit__2301()

	// _TtcpNodeDataSend
	// _TtcpNodE
	// _FtcpNodeAccept__200401x5__dataReceiveMsg01_default
	// _FtcpNode__200301x_send__default
	go _Frun(&_VtcpDebugLog__Dn, 200101) // _FtcpNode__200101x__init_default
	go _Frun(&_VtcpDebugLog__Dn, 200801) // _FtcpNode__200801x_send__tester
	//go _FtcpNode__200801x_send__tester( &_VtcpDebugLog__Dn )

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

} // main
