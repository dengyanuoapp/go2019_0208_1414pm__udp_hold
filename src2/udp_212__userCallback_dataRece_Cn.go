// _TserviceUDP
package main

import (
    "net"
    //"encoding/json"
)

type _TcnTdn struct {
    Version     int
    Idx         int
    Port        int
    IP          net.IP
    IpStr       string
}

type _TnodeCn2dn struct {
    cnt             int
    cn2dn          _TcnTdn
}

type _TmapCn2dn     map[string]_TnodeCn2dn

var _VcnIdx int
func _FuserCallback_dataRece_Cn(___VserviceUDP *_TserviceUDP ) {

    _VcnIdx ++

    *___VserviceUDP.Clog <- _Pspf( "0738181 (%d) Cn receMsg |l:%s|r:%s|(%d)\n" ,
    _VcnIdx,
    ___VserviceUDP.VlocalAddr ,
    ___VserviceUDP.VremoteAddr ,
    ___VserviceUDP.Vlen )

    _Ppf( "0738183 (%d) %s \n" , _VcnIdx , ___VserviceUDP.VremoteAddr.IP.String() )
    if nil != ___VserviceUDP.CuOut01  {
        __Vcn2dn := _TcnTdn {
            Version : 1,
            IP      : ___VserviceUDP.VremoteAddr.IP,
            Idx     : _VcnIdx,
            Port    : ___VserviceUDP.VremoteAddr.Port ,
            IpStr   :  ___VserviceUDP.VremoteAddr.IP.String() ,
        }

        // func Marshal(v interface{}) ([]byte, error)
        //__Vbyte , __Verr := json.Marshal( __Vcn2dn )
        __Vbyte , _ := _FencJson( __Vcn2dn )
        _Ppt( " 0738189 : Cn pack msg as " ,  string(__Vbyte)+"\n" )
        *___VserviceUDP.CuOut01  <- __Vbyte
    }

} // _FuserCallback_dataRece_Cn

