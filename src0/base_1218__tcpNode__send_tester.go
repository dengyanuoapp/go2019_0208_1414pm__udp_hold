package main

func (___VtcpNode4 *_TtcpNodE) _FtcpNode__200801x_send__tester01() {
	for {
		_Fsleep(_T10s)
		//_FpfN(" 838188 01 ")
		//		__Vts := _TtcpNodeDataSend{}
		__Vts := _TtcpNodeDataSend{
			//tnsId128: _FgenRand_nByte__(16),
			//tnsK256:  _FgenRand_nByte__(32),
			tnsBuf: []byte(_Spf(" 838188 02 : testing %d\n", _FtimeInt())),
		}

		__Vts.tnsLen = len(__Vts.tnsBuf)

		_FpfN(" 838188 03 : %d : testing : push into send-chan:{%s}", _FtimeInt(), __Vts.String())

		___VtcpNode4.tnCHsendToAllClientI <- __Vts
	}
}

func (___VtcpNode4 *_TtcpNodE) _FtcpNode__200802x_send__tester02() {
	for {
		_Fsleep(_T10s)
		_FpfN(" 838185 01 ")
		_CpfN(" 838185 03 : %d : testing : push into send-chan: _CpfN ", _FtimeInt())
	}
}
