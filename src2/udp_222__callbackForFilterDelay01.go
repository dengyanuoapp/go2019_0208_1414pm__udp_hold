package main

// _FcallbackFilterDelay_main_swap
func _FcallbackFilterDelay_main_swap( ___Vf *_TfilterDelay)    {
    ___Vf . Cswap01 <- "818391 : " + _FtimeNow() // It's time to swap 
} // _FcallbackFilterDelay_main_swap

// _FcallbackFilterDelay_filter
func _FcallbackFilterDelay_filter( ___Vf *_TfilterDelay)    {

    _FpfN( " 818390: filter select start" );
    select {
    case __Vstr := <- ___Vf.  Cswap01 :
        _FpfN( " 818393: filter swap received " + __Vstr )
    case __Vbyte := <- (*___Vf.  Cin01)  :
    //case <- (*___Vf.  Cin01)  :
        _FpfN( " 818395: filter Cin received " + string(__Vbyte) )
        ___Vf.  Cout01 <- []byte( " 818395 : filte out \n" )
    }

} // _FcallbackFilterDelay_filter
