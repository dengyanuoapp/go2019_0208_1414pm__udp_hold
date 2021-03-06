package main

import (
	"flag"
)

var (
	_VtcpDebugLog__Fn                    _TtcpNodE               // m 1
	_VbyteNoteBuf__Fn                    _TbyteNoteBuf           // m 2
	_VudpGroup_Fn                        _TudpGroupSt            // m 3
	_VudpDecode_Fn                       _TuDecode               // m 4
	_VudpEncode_Fn                       _TuEncode               // m 5
	_VdataMachine_Fn                     _TdataMachine           // m 6
	_VloginCheck_FnWaitDun               _TloginCheck            // m 7
	_CHexit                              chan string             = make(chan string, 10)
	_CHpr                                *chan _TtcpNodeDataSend = &_VtcpDebugLog__Fn.tnCHsendToAllClientI
	_Vself                               _Tself
	_Vconfig                             _Tconfig
	_VudpNode__FnWdn                     _TudpNodeSt                // m 8
	_VgapFilter__FnWdn                   _TgapFilter                // m 9  //_TgapFilterX
	_VudpNode__FnWcn                     _TudpNodeSt                // m 10
	_VgapFilter__FnWcn                   _TgapFilter                // m 11  //_TgapFilterX
	_VrecePackThenEncodeAsLoad__FnWaitCn _TrecePackThenEncodeAsLoad // m 12
)

func _Finit_2201() {

	_Vself.RoleName = "Fn"

	_Fbase_1203__init_self_All()

	_Fbase_104z__try_to_read_json_config_top()
	_Fbase_107__rand_init()

	_FPargs()

	_VtcpDebugLog__Fn = _TtcpNodE{
		tnName:          " tcp_debug_Fn ",
		tnHostPortStr:   "127.0.0.1:50002",
		tnAmount:        10,
		tnCHdebugInfoLO: &_VbyteNoteBuf__Fn.bnbCHinI,
	}
	_VbyteNoteBuf__Fn = _TbyteNoteBuf{
		bnbCHoutLO1: &_VdataMachine_Fn.dmCHdebugInfoI,
	}

	_VudpNode__FnWdn = _TudpNodeSt{
		unName:         "_VudpNode__FnWdn",
		unLoopGap:      _T10s,
		unRKeyLP:       &_Vpasswd_udp_Fn_waitForCliens01_Dn,
		unCHreceByteLO: &_VgapFilter__FnWdn.gfCHbyteI,
	}

	_VgapFilter__FnWdn = _TgapFilter{ //_TgapFilterX
		gfCHbyteLO: &_VudpDecode_Fn.uDeCHreceUnByteI,
	}

	_VudpNode__FnWcn = _TudpNodeSt{
		unName:         "_VudpNode__FnWcn",
		unLoopGap:      _T10s,
		unRKeyLP:       &_Vpasswd_udp_Fn_waitForCliens02_Cn,
		unCHreceByteLO: &_VgapFilter__FnWcn.gfCHbyteI,
	}

	_VgapFilter__FnWcn = _TgapFilter{ //_TgapFilterX
		gfCHbyteLO: &_VrecePackThenEncodeAsLoad__FnWaitCn.pelCHudpNodeDataReceBI,
	}

	_VrecePackThenEncodeAsLoad__FnWaitCn = _TrecePackThenEncodeAsLoad{
		pelCHc2sEncodeBLO: &_VdataMachine_Fn.dmCHencodeDataSpecFnWaitCnBI, // *chan []byte // byte of _TencodeX
	}

	_VudpDecode_Fn = _TuDecode{
		uDeCHdecodeCkLO:   &_VloginCheck_FnWaitDun.ulCHdecodeCkI, // _TloginCheck _Tdecode
		uDeCHdecodeDataLO: &_VdataMachine_Fn.dmCHdecodeDataI,     // dmCHdecodeDataI _TdecodeX

	}

	_VudpEncode_Fn = _TuEncode{
		enCHuDataSendLO: &_VudpGroup_Fn.ugCHSendI, // _TudpGroupSt _TudpNodeDataSendX
	}

	_VloginCheck_FnWaitDun = _TloginCheck{
		ulCHdataMachineIdLO: &_VdataMachine_Fn.dmCHdataMachineIdI,
		ulCHencodeCkLO:      &_VudpEncode_Fn.enCHencodeCkI, // *chan _Tencode

	}

	_VdataMachine_Fn = _TdataMachine{
		dmCBprReceKey: _FdmCBprReceKey__Fn,
		//dmCHloginGenMachineIdLO: &_VloginGenerator_Fn.lgCHdataMachineIdI,
		dmCHencodeIdleLO: &_VudpEncode_Fn.enCHencodeDataCommI, // *chan _Tencode
	}

	_VudpGroup_Fn = _TudpGroupSt{
		ugName:         "udpGroup_Fn",
		ugAmount:       20,
		ugHostPortStr:  []string{":0"},
		ugCHreceByteLO: &_VudpDecode_Fn.uDeCHreceUgByteI, // *chan _TudpNodeDataRece
	}
	flag.StringVar(&_VudpGroup_Fn.ugHostPortStr[0], "cn", ":0", _VudpGroup_Fn.ugName)

	flag.StringVar(&_VudpNode__FnWdn.unHostPortStr, "FnWdn", ":32001", _VudpNode__FnWdn.unName)
	flag.StringVar(&_VudpNode__FnWcn.unHostPortStr, "FnWcn", ":48881", _VudpNode__FnWcn.unName)

	flag.Parse()

	// _FdebugPrintTest()
}

