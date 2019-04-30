package main

import "sync"

func (___Vdmdc *_TdataMachinEdataClient) String() string {
	__Vlen := len(___Vdmdc.ddcConnPortArr)
	__Vs := _Spf("%s eid:{%s} (%d)[", ___Vdmdc.ddcRole, ___Vdmdc.ddcID.String(), __Vlen)
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
	ddcRole           string
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

func (___Vdmds *_TdataMachinEdataSt) String() string {
	__Vs := _Spf("idxArr(%d)[", len(___Vdmds.ddsMidx))
	for __Vk2, __Vv2 := range ___Vdmds.ddsMidx {
		__Vs += _Spf(" %x,%d", __Vk2, __Vv2)
	}

	__Vs += _Spf("] used:%d free:%d lastIdx:%d (",
		___Vdmds.ddsMusedAmount,
		___Vdmds.ddsMfreeAmount,
		___Vdmds.ddsMlastInsIdx)

	for _, __Vv2 := range ___Vdmds.ddsMidx {
		__Vs += _Spf(" (%d)<%s>", __Vv2, ___Vdmds.ddsMm[__Vv2].String())
	}

	__Vs += ")"
	return __Vs
}
