// _TacceptTCP 
// _TserviceTCP 
package main

//import "io"

// _FhandleTcp_accept_dataReceiveMsg01
func ( ___VacceptTcp *_TacceptTCP ) _FhandleTcp__accept_dataChan__main_top ( ) {

    for {
        _Fsleep_1ms()
        ___VacceptTcp . _FhandleTcp__accept_dataChan__loop()

        //_FhandleTcp_accept_dataReceiveMsg01__loop( ___VacceptTCP )
    }

} // _FhandleTcp__accept_dataChan__main_top

func ( ___VacceptTcp *_TacceptTCP ) _FhandleTcp__accept_dataChan__loop ( ) {
    _FnotNullRunTcp02_accept(  ___VacceptTcp.serverTCP .TcallbackAccDataChan , ___VacceptTcp )
} // _FhandleTcp__accept_dataChan__loop

