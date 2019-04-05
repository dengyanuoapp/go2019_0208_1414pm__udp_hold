package main

func _FdataMachin__1000201x__receive__default(___Vdm *_TdataMachine) {
	var __VdmID _TdataMachinEid
	for {
		select {
		case __VdmID = <-___Vdm.dmCHdataMachineIdI: // chan _TdataMachinEid
			_FpfNdb(" 839192 01 : reset-MachineID {%s}", __VdmID.String())
			___Vdm._FdataMachin__1000201x11__connMap_insertId(&__VdmID)

		}
	}
}
func (___Vdm *_TdataMachine) _FdataMachin__1000201x11__connMap_insertId(___Vid *_TdataMachinEid) {

	__Vk := _FgenB16(&__VdmID.diIdx128)
	___Vdm.dmMconNmux.Lock()
	__Vold, __Vok := ___Vdm.dmMconn[__Vk] // map[[16]byte]_TdataMachinEconnMap
	if __Vok {
		if //
		false == bytes.Equal(__Vold.dmmID.diIdx128, __VdmID.diIdx128) ||
			false == bytes.Equal(__Vold.dmmID.diSeq128, __VdmID.diSeq128) ||
			false == bytes.Equal(__Vold.dmmID.diToken, __VdmID.diToken) {
			___Vdm.dmMconn[__Vk] = _TdataMachinEconnMap{
				dmmID:          __VdmID,
				dmmLastAccTime: _FtimeInt(),
				dmmConnPortArr: __VdmID.diConnPort._FdataMachin__1000201x12__appendConnPort(&__Vold.dmmConnPortArr),
			}
		}
	} else {
		___Vdm.dmMconn[__Vk] = _TdataMachinEconnMap{
			dmmID:          __VdmID,
			dmmLastAccTime: _FtimeInt(),
			dmmConnPortArr: []_TudpConnPort{__VdmID.diConnPort},
		}
	}
	___Vdm.dmMconNmux.Unock()
}
func (___VnewConnP *_TudpConnPort) _FdataMachin__1000201x12__appendConnPort(___Vid *_TdataMachinEid) []_TudpConnPort {
	if nil == ___Vid || 0 == len(*___Vid) {
		return []_TudpConnPort{*___VnewConnP}
	}
	__Vlen := len(*___Vid)
	__Vncp := make([]_TudpConnPort, __Vlen+1)
	return __Vncp
}
