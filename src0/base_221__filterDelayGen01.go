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

    FcallbackM          func( *_TfilterCn2dn)    // _FcallbackFilterCn2dn_main_swap

    Cexit               *chan string
    Clog                *chan string
} // _TfilterCn2dn 

//    _FhandleWaitForClientMsgUdpTop
func _FfilterCn2dn01( ___VfC2D *_TfilterCn2dn ) {

    if ( 1 > ___VfC2D.sleepGap ) { _Fex( " 811818 : error sleep gap " , nil ) }

    for {
        _Fsleep_10sX( ___VfC2D.sleepGap )
        _FfilterCn2dn01_loop( ___VfC2D )
    }
} // _FfilterCn2dn01

func _FfilterCn2dn01_loop( ___VfC2D *_TfilterCn2dn ) {
} // _FfilterCn2dn01_loop
