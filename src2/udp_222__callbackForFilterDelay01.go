package main

// _FcallbackFilterDelay_main_swap
func _FcallbackFilterDelay_main_swap( ___Vf *_TfilterDelay)    {
    _FpfN( " 818391: filter swap user start" );
    ___Vf . Cswap01 <- "818392 : " + _FtimeNow() // It's time to swap 
} // _FcallbackFilterDelay_main_swap

// _FcallbackFilterDelay_filter
func _FcallbackFilterDelay_filter( ___Vf *_TfilterDelay)    {

    _FpfN( " 818395: filter select start" );
    select {
    case __Vstr := <- ___Vf.  Cswap01 :
        _FpfN( " 818396: filter swap received " + __Vstr )
    case __Vbyte := <- (*___Vf.  Cin01)  :
        //case <- (*___Vf.  Cin01)  :
        _FpfN( " 818397: filter Cin received " + string(__Vbyte) )
        ___Vf.  Cout01 <- []byte( " 818398 : filte out \n" )
        //default:
        //_Fex( "818399 : what happens ?" , nil)

    }
    _FpfN( " 818399: filter select end" );

} // _FcallbackFilterDelay_filter
