// _TacceptTCP
// _TtcpNodE
package main

import "io"

func _FuserCallback__accept_dataChan__Log_Fn(___VacceptTcp *_TacceptTCP) {

	var __Vbyte []byte
	var __Vstr string

	//_FpfN( "181180 (%d): log accept dataChan Select start " , ___VacceptTcp . idx )
	select {
	case __Vstr = <-___VacceptTcp.CreceiveErr:
		_FpfN(" idx:%d : 181181 accept_dataChan receERR : len:%d , %s", ___VacceptTcp.idx, len(__Vstr), __Vstr)
	case __Vbyte = <-___VacceptTcp.CreceiveMsg:
		_FpfN(" idx:%d : 181183 accept_dataChan receMsg : len:%d cap:%d , %s", ___VacceptTcp.idx, len(__Vbyte), cap(__Vbyte), __Vbyte)
	case __Vbyte = <-___VacceptTcp.CchanMsg:
		//_FpfN( " %d : 181185 accept_dataChan chanMsg : %d %d , %s"  , ___VacceptTcp.idx , len(__Vbyte) , cap(__Vbyte) , __Vbyte )
		if true == ___VacceptTcp.enabled {
			___VacceptTcp.cW.try++
			//  func (c *TCPConn) Write(b []byte) (int, error)
			_, __Verr := ___VacceptTcp.connTCP.Write(__Vbyte)
			if __Verr == nil {
				___VacceptTcp.cW.ok++
			} else {
				if __Verr == io.EOF {
					___VacceptTcp.cW.eofErr++
				}
			}
		}
		//default:
		//_Fex( "181187 : what happens ?" , nil)
	}
	//_FpfN( "181189 : log accept dataChan Select end " )

} // _FuserCallback__accept_dataChan__Log_Fn
