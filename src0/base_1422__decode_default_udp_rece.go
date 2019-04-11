package main

import "sync"

var ___VudpDecodeMux sync.Mutex

func _FudpDecode__700201x__receive__default(___Vutm *_TuDecode) {
	go _FudpDecode__700201x10__receive__default(___Vutm)
}

func _FudpDecode__700201x10__receive__default(___Vutm *_TuDecode) {
	//_FpfNdb(" 388195 01 : filte received start ")
	var __VunRece _TudpNodeDataRece
	var __Vdecode _Tdecode
	for {
		//_Fsleep_100s()
		//__VunRece = _TudpNodeDataRece{}
		//__Vdecode = _Tdecode{}
		select {
		case __VunReceB := <-___Vutm.uTmCHunDataReceI: // _TudpNodeDataReceX
			__VunRece = _TudpNodeDataRece{}
			__Verr0 := _FdecGob___(" 388195 02 err ", __VunReceB, &__VunRece)
			___VtraceIntDE := __VunRece.Ti
			if nil != __Verr0 {
				_FpfN(" 388195 03 , decGob error :%v ", __Verr0)
				continue
			}
			_CpfN(" 388195 04 Ti:%d : before decoder : __VunReceB (%d){%x}[%s} __VunRece{%s} ",
				___VtraceIntDE,
				len(__VunReceB),
				_FgenMd5__5(&__VunReceB),
				String9(&__VunReceB),
				__VunRece.String())

			if 0 != len(__VunRece.UrrReceiveKey.Bkey) { // _Tkey256X
				__Vtmp3in := __VunRece.UrrBuf
				___VudpDecodeMux.Lock()
				__Vtmp3out, __Verr2 :=
					_FdecAesRand__only(__VunRece.UrrReceiveKey.Bkey, __Vtmp3in, ___VtraceIntDE)
				___VudpDecodeMux.Unlock()
				if nil != __Verr2 {
					_CpfN(" 388195 08 Ti:%d AesDec error {%v} {%s}", ___VtraceIntDE, __Verr2, __VunRece.String())
					_FpfN(" 388195 09 Ti:%d AesDec error {%v} {%s}", ___VtraceIntDE, __Verr2, __VunRece.String())
					continue
				}
				__VunRece.UrrBuf = __Vtmp3out
			}

			__VunRece.
				_FdataPack__decode_from_udpNodeDataRece(&__Vdecode) // _TdecodeX
			__Vdecode.remoteAddr = __VunRece.UrrRemoteAddr

			if 3 == 2 {
				_CpfN(" 388196 04 Ti:%d : after decoder  : __Vdecode {%s} (from %d:%x) ::: __VunRece {%s} (from %d:%x)", // _TudpNodeDataReceX
					___VtraceIntDE,
					__Vdecode.String(), len(__VunRece.UrrBuf), _FgenMd5__5(&__VunRece.UrrBuf), // _TdecodeX
					__VunRece.String(), len(__VunRece.UrrBuf), _FgenMd5__5(&__VunRece.UrrBuf))
			}

			switch __Vdecode.Type {
			case Cmd__loginS01genReplyTokenA, Cmd__loginS02genReplyTokenB,
				Cmd__loginS03acceptWithToken: // , Cmd__loginS04acceptWithToken: // 15540362231554036223
				//_FpfN(" 388196 06 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
				if nil == ___Vutm.uTmCHdecodeCmdLO {
					_FpfN(" 388196 07 : outChan null , ignore:%s", __Vdecode.String())
				} else {
					//_FpfN(" 388196 08 real outChain : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
					//_FpfN(" 388196 09 real outChain : %s", __Vdecode.String())
					(*___Vutm.uTmCHdecodeCmdLO) <- __Vdecode // 15540463611554046361
				}
			case Cmd__data_01_idle:
				_FpfN(" 388197 01 : type %d, tokenA %x", __Vdecode.Type, __Vdecode.Dlogin.TokenL)
				if nil == ___Vutm.uTmCHdecodeDataLO {
					_FpfN(" 388197 03 : outChan null , ignore ")
				} else {
					_FpfN(" 388197 05 real outChain : type %d", __Vdecode.Type)
					(*___Vutm.uTmCHdecodeDataLO) <- __Vdecode //
				}
			default:
				_FpfN(" 388197 09 : type %d : unknow how to deal with.", __Vdecode.Type)
			}
		}
	}
}
