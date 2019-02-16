package main

//import "fmt"

func _FerrExit(___VerrMsg string , ___Verr error) {
    if ___Verr  != nil {
        _Fex("Error: " + ___VerrMsg , ___Verr )
    }
} // _FerrExit

func _FnullExit(___VerrMsg string , ___Vck interface{}) {
    if ___Vck  == nil {
        _Fex("Error: " + ___VerrMsg , nil )
    }
} // _FnullExit

func _FnotNullExit(___VerrMsg string , ___Vck interface{}) {
    if ___Vck  != nil {
        _Fex("Error: " + ___VerrMsg , ___Vck )
    }
} // _FnotNullExit

func _FnotNullRunUdp01( ___Vrun func( *_TserviceUDP) , ___Vpara *_TserviceUDP ) {
    //_FpfN( "---" )
    if ( ___Vrun != nil ) {
        ___Vrun( ___Vpara )
    }
} // _FnotNullRunUdp01

func _FnotEqExit(___VerrMsg string , ___Va interface{} , ___Vb interface{} ) {
    if ___Va  != ___Vb {
        _Fex("Error: " + ___VerrMsg , nil )
    }
} // _FnotEqExit

func _FeqExit(___VerrMsg string , ___Va interface{} , ___Vb interface{} ) {
    if ___Va  == ___Vb {
        _Fex("Error: " + ___VerrMsg , nil )
    }
} // _FeqExit
