// _TuExtMRead 
// _TserviceUDP 
package main

import (
    //"net"
    "time"
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

    if ( _VuExtTimer_Dn . enabled ) {
        _Fsleep_1s() ;
    } else {
        _Fsleep_1s() ;
    }
} // _FuserCallback_u03TM__timer_Dn 

