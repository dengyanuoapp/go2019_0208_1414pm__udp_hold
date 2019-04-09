package main

type _Tkey256X struct{}
type _Tkey256 struct {
	// the input password , all data try AES decode, if faild , then drop it.
	// if nil , then ignore(no filter , directly receive)
	Bkey    []byte
	B32     [32]byte
	disable bool
}

func (___Vk *_Tkey256) String() string {
	return _Spf("k1:%s k2:%x dis:%t",
		String5(&___Vk.Bkey),
		___Vk.B32[:5],
		___Vk.disable,
	)
}

func (___Vkey *_Tkey256) initKey(___VinBlp *[]byte) {
	if nil == ___VinBlp {
		_FpfN(" 893818 02 using key : %x (new gen random) ", ___VinBlp)
		___Vkey.Bkey = _FgenRand_nByte__(32)
		copy(___Vkey.B32[:], ___Vkey.Bkey)
	} else {
		//_FpfN(" 893818 03 using key : %x ", ___VinBlp)
		___Vkey.Bkey = *___VinBlp
		copy(___Vkey.B32[:], ___Vkey.Bkey)
	}
	//_FpfN(" 893818 05 using key : BinLP %x, B %x, F %x ", ___VinBlp, ___Vkey.Bkey, ___Vkey.B32)
}
