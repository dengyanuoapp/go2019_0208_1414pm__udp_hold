package main

//import "fmt"

func _FnotNullRunUdp02(___Vrun func(*_TUreqNewSession), ___Vpara *_TUreqNewSession) {
	//_FpfN( "---" )
	if ___Vrun != nil {
		___Vrun(___Vpara)
	} else {
		_Fsleep_10s()
	}
} // _FnotNullRunUdp02

func _FnotNullRunUdp01(___Vrun func(*_TserviceUDP), ___Vpara *_TserviceUDP) {
	//_FpfN( "---" )
	if ___Vrun != nil {
		___Vrun(___Vpara)
	} else {
		_Fsleep_10s()
	}
} // _FnotNullRunUdp01

func _FnotNullRun011_tcp_service_chan(___Vrun func(*_TserviceTCP), ___Vpara *_TserviceTCP) {
	//_FpfN( "---" )
	if ___Vrun != nil {
		___Vrun(___Vpara)
	} else {
		_Fsleep_10s()
	}
} // _FnotNullRun011_tcp_service_chan

func _FnotNullRunTcp02_accept(___Vrun func(*_TacceptTCP), ___Vpara *_TacceptTCP) {
	//_FpfN( "---" )
	if ___Vrun != nil {
		___Vrun(___Vpara)
	} else {
		_Fsleep_10s()
	}
} // _FnotNullRunTcp02_accept
