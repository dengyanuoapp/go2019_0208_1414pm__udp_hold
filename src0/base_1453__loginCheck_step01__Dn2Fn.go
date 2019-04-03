package main

/*
the Dn2Fn , step 01 , "Dn" gen a tokenL(A) and send to "Fn" , record in CmdMap(S) , sending : tokenL(A),-
the Fn2Dn , step 02 , "Fn" gen a tokenR(B) and send to "Dn" , record in CmdMap(C) , sending : tokenL(B),tokenR(A)
the Dn2Fn , step 03 , "Dn" check CmdMap(S) tokenR == A and delay less than 10s, then accept it : clear DataMap, gen new DataMap[ID128(Fn)]
the Fn2Dn , step 04 , "Fn" check CmdMap(C) tokenR == B and delay less than 10s, then accept it : clear DataMap, gen new DataMap[ID128(Dn)]

note :
1. setp 01 precess in :

2. CmdMap(S) and CmdMap(C) is the same, only difference working in Client and Server

*/

var __Vstep02_LoginLenArr []int = []int{0, 0, 16, 16, 16, 0}
var __Vstep03_LoginLenArr []int = []int{0, 0, 16, 16, 16, 16}
var __VmaxCmdPerid int = 100

func _FdeleteOld_cmdStack(___Vm *map[[16]byte]_Tdecode) {
	var __Vdel [][16]byte
	__Vnow := _FtimeInt()
	for __k, __v := range *___Vm {
		if __Vnow-__v.receiveTime > __VmaxCmdPerid {
			__Vdel = append(__Vdel, __k)
		}
	}
	for _, __v := range __Vdel {
		delete(*___Vm, __v)
	}
}
