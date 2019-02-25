// _TuExtMRead 
// _TserviceUDP 
package main

import (
    //"net"
    "time"
    "os"
    "strconv"
)


var (
    _VuExtTimer_Dn      _TuExtTimer = _TuExtTimer {
        name            : "timer_Dn",
        timgGap1        : 25 * time.Second,          // every try must more than 2 gap(10*2==20S) 
        c641            : _FnPasswd,
    }
)

func _FuserCallback_u03TM__timer_Dn ( ___Vsvr *_TserviceUDP )  {
    _VuExtTimer_Dn . idx ++

    _FpfN( " 839111 : %d : trying to Connect to Fn using key 0x%x " , _VuExtTimer_Dn . idx , _VuExtTimer_Dn . c641 )

    // func strconv.ParseUint(s string, base int, bitSize int) (uint64, error)
    //__Vu64 , __Verr := strconv.ParseUint( os.Getenv("idhex") , 16, 64 ) // : invalid syntax : when 0x prefix
    __Vu64 , __Verr := strconv.ParseUint( os.Getenv("idhex") , 0, 0 )
    if ( __Verr == nil ) {
        _FpfN( " 839112 : var idhex : %x , %s , %x" , os.Getenv("idhex") , os.Getenv("idhex") , __Vu64 )
    } else {
        _FpfN( " 839113 : var idhex : %x , %s , convert to u64 error : %v" , os.Getenv("idhex") , os.Getenv("idhex") , __Verr )
    }

    if ( _VuExtTimer_Dn . enabled ) {
        _Fsleep_1s() ;
    } else {
        _Fsleep_1s() ;
    }
} // _FuserCallback_u03TM__timer_Dn 