func _FdmCBprReceKey__Fn(___Vdm *_TdataMachine) {
	_VudpGroup_Fn._FprKey()
}

func main() {

	_Fsleep(_T1s)

	_Finit_2201()

	// ------------------- tcp for debug monitor log --- begin
	// IRun _FtcpNode__200101x__init_default()
	go _Frun(&_VtcpDebugLog__Fn, 200101)
	// ------------------- tcp for debug monitor log --- end

	// _FbyteNoteBuf__1300201x__chan_rece__default
	go _Frun(&_VbyteNoteBuf__Fn, 1300101) // _FbyteNoteBuf__1300101x__init

	// _FudpNode__540211z__receiveCallBack_withTimeGap
	// _FdataPack__dePack_decode__from_udpNodeDataRece
	// _FudpNode__500201y__receive__default
	go _Frun(&_VudpNode__FnWdn, 500101) // IRun _FudpNode__500101__main_init__default
	go _Frun(&_VudpNode__FnWcn, 500101) // IRun _FudpNode__500101__main_init__default

	// _FudpDecode__700201x__receive__default
	go _Frun(&_VudpDecode_Fn, 700101) // IRun _FudpDecode__700101x__init__default

	// _FloginCheck__900201x__standardCheck
	// _FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn
	go _Frun(&_VloginCheck_FnWaitDun, 900101) // _FloginCheck__900101x__init__default

	// _TudpNodeSt _TudpGroupSt
	go _Frun(&_VudpGroup_Fn, 600101) // IRun _FudpGroup__600101__main_init__default
	// _FudpGroup__600201__CHin_selecT_send__default
	// _FdataPack__101__udpConnPort

	// _FdataMachin__1000501x__swapLoginCkInfoForLock__timeoutGen
	// _FdataMachin__1000201x__receive__default
	// _FdataMachin__1000201x11__rece_machineId_check_and_insertConnClient
	go _Frun(&_VdataMachine_Fn, 1000101) // IRun _FdataMachin__1000101__main_init__default

	// _FuEncode__1100201x__chanRece__default
	go _Frun(&_VudpEncode_Fn, 1100101) // _FuEncode__1100101__main_init__default

	// _FgapFilter__1200301x1__uDataSwap
	// _FgapFilter__1200301x2__uData_rece
	// _FgapFilter__1200301x__Chan_rece
	go _Frun(&_VgapFilter__FnWdn, 1200101) // _FgapFilter__1200101x__init_default
	go _Frun(&_VgapFilter__FnWcn, 1200101) // _FgapFilter__1200101x__init_default

	// _FrecePackThenEncodeAsLoad__1400201x__chan_rece__default
	go _Frun(&_VrecePackThenEncodeAsLoad__FnWaitCn, 1400101) // _FrecePackThenEncodeAsLoad__1400101x__init

	<-_CHexit
} // main
