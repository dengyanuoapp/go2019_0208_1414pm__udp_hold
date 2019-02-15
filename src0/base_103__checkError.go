package main

//import "fmt"

func _FerrExit(___Verr error) {
    if ___Verr  != nil {
        _Fex("Error: " , ___Verr )
    }
} // _FerrExit

func _FnullExit(___VerrMsg string , ___Vck interface{}) {
    if ___Vck  == nil {
        _Fex("Error: " + ___VerrMsg )
    }
} // _FnullExit

func _FnotnullExit(___VerrMsg string , ___Vck interface{}) {
    if ___Vck  != nil {
        _Fex("Error: " + ___VerrMsg , ___Vck )
    }
} // _FnotnullExit

