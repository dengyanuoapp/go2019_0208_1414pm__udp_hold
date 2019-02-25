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
        pw2             : _Taes { "password__used_to_connect_to_Fn" , []byte(_FnPasswd) },
    }
)

func _FuserCallback_u03TM__timer_Dn ( ___Vsvr *_TserviceUDP )  {
    _VuExtTimer_Dn . idx ++

    _FpfN( " 839111 : %d : trying to Connect to Fn using key (%s) \n                 %d : <%0x> " , 
    _VuExtTimer_Dn . idx ,
    _VuExtTimer_Dn.pw2.name ,
    len(_VuExtTimer_Dn.pw2.key)  ,
    _VuExtTimer_Dn.pw2.key  )

    // func strconv.ParseUint(s string, base int, bitSize int) (uint64, error)
    //__Vu64 , __Verr := strconv.ParseUint( os.Getenv("id128") , 16, 64 ) // : invalid syntax : when 0x prefix
    __Vu64 , __Verr := strconv.ParseUint( os.Getenv("id128") , 0, 0 )
    if ( __Verr == nil ) {
        _FpfN( " 839113 : var id128 : %x , %s , %x" , os.Getenv("id128") , os.Getenv("id128") , __Vu64 )
    } else {
        _FpfN( " 839115 : var id128 : %x , %s , convert to u64 error : %v" , os.Getenv("id128") , os.Getenv("id128") , __Verr )
    }

    if ( _VuExtTimer_Dn . enabled ) {
        _Fsleep_100s() ;
    } else {
        _Fsleep_100s() ;
    }
} // _FuserCallback_u03TM__timer_Dn 

