package main

import "sync"

var _VrecePackThenEncodeAsLoad__1400201x__mux sync.Mutex

func (___Vpel *_TrecePackThenEncodeAsLoad) _FrecePackThenEncodeAsLoad__1400201x__chan_rece__default() {
	for {

		select {

		case __VpB := <-___Vpel.pelCHudpNodeDataReceBI:
			_VrecePackThenEncodeAsLoad__1400201x__mux.Lock()

			_CFpfN(" 638191 01 _TrecePackThenEncodeAsLoad rece Bytes From Chan :{%s}", String9(&__VpB))
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
	//(*___VbyteOut) = []byte{}
	//_FpfN(" 381917 23 : gfCHbyteI :[%x]", ___VbyteIn)

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
	__Vtmp3out, __Verr2 :=
		_FdecAesRand__only(__VunRece.UrrReceiveKey.Bkey, __Vtmp3in, __VtraceIntDE)
	if nil != __Verr2 {
		_CFpfN(" 638196 04 Ti:%d AesDec error {%v} {%s}", __VtraceIntDE, __Verr2, __VunRece.String()) // keykey
		return
	}

	__Vdecode := _Tdecode{} // _TdecodeX
	//__VunRece.
	//	_FdataPack__decode_from_udpNodeDataRece(&__Vdecode) // _TdecodeX

	_FdataPack__dataDecode_common(&__Vdecode, __Vtmp3out)
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

	__Venc := _Tencode{ // _TencodeX
		Ti:         __Vdecode.Ti,              // _TdecodeX
		enToId128:  __Vdecode.Dlogin.MeIdx128, // _TloginReq
		enLoadType: LoadT__data_01_special,    //byte
		enData:     _TdataTran{},
		//enToConnPort _TudpConnPort // another of to addr
		//enLogin      _TloginReq
		//enDelay      int // a delay resend if not zero
		//enMultiSend  int // send multi timeS if not zero
	}

	_CFpfN(" 638196 08 _TrecePackThenEncodeAsLoad: encOut{%s}", __Venc.String())

	//(*___VbyteOut), __Verr := _FencGob__only(___V)

}
