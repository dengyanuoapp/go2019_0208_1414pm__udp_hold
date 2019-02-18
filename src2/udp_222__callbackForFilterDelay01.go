package main

// _FcallbackFilterDelay_main_swap
func _FcallbackFilterDelay_main_swap( ___Vf *_TfilterDelay)    {
    //_FpfN( " 818391: filter swap user start" );
    ___Vf . CfSwap01 <- "818392 : " + _FtimeNow() // It's time to swap 
} // _FcallbackFilterDelay_main_swap

// _FcallbackFilterDelay_filter
func _FcallbackFilterDelay_filter( ___Vf *_TfilterDelay)    {

    //_FpfN( " 818395: filter select start" );
    select {
    case __Vstr := <- ___Vf.  CfSwap01 :
        //_FpfN( " 818396: filter swap received " + __Vstr )
        ___Vf . _Ftry_tran_the_ok_task_to_Dn( __Vstr )

    case __Vbyte := <- ___Vf.  CfIn01  :
        _FpfN( " 818397: filter Cin received " + string(__Vbyte) )
        ___Vf . _Ftry_update_filter_array( __Vbyte )
    }
    //_FpfN( " 818399: filter select end" );

} // _FcallbackFilterDelay_filter

func ( ___Vf *_TfilterDelay ) _Ftry_tran_the_ok_task_to_Dn(___Vstr string ) {
    _FpfN( " 828391: _Ftry_tran_the_ok_task_to_Dn : " )
    *___Vf.  CfOut01 <- []byte( " 828392 : filte out \n" )
} // _Ftry_tran_the_ok_task_to_Dn

func ( ___Vf *_TfilterDelay ) _Ftry_update_filter_array( ___Vbyte []byte ) {
    _FpfN( " 838391: update table with Cin received :" + string(___Vbyte) )
} // _Ftry_update_filter_array
