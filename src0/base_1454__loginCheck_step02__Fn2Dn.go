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

	__Vk := _FgenB16(&___Vdecode.Dlogin.MeIdx128)

	if len(___Vlc.ulCmd.M) > 300 {
		_FdeleteOld_cmdStack(&___Vlc.ulCmd.M)
	}

	___Vdecode.Dlogin.TokenR = _FgenRand_nByte__(16)

	___Vlc.ulCmd.M[__Vk] = *___Vdecode

	*___Venc = _Tencode{
		EnToConnPort: _TudpConnPort{
			DstAddr: ___Vdecode.RemoteAddr,    // net.UDPAddr
			K256:    ___Vdecode.RemotePortKey, // []byte
		},
		EnLoadType: LoadT__loginS02genReplyTokenB,
		EnLogin: _TloginReq{
			MeRand5:  _FgenRand_nByte__(5),
			MeTime:   _FtimeInt(),
			ReqStr:   " step102y__sReply_tokenB ",
			MeName:   _VC.Name,
			MeIdx128: _VC.MyId128,
			MeSeq128: _VS.MySeq128,
			ToIdx128: ___Vdecode.Dlogin.MeIdx128, // []byte
			ToSeq128: ___Vdecode.Dlogin.MeSeq128, // []byte
			TokenL:   ___Vdecode.Dlogin.TokenR,   // []byte
			TokenR:   ___Vdecode.Dlogin.TokenL,   // []byte
		},
		EnMultiSend: 4 + (_FgenRand_int() % 7), // gen 0 -- 6 , + 4 : so total 5 - 10
	}
}
