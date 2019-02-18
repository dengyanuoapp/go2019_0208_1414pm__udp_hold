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
    Cout01              *chan []byte

    FcallbackM          func( *_TfilterDelay)    // _FcallbackFilterDelay_main_swap
    FcallbackF          func( *_TfilterDelay)    // _FcallbackFilterDelay_filter

    Cexit               *chan string
    Clog                *chan string
} // _TfilterDelay 

//    _FhandleWaitForClientMsgUdpTop
func ( ___Vf *_TfilterDelay ) _FfilterDelayGen01_top() {

    if ( 1 > ___Vf.sleepGap ) { _Fex( " 811818 : error sleep gap " , nil ) }

    //CuByteIn01  chan []byte
    //CuByteOut01 *chan []byte

    for {
        _Fsleep_10sX( ___Vf.sleepGap )
        ___Vf . _FfilterDelayGen01_loop()
    }
} // _FfilterDelayGen01_top

func ( ___Vf *_TfilterDelay ) _FfilterDelayGen01_loop (){
    _FpfN( " 311191 : filter main " )
    if ( nil != ___Vf.FcallbackM ) {
        ___Vf.FcallbackM ( ___Vf ) // _FcallbackFilterDelay_main_swap
    }
} // _FfilterDelayGen01_loop
