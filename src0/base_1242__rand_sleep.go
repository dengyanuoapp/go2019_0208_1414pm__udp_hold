package main

func _FsleepRand_12_to_14s() {
	__Vi := _FgenRand_uint32()
	__Vi %= 3  // gen 0,1,2
	__Vi += 12 // ok , 10 to 14s
	_Fsleep_1sX(int(__Vi))
} //

func _FsleepRand_10_to_15s() {
	__Vi := _FgenRand_uint32()
	__Vi %= 6  // gen 0,1,2,3,4,5
	__Vi += 10 // ok , 10 to 14s
	_Fsleep_1sX(int(__Vi))
} //
