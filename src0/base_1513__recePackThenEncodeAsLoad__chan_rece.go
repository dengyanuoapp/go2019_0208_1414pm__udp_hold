package main

import "sync"

var _VrecePackThenEncodeAsLoad__1400201x__mux sync.Mutex

func (___Vpel *_TrecePackThenEncodeAsLoad) _FrecePackThenEncodeAsLoad__1400201x__chan_rece__default() {
	for {

		select {

		case __VpB := <-___Vpel.pelCHudpNodeDataReceBI:
			_VrecePackThenEncodeAsLoad__1400201x__mux.Lock()

			__FpfN(" 638191 01 _TrecePackThenEncodeAsLoad rece Bytes From Chan :{%s}", String9(&__VpB))
			___Vpel.
				_FrecePackThenEncodeAsLoad__1400201y__decode_and_check_and_repack(&__VpB)

			//			if nil != ___Vpel.pelCHoutLO {
			//				(*___Vpel.pelCHoutLO) <- __Vp
			//			}
		case <-___Vpel.pelChOutputGenGap:
			_VrecePackThenEncodeAsLoad__1400201x__mux.Lock()
			_CFpfN(" 638191 08 _TrecePackThenEncodeAsLoad : genOutput")
		}
		_VrecePackThenEncodeAsLoad__1400201x__mux.Unlock()
	}
}

//_FgapFilter__1200301x3__Byte_rece
// _FudpDecode__700201x11__receive__default
func (___Vpel *_TrecePackThenEncodeAsLoad) _FrecePackThenEncodeAsLoad__1400201y__decode_and_check_and_repack(___VbyteIn *[]byte) {
	if nil == ___VbyteIn {
		_FpfNex(" 638196 00 : _TrecePackThenEncodeAsLoad : why input byte Chan null ?")
		return
	}

	if nil == ___Vpel.pelCHc2sEncodeBLO {
		_CFpfN(" 638196 01 : _TrecePackThenEncodeAsLoad : why out-Chan null <%s> ?", String9(___VbyteIn))
		return
	}

	__VunRece := _TudpNodeDataRece{}
	__Verr4 := _FdecGob___(" 638196 02 ", *___VbyteIn, &__VunRece)
	if nil != __Verr4 {
		_CFpfN(" 638196 02 : why error <%v> ?", __Verr4)
		return
	}

	if 32 != len(__VunRece.UrrReceiveKey.Bkey) { // _Tkey256X
		_CFpfN(" 638196 03 : why error <%v> ?", __Verr4)
		return
	}

	__VtraceIntDE := __VunRece.Ti
	__Vtmp3in := __VunRece.UrrBuf
	__Vtmp3out, __Verr2 := // byte of _Tdecode
		_FdecAesRand__only(__VunRece.UrrReceiveKey.Bkey, __Vtmp3in, __VtraceIntDE)
	if nil != __Verr2 {
		_CFpfN(" 638196 04 Ti:%d AesDec error {%v} {%s}", __VtraceIntDE, __Verr2, __VunRece.String()) // keykey
		return
	}

	__Vdecode := _Tdecode{} // _TdecodeX
	//__VunRece.
	//	_FdataPack__dePack_decode__from_udpNodeDataRece(&__Vdecode) // _TdecodeX

	_FdataPack__dePackUdpNodeRece__decode(&__Vdecode, __Vtmp3out)
	if __Vdecode.Type != LoadT__loginS01genReplyTokenA {
		_CFpfN(" 638196 05 Ti:%d decode error {%v} error , get type %d  , want type %d ",
			__VtraceIntDE, __Verr2, __Vdecode.Type, LoadT__loginS01genReplyTokenA)
		return
	}

	_CFpfN(" 638196 06 _TrecePackThenEncodeAsLoad: dec{%s} ====####==== unRece{%s}", __Vdecode.String(), __VunRece.String())
	/*

		    Fn:1556005752  638191 x1 _TrecePackThenEncodeAsLoad rece Bytes From Chan :{(656){fb5a3b0ccc}[68ff81030101115f54]}

		    Fn:1556005752  638196 x6 _TrecePackThenEncodeAsLoad: dec{Ti:0 ok:bool rm::0 rmk:e3a10d3210 type:LoadT__loginS01genReplyTokenA
		    {Dlogin:rand5:eecf0ef53d ti:1556005735 req: loginS01genReplyTokenA  mn:Cn mid:ff11afd83a,fc0495fdda tid:, tkAB:504387d547,}
		    t:1556005752 } ====####==== unRece{Ti:57349055 mp:48881 ra:127.0.0.1:42228 buf(321){a96049b0a9}[(321){a96049b0a9}[2f7a54b6c9f136dd50]]
		    k[k1:321893a732 k2:321893a732 dis:F]}

		    this package is set from the Cn's  port ra:127.0.0.1:42228 , the 42228's receKey is rmk:e3a10d3210, Cn's id is mid:ff11afd83a,fc0495fdda

		    _TencodeX
			_TdataTran

	*/

	__Vc2s := _Trepack__c2s{
		C2sDe:   __Vdecode,               // _Tdecode
		C2sAddr: __VunRece.UrrRemoteAddr, // net.UDPAddr _TudpNodeDataReceX
	}

	__VoutC2s, __Verr3 := _FencGob__only(&__Vc2s)
	if nil != __Verr3 {
		_CFpfN(" 638196 07 why encGob error ? <%v> , {%s} ", __Verr3, __Vc2s.String())
		return
	}

	__Venc4 := _Tencode{ // _TencodeX
		Ti:         __VunRece.Ti,              // _TdecodeX
		EnToId128:  __Vdecode.Dlogin.MeIdx128, // _TloginReq
		EnLoadType: LoadT__data_01_special,    //byte
		EnData: _TdataTran{
			DDcmd: DDType__c2s, // byte
			DDbuf: __VoutC2s,   // __Vtmp3out , byte of _Tdecode ; __VunRece _TudpNodeDataReceX
			// MEidx128 []byte
			// MYseq128 []byte
			// TOidx128 []byte
			// TOseq128 []byte
			// TTokenD  []byte
			// DDoffset uint64
		},
		//EnToConnPort _TudpConnPort // another of to addr
		//EnLogin      _TloginReq
		//EnDelay      int // a delay resend if not zero
		//EnMultiSend  int // send multi timeS if not zero
	}

	_CFpfN("\n\n\n 638196 08 _TrecePackThenEncodeAsLoad: encOut{%s}", __Venc4.String())

	__VoutC2sB, __Verr4 := _FencGob__only(&__Venc4)
	if nil != __Verr4 {
		_CFpfN(" 638196 09 why encGob error ? <%v> , {%s} ", __Verr4, __Venc4.String())
		return
	}
	_CFpfN(" 638196 10 _TrecePackThenEncodeAsLoad: encOutB{%s} {%s}", String9(&__VoutC2sB), _Fmd5__5s(&__VoutC2sB))

	// (*(___Vpel.pelCHc2sEncodeBLO)) <- __VoutC2sB

	__Venc5 := _Tencode{} // _TencodeX
	__Verr5 := _FdecGob___(" 638196 11 ", __VoutC2sB, &__Venc5)
	if nil != __Verr5 {
		_CFpfN(" 638196 12 _TrecePackThenEncodeAsLoad: decError{%v} ", __Verr5)
	}
	_CFpfN(" 638196 13 _TrecePackThenEncodeAsLoad: recoverEc{%s} ", __Venc5.String())
	_FpfNex(" 638196 14")
}
