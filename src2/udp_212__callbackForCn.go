// _TserviceUDP
package main

import (
    "net"
    "encoding/json"
)

type _Tcn2dn struct {
    version                 int
    IP                      net.IP
    Port                    int
}

func _FcallbackInFnForCn(___VserviceUDP *_TserviceUDP ) {

    //_Fpf( "341011 udpRdump |l:%s|r:%s|" , (*___VserviceUDP.VlocalAddr , (*___VserviceUDP.VremoteAddr )
    //_PpdN( (*___VserviceUDP.Vlen , &(*___VserviceUDP.Vbuf )
    //(*(*___VserviceUDP.Clog) <- _Pspf( "341012 udpRdump |l:%s|r:%s|(%d)\n" , 
    //(*(___VserviceUDP.Clog)) <- _Pspf( "341012 udpRdump |l:%s|r:%s|(%d)\n" , 
    //___VserviceUDP.Clog <- _Pspf( "341012 udpRdump |l:%s|r:%s|(%d)\n" , 
    *___VserviceUDP.Clog <- _Pspf( "341012 udpRdump |l:%s|r:%s|(%d)\n" ,
    ___VserviceUDP.VlocalAddr ,
    ___VserviceUDP.VremoteAddr ,
    ___VserviceUDP.Vlen )

    if nil != ___VserviceUDP.CuOut01  {
        __Vcn2dn := _Tcn2dn { 1, ___VserviceUDP.VremoteAddr.IP , ___VserviceUDP.VremoteAddr.Port }
        __Vbyte , __Verr := json.Marshal( __Vcn2dn )
        if nil == __Verr {
            *___VserviceUDP.CuOut01  <- __Vbyte
        }
    }

} // _FcallbackInFnForCn

