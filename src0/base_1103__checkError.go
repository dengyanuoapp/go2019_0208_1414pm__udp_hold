package main

import (
//"fmt"
)

func _FerrExit(___VerrMsg string, ___Verr error) {
	if ___Verr != nil {
		_Fex1("Error: "+___VerrMsg, ___Verr)
	}
} // _FerrExit

func _FnullExit(___VerrMsg string, ___Vck interface{}) {
	if ___Vck == nil {
		_Fex("Error: " + ___VerrMsg)
	}
} // _FnullExit

func _FnotNullExit(___VerrMsg string, ___Vck interface{}) {
	if ___Vck != nil {
		_Fex1("Error: "+___VerrMsg, ___Vck)
	}
} // _FnotNullExit

func _FnotEqExit(___VerrMsg string, ___Va interface{}, ___Vb interface{}) {
	if ___Va != ___Vb {
		__Vstr := _Spf(" Error : %s : not equal : <%v> <%v>\n", ___VerrMsg, ___Va, ___Vb)
		_Fex(__Vstr)
	}
} // _FnotEqExit

func _FeqExit(___VerrMsg string, ___Va interface{}, ___Vb interface{}) {
	if ___Va == ___Vb {
		_Fex("Error: " + ___VerrMsg)
	}
} // _FeqExit

func _FtrueExit(___VerrMsg string, ___Vbool bool) {
	if true == ___Vbool {
		_Fex("Error: " + ___VerrMsg)
	}
} // _FtrueExit

func _FfalseExit(___VerrMsg string, ___Vbool bool) {
	if false == ___Vbool {
		_Fex("Error: " + ___VerrMsg)
	}
} // _FfalseExit
