package main

import "sync"

func (___Vdmcc *_TdataMachinEconnectClient) Clear() {
	(*___Vdmcc) = _TdataMachinEconnectClient{}
	___Vdmcc.dccConnPortStrMap = make(map[string]byte)
}
func (___Vdmcc *_TdataMachinEconnectClient) String() string {
	__Vlen := len(___Vdmcc.dccConnPortArr)
	__Vs := _Spf("eid:{%s} (%d)[", ___Vdmcc.dccID.String(), __Vlen)
	for __Vi := 0; __Vi < __Vlen; __Vi++ {
		if 0 == __Vi {
			__Vs += _Spf("%s", ___Vdmcc.dccConnPortArr[__Vi].String())
		} else {
			__Vs += _Spf(" | %s", ___Vdmcc.dccConnPortArr[__Vi].String())
		}
	}
	__Vs += _Spf("] amount:%d lastRead:%d lastPr:%d cntLast:%d cntNow:%d lastSendIdx:%d ",
		___Vdmcc.dccConnPortAmount,
		___Vdmcc.dccLastReceTime,
		___Vdmcc.dccLastPrTime,
		___Vdmcc.dccLockCntLast,
		___Vdmcc.dccLockCntLast,
		___Vdmcc.dccLastSendIdx,
	)
	return __Vs
}

type _TdataMachinEconnectClient struct {
	dccID             _TdataMachinEid
	dccConnPortArr    []_TudpConnPort
	dccConnPortStrMap map[string]byte
	dccConnPortAmount int
	dccLastReceTime   int
	dccLastPrTime     int
	dccLockCntLast    int
	dccLockCntNow     int
	dccLastSendIdx    int
}
type _TdataMachinEconnectSt struct {
	dcsMidx        map[[16]byte]int
	dcsMm          [_VallowClientMax]_TdataMachinEconnectClient
	dcsMusedAmount int
	dcsMfreeAmount int
	dcsMlastInsIdx int // the last insert place , the next insert can be start here
	dcsMux         sync.Mutex
}

func (___Vdmes *_TdataMachinEconnectSt) String() string {
	__Vs := _Spf("idxArr(%d)[", len(___Vdmes.dcsMidx))
	for __Vk2, __Vv2 := range ___Vdmes.dcsMidx {
		__Vs += _Spf(" %x,%d", __Vk2, __Vv2)
	}

	__Vs += _Spf("] used:%d free:%d lastIdx:%d (",
		___Vdmes.dcsMusedAmount,
		___Vdmes.dcsMfreeAmount,
		___Vdmes.dcsMlastInsIdx)

	for _, __Vv2 := range ___Vdmes.dcsMidx {
		__Vs += _Spf(" (%d)<%s>", __Vv2, ___Vdmes.dcsMm[__Vv2].String())
	}

	__Vs += ")"
	return __Vs
}
