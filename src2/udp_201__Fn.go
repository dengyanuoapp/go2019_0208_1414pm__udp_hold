package main

import (
    //"encoding/json"
    "flag"
    //"rand"
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

    _Frand_init()

    _VserviceUdpCn.name             = "servicePortForCn"
    _VserviceUdpDn.name             = "servicePortForDn"
    _VserviceUdpSn.name             = "servicePortForSn"

    flag.StringVar(&_VserviceUdpCn.hostPortStr, "c", ":5353",  _VserviceUdpCn.name )
    flag.StringVar(&_VserviceUdpDn.hostPortStr, "d", ":32001", _VserviceUdpDn.name )
    flag.StringVar(&_VserviceUdpSn.hostPortStr, "s", ":32003", _VserviceUdpSn.name )

    flag.Parse()

    _VserviceTcpMo.name             = "servicePortDebugLog"
    _VserviceTcpMo.hostPortStr      = "127.0.0.1:56781"

    // _FdebugPrintTest()
    _VprojectName = "Fn"
    _Fbase_101__get_self_md5_sha()
    _FPargs()

    _Cexit          = make (chan string, 3   )
    _Clog           = make (chan string, 100 )
}

func main() {

    // ------------------- tcp for debug monitor log --- begin
    _VserviceTcpMo.Cexit        = &_Cexit
    _VserviceTcpMo.Clog         = &_Clog
    _VserviceTcpMo.cAmount      = 10
    _FtryListenToTCP01( &_VserviceTcpMo )
    // _TserviceTCP 

    // _FtcpAccept01
    // _FhandleTcpReceiveMsg01
    go _FhandleWaitForClientMsgTcpTop( &_VserviceTcpMo )
    // ------------------- tcp for debug monitor log --- end

    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- begin
    _VserviceUdpCn.callbackR  = _FcallbackInFnForCn

    _VserviceUdpCn.Cexit = &_Cexit
    _VserviceUdpDn.Cexit = &_Cexit
    _VserviceUdpSn.Cexit = &_Cexit
    _VserviceUdpCn.Clog  = &_Clog
    _VserviceUdpDn.Clog  = &_Clog
    _VserviceUdpSn.Clog  = &_Clog
    // _TserviceUDP
    _FtryListenToUDP01( &_VserviceUdpCn )
    _FtryListenToUDP01( &_VserviceUdpDn )
    _FtryListenToUDP01( &_VserviceUdpSn )

    // _TserviceUDP
    go _FhandleWaitForClientMsgUdpTop( &_VserviceUdpCn )
    go _FhandleWaitForClientMsgUdpTop( &_VserviceUdpDn )
    go _FhandleWaitForClientMsgUdpTop( &_VserviceUdpSn )
    // _FnotNullRunUdp01
    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- end

    _Fex( " the reason exit : " + <-_Cexit , nil )
} // main

