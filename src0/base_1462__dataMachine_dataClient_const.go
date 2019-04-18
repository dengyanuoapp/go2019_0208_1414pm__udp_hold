package main

import "sync"

func (___Vdmdc *_TdataMachinEdataClient) String() string {
	__Vlen := len(___Vdmdc.ddcConnPortArr)
	__Vs := _Spf("eid:{%s} (%d)[", ___Vdmdc.ddcID.String(), __Vlen)
	for __Vi := 0; __Vi < __Vlen; __Vi++ {
		if 0 == __Vi {
			__Vs += _Spf("%s", ___Vdmdc.ddcConnPortArr[__Vi].String())
		} else {
			__Vs += _Spf(" | %s", ___Vdmdc.ddcConnPortArr[__Vi].String())
		}
	}
	__Vs += _Spf("] amount:%d lastRead:%d lastSend:%d lastSendIdx:%d ",
		___Vdmdc.ddcConnPortAmount,
		___Vdmdc.ddcLastReceTime,
		___Vdmdc.ddcLastSendTime,
		___Vdmdc.ddcLastSendIdx,
	)
	return __Vs
}
func (___Vdmdc *_TdataMachinEdataClient) Clear() {
	(*___Vdmdc) = _TdataMachinEdataClient{}
	___Vdmdc.ddcConnPortStrMap = make(map[string]byte)
}

type _TdataMachinEdataClient struct {
	ddcID             _TdataMachinEid
	ddcConnPortArr    []_TudpConnPort
	ddcConnPortStrMap map[string]byte
	ddcConnPortAmount int
	ddcLastReceTime   int
	ddcLastSendTime   int
	ddcLastSendIdx    int
}
type _TdataMachinEdataSt struct {
	ddsMidx        map[[16]byte]int
	ddsMm          [_VallowClientMax]_TdataMachinEdataClient
	ddsMusedAmount int
	ddsMfreeAmount int
	ddsMlastInsIdx int // the last insert place , the next insert can be start here
	ddsMux         sync.Mutex
}
