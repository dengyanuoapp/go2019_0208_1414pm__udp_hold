package main

import (
    //"encoding/json"
    "flag"
    //"fmt"
    //"log"
    //"net"
)

var (
    //_VuserIpList map[string]string
    //_VuserIpList = map[string]string{}
    _VserviceCn  _TserviceUDP   
    _VserviceDn  _TserviceUDP   
    _VserviceSn  _TserviceUDP   
)

func init() {
    _VserviceCn.name    = "servicePortForCn" 
    _VserviceDn.name    = "servicePortForDn" 
    _VserviceSn.name    = "servicePortForSn" 

    flag.StringVar(&_VserviceCn.port, "c", ":5353",  _VserviceCn.name )
    flag.StringVar(&_VserviceDn.port, "d", ":32001", _VserviceDn.name )
    flag.StringVar(&_VserviceSn.port, "s", ":32003", _VserviceSn.name )
    flag.Parse()

    // _FdebugPrintTest()
    _VprojectName = "Fn"
    _Fbase_101__get_self_md5_sha()
    _FPargs()

}

func main() {

    _FtryListenToUDP01( &_VserviceCn ) 

    for {
        _FhandleFnClientCn( &_VserviceCn )
    }
} // main

// https://golang.org/pkg/net/#UDPConn.ReadFromUDP
// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
func _FhandleFnClientCn(___VserviceUDP *_TserviceUDP) {
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

} // _FhandleFnClientCn
