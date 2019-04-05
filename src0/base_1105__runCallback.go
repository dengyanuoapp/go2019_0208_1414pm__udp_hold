package main

//import "fmt"

type Trun interface {
	IRun(int)
} // Trun

func _Frun(___Vrun Trun, ___Vidx int) {
	if nil != ___Vrun {
		//(*___Vrun).IRun()
		//_FpfN(" 381931 01 : IR run ")
		___Vrun.IRun(___Vidx)
	}
} // _Frun

func _FnotNullRun011_tcp_service_chan(___Vrun func(*_TtcpNodE), ___Vpara *_TtcpNodE) {
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
