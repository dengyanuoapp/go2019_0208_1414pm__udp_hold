package main

//var ___VinitLoginIdx01 [16]byte = [16]byte{1, 1, 1, 1, 1, 2, 3, 4, 4, 3, 2, 1, 1, 1, 1, 1}

import "sync"

var ___VloginCheck__900201x__mux sync.Mutex

func (___Vlc *_TloginCheck) _FloginCheck__900201x__standardCheck() {
	_Fsleep(_T1s)

	_FpfN(" 138182 01 start")

	//var __Vdecode _Tdecode
	//var __VuConnPort _TudpConnPort
	for {
		//_Fsleep_100s()
		__Venc := _Tencode{}
		select {
		//case __Vdecode := <-___Vlc.ulCHdecodeCkI: // _Tdecode
		case __Vdecode := <-___Vlc.ulCHdecodeCkI: // _Tdecode
			___VloginCheck__900201x__mux.Lock()

			___Vlc.
				_FloginCheck__900201xC1__standardCheck(&__Vdecode, &__Venc)
			if false == __Vdecode.DEok {
				___VloginCheck__900201x__mux.Unlock()
				continue
			}

		case __Vdecode := <-___Vlc.ulCHrepackDecodeC2sI: // _TloginCheck _TdecodeX

			___VloginCheck__900201x__mux.Lock()

			___Vlc.
				_FloginCheck__900201xC1__standardCheck(&__Vdecode, &__Venc)
			_CFpfN(" 138182 02 : {%s}", __Venc.String())

			if false == __Vdecode.DEok {
				___VloginCheck__900201x__mux.Unlock()
				continue
			}

		case __VuConnPort := <-___Vlc.ulCHconnPortI: // _TudpConnPort

			___VloginCheck__900201x__mux.Lock()

			___Vlc.
				_FloginCheck__900201xC2__standardCheck(&__Venc, &__VuConnPort)
		} // end select

		if nil == ___Vlc.ulCHencodeCkLO { //     *chan _Tencode
			_FpfNonce(" 138184 04 : why CMD-en not Chan ? {%s}", __Venc.String())
		} else {
			//_FpfN(" 138184 05 CMD-en-chan push {%s}", __Venc.String())
			switch __Venc.EnLoadType {
			case 0:
				_FpfN(" 138184 06 zero Type , nothing need to send. %s ", __Venc.String()) // _TencodeX
			case LoadT__loginEnd:
				// no use , but told the debug disable only
			default:
				(*___Vlc.ulCHencodeCkLO) <- __Venc // _Tencode
				_CpfN(" 138184 08 My info : _VC.MyId128 %s , _VS.MySeq128 %s, __Venc.EnLogin.LgTokenL %s",
					String5s(&_VC.MyId128),
					String5s(&_VS.MySeq128),
					String5s(&__Venc.EnLogin.LgTokenL))
			}
		}
		___VloginCheck__900201x__mux.Unlock()
	} // end for
} // end func
func (___Vlc *_TloginCheck) _FloginCheck__900201xC1__standardCheck(___Vdecode *_Tdecode, ___Venc *_Tencode) {
	//_FpfNdb(" 138183 03 : %s", ___Vdecode.String()) // 15540463611554046361
	switch ___Vdecode.DEtype {
	case LoadT__loginS01genReplyTokenA:
		// ============================ step 02 : Fn gen tokenB, to Dn, cmd fill 02 ====================
		//_FpfN(" 138183 04 : %s", ___Vdecode.String())
		___Vlc.
			_FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn(___Vdecode, ___Venc)
	case LoadT__loginS02genReplyTokenB:
		// ============================ step 03 : Dn check tokenA,id128,seq128 ,ACCEPT --> cmd fill 03 ====================
		___Vlc.
			_FloginCheck_step900201y__s3accept_tokenA_fill03send_Dn(___Vdecode, ___Venc)
	case LoadT__loginS03acceptWithToken:
		// ============================ step 04 : Fn check tokenB,id128,seq128 ,ACCEPT only,no reply
		___Vlc.
			_FloginCheck_step900201y__s4accept_tokenB_resetData_Fn(___Vdecode)
		___Venc.EnLoadType = LoadT__loginEnd // no use , but told the following debug disable only
	default:
		_CFpfN(" 138183 08 : unknow how to deal with : type %d,", ___Vdecode.DEtype)
		//continue // next Select
	}
}
func (___Vlc *_TloginCheck) _FloginCheck__900201xC2__standardCheck(___Venc *_Tencode, ___VuConnPort *_TudpConnPort) {
	// ============================ step 01 : Dn gen tokenA, to anyhost, cmd fill 01 ====================
	_CFpfN(" 138184 02 : connPort-Chan pop {%s}", ___VuConnPort.String())
	___Vlc.ulTokenA = _FgenRand_nByte__(16) // tokenA / Lo
	___Vlc.ulGenTime = _FtimeInt()

	*___Venc = _Tencode{
		EnToConnPort: *___VuConnPort, // _TudpConnPort
		EnLoadType:   LoadT__loginS01genReplyTokenA,
		EnLogin: _TloginReq{
			LgMeRand5:  _FgenRand_nByte__(5),
			LgMeTime:   _FtimeInt(),
			LgReqStr:   " loginS01genReplyTokenA ",
			LgMeName:   _VC.Name,
			LgMeIdx128: _VC.MyId128,
			LgMeSeq128: _VS.MySeq128,
			LgTokenL:   ___Vlc.ulTokenA,
		},
		EnDelay: (12 + (_FtimeInt() % 3)), // 12 + (0--2) == 12--14
	}
}
