// _TserviceUDP
package main

import (
	//"encoding/json"
	"net"
)

type _Tdn struct {
	version int
	IP      net.IP
	Port    int
}

func _FuserCallback__211_dataRece__main_top__FnWaitDn(___VserviceUDP *_TserviceUDP) {

	*___VserviceUDP.Clog <- _Pspf("1738181 01 : Fn receMsg |l:%s|r:%s|(%d)\n",
		___VserviceUDP.VulocalAddr,
		___VserviceUDP.VuremoteAddr,
		___VserviceUDP.Vulen)

	if nil != ___VserviceUDP.CuOut01 {
		//		__Vcn2dn := _Tdn{1, ___VserviceUDP.VuremoteAddr.IP, ___VserviceUDP.VuremoteAddr.Port}
		//		__Vbyte2, __Verr2 := json.Marshal(__Vcn2dn)
		//		if nil == __Verr2 {
		//			*___VserviceUDP.CuOut01 <- __Vbyte2
		//		}
		_FpfN(" 1738181 02 : %s : deal with CuOut01. ", ___VserviceUDP.name)
	} else {
		_FpfN(" 1738181 03 : %s : no CuOut01 now , skip. ", ___VserviceUDP.name)
		___VserviceUDP._FuserCallback__211x__client_req_accept_FnWaitDn()
	}

} // _FuserCallback__211_dataRece__main_top__FnWaitDn _TserviceUDP

var (
	_VdnReceCnt int
)

func _FuserCallback__221_chanIn__main_top__FnWaitDn(___VserviceUDP *_TserviceUDP) {
	select {
	case __VdnIn := <-___VserviceUDP.CuIn01:
		_VdnReceCnt++
		if 2 == len(__VdnIn) {
			__Vdb := false
			if _VdnReceCnt >= 128 {
				__Vdb = (1 == (_VdnReceCnt % 100))
			} else {
				__Vdb = (1 == (_VdnReceCnt % 10))
			}
			if __Vdb {
				_FpfN(" 2738181 01 (idx:%d) : rece from Chan : %s : empty{}: (len:%d)", _VdnReceCnt, ___VserviceUDP.name, len(__VdnIn))
			}
		} else {
			___VserviceUDP._FuserCallback__211x__chanIn_Fn__ok(&__VdnIn)
		}
		return
	}
} // _FuserCallback__221_chanIn__main_top__FnWaitDn

func (___VserviceUDP *_TserviceUDP) _FuserCallback__211x__chanIn_Fn__ok(___VdnIn *[]byte) {
	_FpfN(" 3738181 01 (idx:%d) : rece from Chan : Fn : ok: (len:%d)"+string(*___VdnIn), _VdnReceCnt, len(*___VdnIn))
} // _FuserCallback__211x__chanIn_Fn__ok
