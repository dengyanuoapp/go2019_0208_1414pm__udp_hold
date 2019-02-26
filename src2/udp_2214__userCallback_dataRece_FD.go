// _TserviceUDP
package main

import (
	"encoding/json"
	"net"
)

type _Tdn struct {
	version int
	IP      net.IP
	Port    int
}

func _FuserCallback_u01M__dataRece_Fn__main_top(___VserviceUDP *_TserviceUDP) {

	*___VserviceUDP.Clog <- _Pspf("1738181 Fn receMsg |l:%s|r:%s|(%d)\n",
		___VserviceUDP.VulocalAddr,
		___VserviceUDP.VuremoteAddr,
		___VserviceUDP.Vulen)

	if nil != ___VserviceUDP.CuOut01 {
		__Vcn2dn := _Tdn{1, ___VserviceUDP.VuremoteAddr.IP, ___VserviceUDP.VuremoteAddr.Port}
		__Vbyte, __Verr := json.Marshal(__Vcn2dn)
		if nil == __Verr {
			*___VserviceUDP.CuOut01 <- __Vbyte
		}
	}

} // _FuserCallback_u01M__dataRece_Fn__main_top

var (
	_VdnReceCnt int
)

func _FuserCallback_chanIn_Fn__main_top(___VserviceUDP *_TserviceUDP) {
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
				_FpfN(" 2738181 (idx:%d) : rece from Chan : Fn : empty{}: (len:%d)", _VdnReceCnt, len(__VdnIn))
			}
		} else {
			___VserviceUDP._FuserCallback_chanIn_Fn__ok(&__VdnIn)
		}
		return
	}
} // _FuserCallback_chanIn_Fn__main_top

func (___VserviceUDP *_TserviceUDP) _FuserCallback_chanIn_Fn__ok(___VdnIn *[]byte) {
	_FpfN(" 3738181 (idx:%d) : rece from Chan : Fn : ok: (len:%d)"+string(*___VdnIn), _VdnReceCnt, len(*___VdnIn))
} // _FuserCallback_chanIn_Fn__ok
