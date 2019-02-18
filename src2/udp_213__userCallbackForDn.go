// _TserviceUDP
package main

import (
    "net"
    "encoding/json"
)

type _Tdn struct {
    version                 int
    IP                      net.IP
    Port                    int
}

func _FcallbackInFnForDn(___VserviceUDP *_TserviceUDP ) {

    *___VserviceUDP.Clog <- _Pspf( "1738181 Dn receMsg |l:%s|r:%s|(%d)\n" ,
    ___VserviceUDP.VlocalAddr ,
    ___VserviceUDP.VremoteAddr ,
    ___VserviceUDP.Vlen )

    if nil != ___VserviceUDP.CuOut01  {
        __Vcn2dn := _Tdn { 1, ___VserviceUDP.VremoteAddr.IP , ___VserviceUDP.VremoteAddr.Port }
        __Vbyte , __Verr := json.Marshal( __Vcn2dn )
        if nil == __Verr {
            *___VserviceUDP.CuOut01  <- __Vbyte
        }
    }

} // _FcallbackInFnForDn

