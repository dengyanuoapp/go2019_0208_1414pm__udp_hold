// _TuExtMRead 
// _TserviceUDP 
package main

import (
    //"net"
    "time"
    "sync"
)

type _TuExtMRead struct {
    name            string

} // _TuExtMRead 

type _TuExtChanI struct {
} // _TuExtChanI

// func time.Sleep(d Duration)
type _TuExtTimer struct {
    name            string
    idx             uint64
    enabled         bool
    timgGap1        time.Duration
    timgGap2        time.Duration
    c641            uint64
    c642            uint64
    mux             sync.Mutex
} // _TuExtTimer


var (
    _VuExtMR_Dn _TuExtMRead
)

func _FuserCallback_UdataMain_Dn ( ___Vsvr *_TserviceUDP ) {
} // _FuserCallback_UdataMain_Dn 

