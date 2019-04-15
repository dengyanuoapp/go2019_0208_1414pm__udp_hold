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
		case __VunReceB = <-___Vude.uTmCHreceUgByteI: // _TudpNodeDataReceX
			___VudpDecodeMux.Lock()
		case __VunReceB = <-___Vude.uTmCHreceUnByteI: // _TudpNodeDataReceX
			___VudpDecodeMux.Lock()

		} // end Select

		_FpfN(" 388195 01 rece : (%d)<%s>", __VunReceB, String9(&__VunReceB))
		_CpfN(" 388195 02 rece : (%d)<%s>", __VunReceB, String9(&__VunReceB))

		//func _FdecGob___(___VeMsg string, ___VbyteIn []byte, ___VoutObjLp interface{}) error {
		__VunRece := _TudpNodeDataRece{}
		__Verr4 := _FdecGob___(" 388195 02 ", __VunReceB, &__VunRece)
		if nil != __Verr4 {
			_FpfN(" 388195 03 err : <%v>", __Verr4)
		} else {
			_FpfN(" 388195 04 udpDecoder receive: <%s>", __VunRece.String()) // _TudpNodeDataReceX
			___Vude.
				_FudpDecode__700201x11__receive__default(&__VunRece)
		}

		___VudpDecodeMux.Unlock()
	} // end For
}

func (___Vude *_TuDecode) _FudpDecode__700201x11__receive__default(___VunRece *_TudpNodeDataRece) {
	var __Vdecode _Tdecode
	___VtraceIntDE := ___VunRece.Ti

	_CpfN(" 388195 04 Ti:%d : before decoder : ___VunRece{%s} ",
		___VtraceIntDE,
		___VunRece.String())

	if 0 != len(___VunRece.UrrReceiveKey.Bkey) { // _Tkey256X
		__Vtmp3in := ___VunRece.UrrBuf
		__Vtmp3out, __Verr2 :=
			_FdecAesRand__only(___VunRece.UrrReceiveKey.Bkey, __Vtmp3in, ___VtraceIntDE)
		if nil != __Verr2 {
			_CpfN(" 388195 08 Ti:%d AesDec error {%v} {%s}", ___VtraceIntDE, __Verr2, ___VunRece.String()) // keykey
			_FpfN(" 388195 09 Ti:%d AesDec error {%v} {%s}", ___VtraceIntDE, __Verr2, ___VunRece.String())
			return
		}
		___VunRece.UrrBuf = __Vtmp3out
	}

	___VunRece.
		_FdataPack__decode_from_udpNodeDataRece(&__Vdecode) // _TdecodeX
	__Vdecode.remoteAddr = ___VunRece.UrrRemoteAddr

	_CpfN(" 388196 04 Ti:%d : after decoder  : __Vdecode {%s} (from %d:%x) ::: ___VunRece {%s} (from %d:%x)", // _TudpNodeDataReceX
		___VtraceIntDE,
		__Vdecode.String(), len(___VunRece.UrrBuf), _FgenMd5__5(&___VunRece.UrrBuf), // _TdecodeX
		___VunRece.String(), len(___VunRece.UrrBuf), _FgenMd5__5(&___VunRece.UrrBuf))

	switch __Vdecode.Type {
	case Cmd__loginS01genReplyTokenA, Cmd__loginS02genReplyTokenB,
		Cmd__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken: // 15540362231554036223
		//_FpfN(" 388196 06 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
		if nil == ___Vude.uTmCHdecodeCkLO {
			_FpfN(" 388196 07 : outChan null , ignore:%s", __Vdecode.String())
		} else {
			//_FpfN(" 388196 08 real outChain : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
			//_FpfN(" 388196 09 real outChain : %s", __Vdecode.String())
			(*___Vude.uTmCHdecodeCkLO) <- __Vdecode // 15540463611554046361
		}
	case Cmd__data_01_idle:
		_FpfN(" 388197 01 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
		if nil == ___Vude.uTmCHdecodeDataLO {
			_FpfN(" 388197 03 : outChan null , ignore ")
		} else {
			_FpfN(" 388197 05 outChain exist , push into chan: _TuDecode receive : type %d", __Vdecode.Type)
			(*___Vude.uTmCHdecodeDataLO) <- __Vdecode //
		}
	default:
		_FpfN(" 388197 09 : type %d : unknow how to deal with.", __Vdecode.Type)
	}
}
