// _TacceptTCP 
// _TserviceTCP 
package main

import "io"

// _FhandleTcp_accept_dataReceiveMsg01
func _FcallbackForDebugLog_accept_dataChan(___VacceptTcp *_TacceptTCP ) {

    var __Vbyte     []byte
    var __Vstr      string

    for {
        _Fsleep_1ms()
        _FpfN( "181180 (%d): log accept dataChan select start " , ___VacceptTcp . idx )
        select {
        case __Vstr = <- ___VacceptTcp.CreceiveErr :
            _FpfN( " idx:%d : 181181 accept_dataChan receERR : len:%d , %s"     , ___VacceptTcp.idx , len(__Vstr) , __Vstr )
        case __Vbyte = <- ___VacceptTcp.CreceiveMsg :
            _FpfN( " idx:%d : 181183 accept_dataChan receMsg : len:%d cap:%d , %s"  , ___VacceptTcp.idx , len(__Vbyte) , cap(__Vbyte) , __Vbyte )
        case __Vbyte = <- ___VacceptTcp.CchanMsg :
            //_FpfN( " %d : 181185 accept_dataChan chanMsg : %d %d , %s"  , ___VacceptTcp.idx , len(__Vbyte) , cap(__Vbyte) , __Vbyte )
            if true == ___VacceptTcp.enabled {
                ___VacceptTcp.           w64try ++
                //  func (c *TCPConn) Write(b []byte) (int, error)
                _ , __Verr := ___VacceptTcp.connTCP. Write( __Vbyte )
                if ( __Verr == nil ) {
                    ___VacceptTcp.       w64ok  ++
                } else {
                    if ( __Verr == io.EOF ) {
                        ___VacceptTcp.   w64eof  ++
                    }
                }
            }
            //default:
            //_Fex( "181185 : what happens ?" , nil)
        }
        _FpfN( "181185 : log accept dataChan select end " )

        //_FhandleTcp_accept_dataReceiveMsg01__loop( ___VacceptTCP )
    }

} // _FcallbackForDebugLog_accept_dataChan
