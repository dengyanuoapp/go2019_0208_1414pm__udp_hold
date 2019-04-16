package main

import (
	"bytes"
)

func (___Vdm *_TdataMachine) _FdataMachin__1000201x21__rece_encodeData(___Vdec *_Tdecode) {
	/*
		__Vid := _TdataMachinEid{
			diConnPort: _TudpConnPort{
				___Vdec.remoteAddr,     // net.UDPAddr
				___Vdec.remotePortKey}, // []byte
			diIdx128: ___Vdec.Ddata.MeIdx128, // []byte
			diSeq128: ___Vdec.Ddata.MeSeq128, // []byte
			diToken:  ___Vdec.Ddata.TokenL,   // []byte
		}
	*/

	if nil == ___Vdec {
		_Fex(" 839195 01 : why NULL ?")
	}

	if false == ___Vdec.Ok { // _TdecodeX
		_FpfN(" 839195 02 : _TdataMachine : received not OK:{%s}", ___Vdec.String())
		return
	}

	switch ___Vdec.Type {
	case Cmd__data_01_idle:
		// do continue
	default:
		_FpfN(" 839195 02 : _TdataMachine : unknown type :{%s}", ___Vdec.String())
		return
	}

	defer ___Vdm.dmMconn.mux.Unlock() // _TdataMachinEconnectSt
	___Vdm.dmMconn.mux.Lock()         // _TdataMachinEconnectSt

	__Vk := _FgenB16(&___Vdec.Ddata.MEidx128)
	//___Vdm.dmMconn.mux.Lock()
	__Vold, __Vok := ___Vdm.dmMconn.M[__Vk] // map[[16]byte]_TdataMachinEconnectClient

	if false == __Vok { // if not exist ,
		_CpfN(" 839195 03 : _TdataMachine : no such id exist . {%s}", ___Vdec.String())
		return
	}

	if                                                                     // if exist , but  token / sequence not match , replace the old one
	false == bytes.Equal(__Vold.dmmID.diIdx128, ___Vdec.Ddata.MEidx128) || // _TdataTran
		false == bytes.Equal(__Vold.dmmID.diSeq128, ___Vdec.Ddata.MYseq128) {
		_CpfN(" 839195 05 : _TdataMachine : error id/seq . {%s}", ___Vdec.String())
		return
	}

	__VconnPort := _TudpConnPort{
		DstAddr: ___Vdec.RemoteAddr, // net.UDPAddr
		K256:    ___Vdec.RemotePortKey,
	} // []byte

	// _TdataMachinEconnectSt
	__VcpArr := __VconnPort._FdataMachin__1000201x12__appendConnPort(&__Vold.dmmConnPortArr)
	___Vdm.dmMconn.M[__Vk] = _TdataMachinEconnectClient{ // _TdataMachinEconnectClient
		dmmID:             __Vold.dmmID,
		dmmLastReceTime:   _FtimeInt(),
		dmmConnPortArr:    __VcpArr,
		dmmConnPortAmount: len(__VcpArr),
	}
}
