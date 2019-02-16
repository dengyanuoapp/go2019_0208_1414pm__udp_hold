package main

//import "fmt"


func _FnotNullRunUdp01( ___Vrun func( *_TserviceUDP) , ___Vpara *_TserviceUDP ) {
    //_FpfN( "---" )
    if ( ___Vrun != nil ) {
        ___Vrun( ___Vpara )
    }
} // _FnotNullRunUdp01

func _FnotNullRunTcp01_service( ___Vrun func( *_TserviceTCP) , ___Vpara *_TserviceTCP ) {
    //_FpfN( "---" )
    if ( ___Vrun != nil ) {
        ___Vrun( ___Vpara )
    }
} // _FnotNullRunTcp01_service

func _FnotNullRunTcp02_accept( ___Vrun func( *_TacceptTCP) , ___Vpara *_TacceptTCP ) {
    //_FpfN( "---" )
    if ( ___Vrun != nil ) {
        ___Vrun( ___Vpara )
    }
} // _FnotNullRunTcp02_accept

