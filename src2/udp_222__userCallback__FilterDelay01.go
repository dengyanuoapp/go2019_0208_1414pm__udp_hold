package main

import (
    "encoding/json"
)

func _Fcallback_user_FilterDelay__main_swap_signal_gen( ___Vf *_TfilterDelay)    {
    //_FpfN( " 818391: filter swap user start" );
    ___Vf . CfSwap01 <- "818392 : " + _FtimeNow() // It's time to swap 
} // _Fcallback_user_FilterDelay__main_swap_signal_gen

func _Fcallback_user_FilterDelay__chan_filter( ___Vf *_TfilterDelay)    {

    //_FpfN( " 818395: filter select start" );
    select {
    case __Vstr := <- ___Vf.  CfSwap01 :
        //_FpfN( " 818396: filter swap received " + __Vstr )
        ___Vf . _Ftry_update_task_list__main_top( __Vstr )

    case __Vbyte := <- ___Vf.  CfIn01  :
        //_FpfN( " 818397: filter Cin received " + string(__Vbyte) )
        ___Vf . _Ftry_insert_new_client_req__main_top( __Vbyte )
    }
    //_FpfN( " 818399: filter select end" );

} // _Fcallback_user_FilterDelay__chan_filter

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

// _Tcn2dn 
//var _VmapCn2dn_now  _TmapCn2dn
//var _VmapCn2dn_now  map[string]_TnodeCn2dn = make( 
var _VmapCn2dn_now  _TmapCn2dn = make(_TmapCn2dn)
func ( ___Vf *_TfilterDelay ) _Ftry_insert_new_client_req__main_top( ___Vbyte []byte ) {
    _FpfN( " 838391: update table with Cin received :" + string(___Vbyte) )

    var __Vcn2dn   _Tcn2dn
    // func Unmarshal(data []byte, v interface{}) error
    __Verr := json.Unmarshal( ___Vbyte , &__Vcn2dn )
    if __Verr != nil {
        _Ppf( " 838393: update table with Cin received , met err :" ) ; _Ppt( __Verr ) ; _Pn()
        return
    }
    __VipStr := __Vcn2dn.ipStr
    _VmapCn2dn_now[__VipStr] = _TnodeCn2dn{ cnt : _VmapCn2dn_now[__VipStr].cnt + 1 , cn2dn : __Vcn2dn }

} // _Ftry_insert_new_client_req__main_top
