package main

// _TtcpBufCellX
// tbmCHtcpReceBI chan []byte // _TtcpNodE tnCHtcpReceBLO *chan []byte // byte of _TtcpNodeDataRece
func (___Vtbm *_TtcpBufMachine) _FtcpBufMachine__1500201y1__chan_rece__Local2Remote(___VbIN *[]byte) {

	__VtnRece := _TtcpNodeDataRece{}
	// _FdecGob___(___VeMsg string, ___VbyteIn []byte, ___VoutObjLp interface{}) error {
	__Verr2 := _FdecGob___(" 381811 01 ", *___VbIN, &__VtnRece)
	_FerrExit(" 381811 02 : ", __Verr2)

	_CFpfN(" 381811 03 : tcpBuf rece from local {%s} ", __VtnRece.String()) // _TtcpNodeDataRece
	if false == ___Vtbm._Finsert_local2remote_buf(&__VtnRece) {
		return
	}
}

func (___Vtbm *_TtcpBufMachine) _Finsert_local2remote_buf(___VtnRece *_TtcpNodeDataRece) bool {
	return true
}
