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

/*
func (___Vdmcc *_TdataMachinEconnectClient) _FdataMachin__clear__TdataMachinEconnectClient() {
	(*___Vdmcc) = _TdataMachinEconnectClient{}
	___Vdmcc.dccLastReceTime = _FtimeInt()
	___Vdmcc.dccConnPortStrMap = make(map[string]byte)
}
*/
