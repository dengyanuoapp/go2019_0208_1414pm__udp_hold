package main

import (
	"strconv"
)

var ()

// func strconv.ParseUint(s string, base int, bitSize int) (uint64, error) {
func _Fhexstr2uint64(___Vstr string) uint64 {
	var __Vu64 uint64
	var __Verr error
	__Vu64, __Verr = strconv.ParseUint(___Vstr, 16, 64) // try to read as hex without 0x prefix
	if nil != __Verr {
		//_FpfN("893893891 err : %v" , __Verr )
		__Vu64, __Verr = strconv.ParseUint(___Vstr, 0, 64) // try to read as hex with 0x prefix
	}
	if nil == __Verr {
		//_FpfN("893893893 ok : %x" , __Vu64 )
		return __Vu64
	}
	_FpfN("893893895 err : %v", __Verr)
	return 0
} // _Fhexstr2uint64
