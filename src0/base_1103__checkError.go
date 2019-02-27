package main

import (
	"fmt"
)

func _FerrExit(___VerrMsg string, ___Verr error) {
	if ___Verr != nil {
		_Fex("Error: "+___VerrMsg, ___Verr)
	}
} // _FerrExit

func _FnullExit(___VerrMsg string, ___Vck interface{}) {
	if ___Vck == nil {
		_Fex("Error: "+___VerrMsg, nil)
	}
} // _FnullExit

func _FnotNullExit(___VerrMsg string, ___Vck interface{}) {
	if ___Vck != nil {
		_Fex("Error: "+___VerrMsg, ___Vck)
	}
} // _FnotNullExit

func _FnotEqExit(___VerrMsg string, ___Va interface{}, ___Vb interface{}) {
	if ___Va != ___Vb {
		__Vstr := fmt.Sprintf(" Error : %s : not equal : <%v> <%v>\n", ___VerrMsg, ___Va, ___Vb)
		_Fex(__Vstr, nil)
	}
} // _FnotEqExit

func _FeqExit(___VerrMsg string, ___Va interface{}, ___Vb interface{}) {
	if ___Va == ___Vb {
		_Fex("Error: "+___VerrMsg, nil)
	}
} // _FeqExit
