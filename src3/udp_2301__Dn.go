package main

import (
	"flag"
)

var (
	_VserviceTcpMd     _TserviceTCP
	_VudpGroup_Dn2Fn   _TudpGroupSt
	_VtmUdpConn__Dn2Fn _TgapTimer
	_Cexit             chan string
	_Clog              chan string
	_Vself             _Tself
	_Vconfig           _Tconfig
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

		Cexit:   &_Cexit,
		Clog:    &_Clog,
		cAmount: 10,
	}

	_VudpGroup_Dn2Fn = _TudpGroupSt{
		ugName:        "udpGroup_Dn2Fn",
		ugAmount:      10,
		ugHostPortStr: make([]string, 1),
	}
	flag.StringVar(&_VudpGroup_Dn2Fn.ugHostPortStr[0], "cn", ":0", _VudpGroup_Dn2Fn.ugName)

	flag.Parse()

	_VtmUdpConn__Dn2Fn = _TgapTimer{
		gap01: _T10s,
		uTmSrvInfo01: &_TsrvInfo{
			Name: "srvDn2Fn",
			Guri: "https://raw.githubusercontent.com/jasas78/jsonOnly/master/json/FnWaitDn.json.rand",
		},
	}

	// _FdebugPrintTest()

	_Cexit = make(chan string, 3)
	_Clog = make(chan string, 100)

}

func main() {

	_Finit__2301()

	go _Frun(&_VserviceTcpMd, 101)

	// _TudpNodeSt _TudpGroupSt
	go _Frun(&_VudpGroup_Dn2Fn, 150201) // IRun _FudpGroup__150201__main_init__default

	go _Frun(&_VtmUdpConn__Dn2Fn, 150102) // IRun _FudpTimer__150102x__init__tryUdpConn__default

	_Fex(" the reason exit : "+<-_Cexit, nil)
} // main
