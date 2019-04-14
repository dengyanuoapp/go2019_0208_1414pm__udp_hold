package main

var (
	_Vgap_nothingToLost  = 39 // 3*(12+1) == 39
	_Vgap_skip_idle_send = 7
)

func (___Vdm *_TdataMachine) _FdataMachin__1000502x__time_gap_dataSendIdle() {

	for {
		_FsleepRand_8_to_12s()
		//dmmLastReadTime
		__VkDelArr := [][16]byte{}
		__Vnow2 := _FtimeInt()

		___Vdm.dmMconn.mux.Lock()
		for __Vk2, __Vv2 := range ___Vdm.dmMconn.M { // M map[[16]byte]_TdataMachinEconnMap
			if __Vnow2-__Vv2.dmmLastReadTime > _Vgap_nothingToLost {
				__VkDelArr = append(__VkDelArr, __Vk2)
			} else {
				if __Vnow2-__Vv2.dmmLastReadTime > _Vgap_skip_idle_send {
					// pack as _TdataTran -->  _TdecodeX .  Ddata // _TencodeX
					_FpfN(" 381921 01 : %11d : try send idle %x %d,%d", _FtimeInt(), __Vk2, __Vnow2, __Vv2.dmmLastReadTime)

					___Vdm.
						_FdataMachin__1000502x2__time_gap_dataSendIdle(__Vk2[:], __Vv2)

				} else {
					_FpfN(" 381921 02 : %11d : try send idle %x , but in 10s sent already. Skip. %d,%d",
						_FtimeInt(), __Vk2, __Vnow2, __Vv2.dmmLastReadTime)
				}
				//if nil != ___Vdm.dmCBprSendKey {
				//	___Vdm.dmCBprSendKey(___Vdm)
				//}
				// _TdataMachinEconnMap
				if __Vnow2-__Vv2.dmmLastPrTime > 20 {
					for __Vidx, __Vconn := range __Vv2.dmmConnPortArr {
						_CpfN(" 381921 03 sendingKeyArray: %x : %2d : to [%s]", // check-important keykey
							__Vconn.K256, // []byte
							__Vidx,
							__Vconn.DstAddr.String(), // net.UDPAddr
						)
					}
					__Vv2.dmmLastPrTime = __Vnow2
				}
			}
		}

		for _, __Vk3 := range __VkDelArr {
			_FpfN(" 381921 05 : %11d : try delete lost connect %x in dataMachine", _FtimeInt(), __Vk3)

			if nil == ___Vdm.dmCHloginGenMachineIdLO {
				_FpfNonce(" 381921 07 : %11d : try to stop lost connect %x in loginGen , but NULL out-chain", _FtimeInt(), __Vk3)
			} else {
				_FpfNonce(" 381921 08 : %11d : try to stop lost connect %x in loginGen, ok ", _FtimeInt(), __Vk3)

				___Vout__dmCHloginGenMachineIdLO__mux.Lock()

				(*___Vdm.dmCHloginGenMachineIdLO) <- _TdataMachinEid{
					diIdx128: ___Vdm.dmMconn.M[__Vk3].dmmID.diIdx128,
				}

				___Vout__dmCHloginGenMachineIdLO__mux.Unlock()
			}

			delete(___Vdm.dmMconn.M, __Vk3)
			delete(___Vdm.dmMconn.LockLast, __Vk3)
			delete(___Vdm.dmMconn.LockNow, __Vk3)
		}
		___Vdm.dmMconn.mux.Unlock()
	}
}

func (___Vdm *_TdataMachine) _FdataMachin__1000502x2__time_gap_dataSendIdle(___Vid []byte, ___Vdmem _TdataMachinEconnMap) {
	__Vlen := len(___Vdmem.dmmConnPortArr) // _TdataMachinEconnMap

	if __Vlen < 1 {
		_FpfN(" 381922 01 len error :%d ", __Vlen)
		return
	}

	__Vidx := _FtimeInt() % __Vlen
	__Vconn := ___Vdmem.dmmConnPortArr[__Vidx]

	__Venc := _Tencode{ // _TencodeX // Cmd__loginS01genReplyTokenA
		Ti:           _FtimeInt(),
		enType:       Cmd__data_01_idle,
		enToId128:    ___Vdmem.dmmID.diIdx128, // _TdataMachinEid
		enToConnPort: __Vconn,                 // _TudpConnPort
		enData: _TdataTran{
			meIdx128: _VC.MyId128,
			MySeq128: _VS.MySeq128,
			toIdx128: ___Vdmem.dmmID.diIdx128,
			toSeq128: ___Vdmem.dmmID.diSeq128,
			Dcmd:     Cmd__data_01_idle,
			tokenD:   []byte{},
			Doffset:  0,
			Dbuf:     []byte{},
		},
	}

	_FpfN(" 381922 02 {%#v} ", ___Vdmem) // _TdataMachinEconnMap
	_FpfN(" 381922 03 {%s} ", ___Vdmem.String())
	_FpfN(" 381922 04 : myID:%x mySeq:%x id:%x", _VC.MyId128, _VS.MySeq128, ___Vid)
	_FpfN(" 381922 05 {%s}", __Venc.String()) // _TencodeX
	_Fex("381922 06")

	___Vdm.
		_FdataMachin__1000601x__encodeData_sendMux(&__Venc)
}
