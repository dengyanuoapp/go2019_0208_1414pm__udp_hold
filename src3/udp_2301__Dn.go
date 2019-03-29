package main

import (
	"flag"
)

var (
	_VserviceTcpMd   _TserviceTCP
	_VudpGroup_Dn    _TudpGroupSt
	_VloginGenerator _TloginGenerator
	_Cexit           chan string
	_Clog            chan string
	_Vself           _Tself
	_Vconfig         _Tconfig
)

func _Finit__2301() {

	_Vself.ProjName = "Dn"

	_Fbase_1203__init_self_All()

	_Fbase_104z__try_to_read_json_config_top()

	_Fbase_107__rand_init()

	_FPargs()

	_VserviceTcpMd = _TserviceTCP{
		name:        "TcpService__DebugLog__Md",
		hostPortStr: "127.0.0.1:56782",
		Cexit:       &_Cexit,
		Clog:        &_Clog,
		cAmount:     10,
	}

	_VudpGroup_Dn = _TudpGroupSt{
		ugName:        "udpGroup_Dn",
		ugAmount:      10,
		ugHostPortStr: []string{":0"},
	}
	flag.StringVar(&_VudpGroup_Dn.ugHostPortStr[0], "cn", ":0", _VudpGroup_Dn.ugName)

	flag.Parse()

	_VloginGenerator = _TloginGenerator{
		ulSrvDownInfoLX: &_TsrvDownInfo{
			name:         "srvDn2Fn",
			updateUri:    "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/FnWaitDn.gob.rand", // gob  json
			updatePasswd: _Vpasswd_udp_FnWaitDn_download_config,
		},
		//ulCHugConnPortLO: &_VudpGroup_Dn.ugCHuConnPortI,
		ulCHunSendLO: &_VudpGroup_Dn.ugCHSendI,
	}

	// _FdebugPrintTest()

	_Cexit = make(chan string, 3)
	_Clog = make(chan string, 100)

}

func main() {

	_Finit__2301()

	go _Frun(&_VserviceTcpMd, 200101)

	// _TudpNodeSt _TudpGroupSt
	go _Frun(&_VudpGroup_Dn, 600101) // IRun _FudpGroup__600101__main_init__default
	// _FudpGroup__600201__CHin_select_send__default
	// _FdataPack__101__udpConnPort

	go _Frun(&_VloginGenerator, 850102) // IRun _FudpTimer__850102x__init__tryUdpLogin__default

	_Fex(" the reason exit : "+<-_Cexit, nil)
} // main
