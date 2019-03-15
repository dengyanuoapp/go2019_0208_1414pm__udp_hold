package main

type _Tkey256 struct {
	// the input password , all data try AES decode, if faild , then drop it.
	// if nil , then ignore(no filter , directly receive)
	LP      *[]byte // if nil , use random [32]byte , or , copy it.
	Bkey    []byte
	B32     [32]byte
	disable bool
}

func (___Vkey *_Tkey256) initKey() {
	if nil == ___Vkey.LP {
		_FpfN(" 893818 02 : %x ", ___Vkey.LP)
		___Vkey.Bkey = _FgenRand_nByte__(32)
		copy(___Vkey.B32[:], ___Vkey.Bkey)
	} else {
		_FpfN(" 893818 03 : %x ", ___Vkey.LP)
		___Vkey.Bkey = *___Vkey.LP
		copy(___Vkey.B32[:], ___Vkey.Bkey)
	}
	_FpfN(" 893818 05 : LP %x, B %x, F %x ", ___Vkey.LP, ___Vkey.Bkey, ___Vkey.B32)
}
