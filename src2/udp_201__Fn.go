package main

import (
    //"encoding/json"
    "flag"
    //"fmt"
    //"log"
    "net"
)

var (
    _VuserIpList map[string]string
    _VserviceCn  string
    _VserviceDn  string
    _VserviceSn  string
)

func init() {

    flag.StringVar(&_VserviceCn, "p", ":5353", "set the server listen port")
    flag.StringVar(&_VserviceDn, "d", ":32001", "set the server listen port")
    flag.StringVar(&_VserviceSn, "s", ":32003", "set the server listen port")
    flag.Parse()

    // _FdebugPrintTest() 
    _VprojectName = "Fn"
    _Fbase_101__get_self_md5_sha()
    _FPargs()

}

func main() {

    _VuserIpList = map[string]string{}
    __VudpAddr, __Verr := net.ResolveUDPAddr("udp4", _VserviceCn)
    if __Verr != nil {
        _Fex( "err13811" , __Verr)
    }

    __Vconn, __Verr := net.ListenUDP("udp", __VudpAddr)
    if __Verr != nil {
        _Fex( "err13811" , __Verr)
    }

    for {
        _FhandleFnClient(__Vconn)
    }
} // main

// https://golang.org/pkg/net/#UDPConn.ReadFromUDP
// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
func _FhandleFnClient(___Vconn *net.UDPConn) {
    //var __Vbuf [2048]byte
    //__Vlen, __Vaddr, __Verr := ___Vconn.ReadFromUDP(__Vbuf[0:])
    //var __Vbuf []byte
    __Vbuf := make( []byte , 2048 )
    __Vlen, __Vaddr, __Verr := ___Vconn.ReadFromUDP(__Vbuf)
    //__Vlen, __Vaddr, __Verr := ___Vconn.ReadFromUDP(__Vbuf[0:])
    //_, __Vaddr, __Verr := ___Vconn.ReadFromUDP(__Vbuf[0:])

    if __Verr != nil {
        return
    }
    if __Vaddr == nil {
        _Fex( " 183811 : why ___Vconn.ReadFromUDP addr error ?" , nil )
    }

    _Fpd( __Vlen , &__Vbuf )
    _Ppn( __Vlen , __Vbuf )

} // _FhandleFnClient
