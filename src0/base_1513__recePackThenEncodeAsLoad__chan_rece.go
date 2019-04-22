package main

import "sync"

var _VrecePackThenEncodeAsLoad__1400201x__mux sync.Mutex

func (___Vpel *_TrecePackThenEncodeAsLoad) _FrecePackThenEncodeAsLoad__1400201x__chan_rece__default() {
	var __VbOut []byte
	for {

		select {

		case __VpB := <-___Vpel.pelCHudpNodeDataReceBI:
			_VrecePackThenEncodeAsLoad__1400201x__mux.Lock()

			_CFpfN(" 638191 01 _TrecePackThenEncodeAsLoad rece Bytes From Chan :{%2}", String9(__VpB))
			___Vgf.
				_FrecePackThenEncodeAsLoad__1400201y__decode_and_check_and_repack(&__VpB, &__VbOut)

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
func (___Vgf *_TrecePackThenEncodeAsLoad) _FrecePackThenEncodeAsLoad__1400201y__decode_and_check_and_repack(___VbyteIn *[]byte, ___VbyteOut *[]byte) {
	(*___VbyteOut) = []byte{}
	//_FpfN(" 381917 23 : gfCHbyteI :[%x]", __VbyteIn)

	__VunRece := _TudpNodeDataRece{}
	__Verr4 := _FdecGob___(" 638196 02 ", *__VbyteIn, &__VunRece)
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
	__VunRece.
		_FdataPack__decode_from_udpNodeDataRece(&__Vdecode) // _TdecodeX
	if __Vdecode.Type != LoadT__loginS01genReplyTokenA {
		_CFpfN(" 638196 0r Ti:%d decode error {%v} error , get type %d  , want type %d ",
			__VtraceIntDE, __Verr2, __Vdecode.Type, LoadT__loginS01genReplyTokenA)
		return
	}

	__Venc := _Tencode{ // _TencodeX
	}

	//(*___VbyteOut), __Verr := _FencGob__only(___V)

}
