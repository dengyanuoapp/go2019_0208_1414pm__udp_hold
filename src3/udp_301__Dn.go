package main

import (
    "flag"
)

var (
    _VserviceUdpDn  _TserviceUDP
    _VserviceUdpDp  _TserviceUDP
    _VserviceUdpDC  _TserviceUDP
    _VserviceUdpDS  _TserviceUDP

    _VfilterCn2dn   _TfilterDelay

    _VserviceTcpMd  _TserviceTCP

    _Cexit       chan string
    _Clog        chan string
)

func init() {

    _Frand_init()
    _VprojectName = "Fn"
    _Fbase_101__get_self_md5_sha()
    _FPargs()

    _VserviceTcpMd  = _TserviceTCP {
        name        : "servicePortDebugLog",
        hostPortStr : "127.0.0.1:56781",
        TcallbackS  : _FcallbackForDebugLog_service_dataChan,
        TcallbackR  : _FcallbackForDebugLog_accept_dataReceive,
        TcallbackC  : _FcallbackForDebugLog_accept_dataChan,

        Cexit       : &_Cexit,
        Clog        : &_Clog,
        cAmount     : 10,
    }
    _VserviceUdpDn = _TserviceUDP  {
        name        : "servicePortForCn",
        UcallbackR  : _FuserCallback_dataRece_Cn,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }
    _VserviceUdpDp = _TserviceUDP  {
        name        : "servicePortForCp",
        //UcallbackR  : _FuserCallback_dataRece_Cp,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }

    _VserviceUdpDC = _TserviceUDP  {
        name        : "servicePortForCD",
        UcallbackR  : _FuserCallback_dataRece_Dn__main_top,
        UcallbackC  : _FuserCallback_chanIn_Dn__main_top,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }

    _VserviceUdpDS = _TserviceUDP  {
        name        : "servicePortForCS",
        UcallbackR  : _FuserCallback_dataRece_Sn,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }

    flag.StringVar(&_VserviceUdpDn.hostPortStr, "cn", ":5353",  _VserviceUdpDn.name )
    flag.StringVar(&_VserviceUdpDp.hostPortStr, "cp", ":32001", _VserviceUdpDp.name )
    flag.StringVar(&_VserviceUdpDC.hostPortStr, "cd", ":32002", _VserviceUdpDC.name )
    flag.StringVar(&_VserviceUdpDS.hostPortStr, "cs", ":32003", _VserviceUdpDS.name )

    flag.Parse()

    // _FdebugPrintTest()

    _Cexit          = make (chan string, 3   )
    _Clog           = make (chan string, 100 )
}

func main() {

    // ------------------- tcp for debug monitor log --- begin
    // _FtcpAccept01
    // _FhandleTcp_accept_dataReceiveMsg01
    go _VserviceTcpMd . _Fhandle_udpListen_Tcp__main_top( )
    // ------------------- tcp for debug monitor log --- end

    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- begin
    // _TserviceUDP
    go _VserviceUdpDn . _Fhandle_udpListen_Udp__read_main_top( )
    go _VserviceUdpDp . _Fhandle_udpListen_Udp__read_main_top( )
    go _VserviceUdpDC . _Fhandle_udpListen_Udp__read_main_top( )
    go _VserviceUdpDS . _Fhandle_udpListen_Udp__read_main_top( )
    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- end

    // ------------------- filter between workers --------- begin
    _VfilterCn2dn = _TfilterDelay {
        sleepGap      : 1,
        udpIn         : &_VserviceUdpDn,
        udpOut        : &_VserviceUdpDC,
        FcallbackM    : _Fcallback_user_FilterDelay__main_swap_signal_gen,
        FcallbackF    : _Fcallback_user_FilterDelay__chan_filter,
    }

    go _VfilterCn2dn . _FfilterDelayGen01_main_top()
    // ------------------- filter between workers --------- end

    _Fex( " the reason exit : " + <-_Cexit , nil )
} // main

