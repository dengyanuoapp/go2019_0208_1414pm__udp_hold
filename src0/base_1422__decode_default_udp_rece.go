package main

import "sync"

var ___VudpDecodeMux sync.Mutex

func _FudpDecode__700201x__receive__default(___Vude *_TuDecode) {
	go _FudpDecode__700201x10__receive__default(___Vude)
}

func _FudpDecode__700201x10__receive__default(___Vude *_TuDecode) {
	//_FpfNdb(" 388195 01 : filte received start ")
	//var __VunRece _TudpNodeDataRece
	for {
		//__VunRece := _TudpNodeDataRece{}
		__VunReceB := []byte{}
		select {
		case __VunReceB = <-___Vude.uDeCHreceUgByteI: // _TudpNodeDataReceX
			___VudpDecodeMux.Lock()
		case __VunReceB = <-___Vude.uDeCHreceUnByteI: // _TudpNodeDataReceX
			___VudpDecodeMux.Lock()

		} // end Select

		_FpfNonce(" 388195 02 udpDecode rece : <%s>", String9(&__VunReceB))
		___CpfN(" 388195 03 udpDecode rece : <%s>", String9(&__VunReceB))

		//func _FdecGob___(___VeMsg string, ___VbyteIn []byte, ___VoutObjLp interface{}) error {
		__VunRece := _TudpNodeDataRece{}
		__Verr4 := _FdecGob___(" 388195 05 ", __VunReceB, &__VunRece)
		if nil != __Verr4 {
			_CpfN(" 388195 06 udpDecode err : <%v>", __Verr4)
			_FpfN(" 388195 07 udpDecode err : <%v>", __Verr4)
		} else {
			___CpfN(" 388195 08 udpDecode receive: <%s>", __VunRece.String())   // _TudpNodeDataReceX
			_FpfNonce(" 388195 09 udpDecode receive: <%s>", __VunRece.String()) // _TudpNodeDataReceX

			___Vude.
				_FudpDecode__700201x11__receive__default(&__VunRece)
		}

		___VudpDecodeMux.Unlock()
	} // end For
}

func (___Vude *_TuDecode) _FudpDecode__700201x11__receive__default(___VunRece *_TudpNodeDataRece) {
	var __Vdecode _Tdecode
	___VtraceIntDE := ___VunRece.Ti

	_FpfNonce(" 388196 01 Ti:%d : before decoder : ___VunRece{%s} ", ___VtraceIntDE, ___VunRece.String())
	//_CpfN(" 388196 04 Ti:%d : before decoder : ___VunRece{%s} ", ___VtraceIntDE, ___VunRece.String())

	if 0 != len(___VunRece.UrrReceiveKey.Bkey) { // _Tkey256X
		__Vtmp3in := ___VunRece.UrrBuf
		__Vtmp3out, __Verr2 :=
			_FdecAesRand__only(___VunRece.UrrReceiveKey.Bkey, __Vtmp3in, ___VtraceIntDE)
		if nil != __Verr2 {
			_CpfN(" 388196 08 Ti:%d AesDec error {%v} {%s}", ___VtraceIntDE, __Verr2, ___VunRece.String()) // keykey
			_FpfN(" 388196 09 Ti:%d AesDec error {%v} {%s}", ___VtraceIntDE, __Verr2, ___VunRece.String())
			return
		}
		___VunRece.UrrBuf = __Vtmp3out
	}

	___VunRece.
		_FdataPack__decode_from_udpNodeDataRece(&__Vdecode) // _TdecodeX
	__Vdecode.RemoteAddr = ___VunRece.UrrRemoteAddr

	___CpfN(" 388197 04 Ti:%d : after decoder  : Type:%d __Vdecode {%s} (from %d:%x) ::: ___VunRece {%s} (from %d:%x)", // _TudpNodeDataReceX
		___VtraceIntDE, __Vdecode.Type,
		__Vdecode.String(), len(___VunRece.UrrBuf), _FgenMd5__5(&___VunRece.UrrBuf), // _TdecodeX
		___VunRece.String(), len(___VunRece.UrrBuf), _FgenMd5__5(&___VunRece.UrrBuf))

	switch __Vdecode.Type {
	case Cmd__loginS01genReplyTokenA, Cmd__loginS02genReplyTokenB,
		Cmd__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken: // 15540362231554036223
		___CpfN(" 388197 06 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
		if nil == ___Vude.uDeCHdecodeCkLO {
			_CpfN(" 388197 07 : uDecode outChan null , ignore:%s", __Vdecode.String())
		} else {
			___CpfN(" 388197 08 uDecode real outChain : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
			___CpfN(" 388197 09 uDecode real outChain : %s", __Vdecode.String())
			(*___Vude.uDeCHdecodeCkLO) <- __Vdecode // 15540463611554046361
		}
	case Cmd__data_01_idle:
		_CpfN(" 388199 01 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
		if nil == ___Vude.uDeCHdecodeDataLO {
			_CpfN(" 388199 03 : outChan null , ignore ")
		} else {
			_CpfN(" 388199 05 outChain exist , push into chan: _TuDecode receive : type %d :<%s>", __Vdecode.Type, __Vdecode.String())
			(*___Vude.uDeCHdecodeDataLO) <- __Vdecode // to &_VdataMachine_Fn.dmCHdecodeDataI,     // dmCHdecodeDataI _TdecodeX
		}
	default:
		_CpfN(" 388199 09 : type %d : unknow how to deal with.", __Vdecode.Type)
	}
}
