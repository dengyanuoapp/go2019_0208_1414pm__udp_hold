// _TacceptTCP
package main

import (
    //"encoding/json"
    //"flag"
    //"fmt"
    //"log"
    //"net"
    //"sync"
)

type _TfilterCn2dn struct {

    sleepGap            int
    udpIn               *_TserviceUDP
    udpOut              *_TserviceUDP

    Cswap01             chan string
    Cin01               *chan []byte
    Cout01              *chan []byte

    FcallbackM          func( *_TfilterCn2dn)    // _FcallbackFilterDelay_main_swap
    FcallbackF          func( *_TfilterCn2dn)    // _FcallbackFilterDelay_filter

    Cexit               *chan string
    Clog                *chan string
} // _TfilterCn2dn 

//    _FhandleWaitForClientMsgUdpTop
func ( ___Vf *_TfilterCn2dn ) _FfilterDelayGen01_top() {

    if ( 1 > ___Vf.sleepGap ) { _Fex( " 811818 : error sleep gap " , nil ) }

    for {
        _Fsleep_10sX( ___Vf.sleepGap )
        ___Vf . _FfilterDelayGen01_loop()
    }
} // _FfilterDelayGen01_top

func ( ___Vf *_TfilterCn2dn ) _FfilterDelayGen01_loop (){
} // _FfilterDelayGen01_loop
