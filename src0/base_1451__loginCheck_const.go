package main

import (
	"sync"
)

/*
the Dn2Fn , step 01 , "Dn" gen a tokenL(A) and send to "Fn" , rec in tokenA(S) , sending : tokenL(A),-
the Fn2Dn , step 02 , "Fn" gen a tokenR(B) and send to "Dn" , rec in CmdMap(C) , sending : tokenL(B),tokenR(A)
the Dn2Fn , step 03 , "Dn" check CmdMap(S) tokenR == A and delay less than 10s, then accept it : clear DataMap, gen new DataMap[ID128(Fn)]
the Fn2Dn , step 04 , "Fn" check CmdMap(C) tokenR == B and delay less than 10s, then accept it : clear DataMap, gen new DataMap[ID128(Dn)]

Dn 01 [ no check ]                             MeIdC,MeSeqC -----,------ tokenLc,------- ( gen MeIdC MeSeqC tokenLc ) reC - send 1      : no data
Fn 02 [ check byte len only : 0,0,16,16,16,0]  MeIdS,MeSeqS MeIdC,MeSeqC tokenLs,tokenLc ( gen MeIdS MeSeqS tokenLs ) reC 1 send 5-10   : no data
Dn 03 [ check time,MeIdC,MeSeqC,tokenLc]       MeIdS,MeSeqS MeIdC,MeSeqC tokenLs,tokenLc ( gen MeIdS MeSeqS tokenLs ) reC - send 1*rece : reset data
Fn 04 [ check key,time,MeIdC,MeSeqC,tokenLc,MeIdS,MeSeqS,tokenLs]     -,- -,- -,-        ( gen nothing              ) reC - send 0      : reset data

note :
1. setp 01 precess in : _FdataPack__101__udpConnPort

2. CmdMap(S) and CmdMap(C) is the same, only difference working in Client and Server

*/

// _TloginReq
//type _TconnInfo struct {
//	ciId128 []byte
//}

type _TcmdMap struct {
	M   map[[16]byte]_Tdecode // key by 16byte : MeIdx128
	mux sync.Mutex
}

func (___Vcm *_TcmdMap) String() string {
	if 0 == len((*___Vcm).M) {
		return ""
	}

	var __VstrO string

	__Vi := 0
	for ___Vk := range (*___Vcm).M {
		__VstrO += _Pspf("%x ", ___Vk[:5])
		__Vi++
	}
	return __VstrO
}

//type _TdataMap struct {
//	M   map[[16]byte]_Tdecode // key by 16byte : MeIdx128
//	mux sync.Mutex
//}

type _TloginCheck struct {
	ulDecodeI           chan _Tdecode
	ulCHconnPortI       chan _TudpConnPort      // all data need to be sent by nodeS send here , then  will distribute to one of node
	ulCHSendLO          *chan _TudpNodeDataSend // ugCHSendI
	ulCHdataMachineIdLO *chan _TdataMachinEid
	ulCB900101init      func(*_TloginCheck) // if nil , use the default init procedure
	ulCB900201stCheck   func(*_TloginCheck) // if nil , use the default receive
	ulCmd               _TcmdMap
	//ulData              _TdataMap
	ulTokenA  []byte
	ulGenTime int
}

func String5(___Vb *[]byte) string {
	if (nil == ___Vb) || (0 == len(*___Vb)) {
		return ""
	}
	if len(*___Vb) < 5 {
		return _Pspf("%x", *___Vb)
	}
	return _Pspf("%x", (*___Vb)[:5])
}

func (___Vlc *_TloginCheck) String() string {
	return _Pspf(
		"loginChck-cmdMap:{%s} key:%x,%d",
		___Vlc.ulCmd.String(),
		String5(&___Vlc.ulTokenA),
		___Vlc.ulGenTime)
}

func _FcheckDecodeType(___Vdecode *_Tdecode, ___VwantType byte) bool { // match --> return false , others -> return true
	if nil == ___Vdecode {
		return true
	}

	if ___VwantType != ___Vdecode.Type {
		return true
	}
	return false
}
