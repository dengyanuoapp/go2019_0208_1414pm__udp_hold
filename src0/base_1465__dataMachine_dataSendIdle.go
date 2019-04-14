package main

var (
	_Vgap_nothingToLost  = 39 // 3*(12+1) == 39
	_Vgap_skip_idle_send = 7
)

func _FdataMachin__1000502x__time_gap_dataSendIdle(___Vdm *_TdataMachine) {

	for {
		_FsleepRand_8_to_12s()
		//dmmLastReadTime
		__VkDelArr := [][16]byte{}
		__Vnow2 := _FtimeInt()

		___Vdm.dmMconn.mux.Lock()
		for __Vk2, __Vcm := range ___Vdm.dmMconn.M { // M map[[16]byte]_TdataMachinEconnMap
			if __Vnow2-__Vcm.dmmLastReadTime > _Vgap_nothingToLost {
				__VkDelArr = append(__VkDelArr, __Vk2)
			} else {
				if __Vnow2-__Vcm.dmmLastReadTime > _Vgap_skip_idle_send {
					_FpfN(" 381921 01 : %11d : try send idle %x %d,%d", _FtimeInt(), __Vk2, __Vnow2, __Vcm.dmmLastReadTime)
					// pack as _TdataTran -->  _TdecodeX .  Ddata // _TencodeX
				} else {
					_FpfN(" 381921 02 : %11d : try send idle %x , but in 10s sent already. Skip. %d,%d",
						_FtimeInt(), __Vk2, __Vnow2, __Vcm.dmmLastReadTime)
				}
				//				if nil != ___Vdm.dmCBprSendKey {
				//					___Vdm.dmCBprSendKey(___Vdm)
				//				}
				// _TdataMachinEconnMap
				if __Vnow2-__Vcm.dmmLastPrTime > 20 {
					for __Vidx, __Vconn := range __Vcm.dmmConnPortArr {
						_CpfN(" 381921 03 sendingKeyArray: %x : %2d : to [%s]", // check-important keykey
							__Vconn.K256, // []byte
							__Vidx,
							__Vconn.DstAddr.String(), // net.UDPAddr
						)
					}
					__Vcm.dmmLastPrTime = __Vnow2
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
