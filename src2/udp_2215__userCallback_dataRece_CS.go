// _TserviceUDP
package main

import (
    "net"
    "encoding/json"
)

type _Tsn struct {
    version                 int
    IP                      net.IP
    Port                    int
}

func _FuserCallback_u01M__dataRece_Sn(___VserviceUDP *_TserviceUDP ) {

    *___VserviceUDP.Clog <- _Pspf( "3738181 Sn receMsg |l:%s|r:%s|(%d)\n" ,
    ___VserviceUDP.VulocalAddr ,
    ___VserviceUDP.VuremoteAddr ,
    ___VserviceUDP.Vulen )

    if nil != ___VserviceUDP.CuOut01  {
        __Vcn2dn := _Tsn { 1, ___VserviceUDP.VuremoteAddr.IP , ___VserviceUDP.VuremoteAddr.Port }
        __Vbyte , __Verr := json.Marshal( __Vcn2dn )
        if nil == __Verr {
            *___VserviceUDP.CuOut01  <- __Vbyte
        }
    }

} // _FuserCallback_u01M__dataRece_Sn

