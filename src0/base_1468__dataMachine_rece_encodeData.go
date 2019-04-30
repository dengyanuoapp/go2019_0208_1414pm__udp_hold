package main

import (
	"bytes"
)

func (___Vdm *_TdataMachine) _FdataMachin__1000201x21__rece_decodeData(___Vdec *_Tdecode) {
	/*
		__Vid := _TdataMachinEid{
			diConnPort: _TudpConnPort{
				___Vdec.remoteAddr,     // net.UDPAddr
				___Vdec.remotePortKey}, // []byte
			diIdx128: ___Vdec.DEdata.LgMeIdx128, // []byte
			diSeq128: ___Vdec.DEdata.LgMeSeq128, // []byte
			diToken:  ___Vdec.DEdata.LgTokenL,   // []byte
		}
	*/

	if nil == ___Vdec {
		_Fex(" 839195 01 : why NULL ?")
	}

	defer ___Vdm.dmMdata.ddsMux.Unlock() // _TdataMachinEdataSt
	___Vdm.dmMdata.ddsMux.Lock()         // _TdataMachinEdataSt

	if false == ___Vdec.DEok { // _TdecodeX
		_FpfN(" 839195 02 : _TdataMachine : received not OK:{%s}", ___Vdec.String())
		return
	}

	switch ___Vdec.DEtype { // _TdecodeX
	case LoadT__data_01_special:
		// do continue
		__FpfN(" 839195 03 : _TdataMachine : LoadT__data_01_special :{%s}", ___Vdec.String())
		___Vdm.
			_FdataMachin__rece__dataSpec(___Vdec)
	default:
		_CFpfN(" 839195 05 : _TdataMachine : unknown type :{%s}", ___Vdec.String())
		return
	}

	// the following : insert / renew the remote key & port pairs.

	__Vk := _FgenB16(&___Vdec.DEdata.DtMEidx128)
	//___Vdm.dmMdata.ddsMux.Lock()
	__Vidx4, __Vok4 := ___Vdm.dmMdata.ddsMidx[__Vk] // map[[16]byte]_TdataMachinEdataClient

	if false == __Vok4 { // if not exist ,
		_CpfN(" 839195 06 : _TdataMachine : no such id exist . {%s}", ___Vdec.String())
		return
	}

	// if exist , but  token / sequence not match , replace the old one
	if                                                                                               // _TdataMachinEdataSt
	false == bytes.Equal(___Vdm.dmMdata.ddsMm[__Vidx4].ddcID.diIdx128, ___Vdec.DEdata.DtMEidx128) || // _Tdecode _TdataTran
		false == bytes.Equal(___Vdm.dmMdata.ddsMm[__Vidx4].ddcID.diSeq128, ___Vdec.DEdata.DtMYseq128) { // _TdataMachinEdataClient
		_CpfN(" 839195 07 : _TdataMachine : error id/seq . {%s}", ___Vdec.String())
		return
	}

	__VconnPort := _TudpConnPort{
		DstAddr: ___Vdec.DEremoteAddr, // net.UDPAddr
		K256:    ___Vdec.DEremotePortKey,
	} // []byte

	// _TdataMachinEdataSt
	__VcpArr := __VconnPort._FdataMachin__1000201x12__appendConnPort(&(___Vdm.dmMdata.ddsMm[__Vidx4].ddcConnPortArr))
	// ddcID:             ___Vdec.DEdata.ddsMm[__Vidx4].ddcID, // _TdataTran
	___Vdm.dmMdata.ddsMm[__Vidx4].ddcLastReceTime = _FtimeInt() // _Tdecode
	___Vdm.dmMdata.ddsMm[__Vidx4].ddcConnPortArr = __VcpArr     // _TdataMachinEdataClient
	___Vdm.dmMdata.ddsMm[__Vidx4].ddcConnPortAmount = len(__VcpArr)

	//_CFpfN(" 839195 08 : _TdataMachine : rece{%s}---{%s}", ___Vdec.String(), ___Vdm.dmMdata.String())
}
