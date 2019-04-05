package main

import (
	"bytes"
)

func (___Vdm *_TdataMachine) _FdataMachin__1000201x11__connMap_insertId(___VinsID *_TdataMachinEid) {

	__Vk := _FgenB16(&___VinsID.diIdx128)
	___Vdm.dmMconn.mux.Lock()
	__Vold, __Vok := ___Vdm.dmMconn.M[__Vk] // map[[16]byte]_TdataMachinEconnMap
	if __Vok {
		if //
		false == bytes.Equal(__Vold.dmmID.diIdx128, ___VinsID.diIdx128) ||
			false == bytes.Equal(__Vold.dmmID.diSeq128, ___VinsID.diSeq128) ||
			false == bytes.Equal(__Vold.dmmID.diToken, ___VinsID.diToken) {
			_FpfN(" 839193 03 : delete-old, gen New connPort Map hash.")
			___Vdm.dmMconn.M[__Vk] = _TdataMachinEconnMap{
				dmmID:          *___VinsID,
				dmmLastAccTime: _FtimeInt(),
				dmmConnPortArr: []_TudpConnPort{___VinsID.diConnPort},
			}
		} else {
			//_FpfN(" 839193 05 : with the same token, so , try append to connPort Map hash.")
			___Vdm.dmMconn.M[__Vk] = _TdataMachinEconnMap{
				dmmID:          *___VinsID,
				dmmLastAccTime: _FtimeInt(),
				dmmConnPortArr: ___VinsID.diConnPort._FdataMachin__1000201x12__appendConnPort(&__Vold.dmmConnPortArr),
			}
		}
	} else {
		___Vdm.dmMconn.M[__Vk] = _TdataMachinEconnMap{
			dmmID:          *___VinsID,
			dmmLastAccTime: _FtimeInt(),
			dmmConnPortArr: []_TudpConnPort{___VinsID.diConnPort},
		}
	}
	___Vdm.dmMconn.mux.Unlock()
	_FpfN(" 839193 08 : dmMconn.M len(%d,%d) ", len(___Vdm.dmMconn.M), len(___Vdm.dmMconn.M[__Vk].dmmConnPortArr))
	if 3 == len(___Vdm.dmMconn.M[__Vk].dmmConnPortArr) {
		//_FpfN(" 839193 09 : dmMconn.M {%v}", ___Vdm.dmMconn.M[__Vk].dmmConnPortArr)
	}
}

func (___VnewConnP *_TudpConnPort) _FdataMachin__1000201x12__appendConnPort(___VoldConnParr *[]_TudpConnPort) []_TudpConnPort {
	if nil == ___VoldConnParr || 0 == len(*___VoldConnParr) {
		_FpfN(" 839194 01 : nil , use old connPort Arr ")
		return []_TudpConnPort{*___VnewConnP}
	}

	__Vlen := len(*___VoldConnParr)
	__Vncp := make([]_TudpConnPort, __Vlen+1)
	__Vstr1 := ___VnewConnP.DstAddr.String()

	__Vi := 0
	for __Vi < __Vlen {
		__Vcp := (*___VoldConnParr)[__Vi]
		__Vstr2 := __Vcp.DstAddr.String()
		if __Vstr1 == __Vstr2 {
			_FpfN(" 839194 03 : already exist , use old connPort Arr ")
			return (*___VoldConnParr)
		}
		__Vncp[__Vi] = __Vcp
		__Vi++
	}

	//_FpfN(" 839194 05 : insert 1 into connPort Arr ")
	__Vncp[__Vlen] = (*___VnewConnP)

	return __Vncp
}
