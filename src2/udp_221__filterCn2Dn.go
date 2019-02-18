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

    Cstart              chan string
    CreceiveMsg         chan []byte
    CchanMsg            chan []byte
    CreceiveErr         chan string

    Cexit               *chan string
    Clog                *chan string
} // _TfilterCn2dn 

//    _FhandleWaitForClientMsgUdpTop
func _FfilterCn2dn01( ___VfC2D *_TfilterCn2dn , ___VsC , ___VsD *_TserviceUDP ) {
    for {
        _Fsleep_1s()
        _FfilterCn2dn01_loop( ___VfC2D , ___VsC , ___VsD )
    }
} // _FfilterCn2dn01
func _FfilterCn2dn01_loop( ___VfC2D *_TfilterCn2dn , ___VsC , ___VsD *_TserviceUDP ) {
} // _FfilterCn2dn01_loop
