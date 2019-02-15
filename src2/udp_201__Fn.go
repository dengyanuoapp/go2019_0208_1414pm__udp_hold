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

// https://golang.org/pkg/net/#UDPConn.ReadFromUDP
// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
func _FhandleFnWaitForClientCn(___VserviceUDP *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
// (*___VserviceUDP).udpConn
    //var __Vbuf [2048]byte             // array : with specified len
    __Vbuf := make( []byte , 2048 )   // silice : with var len

    __Vlen, __Vaddr, __Verr := (*___VserviceUDP).udpConn.ReadFromUDP(__Vbuf)
    //_, __Vaddr, __Verr := (*___VserviceUDP).udpConn.ReadFromUDP(__Vbuf[0:])

    if __Verr != nil {
        return
    }
    if __Vaddr == nil {
        _Fex( " 183811 : why ___Vconn.ReadFromUDP addr error ?" , nil )
    }

    //_FpdN( __Vlen , &__Vbuf )

    _Fpf( "|%s|" , __Vaddr )
    _PpdN( __Vlen , &__Vbuf )

} // _FhandleFnWaitForClientCn
