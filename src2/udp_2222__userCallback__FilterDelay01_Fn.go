package main

import (
	//"encoding/json"
	"sync"
)

func _FuserCallback__511_filterDelay_chan_from_FnWaitCn_to_FnWaitDn(___Vf *_TfilterDelay) {

	//_FpfN( " 818395: filter select start" );
	select {
	case __Vstr := <-___Vf.CfSwap01:
		//_FpfN( " 818396: filter swap received " + __Vstr )
		___Vf._Ftry_update_task_list__gen_and_swap_out(__Vstr)

	case __Vbyte := <-___Vf.CfIn01:
		//_FpfN( " 818397: filter Cin received " + string(__Vbyte) )
		___Vf._Ftry_insert_new_client_req__main_top(__Vbyte)
	}
	//_FpfN( " 818399: filter select end" );

} // _FuserCallback__511_filterDelay_chan_from_FnWaitCn_to_FnWaitDn

var _Vcnt_Cn2Dn int

func (___Vf *_TfilterDelay) _Ftry_update_task_list__gen_and_swap_out(___Vstr string) {
	_Vcnt_Cn2Dn++

	//_FpfN( " 828391: (%d) _Ftry_update_task_list__gen_and_swap_out : " , _Vcnt_Cn2Dn )

	if nil == ___Vf.CfOut01 {
		_FpfN(" out Chan is nil. ")
		return
	}

	//_Ppf( " 828398: (%d) %v \n"         , _Vcnt_Cn2Dn , _VmapCn2dn_now )

	_VmapCn2dn_mux.Lock() // ------ lock
	_VmapCn2dn_last = _VmapCn2dn_now
	_VmapCn2dn_now = make(_TmapCn2dn)
	_VmapCn2dn_mux.Unlock() // ------ unlock

	_VmapCn2dn_tmp = make(_TmapCn2dn)
	for __Vkey, __Vvalue := range _VmapCn2dn_last {
		if 1 == __Vvalue.Cnt {
			_VmapCn2dn_tmp[__Vkey] = __Vvalue
		}
	}
	__Vbyte, _ := _FencJson(_VmapCn2dn_tmp)

	//_Ppf( " 828399: (%d) %d , %s \n"    , _Vcnt_Cn2Dn , len(__Vbyte) , __Vbyte )

	*___Vf.CfOut01 <- __Vbyte

} // _Ftry_update_task_list__gen_and_swap_out

// _TcnTdn
//var _VmapCn2dn_now  _TmapCn2dn = make(_TmapCn2dn)
var _VmapCn2dn_now _TmapCn2dn = make(_TmapCn2dn)
var _VmapCn2dn_last _TmapCn2dn = make(_TmapCn2dn)
var _VmapCn2dn_tmp _TmapCn2dn
var _VmapCn2dn_mux sync.Mutex

func (___Vf *_TfilterDelay) _Ftry_insert_new_client_req__main_top(___Vbyte []byte) {
	//_FpfN( " 838391: update table with Cin received :" + string(___Vbyte) )

	var __Vcn2dn _TcnTdn
	// func Unmarshal(data []byte, v interface{}) error
	//__Verr := json.Unmarshal( ___Vbyte , &__Vcn2dn )
	_FdecJson___(" 838394 ", &___Vbyte, &__Vcn2dn)

	//_Ppf( " 838395: unpack the json :%x , %s \n" , __Vcn2dn , __Vcn2dn )
	__VipStr := __Vcn2dn.IpStr

	_VmapCn2dn_mux.Lock() // ------ lock
	_VmapCn2dn_now[__VipStr] = _TnodeCn2dn{Cnt: _VmapCn2dn_now[__VipStr].Cnt + 1, Cn2dn: __Vcn2dn}
	_VmapCn2dn_mux.Unlock() // ------ unlock

} // _Ftry_insert_new_client_req__main_top
