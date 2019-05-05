package main

func _FchanNullCheckOk(___Vmsg string, ___VchanB *chan []byte) bool {
	if nil == ___VchanB {
		return false
	}
	__Vlen := len(*___VchanB)
	__Vcap := cap(*___VchanB)

	__FpfN(___Vmsg+" 281812 01 : len %d , cap %d ", __Vlen, __Vcap)

	if 0 == __Vcap {
		_FpfNex(___Vmsg+" 281812 03 : why the cap of the output chan ZERO ? %d : %d ", __Vlen, __Vcap)
		return false
	}

	if __Vlen == __Vcap {
		_CFpfN(___Vmsg+" 281812 05 : the output chan FULL (%d) , skip ", __Vlen)
		return false
	}

	return true
}
