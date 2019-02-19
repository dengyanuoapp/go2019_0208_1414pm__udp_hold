package main

// _FcallbackFilterDelay_main_swap
func _FcallbackFilterDelay_main_swap( ___Vf *_TfilterDelay)    {
    //_FpfN( " 818391: filter swap user start" );
    ___Vf . CfSwap01 <- "818392 : " + _FtimeNow() // It's time to swap 
} // _FcallbackFilterDelay_main_swap

func _FcallbackFilterDelay_filter( ___Vf *_TfilterDelay)    {

    //_FpfN( " 818395: filter select start" );
    select {
    case __Vstr := <- ___Vf.  CfSwap01 :
        //_FpfN( " 818396: filter swap received " + __Vstr )
        ___Vf . _Ftry_update_task_list__main_top( __Vstr )

    case __Vbyte := <- ___Vf.  CfIn01  :
        _FpfN( " 818397: filter Cin received " + string(__Vbyte) )
        ___Vf . _Ftry_insert_new_client_req__main_top( __Vbyte )
    }
    //_FpfN( " 818399: filter select end" );

} // _FcallbackFilterDelay_filter

var _Vcnt_Cn2Dn int
func ( ___Vf *_TfilterDelay ) _Ftry_update_task_list__main_top(___Vstr string ) {
    _Fpf( " 828391: (%d) _Ftry_update_task_list__main_top : " , _Vcnt_Cn2Dn ) ; _Vcnt_Cn2Dn ++
    if nil == ___Vf.  CfOut01 {
        _FpfN( " out Chan is nil. " )
    } else {
        _Fpf( " 111 " )
        *___Vf.  CfOut01 <- []byte( " 828392 : filte out \n" )
        _FpfN( " ok. " )
    }
} // _Ftry_update_task_list__main_top

func ( ___Vf *_TfilterDelay ) _Ftry_insert_new_client_req__main_top( ___Vbyte []byte ) {
    _FpfN( " 838391: update table with Cin received :" + string(___Vbyte) )
} // _Ftry_insert_new_client_req__main_top
