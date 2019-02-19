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
        ___Vf . _Ftry_update_task_list__gen_and_swap_out( __Vstr )

    case __Vbyte := <- ___Vf.  CfIn01  :
        //_FpfN( " 818397: filter Cin received " + string(__Vbyte) )
        ___Vf . _Ftry_insert_new_client_req__main_top( __Vbyte )
    }
    //_FpfN( " 818399: filter select end" );

} // _Fcallback_user_FilterDelay__chan_filter

var _Vcnt_Cn2Dn int
func ( ___Vf *_TfilterDelay ) _Ftry_update_task_list__gen_and_swap_out(___Vstr string ) {
    _Vcnt_Cn2Dn ++

    _FpfN( " 828391: (%d) _Ftry_update_task_list__gen_and_swap_out : " , _Vcnt_Cn2Dn )

    if nil == ___Vf.  CfOut01 {
        _FpfN( " out Chan is nil. " )
        return
    }

    __Vbyte , __Verr := json.Marshal( _VmapCn2dn_now )
    if nil != __Verr {
        _Fpf( " 828392: (%d) why error ? " , _Vcnt_Cn2Dn); _Ppt( __Verr ) ; _Pn()
        return
    }

    _Ppn( " 828398: (%d) " , _VmapCn2dn_now )

    *___Vf.  CfOut01 <- __Vbyte

} // _Ftry_update_task_list__gen_and_swap_out

// _TcnTdn
//var _VmapCn2dn_now  _TmapCn2dn
//var _VmapCn2dn_now  map[string]_TnodeCn2dn = make(
var _VmapCn2dn_now  _TmapCn2dn = make(_TmapCn2dn)
func ( ___Vf *_TfilterDelay ) _Ftry_insert_new_client_req__main_top( ___Vbyte []byte ) {
    _FpfN( " 838391: update table with Cin received :" + string(___Vbyte) )

    var __Vcn2dn   _TcnTdn
    // func Unmarshal(data []byte, v interface{}) error
    __Verr := json.Unmarshal( ___Vbyte , &__Vcn2dn )
    if __Verr != nil {
        _Ppf( " 838393: update table with Cin received , met err :" ) ; _Ppt( __Verr ) ; _Pn()
        return
    }
    _Ppt( " 838395: unpack the json :" ) ; _Ppt( __Vcn2dn ) ; _Pn()
    __VipStr := __Vcn2dn.IpStr
    _VmapCn2dn_now[__VipStr] = _TnodeCn2dn{ cnt : _VmapCn2dn_now[__VipStr].cnt + 1 , cn2dn : __Vcn2dn }

} // _Ftry_insert_new_client_req__main_top
