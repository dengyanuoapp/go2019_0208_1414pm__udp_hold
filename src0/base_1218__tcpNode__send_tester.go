package main

func (___VtcpNode4 *_TtcpNodE) _FtcpNode__200801x_send__tester() {
	for {
		_Fsleep(_T10s)
		_FpfN(" 838188 01 ")
		//		__Vts := _TtcpNodeDataSend{}
		__Vts := _TtcpNodeDataSend{
			//tnsId128: _FgenRand_nByte__(16),
			//tnsK256:  _FgenRand_nByte__(32),
			tnsLen: 4,
			tnsBuf: []byte{0x62, 0x63, 0x64, 0x0a},
		}

		_FpfN(" 838188 02 : %d : testing : push into send-chan:{%s}", _FtimeInt(), __Vts.String())

		___VtcpNode4.tnCHsendToAllClientI <- __Vts
	}
}
