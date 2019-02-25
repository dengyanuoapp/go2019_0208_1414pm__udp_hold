package main

import (
    "strconv"
)

var (
)

// func strconv.ParseUint(s string, base int, bitSize int) (uint64, error) {
func _Fhexstr2uint64( ___Vstr string ) uint64 {
        __Vu64 , __Verr := strconv.ParseUint( ___Vstr , 16, 0 ) // try to read as hex without 0x prefix
        if ( nil != __Verr ) {
            __Vu64 , __Verr := strconv.ParseUint( ___Vstr , 0, 0 ) // try to read as hex with 0x prefix
        }
        if ( nil == __Verr ) {
            return __Vu64
        }
        return 0
} // _Fhexstr2uint64
