package main

//var ___VinitLoginIdx01 [16]byte = [16]byte{1, 1, 1, 1, 1, 2, 3, 4, 4, 3, 2, 1, 1, 1, 1, 1}

func (___Vlc *_TloginCheck) _FloginCheck__900201x__standardCheck() {
	_Fsleep(_T1s)
	if nil == ___Vlc.ucCHSendLO {
		_Pn()
		_Pn()
		_FpfNdb(" 838392 01 why no Chan output ? exit loop")
		_Pn()
		_Pn()
		return
	}

	_FpfN(" 838392 02 start")

	var __Vdecode _Tdecode
	var __VuConnPort _TudpConnPort
	for {
		//_Fsleep_100s()
		select {
		case __Vdecode = <-___Vlc.ucDecodeI: // _Tdecode
			//_FpfNdb(" 838392 05 : %s", __Vdecode.String()) // 15540463611554046361
			if true == __Vdecode.ok {
				switch __Vdecode.Type {
				case Cmd__loginS01genReplyTokenA:
					// ============================ step 02 : Fn gen tokenB, to Dn, cmd fill 02 ====================
					___Vlc._FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn(&__Vdecode)
				case Cmd__loginS02genReplyTokenB:
					// ============================ step 03 : Dn check tokenA,id128,seq128 ,ACCEPT --> cmd fill 03 ====================
					___Vlc.
						_FloginCheck_step900201y__s3accept_tokenA_file03send_Dn(&__Vdecode)
					//                case Cmd__loginS03acceptWithToken:
					//                    // ============================ step 04 : Fn check tokenB,id128,seq128 ,ACCEPT only,no reply
					//					___Vlc._FloginCheck_step900201y__s3accept_tokenA_file03send_Dn(&__Vdecode)
					//					//					//_FpfNdb(" 838392 06 : %x", __Vdecode.Dlogin.TokenL)
					//					//					//_FpfNdb(" 838392 07 : %s", __Vdecode.String()) // 15540463611554046361
				default:
					_FpfNdb(" 838392 08 : unknow how to deal with : type %d,", __Vdecode.Type)
				}
			} else {
				_FpfNdb(" 838392 09 : why not ok ?")
			}
		case __VuConnPort = <-___Vlc.ulCHconnPortI: // _TudpConnPort
			// ============================ step 01 : Dn gen tokenA, to anyhost, cmd fill 01 ====================
			//_FpfNdb(" 838392 10 : under constructing ? {%s}", __VuConnPort.String())
			___Vlc.ucTokenA = _FgenRand_nByte__(16) // tokenA / Lo

			var __VusData _TudpNodeDataSend // _TudpConnPort
			__VuConnPort._FdataPack__101__udpConnPort(&___Vlc.ucTokenA, &__VusData.usOutBuf)

			//___Vlc.ucCmd.mux.Lock()
			//___Vlc.ucCmd.M[___VinitLoginIdx01] = *__VloginReq // unknow the srv's id128 , so ,use fixed _TcmdMap
			//___Vlc.ucCmd.mux.Unlock()

			__VusData.usToAddr = __VuConnPort // _TudpConnPort
			*___Vlc.ucCHSendLO <- __VusData   //
			_FsleepRand_12_to_14s()
			*___Vlc.ucCHSendLO <- __VusData // 15540362231554036223
		}
	}
}
