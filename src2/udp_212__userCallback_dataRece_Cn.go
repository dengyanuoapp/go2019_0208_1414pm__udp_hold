// _TserviceUDP
package main

import (
    "net"
    "encoding/json"
)

type _TcnTdn struct {
    version     int
    idx         int
    Port        int
    IP          net.IP
    ipStr       string
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
        __Vcn2dn := _TcnTdn {}
        __Vcn2dn.   version = 1
        __Vcn2dn.   IP      = ___VserviceUDP.VremoteAddr.IP
        __Vcn2dn.   idx     = _VcnIdx
        __Vcn2dn.   Port    = ___VserviceUDP.VremoteAddr.Port 
        __Vcn2dn.   ipStr   =  ___VserviceUDP.VremoteAddr.IP.String() 
        _Ppf( "\n 0738184 : origin msg is: %v\n" ,  __Vcn2dn )
        __Vb2 := _FencGob( __Vcn2dn ) 
        _Ppf( "\n 0738186 : _FencBin msg is: %d , %x\n" , len(__Vb2) ,  __Vb2 )
        var __Vt3 _TcnTdn
        _FdecGob( __Vb2 , &__Vt3)
        _Ppf( "\n 0738187 : _FdecGob msg is: %v\n" ,  __Vt3 )

        // func Marshal(v interface{}) ([]byte, error)
        __Vbyte , __Verr := json.Marshal( __Vcn2dn )
        if nil != __Verr {
            _Ppt( " 0738188 : Cn pack msg error: " ,  __Verr , "\n" )
            return
        }
        _Ppt( " 0738189 : Cn pack msg as " ,  string(__Vbyte)+"\n" )
        *___VserviceUDP.CuOut01  <- __Vbyte
    }

} // _FuserCallback_dataRece_Cn

