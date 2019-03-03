package main

import (
	"encoding/binary"
	"fmt"
	"math/rand"
)

func _Fbase_107__rand_init() {
	_FpfN(" 189301 : _Vself.Id128 %d : %0x ", len(_Vself.Id128.b128), _Vself.Id128.b128)

	// func rand.Seed(seed int64)
	//rand.Seed(time.Now().Unix())
	rand.Seed((int64)(_Vself.startTimEsha.A1 ^
		_Vself.startTimEsha.A2 ^
		_Vself.startTimEsha.A3 ^
		_Vself.startTimEsha.A4 ^
		_Vself.Id256.A1 ^
		_Vself.Id256.A2 ^
		_Vself.Id256.A3 ^
		_Vself.Id256.A4 ^
		binary.BigEndian.Uint64((_Vself.Id128.b128[0:8]))))

	// wall : wall[63] wall[62:30] wall[29:0] : 1bit,33bit, 30bit
	// if hasMonotonic == (1 << 63) bit
	// --> wall[63] --> 0 : no clock counter used : second since Jan 1 year 1 <-- ext : signed 64-bit wall seconds
	// --> wall[63] --> 1 : the clock counter used : second since Jan 1 year 1885 <-- wall[62:30] ,
	//                                               ext <- 64 bit clock-counter value since machine start
	//     func (t *Time) sec() int64 {
	//         if t.wall&hasMonotonic != 0 {
	//             return wallToInternal + int64(t.wall<<1>>(nsecShift+1))
	//         }
	//         return t.ext
	//     }
	//                                               ext <- 64 bit clock-counter value since machine start

	// func rand.Uint64() uint64
	__Vu64 := rand.Uint64
	__Vb := []byte(fmt.Sprintf("%x", __Vu64))
	_Vself.myKey._Fbase_1101__gen_shaT(&__Vb)
	//     _Vself.myKey , _      = sha256.Sum256( Sprintf( "%x" , __Vu64 ) )
	//     _self_rand1 , _     = strconv.ParseUint( _Vself.myKey[0:7]     , 16, 64 )
	//     _self_rand2 , _     = strconv.ParseUint( _Vself.myKey[8:15]    , 16, 64 )
	//     _self_rand3 , _     = strconv.ParseUint( _Vself.myKey[16:23]   , 16, 64 )
	//     _self_rand4 , _     = strconv.ParseUint( _Vself.myKey[24:31]   , 16, 64 )

} // _Fbase_107__rand_init
