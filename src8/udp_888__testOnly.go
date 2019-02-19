package main

import (
    "net"
)

type _TcnTdn struct {
    Version     int
    Idx         int
    Port        int
    IP          net.IP
    IpStr       string
}

func main() {

        __Vcn2dn := _TcnTdn {}
        __Vcn2dn.   Version = 1
        __Vcn2dn.   IP      = net.IPv4(127, 0, 1, 1)
        __Vcn2dn.   Idx     = 5
        __Vcn2dn.   Port    = 4301
        __Vcn2dn.   IpStr   = "127.0.0.1"
        _Ppf( "\n 0738184 : origin msg is: %v\n" ,  __Vcn2dn )
        __Vb2 := _FencGob( __Vcn2dn ) 
        _Ppf( "\n 0738186 : _FencGob msg is: %d , %x\n" , len(__Vb2) ,  __Vb2 )
        var __Vt3 _TcnTdn
        _FdecGob( __Vb2 , &__Vt3)
        _Ppf( "\n 0738187 : _FdecGob msg is: %v\n" ,  __Vt3 )

} // main
