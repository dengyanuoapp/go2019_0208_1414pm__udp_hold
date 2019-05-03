package main

import (
	"flag"
)

var (
	_VtcpDebugLog__Cn      _TtcpNodE
	_VbyteNoteBuf__Cn      _TbyteNoteBuf
	_VudpGroup_Cn          _TudpGroupSt
	_VudpDecode_Cn         _TuDecode
	_VudpEncode_Cn         _TuEncode
	_VdataMachine_Cn       _TdataMachine
	_VloginCheck_Cn        _TloginCheck
	_CHexit                chan string             = make(chan string, 10)
	_CHpr                  *chan _TtcpNodeDataSend = &_VtcpDebugLog__Cn.tnCHsendToAllClientI
	_Vself                 _Tself
	_Vconfig               _Tconfig
	_VloginGenerator_Cn    _TloginGenerator
	_VtcpAccetpClients__Cn _TtcpNodE
	// _TudpNodeSt
	// _TgapFilterX
)

func _Finit__2401() {

	_Vself.RoleName = "Cn"

	_Fbase_1203__init_self_All()

	_Fbase_104z__try_to_read_json_config_top()

	_Fbase_107__rand_init()

	_FPargs()

	_VtcpAccetpClients__Cn = _TtcpNodE{
		tnName:        " tcp_acceptClient_Cn ",
		tnHostPortStr: "127.0.0.1:50080",
		tnAmount:      10,
	}

	_VtcpDebugLog__Cn = _TtcpNodE{
		tnName:          " tcp_debug_Cn ",
		tnHostPortStr:   "127.0.0.1:50005",
		tnAmount:        10,
		tnCHdebugInfoLO: &_VbyteNoteBuf__Cn.bnbCHinI,
	}
	_VbyteNoteBuf__Cn = _TbyteNoteBuf{
		bnbCHoutLO1: &_VdataMachine_Cn.dmCHdebugInfoI,
	}

	_VudpDecode_Cn = _TuDecode{
		uDeCHdecodeCkLO:   &_VloginCheck_Cn.ulCHdecodeCkI,    // _TloginCheck _TdecodeX
		uDeCHdecodeDataLO: &_VdataMachine_Cn.dmCHdecodeDataI, // dmCHdecodeDataI _TdecodeX
	}

	_VudpEncode_Cn = _TuEncode{
		enCHuDataSendLO: &_VudpGroup_Cn.ugCHSendI, // _TudpGroupSt _TudpNodeDataSendX
	}

	_VloginCheck_Cn = _TloginCheck{
		ulCHdataMachineIdLO: &_VdataMachine_Cn.dmCHdataMachineIdI, // *chan _TdataMachinEid
		ulCHencodeCkLO:      &_VudpEncode_Cn.enCHencodeCkI,        // *chan _Tencode
	}

	_VdataMachine_Cn = _TdataMachine{
		dmCBprReceKey:           _FdmCBprReceKey__Cn,
		dmCBprSendKey:           _FdmCBprSendKey__Cn,
		dmCHloginGenMachineIdLO: &_VloginGenerator_Cn.lgCHdataMachineIdI,
		dmCHencodeIdleLO:        &_VudpEncode_Cn.enCHencodeDataCommI, // *chan _Tencode
	}

	_VudpGroup_Cn = _TudpGroupSt{
		ugName:         "udpGroup_Cn",
		ugAmount:       10,
		ugHostPortStr:  []string{":0"},
		ugCHreceByteLO: &_VudpDecode_Cn.uDeCHreceUgByteI, // *chan _TudpNodeDataReceX
	}

	flag.StringVar(&_VudpGroup_Cn.ugHostPortStr[0], "cn", ":0", _VudpGroup_Cn.ugName)

	flag.Parse()

	_VloginGenerator_Cn = _TloginGenerator{
		lgSrvDownInfoLX: &_TsrvDownInfo{
			name:         "srvCn2Fn",
			updateUri:    "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/FnWaitCn.gob.rand", // gob  json
			updatePasswd: _Vpasswd_udp_FnWaitCn_download_config,
		},
		lgCHuConnPortLO: &_VloginCheck_Cn.ulCHconnPortI, // chan _TudpConnPort
	}

	// _FdebugPrintTest()

}

func _FdmCBprReceKey__Cn(___Vdm *_TdataMachine) {
	_VudpGroup_Cn._FprKey()
}
func _FdmCBprSendKey__Cn(___Vdm *_TdataMachine) {
}

func main() {

	_Fsleep(_T1s)

	_Finit__2401()

	// _TtcpNodeDataSend
	// _TtcpNodE
	// _FtcpNodeAccept__200401x5__dataReceiveMsg01_default
	// _FtcpNode__200301x_send__default
	go _Frun(&_VtcpDebugLog__Cn, 200101)      // _FtcpNode__200101x__init_default
	go _Frun(&_VtcpAccetpClients__Cn, 200101) // _FtcpNode__200101x__init_default
	//go _Frun(&_VtcpDebugLog__Cn, 200801) // _FtcpNode__200801x_send__tester01
	//go _Frun(&_VtcpDebugLog__Cn, 200802) // _FtcpNode__200802x_send__tester02

	// _FbyteNoteBuf__1300201x__chan_rece__default
	go _Frun(&_VbyteNoteBuf__Cn, 1300101) // _FbyteNoteBuf__1300101x__init

	// _FudpNode__500201y__receive__default()
	// _FudpNode__500101__main_init__default
	// _FudpNode__500101z__send__default
	// _TudpNodeSt _TudpGroupSt
	go _Frun(&_VudpGroup_Cn, 600101) // IRun _FudpGroup__600101__main_init__default
	// _FudpGroup__600201__CHin_selecT_send__default
	// _FdataPack__101__udpConnPort

	// _FudpDecode__700201x__receive__default
	go _Frun(&_VudpDecode_Cn, 700101) // IRun _FudpDecode__700101x__init__default

	// _FloginCheck__900201x__standardCheck
	// _FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn
	go _Frun(&_VloginCheck_Cn, 900101) // _FloginCheck__900101x__init__default

	// _TloginGenerator
	go _Frun(&_VloginGenerator_Cn, 800101) // IRun _FudpDecode__800101x__init__tryUdpLogin__default

	// _TdataMachinEconnectSt                                     // 1
	// _TdataMachinEdataClient                                    // 2
	// _FdataMachin__search_avaiable__TdataMachinEconnectClient   // 3
	// _FdataMachin__1000501y__swapLoginCkInfoForLock__swap       // 4
	// _FdataMachin__1000502y__dataSendIdle__packAndSendAll       // 5
	// _FdataMachin__1000503x__time_gap_dataResend                // 6
	// _FdataMachin__1000601x__encodeData_sendMux                 // 7
	// _FdataMachin__1000201x21__rece_decodeData                  // 8
	// _FdataMachin__1000501x__swapLoginCkInfoForLock__timeoutGen // 9
	// _FdataMachin__1000504x__checkTimeOutDieClient              // 9
	// _FdataMachin__1000201x__receive__default                   // a
	// _FdataMachin__1000201x11__rece_machineId_check_and_insert  // b
	go _Frun(&_VdataMachine_Cn, 1000101) // IRun _FdataMachin__1000101__main_init__default

	// _FuEncode__1100201x__chanRece__default
	go _Frun(&_VudpEncode_Cn, 1100101) // _FuEncode__1100101__main_init__default

	<-_CHexit
} // main
