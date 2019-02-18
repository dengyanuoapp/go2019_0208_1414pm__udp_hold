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

type _TfilterDelay struct {

    sleepGap            int
    udpIn               *_TserviceUDP
    udpOut              *_TserviceUDP

    Cswap01             chan string
    Cin01               *chan []byte
    Cout01              chan []byte

    FcallbackM          func( *_TfilterDelay)    // _FcallbackFilterDelay_main_swap
    FcallbackF          func( *_TfilterDelay)    // _FcallbackFilterDelay_filter

    Cexit               *chan string
    Clog                *chan string
} // _TfilterDelay 

//    _FhandleWaitForClientMsgUdpTop
func ( ___Vf *_TfilterDelay ) _FfilterDelayGen01_top() {

    if ( 1 > ___Vf.sleepGap ) { _Fex( " 811818 : error sleep gap " , nil ) }

    ___Vf.  Cswap01             = make( chan string , 4  )
    ___Vf.  Cout01              = make( chan []byte , 5  )
    ___Vf.  Cin01               = &(___Vf.udpIn.CuByteIn01)

    go ___Vf. _FfilterDelayGen01_filter_top()

    _Fsleep_1s()
    ___Vf.udpOut.CuByteOut01    = &(___Vf.  Cout01)

    for {
        _Fsleep_10sX( ___Vf.sleepGap )
        ___Vf . _FfilterDelayGen01_loop()
    }
} // _FfilterDelayGen01_top

func ( ___Vf *_TfilterDelay ) _FfilterDelayGen01_loop (){
    _FpfN( " 311191 : filter main " )
    if ( nil != ___Vf.FcallbackM ) {
        _FpfN( " 311192 : filter main " )
        ___Vf.FcallbackM ( ___Vf ) // _FcallbackFilterDelay_main_swap
    }
} // _FfilterDelayGen01_loop

func ( ___Vf *_TfilterDelay ) _FfilterDelayGen01_filter_top (){
    _FpfN( " 311195 : filter main " )
    for {
        go ___Vf. _FfilterDelayGen01_filter_loop()
    }
} // _FfilterDelayGen01_filter_top

func ( ___Vf *_TfilterDelay ) _FfilterDelayGen01_filter_loop (){
    _FpfN( " 311197 : filter main " )
    if ( nil != ___Vf.FcallbackF ) {
        ___Vf.FcallbackF ( ___Vf ) // _FcallbackFilterDelay_filter
    }
} // _FfilterDelayGen01_filter_loop
