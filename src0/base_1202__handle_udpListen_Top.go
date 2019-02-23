// _TserviceUDP
package main

import (
)


// https://golang.org/pkg/net/#UDPConn.ReadFromUDP
// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
//func _Fhandle_udpListen_Udp_main_top(___VserviceUdp *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
func (___VserviceUdp *_TserviceUDP ) _Fhandle_u01x__udpListen_Udp__read_main_top (){

    ___VserviceUdp . _FtryListenToUDP01( )

    ___VserviceUdp.Vubuf         = make(      []byte , 2048 )   // silice : with var len
    ___VserviceUdp.CuIn01       = make( chan []byte , 5    )   // silice : with var len

    go ___VserviceUdp . _Fhandle_udpListen__chanIn_main_top()

    for ; ; {
        ___VserviceUdp . _Fhandle_u01y__udpListen_Udp__read_main_loop( )
    }
    (*___VserviceUdp.Cexit) <- "338191 Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _Fhandle_u01x__udpListen_Udp__read_main_top

func ( ___VserviceUdp *_TserviceUDP ) _Fhandle_u01y__udpListen_Udp__read_main_loop() {
    ___VserviceUdp.Vulen,
    ___VserviceUdp.VuremoteAddr,
    ___VserviceUdp.Vuerr =
    ___VserviceUdp.udpConn.ReadFromUDP(___VserviceUdp.Vubuf)

    _FerrExit( "err 338195 : udp reading : " , ___VserviceUdp.Vuerr )

    _FnullExit( " err 338197 : why ___Vconn.ReadFromUDP addr error ?" , ___VserviceUdp.VuremoteAddr )

    _FnotNullRunUdp01( ___VserviceUdp.UcallbackR , ___VserviceUdp )
} // _Fhandle_u01y__udpListen_Udp__read_main_loop





func (___VserviceUdp *_TserviceUDP ) _Fhandle_udpListen__chanIn_main_top (){

    for ; ; {
        _Fsleep_1s()
        ___VserviceUdp . _Fhandle_udpListen__chanIn_main_loop( )
    }

    (*___VserviceUdp.Cexit) <- "438191 Error : (" + ___VserviceUdp.hostPortStr + ")"
} // _Fhandle_udpListen__chanIn_main_top

func (___VserviceUdp *_TserviceUDP ) _Fhandle_udpListen__chanIn_main_loop (){

    //_FpfN( "338199 : udp rece Chan msg ." )
    _FnotNullRunUdp01( ___VserviceUdp.UcallbackC , ___VserviceUdp )

} // _Fhandle_udpListen__chanIn_main_loop

