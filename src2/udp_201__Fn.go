package main

import (
    //"encoding/json"
    "flag"
    //"fmt"
    "log"
    "net"
)

var (
    _VuserIpList map[string]string
    _Vservice    string
)

func init() {
    flag.StringVar(&_Vservice, "p", ":59999", "set the server listen port")
    flag.Parse()
}

func main() {

    // _FdebugPrintTest() 

    _VprojectName = "Fn"
    _P = true
    _Fbase_101__get_self_md5_sha()

    flag.PrintDefaults()

    _VuserIpList = map[string]string{}
    __VudpAddr, __Verr := net.ResolveUDPAddr("udp4", _Vservice)
    if __Verr != nil {
        log.Fatal(__Verr)
    }

    __Vconn, __Verr := net.ListenUDP("udp", __VudpAddr)
    if __Verr != nil {
        log.Fatal(__Verr)
    }

    for {
        _FhandleFnClient(__Vconn)
    }
} // main

func _FhandleFnClient(___Vconn *net.UDPConn) {
    var __Vbuf [2048]byte

    //__Vn, addr, __Verr := ___Vconn.ReadFromUDP(__Vbuf[0:])
    _, _, __Verr := ___Vconn.ReadFromUDP(__Vbuf[0:])

    if __Verr != nil {
        return
    }

} // _FhandleFnClient
