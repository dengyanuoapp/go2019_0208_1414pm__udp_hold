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

    _VfilterCn2dn   _TfilterDelay

    _VserviceTcpMo  _TserviceTCP

    _Cexit       chan string
    _Clog        chan string
)

func init() {

    _Frand_init()
    _VprojectName = "Fn"
    _Fbase_101__get_self_md5_sha()
    _FPargs()

    _VserviceTcpMo  = _TserviceTCP {
        name        : "servicePortDebugLog",
        hostPortStr : "127.0.0.1:56781",
        TcallbackS  : _FcallbackForDebugLog_service_dataChan,
        TcallbackR  : _FcallbackForDebugLog_accept_dataReceive,
        TcallbackC  : _FcallbackForDebugLog_accept_dataChan,

        Cexit       : &_Cexit,
        Clog        : &_Clog,
        cAmount     : 10,
    }
    _VserviceUdpCn = _TserviceUDP  {
        name        : "servicePortForCn",
        UcallbackR  : _FuserCallback_dataRece_Cn,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }

    _VserviceUdpDn = _TserviceUDP  {
        name        : "servicePortForDn",
        UcallbackR  : _FuserCallback_dataRece_Dn,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }

    _VserviceUdpSn = _TserviceUDP  {
        name        : "servicePortForSn",
        UcallbackR  : _FuserCallback_dataRece_Sn,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }

    flag.StringVar(&_VserviceUdpCn.hostPortStr, "c", ":5353",  _VserviceUdpCn.name )
    flag.StringVar(&_VserviceUdpDn.hostPortStr, "d", ":32001", _VserviceUdpDn.name )
    flag.StringVar(&_VserviceUdpSn.hostPortStr, "s", ":32003", _VserviceUdpSn.name )

    flag.Parse()

    // _FdebugPrintTest()

    _Cexit          = make (chan string, 3   )
    _Clog           = make (chan string, 100 )
}

func main() {

    // ------------------- tcp for debug monitor log --- begin
    //_VserviceTcpMo . _FtryListenToTCP01( )
    // _TserviceTCP 

    // _FtcpAccept01
    // _FhandleTcp_accept_dataReceiveMsg01
    go _VserviceTcpMo . _Fhandle_udpListen_Tcp__main_top( )
    // ------------------- tcp for debug monitor log --- end

    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- begin

    // _TserviceUDP
    //_VserviceUdpCn . _FtryListenToUDP01( )
    //_VserviceUdpDn . _FtryListenToUDP01( )
    //_VserviceUdpSn . _FtryListenToUDP01( )

    // _TserviceUDP
    go _VserviceUdpCn . _Fhandle_udpListen_Udp__read_main_top( )
    go _VserviceUdpDn . _Fhandle_udpListen_Udp__read_main_top( )
    go _VserviceUdpSn . _Fhandle_udpListen_Udp__read_main_top( )
    // _FnotNullRunUdp01

    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- end
    // ------------------- filter between workers --------- begin
    //    _VfilterCn2dn.sleepGap      = 1
    //    _VfilterCn2dn.udpIn         = &_VserviceUdpCn
    //    _VfilterCn2dn.udpOut        = &_VserviceUdpDn
    _VfilterCn2dn = _TfilterDelay {
        sleepGap      : 1,
        udpIn         : &_VserviceUdpCn,
        udpOut        : &_VserviceUdpDn,
        FcallbackM    : _FcallbackFilterDelay_main_swap,
        FcallbackF    : _FcallbackFilterDelay_filter,
    }

    go _VfilterCn2dn . _FfilterDelayGen01_main_top()
    // ------------------- filter between workers --------- end

    _Fex( " the reason exit : " + <-_Cexit , nil )
} // main

