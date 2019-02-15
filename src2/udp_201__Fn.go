package main

import (
    //"encoding/json"
    "flag"
    //"fmt"
    //"log"
    //"net"
)

    //_VuserIpList map[string]string
    //_VuserIpList = map[string]string{}

var (
    _VserviceUdpCn  _TserviceUDP
    _VserviceUdpDn  _TserviceUDP
    _VserviceUdpSn  _TserviceUDP

    _VserviceTcpMo  _TserviceTCP

    _Cexit       chan string
    _Clog        chan string
)

func init() {
    _VserviceUdpCn.name    = "servicePortForCn"
    _VserviceUdpDn.name    = "servicePortForDn"
    _VserviceUdpSn.name    = "servicePortForSn"

    _VserviceTcpMo.name    = "servicePortDebugLog"
    _VserviceTcpMo.port    = "127.0.0.1:56789"

    flag.StringVar(&_VserviceUdpCn.port, "c", ":5353",  _VserviceUdpCn.name )
    flag.StringVar(&_VserviceUdpDn.port, "d", ":32001", _VserviceUdpDn.name )
    flag.StringVar(&_VserviceUdpSn.port, "s", ":32003", _VserviceUdpSn.name )

    flag.Parse()

    // _FdebugPrintTest()
    _VprojectName = "Fn"
    _Fbase_101__get_self_md5_sha()
    _FPargs()

    _Cexit          = make (chan string, 3   )
    _Clog           = make (chan string, 100 )
}

func main() {

    _FtryListenToUDP01( &_VserviceUdpCn )
    _FtryListenToUDP01( &_VserviceUdpDn )
    _FtryListenToUDP01( &_VserviceUdpSn )

    go _FhandleFnWaitForClientCn( &_VserviceUdpCn , _Cexit , _Clog )
    go _FhandleFnWaitForClientCn( &_VserviceUdpDn , _Cexit , _Clog )
    go _FhandleFnWaitForClientCn( &_VserviceUdpSn , _Cexit , _Clog )

    _Fex( " the reason exit : " + <-_Cexit , nil )
} // main

