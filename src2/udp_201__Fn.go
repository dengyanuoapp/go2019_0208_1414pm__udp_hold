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
    _Cexit       chan string
    _Clog        chan string
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

    _Cexit          = make (chan string, 3   )
    _Clog           = make (chan string, 100 )
}

func main() {

    _FtryListenToUDP01( &_VserviceCn ) 
    _FtryListenToUDP01( &_VserviceDn ) 
    _FtryListenToUDP01( &_VserviceSn ) 

    go _FhandleFnClientCn( &_VserviceCn , _Cexit , _Clog )
    go _FhandleFnClientCn( &_VserviceDn , _Cexit , _Clog )
    go _FhandleFnClientCn( &_VserviceSn , _Cexit , _Clog )

    _Fex( " the reason exit : " + <-_Cexit , nil )
} // main

// https://golang.org/pkg/net/#UDPConn.ReadFromUDP
// func (c *UDPConn) ReadFromUDP(b []byte) (int, *UDPAddr, error)
func _FhandleFnClientCn(___VserviceUDP *_TserviceUDP, ___Cexit chan string , ___Clog chan string ) {
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
