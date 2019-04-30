package main

import "reflect"

func (___Vlc *_TloginCheck) _FloginCheck_step900201y__s2Reply_tokenB_fill02send_Fn(___Vdecode *_Tdecode, ___Venc *_Tencode) {
	if _FcheckDecodeType(___Vdecode, LoadT__loginS01genReplyTokenA) {
		_FpfNdb(" 838393 01 type error , ignore ")
		return
	}

	__VlenArr := ___Vdecode.Count128()
	if false == reflect.DeepEqual(__VlenArr, __Vstep02_LoginLenArr) {
		_FpfN(" 838393 05 _TloginCheck :%s", ___Vlc.String())
		_FpfN(" 838393 06 _Tdecode :%s", ___Vdecode.String())
		_FpfNdb(" 838393 07 len error , ignore %d ", __VlenArr)
		return
	}

	__Vk := _FgenB16(&___Vdecode.DElogin.LgMeIdx128)

	if len(___Vlc.ulCmd.M) > 300 {
		_FdeleteOld_cmdStack(&___Vlc.ulCmd.M)
	}

	___Vdecode.DElogin.LgTokenR = _FgenRand_nByte__(16)

	___Vlc.ulCmd.M[__Vk] = *___Vdecode

	*___Venc = _Tencode{
		EnToConnPort: _TudpConnPort{
			DstAddr: ___Vdecode.DEremoteAddr,    // net.UDPAddr
			K256:    ___Vdecode.DEremotePortKey, // []byte
		},
		EnLoadType: LoadT__loginS02genReplyTokenB,
		EnLogin: _TloginReq{
			LgMeRand5:  _FgenRand_nByte__(5),
			LgMeTime:   _FtimeInt(),
			LgReqStr:   " step102y__sReply_tokenB ",
			LgMeName:   _VC.Name,
			LgMeIdx128: _VC.MyId128,
			LgMeSeq128: _VS.MySeq128,
			LgToIdx128: ___Vdecode.DElogin.LgMeIdx128, // []byte
			LgToSeq128: ___Vdecode.DElogin.LgMeSeq128, // []byte
			LgTokenL:   ___Vdecode.DElogin.LgTokenR,   // []byte
			LgTokenR:   ___Vdecode.DElogin.LgTokenL,   // []byte
		},
		EnMultiSend: 4 + (_FgenRand_int() % 7), // gen 0 -- 6 , + 4 : so total 5 - 10
	}
}
