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
    ipStr                   string
}
type _TnodeCn2dn struct {
    cnt             int
    cn2dn           _Tcn2dn
}
type _TmapCn2dn     map[string]_TnodeCn2dn

func _FuserCallback_dataRece_Cn(___VserviceUDP *_TserviceUDP ) {

    *___VserviceUDP.Clog <- _Pspf( "0738181 Cn receMsg |l:%s|r:%s|(%d)\n" ,
    ___VserviceUDP.VlocalAddr ,
    ___VserviceUDP.VremoteAddr ,
    ___VserviceUDP.Vlen )

    if nil != ___VserviceUDP.CuOut01  {
        __Vcn2dn := _Tcn2dn { 1, 
        ___VserviceUDP.VremoteAddr.IP , 
        ___VserviceUDP.VremoteAddr.Port ,
        ___VserviceUDP.VremoteAddr.IP.String() }
        // func Marshal(v interface{}) ([]byte, error)
        __Vbyte , __Verr := json.Marshal( __Vcn2dn )
        if nil == __Verr {
            *___VserviceUDP.CuOut01  <- __Vbyte
        }
    }

} // _FuserCallback_dataRece_Cn

