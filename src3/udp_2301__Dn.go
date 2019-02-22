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
    _VprojectName = "Dn"
    _Fbase_101__get_self_md5_sha()
    _FPargs()

    _VserviceTcpMd  = _TserviceTCP {
        name                    : "TcpService__DebugLog__Md",
        hostPortStr             : "127.0.0.1:56782",
        //TcallbackSvrDataChan    : _FuserCallback__Log_service_dataChan__Dn,
        //TcallbackAccDataRece    : _FuserCallback__LogAccept_dataReceive__Dn,
        //TcallbackAccDataChan    : _FuserCallback__accept_dataChan__Log_Dn,

        Cexit       : &_Cexit,
        Clog        : &_Clog,
        cAmount     : 10,
    }
    // _FuserCallback_UdataRece_Dn 
    // _FuserCallback_UdataMain_Dn 
    _VserviceUdpDn = _TserviceUDP  {
        name        : "UdpService__Dn",
        uExt        : &_VuDataExt_Dn,
        //UcallbackR  : _FuserCallback_dataRece_Dn,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }
    _VserviceUdpDp = _TserviceUDP  {
        name        : "UdpService__Dp",
        //UcallbackR  : _FuserCallback_dataRece_Dp,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }

    _VserviceUdpDC = _TserviceUDP  {
        name        : "UdpService__DC",
        //UcallbackR  : _FuserCallback_dataRece__main_top_DC,
        //UcallbackC  : _FuserCallback_chanIn__main_top_DC,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }

    _VserviceUdpDS = _TserviceUDP  {
        name        : "UdpService__DS",
        //UcallbackR  : _FuserCallback_dataRece_DS,
        Cexit       : &_Cexit,
        Clog        : &_Clog,
    }

    flag.StringVar(&_VserviceUdpDn.hostPortStr, "cn", ":0", _VserviceUdpDn.name )
    flag.StringVar(&_VserviceUdpDp.hostPortStr, "cp", ":0", _VserviceUdpDp.name )
    flag.StringVar(&_VserviceUdpDC.hostPortStr, "cd", ":0", _VserviceUdpDC.name )
    flag.StringVar(&_VserviceUdpDS.hostPortStr, "cs", ":0", _VserviceUdpDS.name )

    flag.Parse()

    // _FdebugPrintTest()

    _Cexit          = make (chan string, 3   )
    _Clog           = make (chan string, 100 )
}

func main() {

    // ------------------- tcp for debug monitor log --- begin
    // _Fhandle_tcpAccept01
    // _FhandleTcp_accept_dataReceiveMsg01
    go _VserviceTcpMd . _Fhandle_udpListen_Tcp__main_top( )
    // ------------------- tcp for debug monitor log --- end

    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- begin
    // _TserviceUDP
    go _VserviceUdpDn . _Fhandle_udpListen_Udp__read_main_top( )
    //go _VserviceUdpDp . _Fhandle_udpListen_Udp__read_main_top( )
    //go _VserviceUdpDC . _Fhandle_udpListen_Udp__read_main_top( )
    //go _VserviceUdpDS . _Fhandle_udpListen_Udp__read_main_top( )
    // ------------------- udp for worker clinet : Cn , Dn , Sn --------- end

    // ------------------- filter between workers --------- begin
    _VfilterCn2dn = _TfilterDelay {
        sleepGap                : 1,
        udpIn                   : &_VserviceUdpDn,
        udpOut                  : &_VserviceUdpDC,
        //FcallbackMainDelayGen   : _FuserCallback__FilterDelay__main_swap_signal_gen__Fn,
        //FcallbackFilterChan     : _FuserCallback__FilterDelay__chan_filter__Fn,
    }

    //go _VfilterCn2dn . _FfilterDelayGen01_main_top()
    // ------------------- filter between workers --------- end

    _Fex( " the reason exit : " + <-_Cexit , nil )
} // main

