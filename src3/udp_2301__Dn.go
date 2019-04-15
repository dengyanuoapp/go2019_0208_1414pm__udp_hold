package main

import (
	"flag"
)

var (
	_VtcpDebugLog__Dn   _TtcpNodE
	_VudpGroup_Dn       _TudpGroupSt
	_VudpDecode_Dn      _TuDecode
	_VudpEncode_Dn      _TuEncode
	_VdataMachine_Dn    _TdataMachine
	_VloginCheck_Dn     _TloginCheck
	_CHexit             chan string             = make(chan string, 10)
	_CHpr               *chan _TtcpNodeDataSend = &_VtcpDebugLog__Dn.tnCHsendToAllClientI
	_Vself              _Tself
	_Vconfig            _Tconfig
	_VloginGenerator_Dn _TloginGenerator
	// _TudpNodeSt
	// _TgapFilterX
)

func _Finit__2301() {

	_Vself.ProjName = "Dn"

	_Fbase_1203__init_self_All()

	_Fbase_104z__try_to_read_json_config_top()

	_Fbase_107__rand_init()

	_FPargs()

	_VtcpDebugLog__Dn = _TtcpNodE{
		tnName:        " tcp_debug_Dn ",
		tnHostPortStr: "127.0.0.1:50003",
		tnAmount:      10,
	}

	_VudpDecode_Dn = _TuDecode{
		uTmCHdecodeCmdLO:  &_VloginCheck_Dn.ulDecodeI,        // _TloginCheck _TdecodeX
		uTmCHdecodeDataLO: &_VdataMachine_Dn.dmCHdecodeDataI, // dmCHdecodeDataI _TdecodeX
	}

	_VudpEncode_Dn = _TuEncode{
		enCHuDataSendLO: &_VudpGroup_Dn.ugCHSendI, // _TudpGroupSt _TudpNodeDataSendX
	}

	_VloginCheck_Dn = _TloginCheck{
		ulCHdataMachineIdLO: &_VdataMachine_Dn.dmCHdataMachineIdI, // *chan _TdataMachinEid
		ulCHencodeCmdLO:     &_VudpEncode_Dn.enCHencodeCmdI,       // *chan _Tencode
	}

	_VdataMachine_Dn = _TdataMachine{
		dmCBprReceKey:           _FdmCBprReceKey__Dn,
		dmCBprSendKey:           _FdmCBprSendKey__Dn,
		dmCHloginGenMachineIdLO: &_VloginGenerator_Dn.lgCHdataMachineIdI,
		dmCHencodeLO:            &_VudpEncode_Dn.enCHencodeDataI, // *chan _Tencode
	}

	_VudpGroup_Dn = _TudpGroupSt{
		ugName:         "udpGroup_Dn",
		ugAmount:       10,
		ugHostPortStr:  []string{":0"},
		ugCHreceByteLO: &_VudpDecode_Dn.uTmCHreceByteI, // *chan _TudpNodeDataReceX
	}

	flag.StringVar(&_VudpGroup_Dn.ugHostPortStr[0], "cn", ":0", _VudpGroup_Dn.ugName)

	flag.Parse()

	_VloginGenerator_Dn = _TloginGenerator{
		lgSrvDownInfoLX: &_TsrvDownInfo{
			name:         "srvDn2Fn",
			updateUri:    "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/FnWaitDn.gob.rand", // gob  json
			updatePasswd: _Vpasswd_udp_FnWaitDn_download_config,
		},
		lgCHuConnPortLO: &_VloginCheck_Dn.ulCHconnPortI, // chan _TudpConnPort
	}

	// _FdebugPrintTest()

}

func _FdmCBprReceKey__Dn(___Vdm *_TdataMachine) {
	_VudpGroup_Dn._FprKey()
}
func _FdmCBprSendKey__Dn(___Vdm *_TdataMachine) {
}

func main() {

	_Fsleep(_T1s)

	_Finit__2301()

	// _TtcpNodeDataSend
	// _TtcpNodE
	// _FtcpNodeAccept__200401x5__dataReceiveMsg01_default
	// _FtcpNode__200301x_send__default
	go _Frun(&_VtcpDebugLog__Dn, 200101) // _FtcpNode__200101x__init_default
	//go _Frun(&_VtcpDebugLog__Dn, 200801) // _FtcpNode__200801x_send__tester01
	//go _Frun(&_VtcpDebugLog__Dn, 200802) // _FtcpNode__200802x_send__tester02

	// _FudpNode__500201y__receive__default()
	// _FudpNode__500101__main_init__default
	// _FudpNode__500101z__send__default
	// _TudpNodeSt _TudpGroupSt
	go _Frun(&_VudpGroup_Dn, 600101) // IRun _FudpGroup__600101__main_init__default
	// _FudpGroup__600201__CHin_selecT_send__default
	// _FdataPack__101__udpConnPort

	// _FudpDecode__700201x__receive__default
	go _Frun(&_VudpDecode_Dn, 700101) // IRun _FudpDecode__700101x__init__default

	// _FloginCheck__900201x__standardCheck
	// _FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn
	go _Frun(&_VloginCheck_Dn, 900101) // _FloginCheck__900101x__init__default

	// _TloginGenerator
	go _Frun(&_VloginGenerator_Dn, 800101) // IRun _FudpDecode__800101x__init__tryUdpLogin__default

	// _FdataMachin__1000501x__time_gap_dataChanLock
	go _Frun(&_VdataMachine_Dn, 1000101) // IRun _FdataMachin__1000101__main_init__default

	// _FuEncode__1100201x__chanRece__default
	go _Frun(&_VudpEncode_Dn, 1100101) // _FuEncode__1100101__main_init__default

	<-_CHexit
} // main
