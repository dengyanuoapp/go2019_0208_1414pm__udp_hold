package main

func (___Vaes *_Taes) _FcbcEncode() {
	defer ___Vaes.aesMux.Unlock()
	___Vaes.aesMux.Lock()

	___Vaes.aesOk = false
	___Vaes.aesEncode = true

	if 32 != len(___Vaes.aesKey)() {
		___Vaes.aesErr = _Spf(" 138141 01 : aes key length error : %d ", len(___Vaes.aesKey))
		return
	}

	if ___Vaes.aesRandMode {
		if false == ___Vaes.
			_FrandPatAddToOutput() {
			return
		}
	}

	___Vaes.aesOk = true
}
