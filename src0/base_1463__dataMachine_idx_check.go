package main

// _TdataMachinEconnectClient _TdataMachinEconnectSt
func (___VdmcT *_TdataMachinEconnectSt) _FdataMachin__search_avaiable__TdataMachinEconnectClient() int {
	if ___VdmcT.dcsMfreeAmount <= 0 { // _TdataMachinEconnectSt
		_FpfN(" 838291 03 : max client amount reach %d/%d/%d",
			___VdmcT.dcsMusedAmount, ___VdmcT.dcsMfreeAmount, _VallowClientMax)
		return -1
	}
	for __Vi := ___VdmcT.dcsMlastInsIdx; __Vi < _VallowClientMax; __Vi++ {
		if 0 == ___VdmcT.dcsMm[__Vi].dccLastReceTime { // _TdataMachinEconnectClient
			___VdmcT.dcsMm[__Vi].
				Clear()
			return __Vi
		}
	}
	for __Vi := 0; __Vi < ___VdmcT.dcsMlastInsIdx; __Vi++ {
		if 0 == ___VdmcT.dcsMm[__Vi].dccLastReceTime { // _TdataMachinEconnectClient
			___VdmcT.dcsMm[__Vi].
				Clear()
			return __Vi
		}
	}

	_FpfN(" 838291 08 : max client amount reach %d/%d/%d",
		___VdmcT.dcsMusedAmount, ___VdmcT.dcsMfreeAmount, _VallowClientMax)
	_Fex(" 838291 09 : error : why not found ? exit ")

	return -1
}

// _TdataMachinEdataClient _TdataMachinEdataSt
func (___VdmdT *_TdataMachinEdataSt) _FdataMachin__search_avaiable__TdataMachinEdataClient() int {
	if ___VdmdT.ddsMfreeAmount <= 0 { // _TdataMachinEdataSt
		_FpfN(" 838292 03 : max client amount reach %d/%d/%d",
			___VdmdT.ddsMusedAmount, ___VdmdT.ddsMfreeAmount, _VallowClientMax)
		return -1
	}
	for __Vi := ___VdmdT.ddsMlastInsIdx; __Vi < _VallowClientMax; __Vi++ {
		if 0 == ___VdmdT.ddsMm[__Vi].ddcLastReceTime { // _TdataMachinEdataClient
			___VdmdT.ddsMm[__Vi].
				Clear()
			return __Vi
		}
	}
	for __Vi := 0; __Vi < ___VdmdT.ddsMlastInsIdx; __Vi++ {
		if 0 == ___VdmdT.ddsMm[__Vi].ddcLastReceTime { // _TdataMachinEdataClient
			___VdmdT.ddsMm[__Vi].
				Clear()
			return __Vi
		}
	}

	_FpfN(" 838292 08 : max client amount reach %d/%d/%d",
		___VdmdT.ddsMusedAmount, ___VdmdT.ddsMfreeAmount, _VallowClientMax)
	_Fex(" 838292 09 : error : why not found ? exit ")

	return -1
}
