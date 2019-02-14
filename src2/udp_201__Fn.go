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
}

func main() {

    // _FdebugPrintTest() 

    _VprojectName = "Fn"
    _P = true
    _Fbase_101__get_self_md5_sha()

    _P.PrintArgs()

    _VuserIpList = map[string]string{}
    __VudpAddr, __Verr := net.ResolveUDPAddr("udp4", _VserviceCn)
    if __Verr != nil {
        _P.X(__Verr)
    }

    __Vconn, __Verr := net.ListenUDP("udp", __VudpAddr)
    if __Verr != nil {
        _P.X(__Verr)
    }

    for {
        _FhandleFnClient(__Vconn)
    }
} // main

func _FhandleFnClient(___Vconn *net.UDPConn) {
    var __Vbuf [2048]byte

    __Vlen, __Vaddr, __Verr := ___Vconn.ReadFromUDP(__Vbuf[0:])

    if __Verr != nil {
        return
    }
    if __Vaddr == nil {
        _P.X( " 183811 : why ___Vconn.ReadFromUDP addr error ?" )
    }

    _P.dint( __Vlen , 2048, __Vbuf[0:] )
    _P.fint( "\n" );

} // _FhandleFnClient
