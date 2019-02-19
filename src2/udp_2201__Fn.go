package main

import (
    "flag"
)

var (
    _VserviceUdpFn  _TserviceUDP
    _VserviceUdpFp  _TserviceUDP
    _VserviceUdpFD  _TserviceUDP
    _VserviceUdpFS  _TserviceUDP

    _VfilterFn2dn   _TfilterDelay

    _VserviceTcpMf  _TserviceTCP

    _Cexit       chan string
    _Clog        chan string
)

func init() {

    _Frand_init()
    _VprojectName = "Fn"
    _Fbase_101__get_self_md5_sha()
    _FPargs()

    _VserviceTcpMf  = _TserviceTCP {
        name                    : "servicePortDebugLog",
        hostPortStr             : "127.0.0.1:56781",
        TcallbackSvrDataChan    : _FuserCallback__service_dataChan__Log_Fn,
        TcallbackAccDataRece    : _FuserCallback__Accept_dataReceive__Log_Fn,
        TcallbackAccDataChan    : _FuserCallback__accept_dataChan__Log_Fn,
        // _FhandleTcp__accept_dataChan__main_top

        Cexit                   : &_Cexit,
        Clog                    : &_Clog,
        cAmount                 : 10,
    }

    _VserviceUdpFn = _TserviceUDP  {
        name                    : "servicePortForCn",
        UcallbackR              : _FuserCallback_dataRece_Cn,
        Cexit                   : &_Cexit,
        Clog                    : &_Clog,
    }
    _VserviceUdpFp = _TserviceUDP  {
        name                    : "servicePortForCp",
        //UcallbackR             : _FuserCallback_dataRece_Cp,
        Cexit                   : &_Cexit,
        Clog                    : &_Clog,
    }

    _VserviceUdpFD = _TserviceUDP  {
        name                    : "servicePortForCD",
        UcallbackR              : _FuserCallback_dataRece_Dn__main_top,
        UcallbackC              : _FuserCallback_chanIn_Dn__main_top,
        Cexit                   : &_Cexit,
        Clog                    : &_Clog,
    }

    _VserviceUdpFS = _TserviceUDP  {
        name                    : "servicePortForCS",
        UcallbackR              : _FuserCallback_dataRece_Sn,
        Cexit                   : &_Cexit,
        Clog                    : &_Clog,
    }

    flag.StringVar(&_VserviceUdpFn.hostPortStr, "cn", ":5353",  _VserviceUdpFn.name )
    flag.StringVar(&_VserviceUdpFp.hostPortStr, "cp", ":32001", _VserviceUdpFp.name )
    flag.StringVar(&_VserviceUdpFD.hostPortStr, "cd", ":32002", _VserviceUdpFD.name )
    flag.StringVar(&_VserviceUdpFS.hostPortStr, "cs", ":32003", _VserviceUdpFS.name )

    flag.Parse()

    // _FdebugPrintTest()

    _Cexit          = make (chan string, 3   )
    _Clog           = make (chan string, 100 )
}

func main() {

    // ------------------- tcp for debug monitor log --- begin
    // _FtcpAccept01
    // _FhandleTcp_accept_dataReceiveMsg01
    go _VserviceTcpMf . _Fhandle_udpListen_Tcp__main_top( )
    // ------------------- tcp for debug monitor log --- end

    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- begin
    // _TserviceUDP
    go _VserviceUdpFn . _Fhandle_udpListen_Udp__read_main_top( )
    go _VserviceUdpFp . _Fhandle_udpListen_Udp__read_main_top( )
    go _VserviceUdpFD . _Fhandle_udpListen_Udp__read_main_top( )
    go _VserviceUdpFS . _Fhandle_udpListen_Udp__read_main_top( )
    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- end

    // ------------------- filter between workers --------- begin
    _VfilterFn2dn = _TfilterDelay {
        sleepGap                : 1,
        udpIn                   : &_VserviceUdpFn,
        udpOut                  : &_VserviceUdpFD,
        FcallbackMainDelayGen   : _FuserCallback__FilterDelay__main_swap_signal_gen__Fn,
        FcallbackFilterChan     : _FuserCallback__FilterDelay__chan_filter__Fn,
    }

    go _VfilterFn2dn . _FfilterDelayGen01_main_top()
    // ------------------- filter between workers --------- end

    _Fex( " the reason exit : " + <-_Cexit , nil )
} // main

