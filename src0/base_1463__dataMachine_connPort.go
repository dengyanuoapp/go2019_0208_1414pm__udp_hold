package main

import (
	"bytes"
)

func (___Vdm *_TdataMachine) _FdataMachin__1000201x11__connMap_insertId(___VinsID *_TdataMachinEid) {

	__Vk := _FgenB16(&___VinsID.diIdx128)
	___Vdm.dmMconNmux.Lock()
	__Vold, __Vok := ___Vdm.dmMconn[__Vk] // map[[16]byte]_TdataMachinEconnMap
	if __Vok {
		if //
		false == bytes.Equal(__Vold.dmmID.diIdx128, ___VinsID.diIdx128) ||
			false == bytes.Equal(__Vold.dmmID.diSeq128, ___VinsID.diSeq128) ||
			false == bytes.Equal(__Vold.dmmID.diToken, ___VinsID.diToken) {
			_FpfN(" 839193 03 : delete-old, gen New connPort Map hash.")
			___Vdm.dmMconn[__Vk] = _TdataMachinEconnMap{
				dmmID:          *___VinsID,
				dmmLastAccTime: _FtimeInt(),
				dmmConnPortArr: []_TudpConnPort{___VinsID.diConnPort},
			}
		} else {
			_FpfN(" 839193 05 : with the same token, so , try append to connPort Map hash.")
			___Vdm.dmMconn[__Vk] = _TdataMachinEconnMap{
				dmmID:          *___VinsID,
				dmmLastAccTime: _FtimeInt(),
				dmmConnPortArr: ___VinsID.diConnPort._FdataMachin__1000201x12__appendConnPort(&__Vold.dmmConnPortArr),
			}
		}
	} else {
		___Vdm.dmMconn[__Vk] = _TdataMachinEconnMap{
			dmmID:          *___VinsID,
			dmmLastAccTime: _FtimeInt(),
			dmmConnPortArr: []_TudpConnPort{___VinsID.diConnPort},
		}
	}
	___Vdm.dmMconNmux.Unlock()
	_FpfN(" 839193 09 : dmMconn len(%d) %d", len(___Vdm.dmMconn), len(___Vdm.dmMconn[__Vk].dmmConnPortArr))
}

func (___VnewConnP *_TudpConnPort) _FdataMachin__1000201x12__appendConnPort(___VoldConnParr *[]_TudpConnPort) []_TudpConnPort {
	if nil == ___VoldConnParr || 0 == len(*___VoldConnParr) {
		return []_TudpConnPort{*___VnewConnP}
	}
	__Vlen := len(*___VoldConnParr)
	__Vncp := make([]_TudpConnPort, __Vlen+1)
	return __Vncp
}
