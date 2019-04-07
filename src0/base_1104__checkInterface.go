package main

func _FinterfaceObjIsPointerOrExit(___Vmsg string, ___Vif interface{}) {
	if nil == ___Vif {
		_FpfNex(___Vmsg + ": 381918 01 : why nil ? ")
	}
	__Vs := []byte(_Spf("%T", ___Vif))
	if '*' != __Vs[0] {
		_FpfNex(___Vmsg+": 381918 03 : why not Pointer ? %T , %v , %s", ___Vif, ___Vif, __Vs)
	}
}
