// _TacceptTCP
// _TtcpNodE
package main

import "io"

func _FuserCallback__accept_dataChan__Log_Fn(___VacceptTcp *_TacceptTCP) {

	var __Vbyte []byte
	var __Vstr string

	//_FpfN( "181180 (%d): log accept dataChan Select start " , ___VacceptTcp . taIdx )
	select {
	case __Vstr = <-___VacceptTcp.taCreceiveErr:
		_FpfN(" taIdx:%d : 181181 accept_dataChan receERR : len:%d , %s", ___VacceptTcp.taIdx, len(__Vstr), __Vstr)
	case __Vbyte = <-___VacceptTcp.taCreceiveMsg:
		_FpfN(" taIdx:%d : 181183 accept_dataChan receMsg : len:%d cap:%d , %s", ___VacceptTcp.taIdx, len(__Vbyte), cap(__Vbyte), __Vbyte)
	case __Vbyte = <-___VacceptTcp.taCchanMsg:
		//_FpfN( " %d : 181185 accept_dataChan chanMsg : %d %d , %s"  , ___VacceptTcp.taIdx , len(__Vbyte) , cap(__Vbyte) , __Vbyte )
		if true == ___VacceptTcp.taEnabled {
			___VacceptTcp.taW.try++
			//  func (c *TCPConn) Write(b []byte) (int, error)
			_, __Verr := ___VacceptTcp.taConnTCP.Write(__Vbyte)
			if __Verr == nil {
				___VacceptTcp.taW.ok++
			} else {
				if __Verr == io.EOF {
					___VacceptTcp.taW.eofErr++
				}
			}
		}
		//default:
		//_Fex( "181187 : what happens ?" , nil)
	}
	//_FpfN( "181189 : log accept dataChan Select end " )

} // _FuserCallback__accept_dataChan__Log_Fn
