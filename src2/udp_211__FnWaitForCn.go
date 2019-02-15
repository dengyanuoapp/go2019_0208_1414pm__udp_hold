package main

import (
    //"encoding/json"
    //"flag"
    //"fmt"
    //"log"
    //"net"
)


// https://golang.org/pkg/net/#UDPConn.ReadFromUDP
// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
func _FhandleFnWaitForClientCn(___VserviceUDP *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
// (*___VserviceUDP).udpConn 
    //var __Vbuf [2048]byte             // array : with specified len
    __Vbuf := make( []byte , 2048 )   // silice : with var len

    __Vlen, __Vaddr, __Verr := (*___VserviceUDP).udpConn.ReadFromUDP(__Vbuf)
    //_, __Vaddr, __Verr := (*___VserviceUDP).udpConn.ReadFromUDP(__Vbuf[0:])

    if __Verr != nil {
        return
    }
    if __Vaddr == nil {
        _Fex( " 183811 : why ___Vconn.ReadFromUDP addr error ?" , nil )
    }

    //_FpdN( __Vlen , &__Vbuf )

    _Fpf( "|%s|" , __Vaddr )
    _PpdN( __Vlen , &__Vbuf )

} // _FhandleFnWaitForClientCn
