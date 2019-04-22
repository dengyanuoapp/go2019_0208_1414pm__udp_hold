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
	case LoadT__data_01_idle:
		// do continue
	default:
		_FpfN(" 839195 02 : _TdataMachine : unknown type :{%s}", ___Vdec.String())
		return
	}

	defer ___Vdm.dmMdata.ddsMux.Unlock() // _TdataMachinEdataSt
	___Vdm.dmMdata.ddsMux.Lock()         // _TdataMachinEdataSt

	__Vk := _FgenB16(&___Vdec.Ddata.MEidx128)
	//___Vdm.dmMdata.ddsMux.Lock()
	__Vidx4, __Vok4 := ___Vdm.dmMdata.ddsMidx[__Vk] // map[[16]byte]_TdataMachinEdataClient

	if false == __Vok4 { // if not exist ,
		_CpfN(" 839195 03 : _TdataMachine : no such id exist . {%s}", ___Vdec.String())
		return
	}

	// if exist , but  token / sequence not match , replace the old one
	if                                                                                            // _TdataMachinEdataSt
	false == bytes.Equal(___Vdm.dmMdata.ddsMm[__Vidx4].ddcID.diIdx128, ___Vdec.Ddata.MEidx128) || // _Tdecode _TdataTran
		false == bytes.Equal(___Vdm.dmMdata.ddsMm[__Vidx4].ddcID.diSeq128, ___Vdec.Ddata.MYseq128) { // _TdataMachinEdataClient
		_CpfN(" 839195 05 : _TdataMachine : error id/seq . {%s}", ___Vdec.String())
		return
	}

	__VconnPort := _TudpConnPort{
		DstAddr: ___Vdec.RemoteAddr, // net.UDPAddr
		K256:    ___Vdec.RemotePortKey,
	} // []byte

	// _TdataMachinEdataSt
	__VcpArr := __VconnPort._FdataMachin__1000201x12__appendConnPort(&(___Vdm.dmMdata.ddsMm[__Vidx4].ddcConnPortArr))
	// ddcID:             ___Vdec.Ddata.ddsMm[__Vidx4].ddcID, // _TdataTran
	___Vdm.dmMdata.ddsMm[__Vidx4].ddcLastReceTime = _FtimeInt() // _Tdecode
	___Vdm.dmMdata.ddsMm[__Vidx4].ddcConnPortArr = __VcpArr     // _TdataMachinEdataClient
	___Vdm.dmMdata.ddsMm[__Vidx4].ddcConnPortAmount = len(__VcpArr)

	//_CFpfN(" 839195 08 : _TdataMachine : rece{%s}---{%s}", ___Vdec.String(), ___Vdm.dmMdata.String())
}
