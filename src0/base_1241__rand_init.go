package main

import (
	//"encoding/binary"
	//"fmt"
	"math/rand"
)

func _Fbase_107__rand_init() {
	// func rand.Seed(seed int64)
	//rand.Seed(time.Now().Unix())
	// 	rand.Seed((int64)(_Vself.startTimEsha.A1 ^
	// 		_Vself.startTimEsha.A2 ^
	// 		_Vself.startTimEsha.A3 ^
	// 		_Vself.startTimEsha.A4 ^
	// 		_Vself.progSha.A1 ^
	// 		_Vself.progSha.A2 ^
	// 		_Vself.progSha.A3 ^
	// 		_Vself.progSha.A4 ^
	// 		binary.BigEndian.Uint64((_Vself.progMd5.b128[0:8]))))
	rand.Seed(_FgenRand_int64__())
} // _Fbase_107__rand_init
