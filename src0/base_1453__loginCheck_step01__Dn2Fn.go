package main

var __Vstep02_LoginLenArr []int = []int{0, 0, 16, 16, 16, 0}
var __Vstep03_LoginLenArr []int = []int{16, 16, 16, 16, 16, 16}
var __VmaxCmdPerid int = 30

func _FdeleteOld_cmdStack(___Vm *map[[16]byte]_Tdecode) {
	var __Vdel [][16]byte
	__Vnow := _FtimeInt()
	for __k, __v := range *___Vm {
		if __Vnow-__v.DEreceiveTime > __VmaxCmdPerid {
			__Vdel = append(__Vdel, __k)
		}
	}
	for _, __v := range __Vdel {
		delete(*___Vm, __v)
	}
}
