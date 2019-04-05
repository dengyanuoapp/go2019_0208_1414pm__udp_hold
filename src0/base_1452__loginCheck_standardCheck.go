package main

//var ___VinitLoginIdx01 [16]byte = [16]byte{1, 1, 1, 1, 1, 2, 3, 4, 4, 3, 2, 1, 1, 1, 1, 1}

func (___Vlc *_TloginCheck) _FloginCheck__900201x__standardCheck() {
	_Fsleep(_T1s)
	if nil == ___Vlc.ulCHSendLO {
		_Pn()
		_Pn()
		_FpfNdb(" 138182 01 why no Chan output ? exit loop")
		_Pn()
		_Pn()
		return
	}

	_FpfN(" 138182 02 start")

	//var __Vdecode _Tdecode
	//var __VuConnPort _TudpConnPort
	for {
		//_Fsleep_100s()
		select {
		case __Vdecode := <-___Vlc.ulDecodeI: // _Tdecode
			//_FpfNdb(" 138183 03 : %s", __Vdecode.String()) // 15540463611554046361
			if true == __Vdecode.ok {
				switch __Vdecode.Type {
				case Cmd__loginS01genReplyTokenA:
					// ============================ step 02 : Fn gen tokenB, to Dn, cmd fill 02 ====================
					_FpfN(" 138183 04 : %s", __Vdecode.String())
					___Vlc.
						_FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn(&__Vdecode)
				case Cmd__loginS02genReplyTokenB:
					// ============================ step 03 : Dn check tokenA,id128,seq128 ,ACCEPT --> cmd fill 03 ====================
					___Vlc.
						_FloginCheck_step900201y__s3accept_tokenA_fill03send_Dn(&__Vdecode)
				case Cmd__loginS03acceptWithToken:
					// ============================ step 04 : Fn check tokenB,id128,seq128 ,ACCEPT only,no reply
					___Vlc.
						_FloginCheck_step900201y__s4accept_tokenB_resetData_Fn(&__Vdecode)
					//					//_FpfNdb(" 138183 06 : %x", __Vdecode.Dlogin.TokenL)
					//					//_FpfNdb(" 138183 07 : %s", __Vdecode.String()) // 15540463611554046361
				default:
					_FpfNdb(" 138183 08 : unknow how to deal with : type %d,", __Vdecode.Type)
				}
			} else {
				_FpfNdb(" 138183 09 : why not ok ?")
			}
		case __VuConnPort := <-___Vlc.ulCHconnPortI: // _TudpConnPort
			// ============================ step 01 : Dn gen tokenA, to anyhost, cmd fill 01 ====================
			_FpfNdb(" 138182 01 : connPort-Chan pop {%s}", __VuConnPort.String())
			___Vlc.ulTokenA = _FgenRand_nByte__(16) // tokenA / Lo
			___Vlc.ulGenTime = _FtimeInt()

			var __VusData _TudpNodeDataSend // _TudpConnPort
			__VuConnPort._FdataPack__101__udpConnPort(&___Vlc.ulTokenA, &__VusData.usOutBuf)

			//___Vlc.ulCmd.mux.Lock()
			//___Vlc.ulCmd.M[___VinitLoginIdx01] = *__VloginReq // unknow the srv's id128 , so ,use fixed _TcmdMap
			//___Vlc.ulCmd.mux.Unlock()

			//__Vk16 := _FgenB16(&___Vdecode.Dlogin.MeIdx128)
			//___Vlc.ulCmd.M[__Vk16] = *___Vdecode

			__VusData.usToAddr = __VuConnPort // _TudpConnPort
			*___Vlc.ulCHSendLO <- __VusData   //
			_FsleepRand_12_to_14s()
			*___Vlc.ulCHSendLO <- __VusData // 15540362231554036223
		}
	}
}
