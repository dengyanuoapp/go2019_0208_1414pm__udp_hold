package main

import (
	"flag"
)

var (
	_VtcpDebugLog__Fn      _TtcpNodE
	_VudpGroup_Fn          _TudpGroupSt
	_VudpDecode_Fn         _TuDecode
	_VudpEncode_Fn         _TuEncode
	_VdataMachine_Fn       _TdataMachine
	_VloginCheck_FnWaitDun _TloginCheck
	_CHexit                chan string             = make(chan string, 10)
	_CHpr                  *chan _TtcpNodeDataSend = &_VtcpDebugLog__Fn.tnCHsendToAllClientI
	_Vself                 _Tself
	_Vconfig               _Tconfig
	_VudpNode_FunWaitDun   _TudpNodeSt
	//_VgapFilter__Dn        _TgapFilter //_TgapFilterX
)

func _Finit_2201() {

	_Vself.ProjName = "Fn"

	_Fbase_1203__init_self_All()

	_Fbase_104z__try_to_read_json_config_top()
	_Fbase_107__rand_init()

	_FPargs()

	_VtcpDebugLog__Fn = _TtcpNodE{
		tnName:        " tcp_debug_Fn ",
		tnHostPortStr: "127.0.0.1:50002",
		tnAmount:      10,
	}

	_VudpNode_FunWaitDun = _TudpNodeSt{
		unName:     "_VudpNode_FunWaitDun",
		unRKeyLP:   &_Vpasswd_udp_Fn_waitForCliens01,
		unLoopGap:  _T10s,
		unCHreceLO: &_VudpDecode_Fn.uTmCHunDataReceI,
		//unCHreceLO: &_VgapFilter__Dn.gfCHbyteI,
	}

	//	_VgapFilter__Dn = _TgapFilter{ //_TgapFilterX
	//		gfCHbyteLO: &_VudpDecode_Fn.uTmCHunDataReceI,
	//	}

	_VudpDecode_Fn = _TuDecode{
		uTmCHdecodeCmdLO:  &_VloginCheck_FnWaitDun.ulDecodeI, // _TloginCheck _Tdecode
		uTmCHdecodeDataLO: &_VdataMachine_Fn.dmCHdecodeDataI, // dmCHdecodeDataI _TdecodeX

	}

	_VdataMachine_Fn = _TdataMachine{
		dmCBprReceKey: _FdmCBprReceKey__Fn,
	}

	_VudpEncode_Fn = _TuEncode{
		enCHuDataSendLO: &_VudpGroup_Fn.ugCHSendI, // _TudpGroupSt _TudpNodeDataSendX
	}

	_VloginCheck_FnWaitDun = _TloginCheck{
		ulCHdataMachineIdLO: &_VdataMachine_Fn.dmCHdataMachineIdI,
		ulCHencodeCmdLO:     &_VudpEncode_Fn.enCHencodeCmdI, // *chan _Tencode

	}

	_VudpGroup_Fn = _TudpGroupSt{
		ugName:        "udpGroup_Fn",
		ugAmount:      20,
		ugHostPortStr: []string{":0"},
		ugCHreceLO:    &_VudpDecode_Fn.uTmCHunDataReceI, // *chan _TudpNodeDataRece
	}
	flag.StringVar(&_VudpGroup_Fn.ugHostPortStr[0], "cn", ":0", _VudpGroup_Fn.ugName)

	flag.StringVar(&_VudpNode_FunWaitDun.unHostPortStr, "FunWdun", ":32001", _VudpNode_FunWaitDun.unName)

	flag.Parse()

	// _FdebugPrintTest()
}

func _FdmCBprReceKey__Fn(___Vdm *_TdataMachine) {
	_VudpGroup_Fn._FprKey()
}

func main() {
	_Finit_2201()

	// ------------------- tcp for debug monitor log --- begin
	// IRun _FtcpNode__200101x__init_default()
	go _Frun(&_VtcpDebugLog__Fn, 200101)
	// ------------------- tcp for debug monitor log --- end

	// _FudpNode__540211z__receiveCallBack_withTimeGap
	// _FdataPack__decode_from_udpNodeDataRece
	go _Frun(&_VudpNode_FunWaitDun, 500101) // IRun _FudpNode__500101__main_init__default

	// _FudpDecode__700201x__receive__default
	go _Frun(&_VudpDecode_Fn, 700101) // IRun _FudpDecode__700101x__init__default

	// _FloginCheck__900201x__standardCheck
	// _FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn
	go _Frun(&_VloginCheck_FnWaitDun, 900101) // _FloginCheck__900101x__init__default

	// _TudpNodeSt _TudpGroupSt
	go _Frun(&_VudpGroup_Fn, 600101) // IRun _FudpGroup__600101__main_init__default
	// _FudpGroup__600201__CHin_selecT_send__default
	// _FdataPack__101__udpConnPort

	// _FdataMachin__1000501x__time_gap_dataChanLock
	go _Frun(&_VdataMachine_Fn, 1000101) // IRun _FdataMachin__1000101__main_init__default

	// _FuEncode__1100201x__packSend__default
	go _Frun(&_VudpEncode_Fn, 1100101) // _FuEncode__1100101__main_init__default

	// _FgapFilter__1200301x__Chan_rece
	//go _Frun(&_VgapFilter__Dn, 1200101) // _FgapFilter__1200101x__init_default

	<-_CHexit
} // main
