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

// _TtcpBufferArrX _TtcpNodeDataRece
func (___Vtbm *_TtcpBufMachine) _Finsert_local2remote_buf(___VtnRece *_TtcpNodeDataRece) bool {
	if nil == ___VtnRece || 16 != len(___VtnRece.TnrId128) || 0 == len(___VtnRece.TnrBuf) {
		_CFpfN(" 381812 01 : rece _TtcpNodeDataRece error {%s} ", ___VtnRece.String())
		return false
	}

	if ___Vtbm.tbmBufArr.tbaCntFree <= 0 {
		_CFpfN(" 381812 02 : no free Buf socket avaiable {%s} ", ___Vtbm.tbmBufArr.String())
		return false
	}

	//var __Vk16 [16]byte = ___VtnRece . TnrId128[:]

	return true
}
