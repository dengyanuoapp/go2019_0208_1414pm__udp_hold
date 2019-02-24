// _TserviceUDP
package main

import (
    "net"
    //"encoding/json"
)

type _TuExtMRead struct {
    name            string

} // _TuExtMRead 

type _TuExtChanI struct {
} // _TuExtChanI

type _TuExtTimer struct {
} // _TuExtTimer


type _TcnTdn struct {
    Version     int
    Idx         int
    Port        int
    IP          net.IP
    IpStr       string
}

type _TnodeCn2dn struct {
    Cnt             int
    Cn2dn          _TcnTdn
}

type _TmapCn2dn     map[string]_TnodeCn2dn

var _VcnIdx int
func _FuserCallback_u01M__dataRece_Cn(___VserviceUDP *_TserviceUDP ) {

    _VcnIdx ++

    *___VserviceUDP.Clog <- _Pspf( "0738181 (%d) Cn receMsg |l:%s|r:%s|(%d)\n" ,
    _VcnIdx,
    ___VserviceUDP.VulocalAddr ,
    ___VserviceUDP.VuremoteAddr ,
    ___VserviceUDP.Vulen )

    //_Ppf( "0738183 (%d) %s \n" , _VcnIdx , ___VserviceUDP.VuremoteAddr.IP.String() )
    if nil != ___VserviceUDP.CuOut01  {
        __Vcn2dn := _TcnTdn {
            Version : 1,
            IP      : ___VserviceUDP.VuremoteAddr.IP,
            Idx     : _VcnIdx,
            Port    : ___VserviceUDP.VuremoteAddr.Port ,
            IpStr   :  ___VserviceUDP.VuremoteAddr.IP.String() ,
        }

        // func Marshal(v interface{}) ([]byte, error)
        //__Vbyte , __Verr := json.Marshal( __Vcn2dn )
        __Vbyte , _ := _FencJson( __Vcn2dn )
        //_Ppt( " 0738189 : Cn pack msg as " ,  string(__Vbyte)+"\n" )
        *___VserviceUDP.CuOut01  <- __Vbyte
    }

} // _FuserCallback_u01M__dataRece_Cn

