package main

import (
    //"rand"
    //"time"
    "math/rand"
    //"crypto/sha256"
    //"strconv"
    "fmt"

    //"bytes"
    //"os"
    //"io/ioutil"
    //"crypto/md5"
    //"log"
    //"encoding/hex"
    //"time"
)

func _Fbase_107__rand_init() {
    // func rand.Seed(seed int64)
    //rand.Seed(time.Now().Unix())
    rand.Seed( (int64)( _self_startTimeSha.A1 ^
     _self_startTimeSha.A2 ^
     _self_startTimeSha.A3 ^
     _self_startTimeSha.A4 ^
    _self_sha.A1          ^
    _self_sha.A2          ^
    _self_sha.A3          ^
    _self_sha.A4          ^
    _self_id128          ))

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
    __Vu64  :=  rand.Uint64 ;
    _self_rand . _Fbase_1101__gen_shaT( []byte(fmt.Sprintf( "%x" , __Vu64 )) )
    //     _self_rand , _      = sha256.Sum256( Sprintf( "%x" , __Vu64 ) )
    //     _self_rand1 , _     = strconv.ParseUint( _self_rand[0:7]     , 16, 64 ) 
    //     _self_rand2 , _     = strconv.ParseUint( _self_rand[8:15]    , 16, 64 ) 
    //     _self_rand3 , _     = strconv.ParseUint( _self_rand[16:23]   , 16, 64 ) 
    //     _self_rand4 , _     = strconv.ParseUint( _self_rand[24:31]   , 16, 64 ) 

} // _Fbase_107__rand_init

