// _TuExtMRead 
// _TserviceUDP 
package main

import (
    //"net"
    "time"
    "os"
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
    _FpfN( " 839112 : var pwhex : %x , %s" , os.Getenv("pwhex") , os.Getenv("pwhex") );

    if ( _VuExtTimer_Dn . enabled ) {
        _Fsleep_1s() ;
    } else {
        _Fsleep_1s() ;
    }
} // _FuserCallback_u03TM__timer_Dn 

